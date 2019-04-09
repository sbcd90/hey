// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package requester provides commands to run load tests and display results.
package requester

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	google_protobuf "github.com/gogo/protobuf/types"
	"golang.org/x/net/http2"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	credentials2 "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"os"
	"sync"
	"tensorflow/core/framework"
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"
	"time"
)

// Max size of the buffer of result channel.
const maxResult = 1000000
const maxIdleConn = 500

type result struct {
	err           error
	statusCode    int
	offset        time.Duration
	duration      time.Duration
	connDuration  time.Duration // connection setup(DNS lookup + Dial up) duration
	dnsDuration   time.Duration // dns lookup duration
	reqDuration   time.Duration // request "write" duration
	resDuration   time.Duration // response "read" duration
	delayDuration time.Duration // delay between response and request
	contentLength int64
}

type Work struct {
	// Request is the request to be made.
	Request *http.Request

	RequestBody []byte

	GrpcUrl string

	ModelName string

	ModelVersion int32

	InputKey string

	InputData []interface{}

	CertBytes []byte

	InputFile string

	// N is the total number of requests to make.
	N int

	// C is the concurrency level, the number of concurrent workers to run.
	C int

	// H2 is an option to make HTTP/2 requests
	H2 bool

	// GRPC is an option to make TF Serving GRPC requests
	GRPC bool

	// Timeout in seconds.
	Timeout int

	// Qps is the rate limit in queries per second.
	QPS float64

	// DisableCompression is an option to disable compression in response
	DisableCompression bool

	// DisableKeepAlives is an option to prevents re-use of TCP connections between different HTTP requests
	DisableKeepAlives bool

	// DisableRedirects is an option to prevent the following of HTTP redirects
	DisableRedirects bool

	// Server Host Override for Tf serving Grpc requests.
	ServerHostOverride string

	OAuthToken string

	// Grpc insecure flag
	InsecureFlag bool

	// Output represents the output type. If "csv" is provided, the
	// output will be dumped as a csv stream.
	Output string

	// ProxyAddr is the address of HTTP proxy server in the format on "host:port".
	// Optional.
	ProxyAddr *url.URL

	// Writer is where results will be written. If nil, results are written to stdout.
	Writer io.Writer

	initOnce sync.Once
	results  chan *result
	stopCh   chan struct{}
	start    time.Duration

	report *report
}

func (b *Work) writer() io.Writer {
	if b.Writer == nil {
		return os.Stdout
	}
	return b.Writer
}

// Init initializes internal data-structures
func (b *Work) Init() {
	b.initOnce.Do(func() {
		b.results = make(chan *result, min(b.C*1000, maxResult))
		b.stopCh = make(chan struct{}, b.C)
	})
}

// Run makes all the requests, prints the summary. It blocks until
// all work is done.
func (b *Work) Run() {
	b.Init()
	b.start = now()
	b.report = newReport(b.writer(), b.results, b.Output, b.N)
	// Run the reporter first, it polls the result channel until it is closed.
	go func() {
		runReporter(b.report)
	}()
	if !b.GRPC {
		b.runWorkers()
	} else {
		b.runTFGrpcWorkers()
	}
	b.Finish()
}

func (b *Work) Stop() {
	// Send stop signal so that workers can stop gracefully.
	for i := 0; i < b.C; i++ {
		b.stopCh <- struct{}{}
	}
}

func (b *Work) Finish() {
	close(b.results)
	total := now() - b.start
	// Wait until the reporter is done.
	<-b.report.done
	b.report.finalize(total)
}

