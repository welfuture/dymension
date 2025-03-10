// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/lightclient/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// verify a client state and its consensus states against the rollapp
// if it matches, set the client as the canonical client for the rollapp
// NOTE: this definition is also copied to the relayer code
type MsgSetCanonicalClient struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// id of ibc client state
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *MsgSetCanonicalClient) Reset()         { *m = MsgSetCanonicalClient{} }
func (m *MsgSetCanonicalClient) String() string { return proto.CompactTextString(m) }
func (*MsgSetCanonicalClient) ProtoMessage()    {}
func (*MsgSetCanonicalClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_86d8e0adcad4a570, []int{0}
}
func (m *MsgSetCanonicalClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetCanonicalClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetCanonicalClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetCanonicalClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetCanonicalClient.Merge(m, src)
}
func (m *MsgSetCanonicalClient) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetCanonicalClient) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetCanonicalClient.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetCanonicalClient proto.InternalMessageInfo

func (m *MsgSetCanonicalClient) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgSetCanonicalClient) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type MsgSetCanonicalClientResponse struct {
}

func (m *MsgSetCanonicalClientResponse) Reset()         { *m = MsgSetCanonicalClientResponse{} }
func (m *MsgSetCanonicalClientResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetCanonicalClientResponse) ProtoMessage()    {}
func (*MsgSetCanonicalClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86d8e0adcad4a570, []int{1}
}
func (m *MsgSetCanonicalClientResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetCanonicalClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetCanonicalClientResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetCanonicalClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetCanonicalClientResponse.Merge(m, src)
}
func (m *MsgSetCanonicalClientResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetCanonicalClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetCanonicalClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetCanonicalClientResponse proto.InternalMessageInfo

type MsgUpdateClient struct {
	Inner *types.MsgUpdateClient `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (m *MsgUpdateClient) Reset()         { *m = MsgUpdateClient{} }
func (m *MsgUpdateClient) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateClient) ProtoMessage()    {}
func (*MsgUpdateClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_86d8e0adcad4a570, []int{2}
}
func (m *MsgUpdateClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateClient.Merge(m, src)
}
func (m *MsgUpdateClient) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateClient) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateClient.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateClient proto.InternalMessageInfo

func (m *MsgUpdateClient) GetInner() *types.MsgUpdateClient {
	if m != nil {
		return m.Inner
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgSetCanonicalClient)(nil), "dymensionxyz.dymension.lightclient.MsgSetCanonicalClient")
	proto.RegisterType((*MsgSetCanonicalClientResponse)(nil), "dymensionxyz.dymension.lightclient.MsgSetCanonicalClientResponse")
	proto.RegisterType((*MsgUpdateClient)(nil), "dymensionxyz.dymension.lightclient.MsgUpdateClient")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/lightclient/tx.proto", fileDescriptor_86d8e0adcad4a570)
}

var fileDescriptor_86d8e0adcad4a570 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4e, 0xea, 0x40,
	0x18, 0xa5, 0xdc, 0x40, 0x2e, 0xc3, 0x4d, 0x6e, 0xd2, 0xdc, 0x1f, 0x6e, 0xb9, 0x56, 0x53, 0x37,
	0x06, 0x92, 0x99, 0x00, 0x1b, 0x71, 0xa7, 0xac, 0x4c, 0x64, 0x53, 0xe3, 0x42, 0x37, 0xa6, 0x3f,
	0xe3, 0x30, 0x49, 0x3b, 0xd3, 0x30, 0x43, 0x43, 0x5d, 0x19, 0x9f, 0xc0, 0xe8, 0x8b, 0xf0, 0x18,
	0x2e, 0x59, 0xba, 0x34, 0xb0, 0xe0, 0x35, 0x0c, 0xed, 0x40, 0xc0, 0x60, 0x24, 0xae, 0x3a, 0xdf,
	0x9c, 0x73, 0xbe, 0x73, 0xbe, 0xce, 0x07, 0xea, 0x7e, 0x12, 0x62, 0x26, 0x28, 0x67, 0xc3, 0xe4,
	0x16, 0x2d, 0x0b, 0x14, 0x50, 0xd2, 0x93, 0x5e, 0x40, 0x31, 0x93, 0x48, 0x0e, 0x61, 0xd4, 0xe7,
	0x92, 0xeb, 0xd6, 0x2a, 0x19, 0x2e, 0x0b, 0xb8, 0x42, 0x36, 0xfe, 0x7a, 0x5c, 0x84, 0x5c, 0xa0,
	0x50, 0x10, 0x14, 0x37, 0xe6, 0x9f, 0x4c, 0x6c, 0xfc, 0x22, 0x9c, 0xf0, 0xf4, 0x88, 0xe6, 0x27,
	0x75, 0xfb, 0x9f, 0x70, 0x4e, 0x02, 0x8c, 0x9c, 0x88, 0x22, 0x87, 0x31, 0x2e, 0x1d, 0x49, 0x39,
	0x13, 0x0a, 0xfd, 0xa7, 0xd0, 0xb4, 0x72, 0x07, 0x37, 0xc8, 0x61, 0x89, 0x82, 0xaa, 0xd4, 0xf5,
	0x90, 0xc7, 0xfb, 0x18, 0xa9, 0x94, 0x71, 0x63, 0x19, 0xd4, 0xba, 0x04, 0xbf, 0xbb, 0x82, 0x9c,
	0x63, 0xd9, 0x71, 0x18, 0x67, 0xd4, 0x73, 0x82, 0x4e, 0x4a, 0xd2, 0xff, 0x80, 0xa2, 0xa0, 0x84,
	0xe1, 0x7e, 0x45, 0xdb, 0xd3, 0x0e, 0x4a, 0xb6, 0xaa, 0xf4, 0x2a, 0x28, 0x65, 0x6d, 0xae, 0xa9,
	0x5f, 0xc9, 0xa7, 0xd0, 0xf7, 0xec, 0xe2, 0xd4, 0x3f, 0x2a, 0xdf, 0xcf, 0x46, 0x35, 0xc5, 0xb4,
	0x76, 0xc1, 0xce, 0xc6, 0xd6, 0x36, 0x16, 0x11, 0x67, 0x02, 0x5b, 0x67, 0xe0, 0x67, 0x57, 0x90,
	0x8b, 0xc8, 0x77, 0x24, 0x56, 0xae, 0x6d, 0x50, 0xa0, 0x6c, 0x61, 0x5a, 0x6e, 0xee, 0x43, 0xea,
	0x7a, 0x70, 0x9e, 0x1d, 0x66, 0x1e, 0x30, 0x6e, 0xc0, 0x77, 0x1a, 0x3b, 0x53, 0x34, 0x9f, 0xf2,
	0xe0, 0x5b, 0x57, 0x10, 0xfd, 0x51, 0x03, 0xfa, 0x86, 0x79, 0xda, 0xf0, 0xf3, 0x27, 0x81, 0x1b,
	0xf3, 0x1a, 0xc7, 0x5f, 0x96, 0x2e, 0x46, 0xd5, 0x23, 0xf0, 0x63, 0x6d, 0xce, 0xd6, 0x96, 0x2d,
	0x57, 0x45, 0x46, 0x7d, 0x9b, 0xbf, 0xa1, 0x1c, 0x8d, 0xc2, 0xdd, 0x6c, 0x54, 0xd3, 0x4e, 0xec,
	0xe7, 0x89, 0xa9, 0x8d, 0x27, 0xa6, 0xf6, 0x3a, 0x31, 0xb5, 0x87, 0xa9, 0x99, 0x1b, 0x4f, 0xcd,
	0xdc, 0xcb, 0xd4, 0xcc, 0x5d, 0x1d, 0x12, 0x2a, 0x7b, 0x03, 0x17, 0x7a, 0x3c, 0x44, 0x1f, 0xac,
	0x76, 0xdc, 0x42, 0xc3, 0xf5, 0xfd, 0x4e, 0x22, 0x2c, 0xdc, 0x62, 0xba, 0x3a, 0xad, 0xb7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8d, 0x8d, 0xe5, 0x8d, 0x12, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SetCanonicalClient(ctx context.Context, in *MsgSetCanonicalClient, opts ...grpc.CallOption) (*MsgSetCanonicalClientResponse, error)
	// The normal IBC update client msg needs to go through our ante handler,
	// which means we block authz wrapping of it, but we still need a way to do it
	// via authz (for portal self-relaying). So, we add a new wrapper message
	// here, and route it through the ante handler logic explicitly.
	//
	// note: cannot reuse already-registered ibc request msg type
	UpdateClient(ctx context.Context, in *MsgUpdateClient, opts ...grpc.CallOption) (*types.MsgUpdateClientResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetCanonicalClient(ctx context.Context, in *MsgSetCanonicalClient, opts ...grpc.CallOption) (*MsgSetCanonicalClientResponse, error) {
	out := new(MsgSetCanonicalClientResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.lightclient.Msg/SetCanonicalClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateClient(ctx context.Context, in *MsgUpdateClient, opts ...grpc.CallOption) (*types.MsgUpdateClientResponse, error) {
	out := new(types.MsgUpdateClientResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.lightclient.Msg/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SetCanonicalClient(context.Context, *MsgSetCanonicalClient) (*MsgSetCanonicalClientResponse, error)
	// The normal IBC update client msg needs to go through our ante handler,
	// which means we block authz wrapping of it, but we still need a way to do it
	// via authz (for portal self-relaying). So, we add a new wrapper message
	// here, and route it through the ante handler logic explicitly.
	//
	// note: cannot reuse already-registered ibc request msg type
	UpdateClient(context.Context, *MsgUpdateClient) (*types.MsgUpdateClientResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetCanonicalClient(ctx context.Context, req *MsgSetCanonicalClient) (*MsgSetCanonicalClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCanonicalClient not implemented")
}
func (*UnimplementedMsgServer) UpdateClient(ctx context.Context, req *MsgUpdateClient) (*types.MsgUpdateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetCanonicalClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetCanonicalClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetCanonicalClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.lightclient.Msg/SetCanonicalClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetCanonicalClient(ctx, req.(*MsgSetCanonicalClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.lightclient.Msg/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateClient(ctx, req.(*MsgUpdateClient))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.lightclient.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetCanonicalClient",
			Handler:    _Msg_SetCanonicalClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _Msg_UpdateClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dymensionxyz/dymension/lightclient/tx.proto",
}

func (m *MsgSetCanonicalClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetCanonicalClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetCanonicalClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetCanonicalClientResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetCanonicalClientResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetCanonicalClientResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Inner != nil {
		{
			size, err := m.Inner.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetCanonicalClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetCanonicalClientResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Inner != nil {
		l = m.Inner.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetCanonicalClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetCanonicalClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetCanonicalClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetCanonicalClientResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetCanonicalClientResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetCanonicalClientResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Inner == nil {
				m.Inner = &types.MsgUpdateClient{}
			}
			if err := m.Inner.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
