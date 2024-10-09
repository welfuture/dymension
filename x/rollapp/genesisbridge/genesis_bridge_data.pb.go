// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/genesisbridge/genesis_bridge_data.proto

package genesisbridge

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	types2 "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
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

// GenesisBridgeData is the data struct that is passed to the hub for the
// genesis bridge flow
type GenesisBridgeData struct {
	// genesis_info is the genesis info of the rollapp. used for hub validation
	GenesisInfo GenesisBridgeInfo `protobuf:"bytes,1,opt,name=genesis_info,json=genesisInfo,proto3" json:"genesis_info"`
	// native_denom is the native denom of the rollapp. registered on the hub
	NativeDenom types.Metadata `protobuf:"bytes,2,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom"`
	// optional genesis transfer packet data
	GenesisTransfer *types1.FungibleTokenPacketData `protobuf:"bytes,3,opt,name=genesis_transfer,json=genesisTransfer,proto3" json:"genesis_transfer,omitempty"`
}

func (m *GenesisBridgeData) Reset()         { *m = GenesisBridgeData{} }
func (m *GenesisBridgeData) String() string { return proto.CompactTextString(m) }
func (*GenesisBridgeData) ProtoMessage()    {}
func (*GenesisBridgeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6038c1fc23369a68, []int{0}
}
func (m *GenesisBridgeData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisBridgeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisBridgeData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisBridgeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisBridgeData.Merge(m, src)
}
func (m *GenesisBridgeData) XXX_Size() int {
	return m.Size()
}
func (m *GenesisBridgeData) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisBridgeData.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisBridgeData proto.InternalMessageInfo

func (m *GenesisBridgeData) GetGenesisInfo() GenesisBridgeInfo {
	if m != nil {
		return m.GenesisInfo
	}
	return GenesisBridgeInfo{}
}

func (m *GenesisBridgeData) GetNativeDenom() types.Metadata {
	if m != nil {
		return m.NativeDenom
	}
	return types.Metadata{}
}

func (m *GenesisBridgeData) GetGenesisTransfer() *types1.FungibleTokenPacketData {
	if m != nil {
		return m.GenesisTransfer
	}
	return nil
}

// The genesis info of the rollapp, that is passed to the hub for validation.
// it's populated on the InitGenesis of the rollapp
type GenesisBridgeInfo struct {
	// checksum used to verify integrity of the genesis file. currently unused
	GenesisChecksum string `protobuf:"bytes,1,opt,name=genesis_checksum,json=genesisChecksum,proto3" json:"genesis_checksum,omitempty"`
	// unique bech32 prefix
	Bech32Prefix string `protobuf:"bytes,2,opt,name=bech32_prefix,json=bech32Prefix,proto3" json:"bech32_prefix,omitempty"`
	// native_denom is the base denom for the native token
	NativeDenom types2.DenomMetadata `protobuf:"bytes,3,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom"`
	// initial_supply is the initial supply of the native token
	InitialSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=initial_supply,json=initialSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_supply"`
	// accounts on the Hub to fund with some bootstrapping transfers
	GenesisAccounts []types2.GenesisAccount `protobuf:"bytes,5,rep,name=genesis_accounts,json=genesisAccounts,proto3" json:"genesis_accounts"`
}

