// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/variable.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a Variable.
type VariableDef struct {
	// Name of the variable tensor.
	VariableName string `protobuf:"bytes,1,opt,name=variable_name,json=variableName,proto3" json:"variable_name,omitempty"`
	// Name of the tensor holding the variable's initial value.
	InitialValueName string `protobuf:"bytes,6,opt,name=initial_value_name,json=initialValueName,proto3" json:"initial_value_name,omitempty"`
	// Name of the initializer op.
	InitializerName string `protobuf:"bytes,2,opt,name=initializer_name,json=initializerName,proto3" json:"initializer_name,omitempty"`
	// Name of the snapshot tensor.
	SnapshotName string `protobuf:"bytes,3,opt,name=snapshot_name,json=snapshotName,proto3" json:"snapshot_name,omitempty"`
	// Support for saving variables as slices of a larger variable.
	SaveSliceInfoDef *SaveSliceInfoDef `protobuf:"bytes,4,opt,name=save_slice_info_def,json=saveSliceInfoDef" json:"save_slice_info_def,omitempty"`
	// Whether to represent this as a ResourceVariable.
	IsResource bool `protobuf:"varint,5,opt,name=is_resource,json=isResource,proto3" json:"is_resource,omitempty"`
}

func (m *VariableDef) Reset()                    { *m = VariableDef{} }
func (m *VariableDef) String() string            { return proto.CompactTextString(m) }
func (*VariableDef) ProtoMessage()               {}
func (*VariableDef) Descriptor() ([]byte, []int) { return fileDescriptorVariable, []int{0} }

func (m *VariableDef) GetVariableName() string {
	if m != nil {
		return m.VariableName
	}
	return ""
}

func (m *VariableDef) GetInitialValueName() string {
	if m != nil {
		return m.InitialValueName
	}
	return ""
}

func (m *VariableDef) GetInitializerName() string {
	if m != nil {
		return m.InitializerName
	}
	return ""
}

func (m *VariableDef) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *VariableDef) GetSaveSliceInfoDef() *SaveSliceInfoDef {
	if m != nil {
		return m.SaveSliceInfoDef
	}
	return nil
}

func (m *VariableDef) GetIsResource() bool {
	if m != nil {
		return m.IsResource
	}
	return false
}

type SaveSliceInfoDef struct {
	// Name of the full variable of which this is a slice.
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// Shape of the full variable.
	FullShape []int64 `protobuf:"varint,2,rep,packed,name=full_shape,json=fullShape" json:"full_shape,omitempty"`
	// Offset of this variable into the full variable.
	VarOffset []int64 `protobuf:"varint,3,rep,packed,name=var_offset,json=varOffset" json:"var_offset,omitempty"`
	// Shape of this variable.
	VarShape []int64 `protobuf:"varint,4,rep,packed,name=var_shape,json=varShape" json:"var_shape,omitempty"`
}

func (m *SaveSliceInfoDef) Reset()                    { *m = SaveSliceInfoDef{} }
func (m *SaveSliceInfoDef) String() string            { return proto.CompactTextString(m) }
func (*SaveSliceInfoDef) ProtoMessage()               {}
func (*SaveSliceInfoDef) Descriptor() ([]byte, []int) { return fileDescriptorVariable, []int{1} }

