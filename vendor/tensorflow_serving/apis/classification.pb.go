// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow_serving/apis/classification.proto

/*
	Package tensorflow_serving is a generated protocol buffer package.

	It is generated from these files:
		tensorflow_serving/apis/classification.proto
		tensorflow_serving/apis/get_model_metadata.proto
		tensorflow_serving/apis/inference.proto
		tensorflow_serving/apis/input.proto
		tensorflow_serving/apis/model.proto
		tensorflow_serving/apis/predict.proto
		tensorflow_serving/apis/prediction_service.proto
		tensorflow_serving/apis/regression.proto

	It has these top-level messages:
		Class
		Classifications
		ClassificationResult
		ClassificationRequest
		ClassificationResponse
		SignatureDefMap
		GetModelMetadataRequest
		GetModelMetadataResponse
		InferenceTask
		InferenceResult
		MultiInferenceRequest
		MultiInferenceResponse
		ExampleList
		ExampleListWithContext
		Input
		ModelSpec
		PredictRequest
		PredictResponse
		Regression
		RegressionResult
		RegressionRequest
		RegressionResponse
*/
package tensorflow_serving

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A single class.
type Class struct {
	// Label or name of the class.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Score for this class (e.g., the probability the item belongs to this
	// class).
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *Class) Reset()                    { *m = Class{} }
func (m *Class) String() string            { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()               {}
func (*Class) Descriptor() ([]byte, []int) { return fileDescriptorClassification, []int{0} }

func (m *Class) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Class) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// List of classes for a single item (tensorflow.Example).
type Classifications struct {
	Classes []*Class `protobuf:"bytes,1,rep,name=classes" json:"classes,omitempty"`
}

func (m *Classifications) Reset()                    { *m = Classifications{} }
func (m *Classifications) String() string            { return proto.CompactTextString(m) }
func (*Classifications) ProtoMessage()               {}
func (*Classifications) Descriptor() ([]byte, []int) { return fileDescriptorClassification, []int{1} }

func (m *Classifications) GetClasses() []*Class {
	if m != nil {
		return m.Classes
	}
	return nil
}

// Contains one result per input example, in the same order as the input in
// ClassificationRequest.
type ClassificationResult struct {
	Classifications []*Classifications `protobuf:"bytes,1,rep,name=classifications" json:"classifications,omitempty"`
}

func (m *ClassificationResult) Reset()         { *m = ClassificationResult{} }
func (m *ClassificationResult) String() string { return proto.CompactTextString(m) }
func (*ClassificationResult) ProtoMessage()    {}
func (*ClassificationResult) Descriptor() ([]byte, []int) {
	return fileDescriptorClassification, []int{2}
}

func (m *ClassificationResult) GetClassifications() []*Classifications {
	if m != nil {
		return m.Classifications
	}
	return nil
}

type ClassificationRequest struct {
	// Model Specification.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec,omitempty"`
	// Input data.
	Input *Input `protobuf:"bytes,2,opt,name=input" json:"input,omitempty"`
}

func (m *ClassificationRequest) Reset()         { *m = ClassificationRequest{} }
func (m *ClassificationRequest) String() string { return proto.CompactTextString(m) }
func (*ClassificationRequest) ProtoMessage()    {}
func (*ClassificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorClassification, []int{3}
}

func (m *ClassificationRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *ClassificationRequest) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