func (m *GenesisBridgeInfo) Reset()         { *m = GenesisBridgeInfo{} }
func (m *GenesisBridgeInfo) String() string { return proto.CompactTextString(m) }
func (*GenesisBridgeInfo) ProtoMessage()    {}
func (*GenesisBridgeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6038c1fc23369a68, []int{1}
}
func (m *GenesisBridgeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisBridgeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisBridgeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisBridgeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisBridgeInfo.Merge(m, src)
}
func (m *GenesisBridgeInfo) XXX_Size() int {
	return m.Size()
}
func (m *GenesisBridgeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisBridgeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisBridgeInfo proto.InternalMessageInfo

func (m *GenesisBridgeInfo) GetGenesisChecksum() string {
	if m != nil {
		return m.GenesisChecksum
	}
	return ""
}

func (m *GenesisBridgeInfo) GetBech32Prefix() string {
	if m != nil {
		return m.Bech32Prefix
	}
	return ""
}

func (m *GenesisBridgeInfo) GetNativeDenom() types2.DenomMetadata {
	if m != nil {
		return m.NativeDenom
	}
	return types2.DenomMetadata{}
}

func (m *GenesisBridgeInfo) GetGenesisAccounts() []types2.GenesisAccount {
	if m != nil {
		return m.GenesisAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisBridgeData)(nil), "dymensionxyz.dymension.genesisbridge.GenesisBridgeData")
	proto.RegisterType((*GenesisBridgeInfo)(nil), "dymensionxyz.dymension.genesisbridge.GenesisBridgeInfo")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/genesisbridge/genesis_bridge_data.proto", fileDescriptor_6038c1fc23369a68)
}

