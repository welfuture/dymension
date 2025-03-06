// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/iro/iro.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Bonding curve represents a bonding curve in the IRO module.
// BondingCurve represents a bonding curve with parameters M, N, and C.
// The price of the token is calculated as follows:
// price = M * x^N + C
type BondingCurve struct {
	M cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=M,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"M"`
	N cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=N,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"N"`
	C cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=C,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"C"`
}

func (m *BondingCurve) Reset()         { *m = BondingCurve{} }
func (m *BondingCurve) String() string { return proto.CompactTextString(m) }
func (*BondingCurve) ProtoMessage()    {}
func (*BondingCurve) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d27cc6b5064d3f, []int{0}
}
func (m *BondingCurve) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BondingCurve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BondingCurve.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BondingCurve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BondingCurve.Merge(m, src)
}
func (m *BondingCurve) XXX_Size() int {
	return m.Size()
}
func (m *BondingCurve) XXX_DiscardUnknown() {
	xxx_messageInfo_BondingCurve.DiscardUnknown(m)
}

var xxx_messageInfo_BondingCurve proto.InternalMessageInfo

// Plan represents a plan in the IRO module.
type Plan struct {
	// The ID of the plan.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the rollapp.
	RollappId string `protobuf:"bytes,2,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	// The module account address to hold the raised DYM tokens.
	ModuleAccAddress string `protobuf:"bytes,3,opt,name=module_acc_address,json=moduleAccAddress,proto3" json:"module_acc_address,omitempty"`
	// The total amount of tokens allocated for the IRO.
	TotalAllocation types.Coin   `protobuf:"bytes,4,opt,name=total_allocation,json=totalAllocation,proto3" json:"total_allocation"`
	BondingCurve    BondingCurve `protobuf:"bytes,5,opt,name=bonding_curve,json=bondingCurve,proto3" json:"bonding_curve"`
	// If set, the plan is settled, and the minted allocated tokens can be claimed
	// for this settled_denom
	SettledDenom string `protobuf:"bytes,6,opt,name=settled_denom,json=settledDenom,proto3" json:"settled_denom,omitempty"`
	// The start time of the plan.
	StartTime time.Time `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// The time before which the rollapp cannot be started.
	PreLaunchTime time.Time `protobuf:"bytes,8,opt,name=pre_launch_time,json=preLaunchTime,proto3,stdtime" json:"pre_launch_time"`
	// The amount of tokens sold so far.
	SoldAmt cosmossdk_io_math.Int `protobuf:"bytes,9,opt,name=sold_amt,json=soldAmt,proto3,customtype=cosmossdk.io/math.Int" json:"sold_amt"`
	// The amount of tokens claimed so far.
	ClaimedAmt cosmossdk_io_math.Int `protobuf:"bytes,10,opt,name=claimed_amt,json=claimedAmt,proto3,customtype=cosmossdk.io/math.Int" json:"claimed_amt"`
	// The incentive plan parameters for the tokens left after the plan is
	// settled.
	IncentivePlanParams IncentivePlanParams `protobuf:"bytes,11,opt,name=incentive_plan_params,json=incentivePlanParams,proto3" json:"incentive_plan_params"`
	// The maximum amount of tokens that can be sold for the plan.
	// This ensures we'll have enough tokens to bootstrap liquidity
	MaxAmountToSell cosmossdk_io_math.Int `protobuf:"bytes,12,opt,name=max_amount_to_sell,json=maxAmountToSell,proto3,customtype=cosmossdk.io/math.Int" json:"max_amount_to_sell"`
	// The part of the liquidity that will be used for liquidity pool
	LiquidityPart cosmossdk_io_math.LegacyDec `protobuf:"bytes,13,opt,name=liquidity_part,json=liquidityPart,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"liquidity_part"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d27cc6b5064d3f, []int{1}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Plan) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *Plan) GetModuleAccAddress() string {
	if m != nil {
		return m.ModuleAccAddress
	}
	return ""
}

func (m *Plan) GetTotalAllocation() types.Coin {
	if m != nil {
		return m.TotalAllocation
	}
	return types.Coin{}
}

func (m *Plan) GetBondingCurve() BondingCurve {
	if m != nil {
		return m.BondingCurve
	}
	return BondingCurve{}
}

func (m *Plan) GetSettledDenom() string {
	if m != nil {
		return m.SettledDenom
	}
	return ""
}

