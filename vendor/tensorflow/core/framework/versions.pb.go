// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensorflow/core/framework/versions.proto

package tensorflow

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Version information for a piece of serialized data
//
// There are different types of versions for each type of data
// (GraphDef, etc.), but they all have the same common shape
// described here.
//
// Each consumer has "consumer" and "min_producer" versions (specified
// elsewhere).  A consumer is allowed to consume this data if
//
//   producer >= min_producer
//   consumer >= min_consumer
//   consumer not in bad_consumers
//
type VersionDef struct {
	// The version of the code that produced this data.
	Producer int32 `protobuf:"varint,1,opt,name=producer,proto3" json:"producer,omitempty"`
	// Any consumer below this version is not allowed to consume this data.
	MinConsumer int32 `protobuf:"varint,2,opt,name=min_consumer,json=minConsumer,proto3" json:"min_consumer,omitempty"`
	// Specific consumer versions which are disallowed (e.g. due to bugs).
	BadConsumers []int32 `protobuf:"varint,3,rep,packed,name=bad_consumers,json=badConsumers" json:"bad_consumers,omitempty"`
}

func (m *VersionDef) Reset()                    { *m = VersionDef{} }
func (m *VersionDef) String() string            { return proto.CompactTextString(m) }
func (*VersionDef) ProtoMessage()               {}
func (*VersionDef) Descriptor() ([]byte, []int) { return fileDescriptorVersions, []int{0} }

func (m *VersionDef) GetProducer() int32 {
	if m != nil {
		return m.Producer
	}
	return 0
}

func (m *VersionDef) GetMinConsumer() int32 {
	if m != nil {
		return m.MinConsumer
	}
	return 0
}

func (m *VersionDef) GetBadConsumers() []int32 {
	if m != nil {
		return m.BadConsumers
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionDef)(nil), "tensorflow.VersionDef")
}
func (m *VersionDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Producer != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintVersions(dAtA, i, uint64(m.Producer))
	}
	if m.MinConsumer != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintVersions(dAtA, i, uint64(m.MinConsumer))
	}
	if len(m.BadConsumers) > 0 {
		dAtA2 := make([]byte, len(m.BadConsumers)*10)
		var j1 int
		for _, num1 := range m.BadConsumers {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintVersions(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func encodeVarintVersions(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VersionDef) Size() (n int) {
	var l int
	_ = l
	if m.Producer != 0 {
		n += 1 + sovVersions(uint64(m.Producer))
	}
	if m.MinConsumer != 0 {
		n += 1 + sovVersions(uint64(m.MinConsumer))
	}
	if len(m.BadConsumers) > 0 {
		l = 0
		for _, e := range m.BadConsumers {
			l += sovVersions(uint64(e))
		}
		n += 1 + sovVersions(uint64(l)) + l
	}
	return n
}

func sovVersions(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVersions(x uint64) (n int) {
	return sovVersions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VersionDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVersions
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
			return fmt.Errorf("proto: VersionDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Producer", wireType)
			}
			m.Producer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Producer |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinConsumer", wireType)
			}
			m.MinConsumer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVersions
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinConsumer |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVersions
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BadConsumers = append(m.BadConsumers, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVersions
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
					return ErrInvalidLengthVersions
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVersions
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BadConsumers = append(m.BadConsumers, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BadConsumers", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVersions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVersions
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
func skipVersions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVersions
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
					return 0, ErrIntOverflowVersions
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
					return 0, ErrIntOverflowVersions
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
				return 0, ErrInvalidLengthVersions
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVersions
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
				next, err := skipVersions(dAtA[start:])
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
	ErrInvalidLengthVersions = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVersions   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tensorflow/core/framework/versions.proto", fileDescriptorVersions) }

var fileDescriptorVersions = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0x2b, 0x4a, 0xcc,
	0x4d, 0x2d, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0xa8, 0x54, 0x2a, 0xe0, 0xe2, 0x0a, 0x83, 0xc8, 0xba,
	0xa4, 0xa6, 0x09, 0x49, 0x71, 0x71, 0x14, 0x14, 0xe5, 0xa7, 0x94, 0x26, 0xa7, 0x16, 0x49, 0x30,
	0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xc1, 0xf9, 0x42, 0x8a, 0x5c, 0x3c, 0xb9, 0x99, 0x79, 0xf1, 0xc9,
	0xf9, 0x79, 0xc5, 0xa5, 0xb9, 0xa9, 0x45, 0x12, 0x4c, 0x60, 0x79, 0xee, 0xdc, 0xcc, 0x3c, 0x67,
	0xa8, 0x90, 0x90, 0x32, 0x17, 0x6f, 0x52, 0x62, 0x0a, 0x5c, 0x49, 0xb1, 0x04, 0xb3, 0x02, 0xb3,
	0x06, 0x6b, 0x10, 0x4f, 0x52, 0x62, 0x0a, 0x4c, 0x4d, 0xb1, 0x93, 0xf5, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xc8, 0x25, 0x91, 0x5f, 0x94, 0xae, 0x87, 0x70,
	0x8f, 0x1e, 0xdc, 0xd1, 0x4e, 0x7c, 0x50, 0x77, 0x15, 0x07, 0x80, 0x1c, 0x5d, 0x1c, 0xc0, 0xf8,
	0x83, 0x91, 0x31, 0x89, 0x0d, 0xec, 0x03, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x2d,
	0xc1, 0x28, 0xed, 0x00, 0x00, 0x00,
}