var fileDescriptor_6038c1fc23369a68 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xdb, 0x6d, 0x20, 0x2d, 0xed, 0xf8, 0x13, 0x71, 0xa8, 0x2a, 0x91, 0x4d, 0x03, 0xa1,
	0xed, 0x50, 0x5b, 0x6d, 0x85, 0xb8, 0x20, 0x24, 0xca, 0x34, 0xb4, 0x03, 0xd2, 0x54, 0xc6, 0x0e,
	0x5c, 0x32, 0xc7, 0x71, 0x53, 0x2b, 0x89, 0x6d, 0xc5, 0x4e, 0xd4, 0xf2, 0x29, 0xf8, 0x06, 0x7c,
	0x9d, 0x1d, 0x77, 0x44, 0x1c, 0x26, 0xd4, 0x7e, 0x04, 0xbe, 0x00, 0x8a, 0xed, 0xb4, 0x5d, 0xa5,
	0x32, 0x4e, 0xed, 0xfb, 0xf6, 0x7d, 0x1f, 0x3f, 0xcf, 0xaf, 0xaf, 0xf3, 0x2e, 0x9c, 0xa6, 0x84,
	0x49, 0xca, 0xd9, 0x64, 0xfa, 0x0d, 0x2e, 0x0a, 0x18, 0x11, 0x46, 0x24, 0x95, 0x41, 0x46, 0xc3,
	0x88, 0x54, 0x95, 0x6f, 0x4a, 0x3f, 0x44, 0x0a, 0x01, 0x91, 0x71, 0xc5, 0xdd, 0x97, 0xab, 0xfb,
	0x60, 0x51, 0x80, 0x3b, 0xfb, 0xed, 0x67, 0x11, 0x8f, 0xb8, 0x5e, 0x80, 0xe5, 0x37, 0xb3, 0xdb,
	0xf6, 0x30, 0x97, 0x29, 0x97, 0x30, 0x40, 0x2c, 0x86, 0x45, 0x37, 0x20, 0x0a, 0x75, 0x75, 0x61,
	0x7f, 0x3f, 0xa6, 0x01, 0x86, 0x48, 0x88, 0x84, 0x62, 0xa4, 0x28, 0x67, 0x12, 0xaa, 0x0c, 0x31,
	0x39, 0x22, 0x19, 0x2c, 0x7a, 0x50, 0x20, 0x1c, 0x13, 0x65, 0x47, 0x3b, 0x1b, 0x62, 0x64, 0x3c,
	0x49, 0x90, 0x10, 0x30, 0x25, 0x0a, 0x2d, 0x5d, 0xb7, 0xbb, 0xf7, 0x8c, 0x57, 0x79, 0x29, 0x1b,
	0x59, 0xb3, 0x87, 0x3f, 0xb6, 0x9c, 0xa7, 0x1f, 0x4d, 0x7b, 0xa0, 0x43, 0x9d, 0x20, 0x85, 0xdc,
	0x2b, 0xa7, 0xb9, 0x3a, 0xdb, 0xaa, 0x1f, 0xd4, 0x8f, 0x1a, 0xbd, 0x37, 0xe0, 0x7f, 0xa8, 0x80,
	0x3b, 0x72, 0x67, 0x6c, 0xc4, 0x07, 0x3b, 0xd7, 0xb7, 0xfb, 0xb5, 0x61, 0xc3, 0x8e, 0x95, 0x2d,
	0xf7, 0xd4, 0x69, 0x32, 0xa4, 0x68, 0x41, 0xfc, 0x90, 0x30, 0x9e, 0xb6, 0xb6, 0xf4, 0x0b, 0xcf,
	0x81, 0x61, 0x07, 0x34, 0x2e, 0xcb, 0x0e, 0x7c, 0xb2, 0x29, 0x2b, 0x1d, 0xb3, 0x78, 0x52, 0xee,
	0xb9, 0x57, 0xce, 0x93, 0xca, 0x69, 0x45, 0xb1, 0xb5, 0xad, 0xb5, 0x5e, 0x03, 0x1a, 0x60, 0xb0,
	0xca, 0x19, 0x54, 0x13, 0xa0, 0xe8, 0x81, 0xd3, 0x9c, 0x45, 0x34, 0x48, 0xc8, 0x05, 0x8f, 0x09,
	0x3b, 0xd7, 0xd0, 0xcb, 0xe8, 0xc3, 0xc7, 0x56, 0xee, 0xc2, 0xce, 0x1e, 0xfe, 0x59, 0x27, 0xa4,
	0xfd, 0x1f, 0x2f, 0xdf, 0xc5, 0x63, 0x82, 0x63, 0x99, 0xa7, 0x9a, 0xd2, 0xee, 0x42, 0xe0, 0x83,
	0x6d, 0xbb, 0x2f, 0x9c, 0xbd, 0x80, 0xe0, 0x71, 0xbf, 0xe7, 0x8b, 0x8c, 0x8c, 0xe8, 0x44, 0x67,
	0xdd, 0x1d, 0x36, 0x4d, 0xf3, 0x5c, 0xf7, 0xdc, 0xcb, 0x35, 0x1e, 0x26, 0x43, 0x67, 0x13, 0x71,
	0xfb, 0x8f, 0x02, 0x0d, 0xe1, 0x5f, 0x7c, 0xbe, 0x38, 0x8f, 0x28, 0xa3, 0x8a, 0xa2, 0xc4, 0x97,
	0xb9, 0x10, 0xc9, 0xb4, 0xb5, 0x53, 0xbe, 0x3e, 0x00, 0xe5, 0xe8, 0xaf, 0xdb, 0xfd, 0x57, 0x11,
	0x55, 0xe3, 0x3c, 0x00, 0x98, 0xa7, 0xd0, 0xde, 0xad, 0xf9, 0xe8, 0xc8, 0x30, 0x86, 0x6a, 0x2a,
	0x88, 0x04, 0x67, 0x4c, 0x0d, 0xf7, 0xac, 0xca, 0x67, 0x2d, 0xe2, 0xfa, 0xcb, 0xf8, 0x08, 0x63,
	0x9e, 0x33, 0x25, 0x5b, 0x0f, 0x0e, 0xb6, 0x8f, 0x1a, 0x3d, 0x70, 0x9f, 0x65, 0xcb, 0xf2, 0xbd,
	0x59, 0xb3, 0x9e, 0x2b, 0x68, 0xb6, 0x2b, 0x07, 0x97, 0xd7, 0x33, 0xaf, 0x7e, 0x33, 0xf3, 0xea,
	0xbf, 0x67, 0x5e, 0xfd, 0xfb, 0xdc, 0xab, 0xdd, 0xcc, 0xbd, 0xda, 0xcf, 0xb9, 0x57, 0xfb, 0xfa,
	0x76, 0xc5, 0xf1, 0x86, 0x7b, 0x2f, 0xfa, 0x70, 0xb2, 0x7e, 0xf4, 0xe6, 0x38, 0x83, 0x87, 0xfa,
	0xec, 0xfb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xc1, 0x44, 0xd3, 0x21, 0x04, 0x00, 0x00,
}

