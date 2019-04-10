// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow_serving/apis/inference.proto

package tensorflow_serving

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Inference request such as classification, regression, etc...
type InferenceTask struct {
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec,omitempty"`
	// Signature's method_name. Should be one of the method names defined in
	// third_party/tensorflow/python/saved_model/signature_constants.py.
	// e.g. "tensorflow/serving/classify".
	MethodName string `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
}

func (m *InferenceTask) Reset()                    { *m = InferenceTask{} }
func (m *InferenceTask) String() string            { return proto.CompactTextString(m) }
func (*InferenceTask) ProtoMessage()               {}
func (*InferenceTask) Descriptor() ([]byte, []int) { return fileDescriptorInference, []int{0} }

func (m *InferenceTask) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *InferenceTask) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

// Inference result, matches the type of request or is an error.
type InferenceResult struct {
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*InferenceResult_ClassificationResult
	//	*InferenceResult_RegressionResult
	Result isInferenceResult_Result `protobuf_oneof:"result"`
}

func (m *InferenceResult) Reset()                    { *m = InferenceResult{} }
func (m *InferenceResult) String() string            { return proto.CompactTextString(m) }
func (*InferenceResult) ProtoMessage()               {}
func (*InferenceResult) Descriptor() ([]byte, []int) { return fileDescriptorInference, []int{1} }

type isInferenceResult_Result interface {
	isInferenceResult_Result()
	MarshalTo([]byte) (int, error)
	Size() int
}

type InferenceResult_ClassificationResult struct {
	ClassificationResult *ClassificationResult `protobuf:"bytes,2,opt,name=classification_result,json=classificationResult,oneof"`
}
type InferenceResult_RegressionResult struct {
	RegressionResult *RegressionResult `protobuf:"bytes,3,opt,name=regression_result,json=regressionResult,oneof"`
}

func (*InferenceResult_ClassificationResult) isInferenceResult_Result() {}
func (*InferenceResult_RegressionResult) isInferenceResult_Result()     {}

func (m *InferenceResult) GetResult() isInferenceResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *InferenceResult) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *InferenceResult) GetClassificationResult() *ClassificationResult {
	if x, ok := m.GetResult().(*InferenceResult_ClassificationResult); ok {
		return x.ClassificationResult
	}
	return nil
}

func (m *InferenceResult) GetRegressionResult() *RegressionResult {
	if x, ok := m.GetResult().(*InferenceResult_RegressionResult); ok {
		return x.RegressionResult
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*InferenceResult) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _InferenceResult_OneofMarshaler, _InferenceResult_OneofUnmarshaler, _InferenceResult_OneofSizer, []interface{}{
		(*InferenceResult_ClassificationResult)(nil),
		(*InferenceResult_RegressionResult)(nil),
	}
}

func _InferenceResult_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*InferenceResult)
	// result
	switch x := m.Result.(type) {
	case *InferenceResult_ClassificationResult:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClassificationResult); err != nil {
			return err
		}
	case *InferenceResult_RegressionResult:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RegressionResult); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("InferenceResult.Result has unexpected type %T", x)
	}
	return nil
}

func _InferenceResult_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*InferenceResult)
	switch tag {
	case 2: // result.classification_result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClassificationResult)
		err := b.DecodeMessage(msg)
		m.Result = &InferenceResult_ClassificationResult{msg}
		return true, err
	case 3: // result.regression_result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegressionResult)
		err := b.DecodeMessage(msg)
		m.Result = &InferenceResult_RegressionResult{msg}
		return true, err
	default:
		return false, nil
	}
}

func _InferenceResult_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*InferenceResult)
	// result
	switch x := m.Result.(type) {
	case *InferenceResult_ClassificationResult:
		s := proto.Size(x.ClassificationResult)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *InferenceResult_RegressionResult:
		s := proto.Size(x.RegressionResult)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Inference request containing one or more requests.
type MultiInferenceRequest struct {
	// Inference tasks.
	Tasks []*InferenceTask `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
	// Input data.
	Input *Input `protobuf:"bytes,2,opt,name=input" json:"input,omitempty"`
}

func (m *MultiInferenceRequest) Reset()                    { *m = MultiInferenceRequest{} }
func (m *MultiInferenceRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiInferenceRequest) ProtoMessage()               {}
func (*MultiInferenceRequest) Descriptor() ([]byte, []int) { return fileDescriptorInference, []int{2} }

func (m *MultiInferenceRequest) GetTasks() []*InferenceTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *MultiInferenceRequest) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

// Inference request containing one or more responses.
type MultiInferenceResponse struct {
	// List of results; one for each InferenceTask in the request, returned in the
	// same order as the request.
	Results []*InferenceResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *MultiInferenceResponse) Reset()                    { *m = MultiInferenceResponse{} }