func (m *Plan) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Plan) GetPreLaunchTime() time.Time {
	if m != nil {
		return m.PreLaunchTime
	}
	return time.Time{}
}

func (m *Plan) GetIncentivePlanParams() IncentivePlanParams {
	if m != nil {
		return m.IncentivePlanParams
	}
	return IncentivePlanParams{}
}

type IncentivePlanParams struct {
	// start_time_after_settlement is the time after IRO settlement when the
	// distribution of the remaining tokens as incentives will start
	StartTimeAfterSettlement time.Duration `protobuf:"bytes,1,opt,name=start_time_after_settlement,json=startTimeAfterSettlement,proto3,stdduration" json:"start_time_after_settlement"`
	// num_epochs_paid_over is the number of total epochs distribution will be
	// completed over
	NumEpochsPaidOver uint64 `protobuf:"varint,2,opt,name=num_epochs_paid_over,json=numEpochsPaidOver,proto3" json:"num_epochs_paid_over,omitempty"`
}

func (m *IncentivePlanParams) Reset()         { *m = IncentivePlanParams{} }
func (m *IncentivePlanParams) String() string { return proto.CompactTextString(m) }
func (*IncentivePlanParams) ProtoMessage()    {}
func (*IncentivePlanParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7d27cc6b5064d3f, []int{2}
}
func (m *IncentivePlanParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentivePlanParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentivePlanParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentivePlanParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentivePlanParams.Merge(m, src)
}
func (m *IncentivePlanParams) XXX_Size() int {
	return m.Size()
}
func (m *IncentivePlanParams) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentivePlanParams.DiscardUnknown(m)
}

var xxx_messageInfo_IncentivePlanParams proto.InternalMessageInfo

func (m *IncentivePlanParams) GetStartTimeAfterSettlement() time.Duration {
	if m != nil {
		return m.StartTimeAfterSettlement
	}
	return 0
}

func (m *IncentivePlanParams) GetNumEpochsPaidOver() uint64 {
	if m != nil {
		return m.NumEpochsPaidOver
	}
	return 0
}

func init() {
	proto.RegisterType((*BondingCurve)(nil), "dymensionxyz.dymension.iro.BondingCurve")
	proto.RegisterType((*Plan)(nil), "dymensionxyz.dymension.iro.Plan")
	proto.RegisterType((*IncentivePlanParams)(nil), "dymensionxyz.dymension.iro.IncentivePlanParams")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/iro/iro.proto", fileDescriptor_e7d27cc6b5064d3f)
}

