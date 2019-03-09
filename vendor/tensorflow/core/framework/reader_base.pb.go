// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/reader_base.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// For serializing and restoring the state of ReaderBase, see
// reader_base.h for details.
type ReaderBaseState struct {
	WorkStarted        int64  `protobuf:"varint,1,opt,name=work_started,json=workStarted,proto3" json:"work_started,omitempty"`
	WorkFinished       int64  `protobuf:"varint,2,opt,name=work_finished,json=workFinished,proto3" json:"work_finished,omitempty"`
	NumRecordsProduced int64  `protobuf:"varint,3,opt,name=num_records_produced,json=numRecordsProduced,proto3" json:"num_records_produced,omitempty"`
	CurrentWork        []byte `protobuf:"bytes,4,opt,name=current_work,json=currentWork,proto3" json:"current_work,omitempty"`
}

func (m *ReaderBaseState) Reset()                    { *m = ReaderBaseState{} }
func (m *ReaderBaseState) String() string            { return proto.CompactTextString(m) }
func (*ReaderBaseState) ProtoMessage()               {}
func (*ReaderBaseState) Descriptor() ([]byte, []int) { return fileDescriptorReaderBase, []int{0} }

func (m *ReaderBaseState) GetWorkStarted() int64 {
	if m != nil {
		return m.WorkStarted
	}
	return 0
}

func (m *ReaderBaseState) GetWorkFinished() int64 {
	if m != nil {
		return m.WorkFinished
	}
	return 0
}

func (m *ReaderBaseState) GetNumRecordsProduced() int64 {
	if m != nil {
		return m.NumRecordsProduced
	}
	return 0
}

func (m *ReaderBaseState) GetCurrentWork() []byte {
	if m != nil {
		return m.CurrentWork
	}
	return nil
}

func init() {
	proto.RegisterType((*ReaderBaseState)(nil), "tensorflow.ReaderBaseState")
}
func (m *ReaderBaseState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReaderBaseState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WorkStarted != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.WorkStarted))
	}
	if m.WorkFinished != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.WorkFinished))
	}
	if m.NumRecordsProduced != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(m.NumRecordsProduced))
	}
	if len(m.CurrentWork) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintReaderBase(dAtA, i, uint64(len(m.CurrentWork)))
		i += copy(dAtA[i:], m.CurrentWork)
	}
	return i, nil
}

func encodeVarintReaderBase(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReaderBaseState) Size() (n int) {
	var l int
	_ = l
	if m.WorkStarted != 0 {
		n += 1 + sovReaderBase(uint64(m.WorkStarted))
	}
	if m.WorkFinished != 0 {
		n += 1 + sovReaderBase(uint64(m.WorkFinished))
	}
	if m.NumRecordsProduced != 0 {
		n += 1 + sovReaderBase(uint64(m.NumRecordsProduced))
	}
	l = len(m.CurrentWork)
	if l > 0 {
		n += 1 + l + sovReaderBase(uint64(l))
	}
	return n
}

func sovReaderBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReaderBase(x uint64) (n int) {
	return sovReaderBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReaderBaseState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReaderBase
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
			return fmt.Errorf("proto: ReaderBaseState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReaderBaseState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkStarted", wireType)
			}
			m.WorkStarted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkStarted |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkFinished", wireType)
			}
			m.WorkFinished = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkFinished |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRecordsProduced", wireType)
			}
			m.NumRecordsProduced = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRecordsProduced |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentWork", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReaderBase
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
				return ErrInvalidLengthReaderBase
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentWork = append(m.CurrentWork[:0], dAtA[iNdEx:postIndex]...)
			if m.CurrentWork == nil {
				m.CurrentWork = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReaderBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReaderBase
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
func skipReaderBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReaderBase
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
					return 0, ErrIntOverflowReaderBase
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
					return 0, ErrIntOverflowReaderBase
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
				return 0, ErrInvalidLengthReaderBase
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReaderBase
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
				next, err := skipReaderBase(dAtA[start:])
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
	ErrInvalidLengthReaderBase = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReaderBase   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("tensorflow/core/framework/reader_base.proto", fileDescriptorReaderBase)
}

var fileDescriptorReaderBase = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x75, 0x14, 0x31, 0xb8, 0x45, 0x20, 0x8b, 0xc1, 0x53, 0x54, 0x60, 0xa9, 0x84, 0x94,
	0x20, 0x31, 0xb3, 0x64, 0x60, 0x8e, 0xd2, 0x81, 0xd1, 0x72, 0xe3, 0x0b, 0x54, 0x25, 0xbe, 0xea,
	0xec, 0xa8, 0x7f, 0xc5, 0x77, 0x30, 0xf2, 0x09, 0x28, 0x5f, 0xc1, 0x88, 0xec, 0x44, 0x64, 0x7d,
	0xf7, 0xce, 0xd6, 0x3b, 0xf1, 0x10, 0xd0, 0x79, 0xe2, 0xf6, 0x83, 0x4e, 0x45, 0x43, 0x8c, 0x45,
	0xcb, 0xa6, 0xc3, 0x13, 0xf1, 0xa1, 0x60, 0x34, 0x16, 0x59, 0xef, 0x8c, 0xc7, 0xfc, 0xc8, 0x14,
	0x48, 0x8a, 0x59, 0xbe, 0xfb, 0x04, 0x71, 0x55, 0x27, 0xa3, 0x34, 0x1e, 0xb7, 0xc1, 0x04, 0x94,
	0xb7, 0x62, 0x15, 0x37, 0xb5, 0x0f, 0x86, 0x03, 0x5a, 0x05, 0x6b, 0xd8, 0x2c, 0xea, 0x65, 0x64,
	0xdb, 0x11, 0xc9, 0x7b, 0x71, 0x99, 0x94, 0x76, 0xef, 0xf6, 0xfe, 0x1d, 0xad, 0x3a, 0x4b, 0x4e,
	0xda, 0x7b, 0x99, 0x98, 0x7c, 0x14, 0x37, 0xae, 0xef, 0x34, 0x63, 0x43, 0x6c, 0xbd, 0x3e, 0x32,
	0xd9, 0xbe, 0x41, 0xab, 0x16, 0xc9, 0x95, 0xae, 0xef, 0xea, 0x71, 0x54, 0x4d, 0x93, 0xf8, 0x73,
	0xd3, 0x33, 0xa3, 0x0b, 0x3a, 0xbe, 0xa4, 0xce, 0xd7, 0xb0, 0x59, 0xd5, 0xcb, 0x89, 0xbd, 0x12,
	0x1f, 0xca, 0xe7, 0xaf, 0x21, 0x83, 0xef, 0x21, 0x83, 0x9f, 0x21, 0x03, 0xa1, 0x88, 0xdf, 0xf2,
	0x39, 0x27, 0xff, 0xcf, 0x2e, 0xaf, 0xe7, 0xaa, 0x2a, 0x56, 0xfb, 0x0a, 0x7e, 0x01, 0x76, 0x17,
	0xe9, 0x04, 0x4f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe1, 0x0a, 0xa8, 0x31, 0x01, 0x00,
	0x00,
}
