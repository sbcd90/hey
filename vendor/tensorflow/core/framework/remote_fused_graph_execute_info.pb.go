// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/remote_fused_graph_execute_info.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RemoteFusedGraphExecuteInfo_NodeType int32

const (
	RemoteFusedGraphExecuteInfo_UNUSED        RemoteFusedGraphExecuteInfo_NodeType = 0
	RemoteFusedGraphExecuteInfo_GRAPH_INPUT   RemoteFusedGraphExecuteInfo_NodeType = 1
	RemoteFusedGraphExecuteInfo_GRAPH_OUTPUT  RemoteFusedGraphExecuteInfo_NodeType = 2
	RemoteFusedGraphExecuteInfo_FUSED_NODE    RemoteFusedGraphExecuteInfo_NodeType = 3
	RemoteFusedGraphExecuteInfo_BORDER_INPUT  RemoteFusedGraphExecuteInfo_NodeType = 4
	RemoteFusedGraphExecuteInfo_BORDER_OUTPUT RemoteFusedGraphExecuteInfo_NodeType = 5
)

var RemoteFusedGraphExecuteInfo_NodeType_name = map[int32]string{
	0: "UNUSED",
	1: "GRAPH_INPUT",
	2: "GRAPH_OUTPUT",
	3: "FUSED_NODE",
	4: "BORDER_INPUT",
	5: "BORDER_OUTPUT",
}
var RemoteFusedGraphExecuteInfo_NodeType_value = map[string]int32{
	"UNUSED":        0,
	"GRAPH_INPUT":   1,
	"GRAPH_OUTPUT":  2,
	"FUSED_NODE":    3,
	"BORDER_INPUT":  4,
	"BORDER_OUTPUT": 5,
}

func (x RemoteFusedGraphExecuteInfo_NodeType) String() string {
	return proto.EnumName(RemoteFusedGraphExecuteInfo_NodeType_name, int32(x))
}
func (RemoteFusedGraphExecuteInfo_NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorRemoteFusedGraphExecuteInfo, []int{0, 0}
}

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type RemoteFusedGraphExecuteInfo struct {
	// Definition of remote graph
	RemoteGraph *GraphDef `protobuf:"bytes,1,opt,name=remote_graph,json=remoteGraph" json:"remote_graph,omitempty"`
	// Remote fused graph input node name
	GraphInputNodeName []string `protobuf:"bytes,2,rep,name=graph_input_node_name,json=graphInputNodeName" json:"graph_input_node_name,omitempty"`
	// Remote fused graph output node name
	GraphOutputNodeName []string `protobuf:"bytes,3,rep,name=graph_output_node_name,json=graphOutputNodeName" json:"graph_output_node_name,omitempty"`
	// Executor's name
	ExecutorName string `protobuf:"bytes,4,opt,name=executor_name,json=executorName,proto3" json:"executor_name,omitempty"`
	// Optional: Parameters given to the executor
	SerializedExecutorParameters []byte `protobuf:"bytes,5,opt,name=serialized_executor_parameters,json=serializedExecutorParameters,proto3" json:"serialized_executor_parameters,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	DefaultGraphInputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,6,rep,name=default_graph_input_tensor_shape,json=defaultGraphInputTensorShape" json:"default_graph_input_tensor_shape,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	// TODO(satok): Remote output tensor shape once shape information is stored
	// in NodeDef
	DefaultGraphOutputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,7,rep,name=default_graph_output_tensor_shape,json=defaultGraphOutputTensorShape" json:"default_graph_output_tensor_shape,omitempty"`
}

func (m *RemoteFusedGraphExecuteInfo) Reset()         { *m = RemoteFusedGraphExecuteInfo{} }
func (m *RemoteFusedGraphExecuteInfo) String() string { return proto.CompactTextString(m) }
func (*RemoteFusedGraphExecuteInfo) ProtoMessage()    {}
func (*RemoteFusedGraphExecuteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptorRemoteFusedGraphExecuteInfo, []int{0}
}