var fileDescriptor_e7d27cc6b5064d3f = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xf6, 0xba, 0x6a, 0x6c, 0xd3, 0x56, 0x9c, 0x32, 0x0e, 0xb0, 0x71, 0x50, 0x29, 0x70, 0x0b,
	0xd4, 0x40, 0x91, 0xdd, 0xba, 0x79, 0x80, 0x42, 0x3f, 0x0d, 0xa0, 0x40, 0x49, 0x85, 0x55, 0x0e,
	0x41, 0x2f, 0x04, 0x97, 0x64, 0x56, 0x44, 0xf9, 0xb3, 0xe5, 0x72, 0x05, 0xa9, 0xaf, 0xd0, 0x4b,
	0x8e, 0x7d, 0x86, 0x9e, 0xf3, 0x10, 0x39, 0xa6, 0x39, 0x15, 0x3d, 0xa4, 0x85, 0xfd, 0x22, 0x05,
	0xb9, 0x94, 0x2a, 0x24, 0x8d, 0x01, 0xfb, 0x20, 0x40, 0xe4, 0xcc, 0xf7, 0xcd, 0xcc, 0x7e, 0xdf,
	0x10, 0x7c, 0x49, 0x97, 0x92, 0xa9, 0x8a, 0x6b, 0xb5, 0x58, 0xfe, 0x92, 0xae, 0x0f, 0x29, 0x37,
	0xda, 0xfd, 0x92, 0xd2, 0x68, 0xab, 0xe1, 0xf1, 0x66, 0x56, 0xb2, 0x3e, 0x24, 0xdc, 0xe8, 0xe3,
	0xa3, 0x42, 0x17, 0xda, 0xa7, 0xa5, 0xee, 0x5f, 0x83, 0x38, 0xee, 0x16, 0x5a, 0x17, 0x82, 0xa5,
	0xfe, 0x94, 0xd7, 0x2f, 0x52, 0xcb, 0x25, 0xab, 0x2c, 0x96, 0x65, 0x48, 0xe8, 0xbc, 0x9f, 0x40,
	0x6b, 0x83, 0xad, 0x23, 0x0d, 0x71, 0xa2, 0x2b, 0xa9, 0xab, 0x34, 0xc7, 0x15, 0x4b, 0xe7, 0x67,
	0x39, 0xb3, 0xf8, 0x2c, 0x25, 0x9a, 0xaf, 0xe2, 0x77, 0x9b, 0x38, 0x6a, 0x2a, 0x37, 0x87, 0x10,
	0xfa, 0xea, 0x92, 0x99, 0x4a, 0x6c, 0xb0, 0x0c, 0x89, 0x27, 0x7f, 0x44, 0xe0, 0xa0, 0xaf, 0x15,
	0xe5, 0xaa, 0x18, 0xd4, 0x66, 0xce, 0xe0, 0x77, 0x20, 0x7a, 0x12, 0x47, 0xf7, 0xa3, 0xd3, 0xbd,
	0xfe, 0xd9, 0xeb, 0x77, 0xdd, 0xad, 0xbf, 0xde, 0x75, 0xef, 0x35, 0xd4, 0x15, 0xfd, 0x29, 0xe1,
	0x3a, 0x95, 0xd8, 0xce, 0x92, 0x31, 0x2b, 0x30, 0x59, 0x0e, 0x19, 0x79, 0xfb, 0xea, 0x01, 0x08,
	0x95, 0x87, 0x8c, 0x64, 0xd1, 0x13, 0x47, 0xf0, 0x34, 0xde, 0xbe, 0x36, 0xc1, 0x53, 0x47, 0x30,
	0x88, 0x3f, 0xb9, 0x36, 0xc1, 0xe0, 0xe4, 0xd7, 0x1d, 0xd0, 0x9a, 0x08, 0xac, 0xe0, 0x4d, 0xb0,
	0xcd, 0xa9, 0x1f, 0xa6, 0x95, 0x6d, 0x73, 0x0a, 0x3f, 0x07, 0xc0, 0x68, 0x21, 0x70, 0x59, 0x22,
	0x4e, 0x9b, 0x1e, 0xb3, 0xbd, 0x70, 0x33, 0xa2, 0xf0, 0x11, 0x80, 0x52, 0xd3, 0x5a, 0x30, 0x84,
	0x09, 0x41, 0x98, 0x52, 0xc3, 0xaa, 0x2a, 0x74, 0x12, 0xbf, 0x7d, 0xf5, 0xe0, 0x28, 0x94, 0xe9,
	0x35, 0x91, 0xa9, 0x35, 0x5c, 0x15, 0xd9, 0xad, 0x06, 0xd3, 0x23, 0x24, 0xdc, 0xc3, 0xc7, 0xe0,
	0x96, 0xd5, 0x16, 0x0b, 0x84, 0x85, 0xd0, 0xc4, 0x2b, 0x1a, 0xb7, 0xee, 0x47, 0xa7, 0xfb, 0xdf,
	0xde, 0x4d, 0x02, 0x85, 0x93, 0x34, 0x09, 0x92, 0x26, 0x03, 0xcd, 0x55, 0xbf, 0xe5, 0x46, 0xcd,
	0x0e, 0x3d, 0xb0, 0xb7, 0xc6, 0xc1, 0x29, 0x68, 0xe7, 0x8d, 0x3c, 0x88, 0x38, 0x7d, 0xe2, 0x4f,
	0x3d, 0xd1, 0x69, 0xf2, 0x71, 0x3b, 0x26, 0x9b, 0x7a, 0x06, 0xde, 0x83, 0x7c, 0x53, 0xe3, 0x2f,
	0x40, 0xbb, 0x62, 0xd6, 0x0a, 0x46, 0x11, 0x65, 0x4a, 0xcb, 0xf8, 0x86, 0xff, 0x14, 0x07, 0xe1,
	0x72, 0xe8, 0xee, 0xe0, 0x00, 0x80, 0xca, 0x62, 0x63, 0x91, 0xb3, 0x6d, 0xbc, 0xe3, 0xcb, 0x1e,
	0x27, 0x8d, 0x65, 0x93, 0x95, 0x65, 0x93, 0x67, 0x2b, 0x4f, 0xf7, 0x77, 0x5d, 0xa1, 0x97, 0x7f,
	0x77, 0xa3, 0x6c, 0xcf, 0xe3, 0x5c, 0x04, 0x8e, 0xc1, 0x61, 0x69, 0x18, 0x12, 0xb8, 0x56, 0x64,
	0xd6, 0x30, 0xed, 0x5e, 0x81, 0xa9, 0x5d, 0x1a, 0x36, 0xf6, 0x58, 0xcf, 0xf6, 0x08, 0xec, 0x56,
	0x5a, 0x50, 0x84, 0xa5, 0x8d, 0xf7, 0xbc, 0x2c, 0x5f, 0x07, 0x83, 0xdc, 0xf9, 0xd0, 0x20, 0x23,
	0x65, 0x37, 0xac, 0x31, 0x52, 0x36, 0xdb, 0x71, 0xe0, 0x9e, 0xb4, 0x70, 0x0c, 0xf6, 0x89, 0xc0,
	0x5c, 0xb2, 0x86, 0x0a, 0x5c, 0x9d, 0x0a, 0x04, 0xbc, 0x63, 0xe3, 0xe0, 0x0e, 0x57, 0x84, 0x29,
	0xcb, 0xe7, 0x0c, 0x95, 0x02, 0x2b, 0xd4, 0x6c, 0x58, 0xbc, 0xef, 0x27, 0x4d, 0x2f, 0x93, 0x6a,
	0xb4, 0x02, 0x3a, 0xbf, 0x4e, 0x3c, 0x2c, 0x28, 0x76, 0x9b, 0x7f, 0x18, 0x82, 0xcf, 0x01, 0x94,
	0x78, 0x81, 0xb0, 0xd4, 0xb5, 0xb2, 0xc8, 0x6a, 0x54, 0x31, 0x21, 0xe2, 0x83, 0xab, 0xf7, 0x7f,
	0x28, 0xf1, 0xa2, 0xe7, 0x59, 0x9e, 0xe9, 0x29, 0x13, 0x02, 0x3e, 0x07, 0x37, 0x05, 0xff, 0xb9,
	0xe6, 0x94, 0xdb, 0xa5, 0xeb, 0xdf, 0xc6, 0xed, 0xeb, 0x6e, 0x60, 0x7b, 0x4d, 0x34, 0xc1, 0xc6,
	0x9e, 0xfc, 0x1e, 0x81, 0xdb, 0xff, 0x33, 0x26, 0xcc, 0xc1, 0xbd, 0xff, 0xfc, 0x85, 0xf0, 0x0b,
	0xcb, 0x0c, 0x6a, 0x0c, 0x28, 0x99, 0xb2, 0x7e, 0x6b, 0xdd, 0xc2, 0xbc, 0x6f, 0x93, 0x61, 0x78,
	0x23, 0x1b, 0x97, 0xfc, 0xe6, 0x5c, 0x12, 0xaf, 0xfd, 0xd6, 0x73, 0x2c, 0xd3, 0x35, 0x09, 0x4c,
	0xc1, 0x91, 0xaa, 0x25, 0x62, 0xa5, 0x26, 0xb3, 0x0a, 0x95, 0x98, 0x53, 0xa4, 0xe7, 0xcc, 0xf8,
	0xd5, 0x6f, 0x65, 0x9f, 0xa9, 0x5a, 0x7e, 0xef, 0x43, 0x13, 0xcc, 0xe9, 0x0f, 0x73, 0x66, 0xfa,
	0x8f, 0x5f, 0x9f, 0x77, 0xa2, 0x37, 0xe7, 0x9d, 0xe8, 0x9f, 0xf3, 0x4e, 0xf4, 0xf2, 0xa2, 0xb3,
	0xf5, 0xe6, 0xa2, 0xb3, 0xf5, 0xe7, 0x45, 0x67, 0xeb, 0xc7, 0x6f, 0x0a, 0x6e, 0x67, 0x75, 0x9e,
	0x10, 0x2d, 0xd3, 0x8f, 0x3c, 0xae, 0xf3, 0x87, 0xe9, 0xc2, 0xbf, 0xb0, 0x76, 0x59, 0xb2, 0x2a,
	0xbf, 0xe1, 0x7b, 0x7e, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x52, 0x49, 0x88, 0x60,
	0x06, 0x00, 0x00,
}