type ClassificationResponse struct {
	// Result of the classification.
	Result *ClassificationResult `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *ClassificationResponse) Reset()         { *m = ClassificationResponse{} }
func (m *ClassificationResponse) String() string { return proto.CompactTextString(m) }
func (*ClassificationResponse) ProtoMessage()    {}
func (*ClassificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorClassification, []int{4}
}

func (m *ClassificationResponse) GetResult() *ClassificationResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Class)(nil), "tensorflow.serving.Class")
	proto.RegisterType((*Classifications)(nil), "tensorflow.serving.Classifications")
	proto.RegisterType((*ClassificationResult)(nil), "tensorflow.serving.ClassificationResult")
	proto.RegisterType((*ClassificationRequest)(nil), "tensorflow.serving.ClassificationRequest")
	proto.RegisterType((*ClassificationResponse)(nil), "tensorflow.serving.ClassificationResponse")
}
func (m *Class) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Class) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Label) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClassification(dAtA, i, uint64(len(m.Label)))
		i += copy(dAtA[i:], m.Label)
	}
	if m.Score != 0 {
		dAtA[i] = 0x15
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Score))))
		i += 4
	}
	return i, nil
}

func (m *Classifications) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Classifications) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Classes) > 0 {
		for _, msg := range m.Classes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintClassification(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClassificationResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassificationResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Classifications) > 0 {
		for _, msg := range m.Classifications {
			dAtA[i] = 0xa
			i++
			i = encodeVarintClassification(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClassificationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassificationRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ModelSpec != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClassification(dAtA, i, uint64(m.ModelSpec.Size()))
		n1, err := m.ModelSpec.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Input != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClassification(dAtA, i, uint64(m.Input.Size()))
		n2, err := m.Input.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ClassificationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassificationResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClassification(dAtA, i, uint64(m.Result.Size()))
		n3, err := m.Result.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintClassification(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Class) Size() (n int) {
	var l int
	_ = l
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovClassification(uint64(l))
	}
	if m.Score != 0 {
		n += 5
	}
	return n
}

func (m *Classifications) Size() (n int) {
	var l int
	_ = l
	if len(m.Classes) > 0 {
		for _, e := range m.Classes {
			l = e.Size()
			n += 1 + l + sovClassification(uint64(l))
		}
	}
	return n
}

func (m *ClassificationResult) Size() (n int) {
	var l int
	_ = l
	if len(m.Classifications) > 0 {
		for _, e := range m.Classifications {
			l = e.Size()
			n += 1 + l + sovClassification(uint64(l))
		}
	}
	return n
}

func (m *ClassificationRequest) Size() (n int) {
	var l int
	_ = l
	if m.ModelSpec != nil {
		l = m.ModelSpec.Size()
		n += 1 + l + sovClassification(uint64(l))
	}
	if m.Input != nil {
		l = m.Input.Size()
		n += 1 + l + sovClassification(uint64(l))
	}
	return n
}

func (m *ClassificationResponse) Size() (n int) {
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovClassification(uint64(l))
	}
	return n
}

func sovClassification(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClassification(x uint64) (n int) {
	return sovClassification(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Class) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassification
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
			return fmt.Errorf("proto: Class: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Class: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassification
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
				return ErrInvalidLengthClassification
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Score = float32(math.Float32frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipClassification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClassification
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
func (m *Classifications) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassification
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
			return fmt.Errorf("proto: Classifications: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Classifications: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Classes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassification
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
				return ErrInvalidLengthClassification
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Classes = append(m.Classes, &Class{})
			if err := m.Classes[len(m.Classes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClassification
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
func (m *ClassificationResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassification
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
			return fmt.Errorf("proto: ClassificationResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassificationResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Classifications", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassification
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
				return ErrInvalidLengthClassification
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Classifications = append(m.Classifications, &Classifications{})
			if err := m.Classifications[len(m.Classifications)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClassification
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
func (m *ClassificationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassification
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
			return fmt.Errorf("proto: ClassificationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassificationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassification
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
				return ErrInvalidLengthClassification
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
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassification
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
				return ErrInvalidLengthClassification
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
			skippy, err := skipClassification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClassification
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
func (m *ClassificationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassification
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
			return fmt.Errorf("proto: ClassificationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassificationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassification
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
				return ErrInvalidLengthClassification
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &ClassificationResult{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClassification
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
func skipClassification(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClassification
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
					return 0, ErrIntOverflowClassification
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
					return 0, ErrIntOverflowClassification
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
				return 0, ErrInvalidLengthClassification
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClassification
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
				next, err := skipClassification(dAtA[start:])
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
	ErrInvalidLengthClassification = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClassification   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow_serving/apis/classification.proto", fileDescriptorClassification)
}

var fileDescriptorClassification = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x4a, 0x2b, 0x9d, 0x1e, 0x0a, 0x4b, 0x95, 0x2a, 0x58, 0x4a, 0x7a, 0xc9, 0x41,
	0x52, 0x68, 0xaf, 0x1e, 0xc4, 0x82, 0xe0, 0xa1, 0x97, 0xf5, 0xe6, 0xa5, 0xa4, 0x71, 0x2a, 0x0b,
	0xdb, 0xec, 0x9a, 0xd9, 0xea, 0x1b, 0xf8, 0x6c, 0x1e, 0x7d, 0x04, 0xc9, 0x53, 0x78, 0x94, 0xec,
	0x26, 0x94, 0xb4, 0x06, 0x6f, 0x99, 0xf0, 0xcd, 0x3f, 0xff, 0xff, 0x2f, 0x5c, 0x5b, 0x4c, 0x49,
	0x67, 0x1b, 0xa5, 0xdf, 0x57, 0x84, 0xd9, 0x9b, 0x4c, 0x5f, 0xa6, 0xb1, 0x91, 0x34, 0x4d, 0x54,
	0x4c, 0x24, 0x37, 0x32, 0x89, 0xad, 0xd4, 0x69, 0x64, 0x32, 0x6d, 0x35, 0xe7, 0x7b, 0x3a, 0x2a,
	0xe9, 0xcb, 0x49, 0x93, 0x82, 0x4c, 0xcd, 0xce, 0xfa, 0xc5, 0x66, 0x68, 0xab, 0x9f, 0x51, 0x79,
	0x28, 0x98, 0x43, 0x7b, 0x51, 0x5c, 0xe5, 0x03, 0x68, 0xab, 0x78, 0x8d, 0x6a, 0xc8, 0xc6, 0x2c,
	0xec, 0x0a, 0x3f, 0x14, 0x7f, 0x29, 0xd1, 0x19, 0x0e, 0x5b, 0x63, 0x16, 0xb6, 0x84, 0x1f, 0x82,
	0x7b, 0xe8, 0x2f, 0x6a, 0x56, 0x89, 0xcf, 0xe1, 0xd4, 0xb9, 0x47, 0x1a, 0xb2, 0xf1, 0x49, 0xd8,
	0x9b, 0x5d, 0x44, 0xc7, 0xbe, 0x23, 0xb7, 0x25, 0x2a, 0x32, 0x40, 0x18, 0xd4, 0x75, 0x04, 0xd2,
	0x4e, 0x59, 0xbe, 0x84, 0x7e, 0xbd, 0x8a, 0x4a, 0x74, 0xd2, 0x28, 0xba, 0x47, 0xc5, 0xe1, 0x6e,
	0xf0, 0xc1, 0xe0, 0xec, 0xf0, 0xce, 0xeb, 0x0e, 0xc9, 0xf2, 0x1b, 0x00, 0x57, 0xc6, 0x8a, 0x0c,
	0x26, 0x2e, 0x79, 0x6f, 0x76, 0xf5, 0xd7, 0x8d, 0x65, 0x41, 0x3d, 0x1a, 0x4c, 0x44, 0x77, 0x5b,
	0x7d, 0xf2, 0x29, 0xb4, 0x5d, 0xdf, 0xae, 0x9c, 0x86, 0xc4, 0x0f, 0x05, 0x20, 0x3c, 0x17, 0x3c,
	0xc1, 0xf9, 0x51, 0x5e, 0xa3, 0x53, 0x42, 0x7e, 0x0b, 0x9d, 0xcc, 0x65, 0x2f, 0x4d, 0x84, 0xff,
	0x07, 0xf5, 0x5d, 0x89, 0x72, 0xef, 0xae, 0xff, 0x99, 0x8f, 0xd8, 0x57, 0x3e, 0x62, 0xdf, 0xf9,
	0x88, 0xfd, 0x30, 0xb6, 0xee, 0xb8, 0x07, 0x9e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x9f,
	0xab, 0x29, 0x6e, 0x02, 0x00, 0x00,
}