func (m *RemoteFusedGraphExecuteInfo) GetRemoteGraph() *GraphDef {
	if m != nil {
		return m.RemoteGraph
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphInputNodeName() []string {
	if m != nil {
		return m.GraphInputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphOutputNodeName() []string {
	if m != nil {
		return m.GraphOutputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetExecutorName() string {
	if m != nil {
		return m.ExecutorName
	}
	return ""
}

func (m *RemoteFusedGraphExecuteInfo) GetSerializedExecutorParameters() []byte {
	if m != nil {
		return m.SerializedExecutorParameters
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphInputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphInputTensorShape
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphOutputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphOutputTensorShape
	}
	return nil
}

type RemoteFusedGraphExecuteInfo_TensorShapeTypeProto struct {
	Dtype DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *TensorShapeProto `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Reset() {
	*m = RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{}
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) String() string {
	return proto.CompactTextString(m)
}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoMessage() {}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptorRemoteFusedGraphExecuteInfo, []int{0, 0}
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteFusedGraphExecuteInfo)(nil), "tensorflow.RemoteFusedGraphExecuteInfo")
	proto.RegisterType((*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto)(nil), "tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto")
	proto.RegisterEnum("tensorflow.RemoteFusedGraphExecuteInfo_NodeType", RemoteFusedGraphExecuteInfo_NodeType_name, RemoteFusedGraphExecuteInfo_NodeType_value)
}
func (m *RemoteFusedGraphExecuteInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteFusedGraphExecuteInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RemoteGraph != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(m.RemoteGraph.Size()))
		n1, err := m.RemoteGraph.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.GraphInputNodeName) > 0 {
		for _, s := range m.GraphInputNodeName {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.GraphOutputNodeName) > 0 {
		for _, s := range m.GraphOutputNodeName {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.ExecutorName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(len(m.ExecutorName)))
		i += copy(dAtA[i:], m.ExecutorName)
	}
	if len(m.SerializedExecutorParameters) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(len(m.SerializedExecutorParameters)))
		i += copy(dAtA[i:], m.SerializedExecutorParameters)
	}
	if len(m.DefaultGraphInputTensorShape) > 0 {
		for _, msg := range m.DefaultGraphInputTensorShape {
			dAtA[i] = 0x32
			i++
			i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DefaultGraphOutputTensorShape) > 0 {
		for _, msg := range m.DefaultGraphOutputTensorShape {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Dtype != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(m.Dtype))
	}
	if m.Shape != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRemoteFusedGraphExecuteInfo(dAtA, i, uint64(m.Shape.Size()))
		n2, err := m.Shape.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintRemoteFusedGraphExecuteInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RemoteFusedGraphExecuteInfo) Size() (n int) {
	var l int
	_ = l
	if m.RemoteGraph != nil {
		l = m.RemoteGraph.Size()
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	if len(m.GraphInputNodeName) > 0 {
		for _, s := range m.GraphInputNodeName {
			l = len(s)
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	if len(m.GraphOutputNodeName) > 0 {
		for _, s := range m.GraphOutputNodeName {
			l = len(s)
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	l = len(m.ExecutorName)
	if l > 0 {
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	l = len(m.SerializedExecutorParameters)
	if l > 0 {
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	if len(m.DefaultGraphInputTensorShape) > 0 {
		for _, e := range m.DefaultGraphInputTensorShape {
			l = e.Size()
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	if len(m.DefaultGraphOutputTensorShape) > 0 {
		for _, e := range m.DefaultGraphOutputTensorShape {
			l = e.Size()
			n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
		}
	}
	return n
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Size() (n int) {
	var l int
	_ = l
	if m.Dtype != 0 {
		n += 1 + sovRemoteFusedGraphExecuteInfo(uint64(m.Dtype))
	}
	if m.Shape != nil {
		l = m.Shape.Size()
		n += 1 + l + sovRemoteFusedGraphExecuteInfo(uint64(l))
	}
	return n
}

func sovRemoteFusedGraphExecuteInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRemoteFusedGraphExecuteInfo(x uint64) (n int) {
	return sovRemoteFusedGraphExecuteInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteFusedGraphExecuteInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
			return fmt.Errorf("proto: RemoteFusedGraphExecuteInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteFusedGraphExecuteInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteGraph", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RemoteGraph == nil {
				m.RemoteGraph = &GraphDef{}
			}
			if err := m.RemoteGraph.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GraphInputNodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GraphInputNodeName = append(m.GraphInputNodeName, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GraphOutputNodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GraphOutputNodeName = append(m.GraphOutputNodeName, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerializedExecutorParameters", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerializedExecutorParameters = append(m.SerializedExecutorParameters[:0], dAtA[iNdEx:postIndex]...)
			if m.SerializedExecutorParameters == nil {
				m.SerializedExecutorParameters = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultGraphInputTensorShape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultGraphInputTensorShape = append(m.DefaultGraphInputTensorShape, &RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{})
			if err := m.DefaultGraphInputTensorShape[len(m.DefaultGraphInputTensorShape)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultGraphOutputTensorShape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultGraphOutputTensorShape = append(m.DefaultGraphOutputTensorShape, &RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{})
			if err := m.DefaultGraphOutputTensorShape[len(m.DefaultGraphOutputTensorShape)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemoteFusedGraphExecuteInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
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
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
			return fmt.Errorf("proto: TensorShapeTypeProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TensorShapeTypeProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dtype", wireType)
			}
			m.Dtype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dtype |= (DataType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shape", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Shape == nil {
				m.Shape = &TensorShapeProto{}
			}
			if err := m.Shape.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemoteFusedGraphExecuteInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemoteFusedGraphExecuteInfo
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
func skipRemoteFusedGraphExecuteInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
					return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
					return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				return 0, ErrInvalidLengthRemoteFusedGraphExecuteInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRemoteFusedGraphExecuteInfo
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
				next, err := skipRemoteFusedGraphExecuteInfo(dAtA[start:])
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
	ErrInvalidLengthRemoteFusedGraphExecuteInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRemoteFusedGraphExecuteInfo   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow/core/framework/remote_fused_graph_execute_info.proto", fileDescriptorRemoteFusedGraphExecuteInfo)
}

var fileDescriptorRemoteFusedGraphExecuteInfo = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0xc7, 0x1d, 0x28, 0x68, 0x1f, 0xb4, 0xe2, 0x58, 0x0d, 0x41, 0x24, 0x6b, 0x8d, 0x09, 0x31,
	0x06, 0x22, 0x3d, 0x78, 0x31, 0x31, 0x92, 0xa5, 0xc8, 0x41, 0x20, 0x53, 0x38, 0x4f, 0xc6, 0xee,
	0xdb, 0x96, 0x08, 0x3b, 0x9b, 0xd9, 0xc1, 0x5a, 0xcf, 0xc6, 0xef, 0xe1, 0xb7, 0xf1, 0xe8, 0x47,
	0x30, 0x7c, 0x0a, 0x8f, 0x66, 0x66, 0x56, 0x58, 0x12, 0xcb, 0xc9, 0xe3, 0xbe, 0xf7, 0xfb, 0xbf,
	0xfd, 0xbf, 0xff, 0xbe, 0x85, 0x37, 0x1a, 0xa3, 0x44, 0xaa, 0x70, 0x2e, 0xaf, 0xda, 0xe7, 0x52,
	0x61, 0x3b, 0x54, 0x62, 0x81, 0x57, 0x52, 0x7d, 0x6c, 0x2b, 0x5c, 0x48, 0x8d, 0x3c, 0x5c, 0x26,
	0x18, 0xf0, 0x0b, 0x25, 0xe2, 0x4b, 0x8e, 0x9f, 0xf1, 0x7c, 0xa9, 0x91, 0xcf, 0xa2, 0x50, 0xb6,
	0x62, 0x25, 0xb5, 0xa4, 0xb0, 0x19, 0x50, 0x7b, 0x76, 0xf3, 0x30, 0xab, 0x77, 0x92, 0xda, 0x8b,
	0x9b, 0x31, 0xd7, 0xe1, 0xc9, 0xa5, 0x88, 0x31, 0xa5, 0x77, 0x0c, 0xd5, 0xd7, 0x31, 0x26, 0x0e,
	0x3b, 0xfe, 0x5e, 0x84, 0x47, 0xcc, 0x3a, 0x3e, 0x35, 0x86, 0xfb, 0xe6, 0x7d, 0x3d, 0x67, 0x77,
	0x10, 0x85, 0x92, 0xbe, 0x82, 0x72, 0xba, 0x90, 0xb5, 0x52, 0x25, 0x1e, 0x69, 0x96, 0x3a, 0x47,
	0xad, 0xcd, 0xf4, 0x96, 0xd5, 0xf8, 0x18, 0xb2, 0x92, 0x23, 0xed, 0x33, 0x7d, 0x09, 0x0f, 0xdc,
	0xf2, 0xb3, 0x28, 0x5e, 0x6a, 0x1e, 0xc9, 0x00, 0x79, 0x24, 0x16, 0x58, 0xcd, 0x79, 0xf9, 0xe6,
	0x3e, 0xa3, 0xb6, 0x39, 0x30, 0xbd, 0xa1, 0x0c, 0x70, 0x28, 0x16, 0x48, 0x4f, 0xe0, 0xa1, 0x93,
	0xc8, 0xa5, 0xde, 0xd6, 0xe4, 0xad, 0xe6, 0xbe, 0xed, 0x8e, 0x6c, 0x73, 0x2d, 0x7a, 0x0a, 0x07,
	0x2e, 0x5e, 0xa9, 0x1c, 0xbb, 0xe7, 0x91, 0xe6, 0x3e, 0x2b, 0xff, 0x2d, 0x5a, 0xc8, 0x87, 0x46,
	0x82, 0x6a, 0x26, 0xe6, 0xb3, 0x2f, 0x18, 0xf0, 0x35, 0x1f, 0x0b, 0x93, 0x89, 0x46, 0x95, 0x54,
	0x0b, 0x1e, 0x69, 0x96, 0x59, 0x7d, 0x43, 0xf5, 0x52, 0x68, 0xbc, 0x66, 0xe8, 0x57, 0x02, 0x5e,
	0x80, 0xa1, 0x58, 0xce, 0x35, 0xcf, 0xee, 0x96, 0x4d, 0xbf, 0x5a, 0xf4, 0xf2, 0xcd, 0x52, 0xe7,
	0x75, 0x36, 0xa0, 0x1d, 0xf9, 0xb6, 0x26, 0x16, 0x3b, 0x33, 0xd2, 0xc9, 0x75, 0x8c, 0x63, 0xf3,
	0x51, 0x58, 0x3d, 0x7d, 0x4b, 0x7f, 0x9d, 0x51, 0x06, 0xa3, 0xdf, 0x08, 0x3c, 0xd9, 0xb6, 0x91,
	0xe6, 0xb5, 0xe5, 0xe3, 0xf6, 0x7f, 0xf0, 0xf1, 0x38, 0xeb, 0xc3, 0xe5, 0x9e, 0xe1, 0x6a, 0x9f,
	0xe0, 0xe8, 0x5f, 0x32, 0xfa, 0x1c, 0x0a, 0x81, 0xb9, 0x31, 0x7b, 0x2c, 0x87, 0xdb, 0xc7, 0xe2,
	0x0b, 0x2d, 0x0c, 0xc9, 0x1c, 0x42, 0x3b, 0x50, 0x70, 0x7e, 0x73, 0xf6, 0xb0, 0xea, 0x59, 0x36,
	0x33, 0xdc, 0xf9, 0x71, 0xe8, 0x71, 0x04, 0x77, 0xcc, 0xe7, 0x37, 0x63, 0x28, 0x40, 0x71, 0x3a,
	0x9c, 0x9e, 0xf5, 0xfc, 0xca, 0x2d, 0x7a, 0x17, 0x4a, 0x7d, 0xf6, 0x76, 0xfc, 0x8e, 0x0f, 0x86,
	0xe3, 0xe9, 0xa4, 0x42, 0x68, 0x05, 0xca, 0xae, 0x30, 0x9a, 0x4e, 0x4c, 0x25, 0x47, 0x0f, 0x01,
	0x4e, 0x0d, 0xcd, 0x87, 0x23, 0xbf, 0x57, 0xc9, 0x1b, 0xa2, 0x3b, 0x62, 0x7e, 0x8f, 0xa5, 0x9a,
	0x3d, 0x7a, 0x0f, 0x0e, 0xd2, 0x4a, 0x2a, 0x2a, 0x74, 0xdf, 0xff, 0x58, 0x35, 0xc8, 0xcf, 0x55,
	0x83, 0xfc, 0x5a, 0x35, 0x08, 0x54, 0xa5, 0xba, 0xc8, 0xba, 0x5c, 0xff, 0x57, 0x5d, 0x6f, 0x47,
	0xc0, 0x76, 0x81, 0x31, 0xf9, 0x4d, 0xc8, 0x87, 0xa2, 0xfd, 0xf3, 0x4e, 0xfe, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xbd, 0x04, 0x0c, 0x16, 0x44, 0x04, 0x00, 0x00,
}