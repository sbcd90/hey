![hey](http://i.imgur.com/szzD9q0.png)

[![Build Status](https://travis-ci.org/rakyll/hey.svg?branch=master)](https://travis-ci.org/rakyll/hey)

hey-for-TF-serving is a tiny program that sends some load to a tensorflow-serving grpc server hosted on Knative-serving/K8S.

## Installation

    go get -u github.com/sbcd90/hey
    
## Generating Tensorflow-serving stubs

A nice documentation on how to generate TF Grpc stubs in Go can be found [here](https://mauri870.github.io/blog/posts/tensorflow-serving-inception-go/).

## Usage

hey runs provided number of requests in the provided concurrency level and prints stats.

This fork of hey is intended to support a tensorflow-serving grpc server hosted on Knative-serving/K8S.

```
Usage: hey [options...] <url>

Options:
  -n  Number of requests to run. Default is 200.
  -c  Number of requests to run concurrently. Total number of requests cannot
      be smaller than the concurrency level. Default is 50.
  -q  Rate limit, in queries per second (QPS). Default is no rate limit.
  -z  Duration of application to send requests. When duration is reached,
      application stops and exits. If duration is specified, n is ignored.
      Examples: -z 10s -z 3m.
  -o  Output type. If none provided, a summary is printed.
      "csv" is the only supported alternative. Dumps the response
      metrics in comma-separated values format.

  -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
  -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
  -A  HTTP Accept header.
  -d  HTTP request body.
  -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
  -T  Content-type, defaults to "text/html".
  -a  Basic authentication, username:password.
  -x  HTTP Proxy address as host:port.
  -h2 Enable HTTP/2.
  -g  Enable GRPC.
  -M  Model Name.
  -v  Model Version.
  -k  Input Key.
  -d  Input Data.

  -host	HTTP Host header.

  -disable-compression  Disable compression.
  -disable-keepalive    Disable keep-alive, prevents re-use of TCP
                        connections between different HTTP requests.
  -disable-redirects    Disable following of HTTP redirects
  -cpus                 Number of used cpu cores.
                        (default for current machine is %d cores)
  -server-host-override Server Host Override option in Grpc
  -insecure             Insecure option in Grpc
```

Note: Requires go 1.7 or greater.

## Examples for Tensorflow Grpc interface

```
-n 10000 -c 100 -mo tensorflow-feature-sum-model -v 1 -k X -d 1,2,3,4 -server-host-override grpc-model-server.default.example.com -insecure -g 192.168.140.140:31380
```
