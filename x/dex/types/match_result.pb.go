// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/match_result.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MatchResult struct {
	Height        int64              `protobuf:"varint,1,opt,name=height,proto3" json:"height"`
	ContractAddr  string             `protobuf:"bytes,2,opt,name=contractAddr,proto3" json:"contract_address"`
	Orders        []*Order           `protobuf:"bytes,3,rep,name=orders,proto3" json:"orders"`
	Settlements   []*SettlementEntry `protobuf:"bytes,4,rep,name=settlements,proto3" json:"settlements"`
	Cancellations []*Cancellation    `protobuf:"bytes,5,rep,name=cancellations,proto3" json:"cancellations"`
}

func (m *MatchResult) Reset()         { *m = MatchResult{} }
func (m *MatchResult) String() string { return proto.CompactTextString(m) }
func (*MatchResult) ProtoMessage()    {}
func (*MatchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9225b122096d4ce7, []int{0}
}
func (m *MatchResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MatchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MatchResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MatchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchResult.Merge(m, src)
}
func (m *MatchResult) XXX_Size() int {
	return m.Size()
}
func (m *MatchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchResult.DiscardUnknown(m)
}

var xxx_messageInfo_MatchResult proto.InternalMessageInfo

func (m *MatchResult) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MatchResult) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func (m *MatchResult) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *MatchResult) GetSettlements() []*SettlementEntry {
	if m != nil {
		return m.Settlements
	}
	return nil
}

func (m *MatchResult) GetCancellations() []*Cancellation {
	if m != nil {
		return m.Cancellations
	}
	return nil
}

func init() {
	proto.RegisterType((*MatchResult)(nil), "seiprotocol.seichain.dex.MatchResult")
}

func init() { proto.RegisterFile("dex/match_result.proto", fileDescriptor_9225b122096d4ce7) }

var fileDescriptor_9225b122096d4ce7 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x6a, 0xfa, 0x40,
	0x10, 0xc6, 0x8d, 0xfe, 0xff, 0x42, 0xd7, 0x16, 0xdb, 0x20, 0x25, 0x78, 0x48, 0xc4, 0x43, 0xb1,
	0x07, 0x13, 0x68, 0x2f, 0xbd, 0x36, 0x52, 0x7a, 0x2a, 0x85, 0xf4, 0x56, 0x0a, 0x12, 0x77, 0x87,
	0x64, 0x21, 0x66, 0x65, 0x77, 0x04, 0x7d, 0x8b, 0x3e, 0x56, 0x8f, 0x1e, 0x7b, 0x0a, 0x45, 0x6f,
	0xb9, 0xf4, 0x15, 0xca, 0x6e, 0xb4, 0xd5, 0x83, 0xa7, 0x99, 0xf9, 0xf2, 0xcd, 0xef, 0xcb, 0xb0,
	0xe4, 0x92, 0xc1, 0x22, 0x98, 0xc6, 0x48, 0xd3, 0xb1, 0x04, 0x35, 0xcf, 0xd0, 0x9f, 0x49, 0x81,
	0xc2, 0x76, 0x14, 0x70, 0xd3, 0x51, 0x91, 0xf9, 0x0a, 0x38, 0x4d, 0x63, 0x9e, 0xfb, 0x0c, 0x16,
	0xdd, 0xb6, 0xde, 0x10, 0x92, 0x81, 0xac, 0xac, 0xdd, 0x8e, 0x16, 0x14, 0x20, 0x66, 0x30, 0x85,
	0x1c, 0x77, 0x6a, 0x22, 0x12, 0x61, 0xda, 0x40, 0x77, 0x95, 0xda, 0xff, 0xae, 0x93, 0xd6, 0x93,
	0x4e, 0x8b, 0x4c, 0x98, 0xdd, 0x27, 0xcd, 0x14, 0x78, 0x92, 0xa2, 0x63, 0xf5, 0xac, 0x41, 0x23,
	0x24, 0x65, 0xe1, 0x6d, 0x95, 0x68, 0x5b, 0xed, 0x3b, 0x72, 0x4a, 0x45, 0x8e, 0x32, 0xa6, 0x78,
	0xcf, 0x98, 0x74, 0xea, 0x3d, 0x6b, 0x70, 0x12, 0x76, 0xca, 0xc2, 0x3b, 0xdf, 0xe9, 0xe3, 0x98,
	0x31, 0x09, 0x4a, 0x45, 0x07, 0x4e, 0x7b, 0x44, 0x9a, 0xe6, 0x47, 0x95, 0xd3, 0xe8, 0x35, 0x06,
	0xad, 0x1b, 0xcf, 0x3f, 0x76, 0x95, 0xff, 0xac, 0x7d, 0x55, 0x7c, 0xb5, 0x12, 0x6d, 0xab, 0xfd,
	0x46, 0x5a, 0x7f, 0xc7, 0x29, 0xe7, 0x9f, 0x21, 0x5d, 0x1f, 0x27, 0xbd, 0xfc, 0x9a, 0x1f, 0x72,
	0x94, 0xcb, 0xb0, 0x5d, 0x16, 0xde, 0x3e, 0x21, 0xda, 0x1f, 0xec, 0x31, 0x39, 0xa3, 0x71, 0x4e,
	0x21, 0xcb, 0x62, 0xe4, 0x22, 0x57, 0xce, 0x7f, 0xc3, 0xbf, 0x3a, 0xce, 0x1f, 0xed, 0xd9, 0xc3,
	0x8b, 0xb2, 0xf0, 0x0e, 0x01, 0xd1, 0xe1, 0x18, 0x3e, 0x7e, 0xac, 0x5d, 0x6b, 0xb5, 0x76, 0xad,
	0xaf, 0xb5, 0x6b, 0xbd, 0x6f, 0xdc, 0xda, 0x6a, 0xe3, 0xd6, 0x3e, 0x37, 0x6e, 0xed, 0x75, 0x98,
	0x70, 0x4c, 0xe7, 0x13, 0x9f, 0x8a, 0x69, 0xa0, 0x80, 0x0f, 0x77, 0x71, 0x66, 0x30, 0x79, 0xc1,
	0x22, 0xd0, 0x6f, 0x8b, 0xcb, 0x19, 0xa8, 0x49, 0xd3, 0x7c, 0xbf, 0xfd, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x40, 0xd8, 0xfe, 0x79, 0x32, 0x02, 0x00, 0x00,
}

