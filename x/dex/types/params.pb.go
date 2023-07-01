// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	PriceSnapshotRetention     uint64                                 `protobuf:"varint,1,opt,name=price_snapshot_retention,json=priceSnapshotRetention,proto3" json:"price_snapshot_retention" yaml:"price_snapshot_retention"`
	SudoCallGasPrice           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=sudo_call_gas_price,json=sudoCallGasPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"sudo_call_gas_price"`
	BeginBlockGasLimit         uint64                                 `protobuf:"varint,3,opt,name=begin_block_gas_limit,json=beginBlockGasLimit,proto3" json:"begin_block_gas_limit" yaml:"begin_block_gas_limit"`
	EndBlockGasLimit           uint64                                 `protobuf:"varint,4,opt,name=end_block_gas_limit,json=endBlockGasLimit,proto3" json:"end_block_gas_limit" yaml:"end_block_gas_limit"`
	DefaultGasPerOrder         uint64                                 `protobuf:"varint,5,opt,name=default_gas_per_order,json=defaultGasPerOrder,proto3" json:"default_gas_per_order" yaml:"default_gas_per_order"`
	DefaultGasPerCancel        uint64                                 `protobuf:"varint,6,opt,name=default_gas_per_cancel,json=defaultGasPerCancel,proto3" json:"default_gas_per_cancel" yaml:"default_gas_per_cancel"`
	MinRentDeposit             uint64                                 `protobuf:"varint,7,opt,name=min_rent_deposit,json=minRentDeposit,proto3" json:"min_rent_deposit" yaml:"min_rent_deposit"`
	GasAllowancePerSettlement  uint64                                 `protobuf:"varint,8,opt,name=gas_allowance_per_settlement,json=gasAllowancePerSettlement,proto3" json:"gas_allowance_per_settlement" yaml:"gas_allowance_per_settlement"`
	MinProcessableRent         uint64                                 `protobuf:"varint,9,opt,name=min_processable_rent,json=minProcessableRent,proto3" json:"min_processable_rent" yaml:"min_processable_rent"`
	OrderBookEntriesPerLoad    uint64                                 `protobuf:"varint,10,opt,name=order_book_entries_per_load,json=orderBookEntriesPerLoad,proto3" json:"order_book_entries_per_load" yaml:"order_book_entries_per_load"`
	ContractUnsuspendCost      uint64                                 `protobuf:"varint,11,opt,name=contract_unsuspend_cost,json=contractUnsuspendCost,proto3" json:"contract_unsuspend_cost" yaml:"contract_unsuspend_cost"`
	MaxOrderPerPrice           uint64                                 `protobuf:"varint,12,opt,name=max_order_per_price,json=maxOrderPerPrice,proto3" json:"max_order_per_price" yaml:"max_order_per_price"`
	MaxPairsPerContract        uint64                                 `protobuf:"varint,13,opt,name=max_pairs_per_contract,json=maxPairsPerContract,proto3" json:"max_pairs_per_contract" yaml:"max_pairs_per_contract"`
	DefaultGasPerOrderDataByte uint64                                 `protobuf:"varint,14,opt,name=default_gas_per_order_data_byte,json=defaultGasPerOrderDataByte,proto3" json:"default_gas_per_order_data_byte" yaml:"default_gas_per_order_data_byte"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e49286500ccff43e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPriceSnapshotRetention() uint64 {
	if m != nil {
		return m.PriceSnapshotRetention
	}
	return 0
}

func (m *Params) GetBeginBlockGasLimit() uint64 {
	if m != nil {
		return m.BeginBlockGasLimit
	}
	return 0
}

func (m *Params) GetEndBlockGasLimit() uint64 {
	if m != nil {
		return m.EndBlockGasLimit
	}
	return 0
}

func (m *Params) GetDefaultGasPerOrder() uint64 {
	if m != nil {
		return m.DefaultGasPerOrder
	}
	return 0
}

func (m *Params) GetDefaultGasPerCancel() uint64 {
	if m != nil {
		return m.DefaultGasPerCancel
	}
	return 0
}

func (m *Params) GetMinRentDeposit() uint64 {
	if m != nil {
		return m.MinRentDeposit
	}
	return 0
}

func (m *Params) GetGasAllowancePerSettlement() uint64 {
	if m != nil {
		return m.GasAllowancePerSettlement
	}
	return 0
}

func (m *Params) GetMinProcessableRent() uint64 {
	if m != nil {
		return m.MinProcessableRent
	}
	return 0
}