func (b *Work) makeRequest(c *http.Client) {
	s := now()
	var size int64
	var code int
	var dnsStart, connStart, resStart, reqStart, delayStart time.Duration
	var dnsDuration, connDuration, resDuration, reqDuration, delayDuration time.Duration
	req := cloneRequest(b.Request, b.RequestBody)
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dnsDuration = now() - dnsStart
		},
		GetConn: func(h string) {
			connStart = now()
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			if !connInfo.Reused {
				connDuration = now() - connStart
			}
			reqStart = now()
		},
		WroteRequest: func(w httptrace.WroteRequestInfo) {
			reqDuration = now() - reqStart
			delayStart = now()
		},
		GotFirstResponseByte: func() {
			delayDuration = now() - delayStart
			resStart = now()
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	resp, err := c.Do(req)
	if err == nil {
		size = resp.ContentLength
		code = resp.StatusCode
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
	t := now()
	resDuration = t - resStart
	finish := t - s
	b.results <- &result{
		offset:        s,
		statusCode:    code,
		duration:      finish,
		err:           err,
		contentLength: size,
		connDuration:  connDuration,
		dnsDuration:   dnsDuration,
		reqDuration:   reqDuration,
		resDuration:   resDuration,
		delayDuration: delayDuration,
	}
}

func (b *Work) makeTFGrpcRequest(c *pb.PredictionServiceClient, statsHandler *ClientStatsHandlerImpl)  {
	s := now()
	var size int64
	var code int
	var resStart time.Duration
	var connDuration, resDuration time.Duration
	req := cloneTFGrpcRequest(&pb.PredictRequest{}, b.ModelName, int64(b.ModelVersion),
		b.InputKey, b.InputData)
	resp, err := (*c).Predict(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while calling Grpc")
		fmt.Fprintf(os.Stderr, "")
		code = 500
		os.Exit(1)
	} else {
		code = 200
	}
	resStart = (*statsHandler).ResStart
	size = int64(resp.Size())
	t := now()
	resDuration = t - resStart
	finish := t - s
	connDuration = (*statsHandler).ConnEnd.Sub((*statsHandler).ConnStart).Round(time.Millisecond)
	b.results <- &result{
		offset:        s,
		statusCode:    code,
		duration:      finish,
		err:           err,
		contentLength: size,
		connDuration:  connDuration,
		resDuration:   resDuration,
	}
}

func (b *Work) runWorker(client *http.Client, n int) {
	var throttle <-chan time.Time
	if b.QPS > 0 {
		throttle = time.Tick(time.Duration(1e6/(b.QPS)) * time.Microsecond)
	}

	if b.DisableRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	for i := 0; i < n; i++ {
		// Check if application is stopped. Do not send into a closed channel.
		select {
		case <-b.stopCh:
			return
		default:
			if b.QPS > 0 {
				<-throttle
			}
			b.makeRequest(client)
		}
	}
}

func (b *Work) runTFGrpcWorker(client *pb.PredictionServiceClient, statsHandler *ClientStatsHandlerImpl, n int) {
	var throttle <-chan time.Time
	for i := 0; i < n; i++ {
		select {
		case <-b.stopCh:
			return
		default:
			if b.QPS > 0 {
				<-throttle
			}
			b.makeTFGrpcRequest(client, statsHandler)
		}
	}
}

func (b *Work) runWorkers() {
	var wg sync.WaitGroup
	wg.Add(b.C)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         b.Request.Host,
		},
		MaxIdleConnsPerHost: min(b.C, maxIdleConn),
		DisableCompression:  b.DisableCompression,
		DisableKeepAlives:   b.DisableKeepAlives,
		Proxy:               http.ProxyURL(b.ProxyAddr),
	}
	if b.H2 {
		http2.ConfigureTransport(tr)
	} else {
		tr.TLSNextProto = make(map[string]func(string, *tls.Conn) http.RoundTripper)
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(b.Timeout) * time.Second}

	// Ignore the case where b.N % b.C != 0.
	for i := 0; i < b.C; i++ {
		go func() {
			b.runWorker(client, b.N/b.C)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (b *Work) runTFGrpcWorkers() {
	var wg sync.WaitGroup
	wg.Add(b.C)

	var opts []grpc.DialOption
	if b.ServerHostOverride != "" {
		opts = append(opts, grpc.WithAuthority(b.ServerHostOverride))
	}
	if b.InsecureFlag {
		opts = append(opts, grpc.WithInsecure())
	}

	perRPC := oauth.NewOauthAccess(b.fetchToken())
	if perRPC != nil {
		opts = append(opts, grpc.WithPerRPCCredentials(perRPC))
	}

	credentials := b.fetchSSLContext()
	if credentials != nil {
		opts = append(opts, grpc.WithTransportCredentials(credentials))
	}

	statsHandler := NewClientStatusHandler()
	ClientStatsHandler, ok := statsHandler.(*ClientStatsHandlerImpl)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error with stats handler")
		fmt.Fprintf(os.Stderr, "")
		os.Exit(1)
	}

	opts = append(opts, grpc.WithStatsHandler(statsHandler))

	if b.GrpcUrl != "" {
		conn, err := grpc.Dial(b.GrpcUrl, opts...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while opening Grpc conn")
			fmt.Fprintf(os.Stderr, "")
			os.Exit(1)
		}

		client := pb.NewPredictionServiceClient(conn)
		for i := 0; i < b.C; i++ {
			go func() {
				b.runTFGrpcWorker(&client, ClientStatsHandler, b.N/b.C)
				wg.Done()
			}()
		}
		wg.Wait()
	} else {
		fmt.Fprintf(os.Stderr, "Error while creating Grpc client")
		fmt.Fprintf(os.Stderr, "")
		os.Exit(1)
	}
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request, body []byte) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	if len(body) > 0 {
		r2.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return r2
}

// cloneTFGrpcRequest returns a clone of the provided *pb.PredictRequest.
// The clone is a shallow copy of the struct and its Header map.
func cloneTFGrpcRequest(r *pb.PredictRequest, ModelName string, ModelVersion int64,
						InputKey string, InputData []interface{}) *pb.PredictRequest {
	// shallow copy of the struct
	r2 := new(pb.PredictRequest)
	*r2 = *r
	// deep copy of the Header
	var Dtype tensorflow.DataType
	tensorShapeProtoDim := tf_core_framework.TensorShapeProto_Dim{
		Size_: int64(len(InputData)),
	}

	tensorShapeProtoDimArr := []*tf_core_framework.TensorShapeProto_Dim{&tensorShapeProtoDim}

	var int16Arr   []int16
	var int32Arr   []int32
	var int64Arr   []int64
	var float32Arr []float32
	var float64Arr []float64
	var stringArr  [][]byte


	for i := 0; i < len(InputData); i++ {
		switch v := InputData[i].(type) {
		case int32:
			Dtype = tf_core_framework.DataType_DT_INT32
			int32Arr = append(int32Arr, v)
		case int16:
			Dtype = tf_core_framework.DataType_DT_INT16
			int16Arr = append(int16Arr, v)
		case int64:
			Dtype = tf_core_framework.DataType_DT_INT64
			int64Arr = append(int64Arr, v)
		case int:
			Dtype = tf_core_framework.DataType_DT_INT32
			int32Arr = append(int32Arr, int32(v))
		case float32:
			Dtype = tf_core_framework.DataType_DT_FLOAT
			float32Arr = append(float32Arr, v)
		case float64:
			Dtype = tf_core_framework.DataType_DT_DOUBLE
			float64Arr = append(float64Arr, v)
		case string:
			Dtype = tf_core_framework.DataType_DT_STRING
			stringArr = append(stringArr, []byte(v))
		default:
			fmt.Fprintf(os.Stderr, "The type of input could not be understood")
			fmt.Fprintf(os.Stderr, "")
			os.Exit(1)
		}
	}

	var inputProto tf_core_framework.TensorProto
	if len(int32Arr) > 0 {
		inputProto = tf_core_framework.TensorProto{
			Dtype: Dtype,
			TensorShape: &tf_core_framework.TensorShapeProto{
				Dim: tensorShapeProtoDimArr,
			},
			IntVal: int32Arr,
		}
	}
	if len(int64Arr) > 0 {
		inputProto = tf_core_framework.TensorProto{
			Dtype: Dtype,
			TensorShape: &tf_core_framework.TensorShapeProto{
				Dim: tensorShapeProtoDimArr,
			},
			Int64Val: int64Arr,
		}
	}
	if len(float32Arr) > 0 {
		inputProto = tf_core_framework.TensorProto{
			Dtype: Dtype,
			TensorShape: &tf_core_framework.TensorShapeProto{
				Dim: tensorShapeProtoDimArr,
			},
			FloatVal: float32Arr,
		}
	}
	if len(float64Arr) > 0 {
		inputProto = tf_core_framework.TensorProto{
			Dtype: Dtype,
			TensorShape: &tf_core_framework.TensorShapeProto{
				Dim: tensorShapeProtoDimArr,
			},
			DoubleVal: float64Arr,
		}
	}
	if len(stringArr) > 0 {
		inputProto = tf_core_framework.TensorProto{
			Dtype: Dtype,
			TensorShape: &tf_core_framework.TensorShapeProto{
				Dim: tensorShapeProtoDimArr,
			},
			StringVal: stringArr,
		}
	}
	modelSpec := pb.ModelSpec{
		Name: ModelName,
		SignatureName: ModelName,
		Version: &google_protobuf.Int64Value{Value: ModelVersion},
	}
	r2.Inputs = map[string]*tf_core_framework.TensorProto{
		InputKey: &inputProto,}
	r2.ModelSpec = &modelSpec
	return r2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (b *Work) fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: b.OAuthToken,
		TokenType: "bearer",
		Expiry: time.Now().Add(43199 * time.Second),
	}
}

func (b *Work) fetchSSLContext() credentials2.TransportCredentials {
	certBytes := b.CertBytes

	cert := &tls.Certificate{
		Certificate:  [][]byte{certBytes},
	}
	credentials := credentials2.NewServerTLSFromCert(cert)
	return credentials
}