func (m *MatchResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MatchResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MatchResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cancellations) > 0 {
		for iNdEx := len(m.Cancellations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Cancellations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMatchResult(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Settlements) > 0 {
		for iNdEx := len(m.Settlements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Settlements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMatchResult(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Orders) > 0 {
		for iNdEx := len(m.Orders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Orders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMatchResult(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ContractAddr) > 0 {
		i -= len(m.ContractAddr)
		copy(dAtA[i:], m.ContractAddr)
		i = encodeVarintMatchResult(dAtA, i, uint64(len(m.ContractAddr)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintMatchResult(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMatchResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovMatchResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MatchResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovMatchResult(uint64(m.Height))
	}
	l = len(m.ContractAddr)
	if l > 0 {
		n += 1 + l + sovMatchResult(uint64(l))
	}
	if len(m.Orders) > 0 {
		for _, e := range m.Orders {
			l = e.Size()
			n += 1 + l + sovMatchResult(uint64(l))
		}
	}
	if len(m.Settlements) > 0 {
		for _, e := range m.Settlements {
			l = e.Size()
			n += 1 + l + sovMatchResult(uint64(l))
		}
	}
	if len(m.Cancellations) > 0 {
		for _, e := range m.Cancellations {
			l = e.Size()
			n += 1 + l + sovMatchResult(uint64(l))
		}
	}
	return n
}

func sovMatchResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMatchResult(x uint64) (n int) {
	return sovMatchResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MatchResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMatchResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MatchResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MatchResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatchResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatchResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMatchResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMatchResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatchResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMatchResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMatchResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orders = append(m.Orders, &Order{})
			if err := m.Orders[len(m.Orders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Settlements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatchResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMatchResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMatchResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Settlements = append(m.Settlements, &SettlementEntry{})
			if err := m.Settlements[len(m.Settlements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancellations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMatchResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMatchResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMatchResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cancellations = append(m.Cancellations, &Cancellation{})
			if err := m.Cancellations[len(m.Cancellations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMatchResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMatchResult
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
func skipMatchResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMatchResult
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
					return 0, ErrIntOverflowMatchResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMatchResult
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
			if length < 0 {
				return 0, ErrInvalidLengthMatchResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMatchResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMatchResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMatchResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMatchResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMatchResult = fmt.Errorf("proto: unexpected end of group")
)