func (m *SaveSliceInfoDef) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *SaveSliceInfoDef) GetFullShape() []int64 {
	if m != nil {
		return m.FullShape
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarOffset() []int64 {
	if m != nil {
		return m.VarOffset
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarShape() []int64 {
	if m != nil {
		return m.VarShape
	}
	return nil
}

func init() {
	proto.RegisterType((*VariableDef)(nil), "tensorflow.VariableDef")
	proto.RegisterType((*SaveSliceInfoDef)(nil), "tensorflow.SaveSliceInfoDef")
}
func (m *VariableDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VariableDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VariableName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintVariable(dAtA, i, uint64(len(m.VariableName)))
		i += copy(dAtA[i:], m.VariableName)
	}
	if len(m.InitializerName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintVariable(dAtA, i, uint64(len(m.InitializerName)))
		i += copy(dAtA[i:], m.InitializerName)
	}
	if len(m.SnapshotName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintVariable(dAtA, i, uint64(len(m.SnapshotName)))
		i += copy(dAtA[i:], m.SnapshotName)
	}
	if m.SaveSliceInfoDef != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintVariable(dAtA, i, uint64(m.SaveSliceInfoDef.Size()))
		n1, err := m.SaveSliceInfoDef.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.IsResource {
		dAtA[i] = 0x28
		i++
		if m.IsResource {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.InitialValueName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintVariable(dAtA, i, uint64(len(m.InitialValueName)))
		i += copy(dAtA[i:], m.InitialValueName)
	}
	return i, nil
}

func (m *SaveSliceInfoDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SaveSliceInfoDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FullName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintVariable(dAtA, i, uint64(len(m.FullName)))
		i += copy(dAtA[i:], m.FullName)
	}
	if len(m.FullShape) > 0 {
		dAtA3 := make([]byte, len(m.FullShape)*10)
		var j2 int
		for _, num1 := range m.FullShape {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintVariable(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.VarOffset) > 0 {
		dAtA5 := make([]byte, len(m.VarOffset)*10)
		var j4 int
		for _, num1 := range m.VarOffset {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintVariable(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	if len(m.VarShape) > 0 {
		dAtA7 := make([]byte, len(m.VarShape)*10)
		var j6 int
		for _, num1 := range m.VarShape {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA7[j6] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j6++
			}
			dAtA7[j6] = uint8(num)
			j6++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintVariable(dAtA, i, uint64(j6))
		i += copy(dAtA[i:], dAtA7[:j6])
	}
	return i, nil
}

func encodeVarintVariable(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VariableDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.VariableName)
	if l > 0 {
		n += 1 + l + sovVariable(uint64(l))
	}
	l = len(m.InitializerName)
	if l > 0 {
		n += 1 + l + sovVariable(uint64(l))
	}
	l = len(m.SnapshotName)
	if l > 0 {
		n += 1 + l + sovVariable(uint64(l))
	}
	if m.SaveSliceInfoDef != nil {
		l = m.SaveSliceInfoDef.Size()
		n += 1 + l + sovVariable(uint64(l))
	}
	if m.IsResource {
		n += 2
	}
	l = len(m.InitialValueName)
	if l > 0 {
		n += 1 + l + sovVariable(uint64(l))
	}
	return n
}

func (m *SaveSliceInfoDef) Size() (n int) {
	var l int
	_ = l
	l = len(m.FullName)
	if l > 0 {
		n += 1 + l + sovVariable(uint64(l))
	}
	if len(m.FullShape) > 0 {
		l = 0
		for _, e := range m.FullShape {
			l += sovVariable(uint64(e))
		}
		n += 1 + sovVariable(uint64(l)) + l
	}
	if len(m.VarOffset) > 0 {
		l = 0
		for _, e := range m.VarOffset {
			l += sovVariable(uint64(e))
		}
		n += 1 + sovVariable(uint64(l)) + l
	}
	if len(m.VarShape) > 0 {
		l = 0
		for _, e := range m.VarShape {
			l += sovVariable(uint64(e))
		}
		n += 1 + sovVariable(uint64(l)) + l
	}
	return n
}

func sovVariable(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVariable(x uint64) (n int) {
	return sovVariable(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VariableDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVariable
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
			return fmt.Errorf("proto: VariableDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VariableDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VariableName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVariable
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
				return ErrInvalidLengthVariable
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VariableName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitializerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVariable
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
				return ErrInvalidLengthVariable
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitializerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVariable
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
				return ErrInvalidLengthVariable
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaveSliceInfoDef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVariable
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
				return ErrInvalidLengthVariable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SaveSliceInfoDef == nil {
				m.SaveSliceInfoDef = &SaveSliceInfoDef{}
			}
			if err := m.SaveSliceInfoDef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsResource", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVariable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsResource = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialValueName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVariable
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
				return ErrInvalidLengthVariable
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialValueName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVariable(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVariable
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
func (m *SaveSliceInfoDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVariable
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
			return fmt.Errorf("proto: SaveSliceInfoDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SaveSliceInfoDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVariable
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
				return ErrInvalidLengthVariable
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FullName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVariable
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FullShape = append(m.FullShape, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVariable
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthVariable
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVariable
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FullShape = append(m.FullShape, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FullShape", wireType)
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVariable
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.VarOffset = append(m.VarOffset, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVariable
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthVariable
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVariable
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.VarOffset = append(m.VarOffset, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field VarOffset", wireType)
			}
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVariable
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.VarShape = append(m.VarShape, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVariable
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthVariable
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVariable
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.VarShape = append(m.VarShape, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field VarShape", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVariable(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVariable
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
func skipVariable(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVariable
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
					return 0, ErrIntOverflowVariable
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
					return 0, ErrIntOverflowVariable
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
				return 0, ErrInvalidLengthVariable
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVariable
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
				next, err := skipVariable(dAtA[start:])
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
	ErrInvalidLengthVariable = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVariable   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tensorflow/core/framework/variable.proto", fileDescriptorVariable) }

var fileDescriptorVariable = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0x87, 0x33, 0x94, 0x4b, 0xe0, 0x70, 0xb9, 0x97, 0xd4, 0x4d, 0x13, 0x15, 0x09, 0x6c, 0x6a,
	0x62, 0x20, 0xd1, 0xa5, 0x3b, 0xc2, 0xc6, 0x98, 0x28, 0x29, 0x09, 0xdb, 0xe6, 0x80, 0x67, 0x64,
	0x62, 0xe9, 0x90, 0x99, 0x32, 0x24, 0x3e, 0x82, 0x0f, 0xe1, 0xf3, 0xb8, 0xf4, 0x11, 0x0c, 0x4f,
	0xe1, 0xd2, 0xcc, 0xb4, 0x15, 0xc2, 0xf6, 0xfb, 0x7d, 0x73, 0xe6, 0xfc, 0x81, 0x30, 0xa3, 0x54,
	0x4b, 0xc5, 0x13, 0xb9, 0x1d, 0x2e, 0xa4, 0xa2, 0x21, 0x57, 0xb8, 0xa2, 0xad, 0x54, 0x2f, 0x43,
	0x83, 0x4a, 0xe0, 0x3c, 0xa1, 0xc1, 0x5a, 0xc9, 0x4c, 0xfa, 0xb0, 0x37, 0x7b, 0xef, 0x15, 0x68,
	0xce, 0x8a, 0x78, 0x4c, 0xdc, 0xef, 0x43, 0xab, 0xb4, 0xe3, 0x14, 0x57, 0x14, 0xb0, 0x2e, 0x0b,
	0x1b, 0xd1, 0xdf, 0x12, 0x3e, 0xe0, 0x8a, 0xfc, 0x4b, 0x68, 0x8b, 0x54, 0x64, 0x02, 0x13, 0xf1,
	0x4a, 0x2a, 0xf7, 0x2a, 0xce, 0xfb, 0x7f, 0xc0, 0x9d, 0xda, 0x87, 0x96, 0x4e, 0x71, 0xad, 0x97,
	0x32, 0xcb, 0x3d, 0x2f, 0xaf, 0x57, 0x42, 0x27, 0xdd, 0xc3, 0x89, 0x46, 0x43, 0xb1, 0x4e, 0xc4,
	0x82, 0x62, 0x91, 0x72, 0x19, 0x3f, 0x11, 0x0f, 0xaa, 0x5d, 0x16, 0x36, 0xaf, 0xcf, 0x06, 0xfb,
	0x76, 0x07, 0x53, 0x34, 0x34, 0xb5, 0xd6, 0x5d, 0xca, 0xe5, 0x98, 0x78, 0xd4, 0xd6, 0x47, 0xc4,
	0xbf, 0x80, 0xa6, 0xd0, 0xb1, 0x22, 0x2d, 0x37, 0x6a, 0x41, 0xc1, 0x9f, 0x2e, 0x0b, 0xeb, 0x11,
	0x08, 0x1d, 0x15, 0xc4, 0xbf, 0x02, 0xbf, 0xe8, 0x32, 0x36, 0x98, 0x6c, 0x8a, 0x39, 0x6b, 0xae,
	0xaf, 0x72, 0xae, 0x99, 0x0d, 0x6c, 0x6f, 0xbd, 0x37, 0x06, 0xed, 0xe3, 0x5f, 0xfd, 0x53, 0x68,
	0xf0, 0x4d, 0x92, 0x1c, 0x6e, 0xa8, 0x6e, 0x81, 0x9b, 0xe6, 0x1c, 0xc0, 0x85, 0x7a, 0x89, 0x6b,
	0xbb, 0x17, 0x2f, 0xf4, 0x22, 0xa7, 0x4f, 0x2d, 0xb0, 0xb1, 0x41, 0x15, 0x4b, 0xce, 0x35, 0x65,
	0x81, 0x97, 0xc7, 0x06, 0xd5, 0xa3, 0x03, 0xb6, 0xb4, 0x8d, 0xf3, 0xc7, 0x55, 0x97, 0xd6, 0x0d,
	0x2a, 0xf7, 0x76, 0x74, 0xfb, 0xb1, 0xeb, 0xb0, 0xcf, 0x5d, 0x87, 0x7d, 0xed, 0x3a, 0x0c, 0x02,
	0xa9, 0x9e, 0x0f, 0x97, 0xf3, 0x7b, 0xf0, 0xd1, 0xbf, 0xf2, 0xa4, 0x13, 0x7b, 0x70, 0x3d, 0x61,
	0xdf, 0x8c, 0xcd, 0x6b, 0xee, 0xfa, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xa8, 0x8e,
	0x87, 0x29, 0x02, 0x00, 0x00,
}