func (m *MultiInferenceResponse) String() string            { return proto.CompactTextString(m) }
func (*MultiInferenceResponse) ProtoMessage()               {}
func (*MultiInferenceResponse) Descriptor() ([]byte, []int) { return fileDescriptorInference, []int{3} }

func (m *MultiInferenceResponse) GetResults() []*InferenceResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*InferenceTask)(nil), "tensorflow.serving.InferenceTask")
	proto.RegisterType((*InferenceResult)(nil), "tensorflow.serving.InferenceResult")
	proto.RegisterType((*MultiInferenceRequest)(nil), "tensorflow.serving.MultiInferenceRequest")
	proto.RegisterType((*MultiInferenceResponse)(nil), "tensorflow.serving.MultiInferenceResponse")
}
func (m *InferenceTask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InferenceTask) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ModelSpec != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInference(dAtA, i, uint64(m.ModelSpec.Size()))
		n1, err := m.ModelSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.MethodName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInference(dAtA, i, uint64(len(m.MethodName)))
		i += copy(dAtA[i:], m.MethodName)
	}
	return i, nil
}

func (m *InferenceResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InferenceResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ModelSpec != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInference(dAtA, i, uint64(m.ModelSpec.Size()))
		n2, err := m.ModelSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Result != nil {
		nn3, err := m.Result.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *InferenceResult_ClassificationResult) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ClassificationResult != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInference(dAtA, i, uint64(m.ClassificationResult.Size()))
		n4, err := m.ClassificationResult.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *InferenceResult_RegressionResult) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.RegressionResult != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintInference(dAtA, i, uint64(m.RegressionResult.Size()))
		n5, err := m.RegressionResult.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *MultiInferenceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiInferenceRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tasks) > 0 {
		for _, msg := range m.Tasks {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInference(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Input != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInference(dAtA, i, uint64(m.Input.Size()))
		n6, err := m.Input.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *MultiInferenceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiInferenceResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			dAtA[i] = 0xa
			i++
			i = encodeVarintInference(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintInference(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InferenceTask) Size() (n int) {
	var l int
	_ = l
	if m.ModelSpec != nil {
		l = m.ModelSpec.Size()
		n += 1 + l + sovInference(uint64(l))
	}
	l = len(m.MethodName)
	if l > 0 {
		n += 1 + l + sovInference(uint64(l))
	}
	return n
}

func (m *InferenceResult) Size() (n int) {
	var l int
	_ = l
	if m.ModelSpec != nil {
		l = m.ModelSpec.Size()
		n += 1 + l + sovInference(uint64(l))
	}
	if m.Result != nil {
		n += m.Result.Size()
	}
	return n
}

func (m *InferenceResult_ClassificationResult) Size() (n int) {
	var l int
	_ = l
	if m.ClassificationResult != nil {
		l = m.ClassificationResult.Size()
		n += 1 + l + sovInference(uint64(l))
	}
	return n
}
func (m *InferenceResult_RegressionResult) Size() (n int) {
	var l int
	_ = l
	if m.RegressionResult != nil {
		l = m.RegressionResult.Size()
		n += 1 + l + sovInference(uint64(l))
	}
	return n
}
func (m *MultiInferenceRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Tasks) > 0 {
		for _, e := range m.Tasks {
			l = e.Size()
			n += 1 + l + sovInference(uint64(l))
		}
	}
	if m.Input != nil {
		l = m.Input.Size()
		n += 1 + l + sovInference(uint64(l))
	}
	return n
}

func (m *MultiInferenceResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovInference(uint64(l))
		}
	}
	return n
}

func sovInference(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInference(x uint64) (n int) {
	return sovInference(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InferenceTask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInference
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InferenceTask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InferenceTask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ModelSpec == nil {
				m.ModelSpec = &ModelSpec{}
			}
			if err := m.ModelSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInference(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInference
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InferenceResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInference
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InferenceResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InferenceResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ModelSpec == nil {
				m.ModelSpec = &ModelSpec{}
			}
			if err := m.ModelSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationResult", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ClassificationResult{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Result = &InferenceResult_ClassificationResult{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegressionResult", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RegressionResult{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Result = &InferenceResult_RegressionResult{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInference(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInference
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiInferenceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInference
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MultiInferenceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiInferenceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tasks = append(m.Tasks, &InferenceTask{})
			if err := m.Tasks[len(m.Tasks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Input == nil {
				m.Input = &Input{}
			}
			if err := m.Input.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInference(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInference
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiInferenceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInference
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MultiInferenceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiInferenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInference
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInference
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &InferenceResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInference(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInference
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInference(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInference
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInference
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInference
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthInference
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInference
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipInference(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthInference = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInference   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tensorflow_serving/apis/inference.proto", fileDescriptorInference) }

var fileDescriptorInference = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0xbf, 0x85, 0xc0, 0x27, 0xdb, 0x18, 0x74, 0x23, 0x06, 0x49, 0x44, 0x2c, 0x26, 0xf6,
	0x60, 0x4a, 0x82, 0x07, 0x2f, 0x7a, 0xc1, 0x8b, 0x1e, 0xf0, 0xb0, 0x98, 0x78, 0x6c, 0x6a, 0x19,
	0xb0, 0xa1, 0xdd, 0xad, 0xbb, 0x5b, 0x3d, 0xfb, 0x04, 0xbe, 0x96, 0x47, 0x1f, 0xc1, 0xf0, 0x14,
	0x1e, 0x8d, 0x5d, 0x5a, 0x04, 0x8b, 0x1e, 0xbc, 0x35, 0x93, 0xdf, 0xfc, 0xfe, 0x33, 0xd3, 0xc5,
	0x87, 0x0a, 0x98, 0xe4, 0x62, 0x14, 0xf0, 0x47, 0x47, 0x82, 0x78, 0xf0, 0xd9, 0xb8, 0xe3, 0x46,
	0xbe, 0xec, 0xf8, 0x6c, 0x04, 0x02, 0x98, 0x07, 0x76, 0x24, 0xb8, 0xe2, 0x84, 0xcc, 0x41, 0x7b,
	0x06, 0x36, 0x8e, 0x56, 0x35, 0x7b, 0x81, 0x2b, 0xa5, 0x3f, 0xf2, 0x3d, 0x57, 0xf9, 0x9c, 0x69,
	0x43, 0xa3, 0xbd, 0x3a, 0x2a, 0x8a, 0xd5, 0x6f, 0x50, 0xc8, 0x87, 0x10, 0xcc, 0x20, 0x6b, 0x15,
	0x24, 0x60, 0x2c, 0x40, 0xca, 0x2c, 0xd3, 0x64, 0x78, 0xfd, 0x32, 0x5d, 0xe4, 0xda, 0x95, 0x13,
	0x72, 0x8a, 0x71, 0x62, 0x72, 0x64, 0x04, 0x5e, 0x1d, 0xb5, 0x90, 0x65, 0x74, 0x77, 0xed, 0xef,
	0xbb, 0xd9, 0xfd, 0x4f, 0x6a, 0x10, 0x81, 0x47, 0x2b, 0x61, 0xfa, 0x49, 0xf6, 0xb0, 0x11, 0x82,
	0xba, 0xe3, 0x43, 0x87, 0xb9, 0x21, 0xd4, 0x0b, 0x2d, 0x64, 0x55, 0x28, 0xd6, 0xa5, 0x2b, 0x37,
	0x04, 0xf3, 0xb9, 0x80, 0xab, 0x59, 0x20, 0x05, 0x19, 0x07, 0xea, 0x8f, 0x91, 0x0e, 0xae, 0x2d,
	0x5e, 0xd3, 0x11, 0x89, 0x36, 0x09, 0x37, 0xba, 0x56, 0x9e, 0xe8, 0x7c, 0xa1, 0x41, 0x8f, 0x71,
	0xf1, 0x8f, 0x6e, 0x79, 0x39, 0x75, 0x32, 0xc0, 0x9b, 0xf3, 0xb3, 0xa5, 0xf2, 0x62, 0x22, 0x3f,
	0xc8, 0x93, 0xd3, 0x0c, 0xce, 0xc4, 0x1b, 0x62, 0xa9, 0xd6, 0x5b, 0xc3, 0x65, 0x6d, 0x32, 0x9f,
	0x10, 0xae, 0xf5, 0xe3, 0x40, 0xf9, 0x5f, 0xce, 0x72, 0x1f, 0x83, 0x54, 0xe4, 0x04, 0x97, 0x94,
	0x2b, 0x27, 0xb2, 0x8e, 0x5a, 0x45, 0xcb, 0xe8, 0xee, 0xe7, 0x85, 0x2d, 0xfc, 0x3c, 0xaa, 0x79,
	0xd2, 0xc1, 0xa5, 0xe4, 0xc9, 0xcc, 0x4e, 0xb0, 0x93, 0xdf, 0x18, 0xc5, 0x8a, 0x6a, 0xce, 0xbc,
	0xc1, 0xdb, 0xcb, 0x23, 0xc8, 0x88, 0x33, 0x09, 0xe4, 0x0c, 0xff, 0xd7, 0x73, 0xa6, 0x53, 0xb4,
	0x7f, 0x9c, 0x42, 0x6f, 0x47, 0xd3, 0x9e, 0x5e, 0xf5, 0x65, 0xda, 0x44, 0xaf, 0xd3, 0x26, 0x7a,
	0x9b, 0x36, 0xd1, 0x3b, 0x42, 0xb7, 0xe5, 0xe4, 0xd9, 0x1d, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xe1, 0x06, 0x95, 0xb3, 0x57, 0x03, 0x00, 0x00,
}