func (m *Params) GetOrderBookEntriesPerLoad() uint64 {
	if m != nil {
		return m.OrderBookEntriesPerLoad
	}
	return 0
}

func (m *Params) GetContractUnsuspendCost() uint64 {
	if m != nil {
		return m.ContractUnsuspendCost
	}
	return 0
}

func (m *Params) GetMaxOrderPerPrice() uint64 {
	if m != nil {
		return m.MaxOrderPerPrice
	}
	return 0
}

func (m *Params) GetMaxPairsPerContract() uint64 {
	if m != nil {
		return m.MaxPairsPerContract
	}
	return 0
}

func (m *Params) GetDefaultGasPerOrderDataByte() uint64 {
	if m != nil {
		return m.DefaultGasPerOrderDataByte
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "seiprotocol.seichain.dex.Params")
}

func init() { proto.RegisterFile("dex/params.proto", fileDescriptor_e49286500ccff43e) }

var fileDescriptor_e49286500ccff43e = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xc1, 0x6f, 0x1b, 0x45,
	0x14, 0xc6, 0xbd, 0x50, 0x42, 0x3b, 0x94, 0xca, 0xda, 0x34, 0xc9, 0x92, 0x16, 0x4f, 0x35, 0x48,
	0x55, 0x2f, 0xb1, 0x0f, 0x08, 0x21, 0x8a, 0x10, 0xc2, 0x49, 0x94, 0x4b, 0x11, 0xd6, 0x54, 0x1c,
	0xe0, 0xb2, 0x1a, 0xef, 0x3e, 0x9c, 0x51, 0x66, 0x67, 0x56, 0x33, 0x63, 0x61, 0x9f, 0xb9, 0x70,
	0x44, 0x9c, 0x38, 0xf6, 0xcf, 0xe9, 0xb1, 0x47, 0xc4, 0x61, 0x84, 0x92, 0x0b, 0xda, 0xe3, 0xfe,
	0x05, 0x68, 0x66, 0x6d, 0x96, 0x24, 0x6b, 0xf7, 0x94, 0xcd, 0xf7, 0xfb, 0xb4, 0xdf, 0x7b, 0xde,
	0x79, 0x6f, 0x50, 0x3f, 0x87, 0xc5, 0xa8, 0x64, 0x9a, 0x15, 0x66, 0x58, 0x6a, 0x65, 0x55, 0x9c,
	0x18, 0xe0, 0xe1, 0x29, 0x53, 0x62, 0x68, 0x80, 0x67, 0xe7, 0x8c, 0xcb, 0x61, 0x0e, 0x8b, 0xc3,
	0x87, 0x33, 0x35, 0x53, 0x01, 0x8d, 0xfc, 0x53, 0xe3, 0x27, 0xf5, 0x7d, 0xb4, 0x33, 0x09, 0x2f,
	0x88, 0x97, 0x28, 0x29, 0x35, 0xcf, 0x20, 0x35, 0x92, 0x95, 0xe6, 0x5c, 0xd9, 0x54, 0x83, 0x05,
	0x69, 0xb9, 0x92, 0x49, 0xf4, 0x24, 0x7a, 0x76, 0x67, 0xfc, 0x75, 0xe5, 0xf0, 0x46, 0x4f, 0xed,
	0x30, 0x5e, 0xb2, 0x42, 0x3c, 0x27, 0x9b, 0x1c, 0x84, 0xee, 0x07, 0xf4, 0x72, 0x45, 0xe8, 0x1a,
	0xc4, 0x16, 0xed, 0x9a, 0x79, 0xae, 0xd2, 0x8c, 0x09, 0x91, 0xce, 0x98, 0x49, 0x83, 0x2f, 0x79,
	0xe7, 0x49, 0xf4, 0xec, 0xde, 0xf8, 0xf4, 0xb5, 0xc3, 0xbd, 0xbf, 0x1c, 0x7e, 0x3a, 0xe3, 0xf6,
	0x7c, 0x3e, 0x1d, 0x66, 0xaa, 0x18, 0x65, 0xca, 0x14, 0xca, 0xac, 0xfe, 0x1c, 0x99, 0xfc, 0x62,
	0x64, 0x97, 0x25, 0x98, 0xe1, 0x09, 0x64, 0x95, 0xc3, 0x5d, 0x2f, 0xa3, 0x7d, 0x2f, 0x1e, 0x33,
	0x21, 0xce, 0x98, 0x99, 0x78, 0x25, 0x16, 0x68, 0x6f, 0x0a, 0x33, 0x2e, 0xd3, 0xa9, 0x50, 0xd9,
	0x45, 0xb0, 0x0a, 0x5e, 0x70, 0x9b, 0xbc, 0x1b, 0xba, 0xfd, 0xa2, 0x72, 0xb8, 0xdb, 0x50, 0x3b,
	0xfc, 0xb8, 0x69, 0xb5, 0x13, 0x13, 0x1a, 0x07, 0x7d, 0xec, 0xe5, 0x33, 0x66, 0x5e, 0x78, 0x31,
	0xce, 0xd1, 0x2e, 0xc8, 0xfc, 0x56, 0xd6, 0x9d, 0x90, 0xf5, 0x99, 0xaf, 0xba, 0x03, 0xd7, 0x0e,
	0x1f, 0x36, 0x49, 0x1d, 0x90, 0xd0, 0x3e, 0xc8, 0xfc, 0x7a, 0x8a, 0x40, 0x7b, 0x39, 0xfc, 0xc4,
	0xe6, 0xc2, 0x36, 0xad, 0x83, 0x4e, 0x95, 0xce, 0x41, 0x27, 0xef, 0xb5, 0x3d, 0x75, 0x1a, 0xda,
	0x9e, 0x3a, 0x31, 0xa1, 0xf1, 0x4a, 0xf7, 0x3f, 0x1f, 0xe8, 0xef, 0xbc, 0x18, 0x97, 0x68, 0xff,
	0xa6, 0x3b, 0x63, 0x32, 0x03, 0x91, 0xec, 0x84, 0xb8, 0x2f, 0x2b, 0x87, 0x37, 0x38, 0x6a, 0x87,
	0x3f, 0xee, 0xce, 0x6b, 0x38, 0xa1, 0xbb, 0xd7, 0x02, 0x8f, 0x83, 0x1a, 0xff, 0x80, 0xfa, 0x05,
	0x97, 0xa9, 0x06, 0x69, 0xd3, 0x1c, 0x4a, 0x65, 0xb8, 0x4d, 0xde, 0x0f, 0x59, 0xa3, 0xca, 0xe1,
	0x5b, 0xac, 0x76, 0xf8, 0xa0, 0x49, 0xb9, 0x49, 0x08, 0x7d, 0x50, 0x70, 0x49, 0x41, 0xda, 0x93,
	0x46, 0x88, 0x7f, 0x8d, 0xd0, 0x63, 0x5f, 0x03, 0x13, 0x42, 0xfd, 0xec, 0xd3, 0x42, 0x35, 0x06,
	0xac, 0x15, 0x50, 0x80, 0xb4, 0xc9, 0xdd, 0x90, 0x73, 0x56, 0x39, 0xbc, 0xd5, 0x57, 0x3b, 0xfc,
	0x49, 0x93, 0xb9, 0xcd, 0x45, 0xe8, 0x47, 0x33, 0x66, 0xbe, 0x59, 0xd3, 0x09, 0xe8, 0x97, 0xff,
	0xb1, 0x98, 0xa3, 0x87, 0xbe, 0xde, 0x52, 0xab, 0x0c, 0x8c, 0x61, 0x53, 0x01, 0xa1, 0xf6, 0xe4,
	0x5e, 0xa8, 0xe0, 0xf3, 0xca, 0xe1, 0x4e, 0x5e, 0x3b, 0xfc, 0xa8, 0xed, 0xf6, 0x26, 0x25, 0x34,
	0x2e, 0xb8, 0x9c, 0xb4, 0xaa, 0x6f, 0x3e, 0xfe, 0x25, 0x42, 0x8f, 0xc2, 0x17, 0x4e, 0xa7, 0x4a,
	0x5d, 0xa4, 0x20, 0xad, 0xe6, 0xd0, 0x7c, 0x08, 0xa1, 0x58, 0x9e, 0xa0, 0x10, 0x79, 0x5a, 0x39,
	0xbc, 0xcd, 0x56, 0x3b, 0x4c, 0x9a, 0xe4, 0x2d, 0x26, 0x42, 0x0f, 0x02, 0x1d, 0x2b, 0x75, 0x71,
	0xda, 0xb0, 0x09, 0xe8, 0x17, 0x8a, 0xe5, 0xf1, 0x1c, 0x1d, 0x64, 0x4a, 0x5a, 0xcd, 0x32, 0x9b,
	0xce, 0xa5, 0x99, 0x9b, 0xd2, 0x9f, 0xf7, 0x4c, 0x19, 0x9b, 0x7c, 0x10, 0x0a, 0xf8, 0xaa, 0x72,
	0x78, 0x93, 0xa5, 0x76, 0x78, 0xd0, 0x84, 0x6f, 0x30, 0x10, 0xba, 0xb7, 0x26, 0xdf, 0xaf, 0xc1,
	0xb1, 0x32, 0x61, 0x26, 0x0b, 0xb6, 0x68, 0x4e, 0x78, 0x28, 0xb3, 0xd9, 0x3b, 0xf7, 0xdb, 0x99,
	0xec, 0xc0, 0xed, 0x4c, 0x76, 0x40, 0x42, 0xfb, 0x05, 0x5b, 0x84, 0xe9, 0x98, 0x80, 0x6e, 0xf6,
	0x4c, 0x89, 0xf6, 0xbd, 0xb3, 0x64, 0x5c, 0xaf, 0x4e, 0xf8, 0xaa, 0x98, 0xe4, 0xc3, 0x76, 0x4a,
	0xba, 0x1d, 0xed, 0x94, 0x74, 0x73, 0x42, 0x7d, 0x85, 0x13, 0xaf, 0xfb, 0x19, 0x59, 0xa9, 0xf1,
	0xef, 0x11, 0xc2, 0x9d, 0x63, 0x9c, 0xe6, 0xcc, 0xb2, 0x74, 0xba, 0xb4, 0x90, 0x3c, 0x08, 0xd9,
	0xdf, 0x56, 0x0e, 0xbf, 0xcd, 0x5a, 0x3b, 0xfc, 0x74, 0xcb, 0x6a, 0x68, 0x8d, 0x84, 0x1e, 0xde,
	0x5e, 0x12, 0x27, 0xcc, 0xb2, 0xf1, 0xd2, 0xc2, 0xf3, 0xbb, 0x7f, 0xbc, 0xc2, 0xbd, 0x7f, 0x5e,
	0xe1, 0x68, 0x7c, 0xf6, 0xfa, 0x72, 0x10, 0xbd, 0xb9, 0x1c, 0x44, 0x7f, 0x5f, 0x0e, 0xa2, 0xdf,
	0xae, 0x06, 0xbd, 0x37, 0x57, 0x83, 0xde, 0x9f, 0x57, 0x83, 0xde, 0x8f, 0x47, 0xff, 0xdb, 0xf1,
	0x06, 0xf8, 0xd1, 0xfa, 0x2a, 0x0b, 0xff, 0x84, 0xbb, 0x6c, 0xb4, 0x18, 0xf9, 0x4b, 0x2f, 0xac,
	0xfb, 0xe9, 0x4e, 0xe0, 0x9f, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xff, 0x27, 0x82, 0x3c, 0x08,
	0x07, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PriceSnapshotRetention != that1.PriceSnapshotRetention {
		return false
	}
	if !this.SudoCallGasPrice.Equal(that1.SudoCallGasPrice) {
		return false
	}
	if this.BeginBlockGasLimit != that1.BeginBlockGasLimit {
		return false
	}
	if this.EndBlockGasLimit != that1.EndBlockGasLimit {
		return false
	}
	if this.DefaultGasPerOrder != that1.DefaultGasPerOrder {
		return false
	}
	if this.DefaultGasPerCancel != that1.DefaultGasPerCancel {
		return false
	}
	if this.MinRentDeposit != that1.MinRentDeposit {
		return false
	}
	if this.GasAllowancePerSettlement != that1.GasAllowancePerSettlement {
		return false
	}
	if this.MinProcessableRent != that1.MinProcessableRent {
		return false
	}
	if this.OrderBookEntriesPerLoad != that1.OrderBookEntriesPerLoad {
		return false
	}
	if this.ContractUnsuspendCost != that1.ContractUnsuspendCost {
		return false
	}
	if this.MaxOrderPerPrice != that1.MaxOrderPerPrice {
		return false
	}
	if this.MaxPairsPerContract != that1.MaxPairsPerContract {
		return false
	}
	if this.DefaultGasPerOrderDataByte != that1.DefaultGasPerOrderDataByte {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DefaultGasPerOrderDataByte != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DefaultGasPerOrderDataByte))
		i--
		dAtA[i] = 0x70
	}
	if m.MaxPairsPerContract != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxPairsPerContract))
		i--
		dAtA[i] = 0x68
	}
	if m.MaxOrderPerPrice != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxOrderPerPrice))
		i--
		dAtA[i] = 0x60
	}
	if m.ContractUnsuspendCost != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ContractUnsuspendCost))
		i--
		dAtA[i] = 0x58
	}
	if m.OrderBookEntriesPerLoad != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.OrderBookEntriesPerLoad))
		i--
		dAtA[i] = 0x50
	}
	if m.MinProcessableRent != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinProcessableRent))
		i--
		dAtA[i] = 0x48
	}
	if m.GasAllowancePerSettlement != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GasAllowancePerSettlement))
		i--
		dAtA[i] = 0x40
	}
	if m.MinRentDeposit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinRentDeposit))
		i--
		dAtA[i] = 0x38
	}
	if m.DefaultGasPerCancel != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DefaultGasPerCancel))
		i--
		dAtA[i] = 0x30
	}
	if m.DefaultGasPerOrder != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DefaultGasPerOrder))
		i--
		dAtA[i] = 0x28
	}
	if m.EndBlockGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockGasLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.BeginBlockGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BeginBlockGasLimit))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.SudoCallGasPrice.Size()
		i -= size
		if _, err := m.SudoCallGasPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PriceSnapshotRetention != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PriceSnapshotRetention))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PriceSnapshotRetention != 0 {
		n += 1 + sovParams(uint64(m.PriceSnapshotRetention))
	}
	l = m.SudoCallGasPrice.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BeginBlockGasLimit != 0 {
		n += 1 + sovParams(uint64(m.BeginBlockGasLimit))
	}
	if m.EndBlockGasLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockGasLimit))
	}
	if m.DefaultGasPerOrder != 0 {
		n += 1 + sovParams(uint64(m.DefaultGasPerOrder))
	}
	if m.DefaultGasPerCancel != 0 {
		n += 1 + sovParams(uint64(m.DefaultGasPerCancel))
	}
	if m.MinRentDeposit != 0 {
		n += 1 + sovParams(uint64(m.MinRentDeposit))
	}
	if m.GasAllowancePerSettlement != 0 {
		n += 1 + sovParams(uint64(m.GasAllowancePerSettlement))
	}
	if m.MinProcessableRent != 0 {
		n += 1 + sovParams(uint64(m.MinProcessableRent))
	}
	if m.OrderBookEntriesPerLoad != 0 {
		n += 1 + sovParams(uint64(m.OrderBookEntriesPerLoad))
	}
	if m.ContractUnsuspendCost != 0 {
		n += 1 + sovParams(uint64(m.ContractUnsuspendCost))
	}
	if m.MaxOrderPerPrice != 0 {
		n += 1 + sovParams(uint64(m.MaxOrderPerPrice))
	}
	if m.MaxPairsPerContract != 0 {
		n += 1 + sovParams(uint64(m.MaxPairsPerContract))
	}
	if m.DefaultGasPerOrderDataByte != 0 {
		n += 1 + sovParams(uint64(m.DefaultGasPerOrderDataByte))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceSnapshotRetention", wireType)
			}
			m.PriceSnapshotRetention = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceSnapshotRetention |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SudoCallGasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SudoCallGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginBlockGasLimit", wireType)
			}
			m.BeginBlockGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BeginBlockGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockGasLimit", wireType)
			}
			m.EndBlockGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultGasPerOrder", wireType)
			}
			m.DefaultGasPerOrder = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultGasPerOrder |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultGasPerCancel", wireType)
			}
			m.DefaultGasPerCancel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultGasPerCancel |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinRentDeposit", wireType)
			}
			m.MinRentDeposit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinRentDeposit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasAllowancePerSettlement", wireType)
			}
			m.GasAllowancePerSettlement = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasAllowancePerSettlement |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinProcessableRent", wireType)
			}
			m.MinProcessableRent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinProcessableRent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderBookEntriesPerLoad", wireType)
			}
			m.OrderBookEntriesPerLoad = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderBookEntriesPerLoad |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractUnsuspendCost", wireType)
			}
			m.ContractUnsuspendCost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractUnsuspendCost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOrderPerPrice", wireType)
			}
			m.MaxOrderPerPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxOrderPerPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPairsPerContract", wireType)
			}
			m.MaxPairsPerContract = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPairsPerContract |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultGasPerOrderDataByte", wireType)
			}
			m.DefaultGasPerOrderDataByte = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultGasPerOrderDataByte |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