func (m *GenesisBridgeData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisBridgeData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisBridgeData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GenesisTransfer != nil {
		{
			size, err := m.GenesisTransfer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesisBridgeData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.NativeDenom.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesisBridgeData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.GenesisInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesisBridgeData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisBridgeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisBridgeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisBridgeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisAccounts) > 0 {
		for iNdEx := len(m.GenesisAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesisBridgeData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size := m.InitialSupply.Size()
		i -= size
		if _, err := m.InitialSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesisBridgeData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.NativeDenom.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesisBridgeData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bech32Prefix) > 0 {
		i -= len(m.Bech32Prefix)
		copy(dAtA[i:], m.Bech32Prefix)
		i = encodeVarintGenesisBridgeData(dAtA, i, uint64(len(m.Bech32Prefix)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GenesisChecksum) > 0 {
		i -= len(m.GenesisChecksum)
		copy(dAtA[i:], m.GenesisChecksum)
		i = encodeVarintGenesisBridgeData(dAtA, i, uint64(len(m.GenesisChecksum)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesisBridgeData(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesisBridgeData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisBridgeData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.GenesisInfo.Size()
	n += 1 + l + sovGenesisBridgeData(uint64(l))
	l = m.NativeDenom.Size()
	n += 1 + l + sovGenesisBridgeData(uint64(l))
	if m.GenesisTransfer != nil {
		l = m.GenesisTransfer.Size()
		n += 1 + l + sovGenesisBridgeData(uint64(l))
	}
	return n
}

func (m *GenesisBridgeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GenesisChecksum)
	if l > 0 {
		n += 1 + l + sovGenesisBridgeData(uint64(l))
	}
	l = len(m.Bech32Prefix)
	if l > 0 {
		n += 1 + l + sovGenesisBridgeData(uint64(l))
	}
	l = m.NativeDenom.Size()
	n += 1 + l + sovGenesisBridgeData(uint64(l))
	l = m.InitialSupply.Size()
	n += 1 + l + sovGenesisBridgeData(uint64(l))
	if len(m.GenesisAccounts) > 0 {
		for _, e := range m.GenesisAccounts {
			l = e.Size()
			n += 1 + l + sovGenesisBridgeData(uint64(l))
		}
	}
	return n
}

func sovGenesisBridgeData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesisBridgeData(x uint64) (n int) {
	return sovGenesisBridgeData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisBridgeData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisBridgeData
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
			return fmt.Errorf("proto: GenesisBridgeData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisBridgeData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NativeDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisTransfer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GenesisTransfer == nil {
				m.GenesisTransfer = &types1.FungibleTokenPacketData{}
			}
			if err := m.GenesisTransfer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisBridgeData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisBridgeData
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
func (m *GenesisBridgeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesisBridgeData
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
			return fmt.Errorf("proto: GenesisBridgeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisBridgeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisChecksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisChecksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NativeDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesisBridgeData
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
				return ErrInvalidLengthGenesisBridgeData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesisBridgeData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisAccounts = append(m.GenesisAccounts, types2.GenesisAccount{})
			if err := m.GenesisAccounts[len(m.GenesisAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesisBridgeData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesisBridgeData
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
func skipGenesisBridgeData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesisBridgeData
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
					return 0, ErrIntOverflowGenesisBridgeData
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
					return 0, ErrIntOverflowGenesisBridgeData
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
				return 0, ErrInvalidLengthGenesisBridgeData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesisBridgeData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesisBridgeData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesisBridgeData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesisBridgeData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesisBridgeData = fmt.Errorf("proto: unexpected end of group")
)