func (m *BondingCurve) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BondingCurve) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BondingCurve) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.C.Size()
		i -= size
		if _, err := m.C.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.N.Size()
		i -= size
		if _, err := m.N.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.M.Size()
		i -= size
		if _, err := m.M.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LiquidityPart.Size()
		i -= size
		if _, err := m.LiquidityPart.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.MaxAmountToSell.Size()
		i -= size
		if _, err := m.MaxAmountToSell.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size, err := m.IncentivePlanParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.ClaimedAmt.Size()
		i -= size
		if _, err := m.ClaimedAmt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.SoldAmt.Size()
		i -= size
		if _, err := m.SoldAmt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.PreLaunchTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.PreLaunchTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintIro(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x42
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintIro(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x3a
	if len(m.SettledDenom) > 0 {
		i -= len(m.SettledDenom)
		copy(dAtA[i:], m.SettledDenom)
		i = encodeVarintIro(dAtA, i, uint64(len(m.SettledDenom)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.BondingCurve.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.TotalAllocation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIro(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ModuleAccAddress) > 0 {
		i -= len(m.ModuleAccAddress)
		copy(dAtA[i:], m.ModuleAccAddress)
		i = encodeVarintIro(dAtA, i, uint64(len(m.ModuleAccAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintIro(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintIro(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IncentivePlanParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentivePlanParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentivePlanParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumEpochsPaidOver != 0 {
		i = encodeVarintIro(dAtA, i, uint64(m.NumEpochsPaidOver))
		i--
		dAtA[i] = 0x10
	}
	n6, err6 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.StartTimeAfterSettlement, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.StartTimeAfterSettlement):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintIro(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIro(dAtA []byte, offset int, v uint64) int {
	offset -= sovIro(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BondingCurve) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.M.Size()
	n += 1 + l + sovIro(uint64(l))
	l = m.N.Size()
	n += 1 + l + sovIro(uint64(l))
	l = m.C.Size()
	n += 1 + l + sovIro(uint64(l))
	return n
}

func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovIro(uint64(m.Id))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovIro(uint64(l))
	}
	l = len(m.ModuleAccAddress)
	if l > 0 {
		n += 1 + l + sovIro(uint64(l))
	}
	l = m.TotalAllocation.Size()
	n += 1 + l + sovIro(uint64(l))
	l = m.BondingCurve.Size()
	n += 1 + l + sovIro(uint64(l))
	l = len(m.SettledDenom)
	if l > 0 {
		n += 1 + l + sovIro(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovIro(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.PreLaunchTime)
	n += 1 + l + sovIro(uint64(l))
	l = m.SoldAmt.Size()
	n += 1 + l + sovIro(uint64(l))
	l = m.ClaimedAmt.Size()
	n += 1 + l + sovIro(uint64(l))
	l = m.IncentivePlanParams.Size()
	n += 1 + l + sovIro(uint64(l))
	l = m.MaxAmountToSell.Size()
	n += 1 + l + sovIro(uint64(l))
	l = m.LiquidityPart.Size()
	n += 1 + l + sovIro(uint64(l))
	return n
}

func (m *IncentivePlanParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.StartTimeAfterSettlement)
	n += 1 + l + sovIro(uint64(l))
	if m.NumEpochsPaidOver != 0 {
		n += 1 + sovIro(uint64(m.NumEpochsPaidOver))
	}
	return n
}

func sovIro(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIro(x uint64) (n int) {
	return sovIro(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BondingCurve) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIro
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
			return fmt.Errorf("proto: BondingCurve: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BondingCurve: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field M", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.M.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field N", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.N.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.C.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIro(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIro
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
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIro
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAccAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleAccAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAllocation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAllocation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondingCurve", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondingCurve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettledDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SettledDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreLaunchTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.PreLaunchTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoldAmt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SoldAmt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClaimedAmt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivePlanParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncentivePlanParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAmountToSell", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxAmountToSell.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityPart", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityPart.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIro(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIro
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
func (m *IncentivePlanParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIro
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
			return fmt.Errorf("proto: IncentivePlanParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentivePlanParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimeAfterSettlement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
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
				return ErrInvalidLengthIro
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIro
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.StartTimeAfterSettlement, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochsPaidOver", wireType)
			}
			m.NumEpochsPaidOver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIro
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEpochsPaidOver |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIro(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIro
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
func skipIro(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIro
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
					return 0, ErrIntOverflowIro
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
					return 0, ErrIntOverflowIro
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
				return 0, ErrInvalidLengthIro
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIro
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIro
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIro        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIro          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIro = fmt.Errorf("proto: unexpected end of group")
)
