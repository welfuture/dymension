// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/sequencer/gov_punish.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type PunishSequencerProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// sequencer address to punish
	PunishSequencerAddress string `protobuf:"bytes,3,opt,name=punish_sequencer_address,json=punishSequencerAddress,proto3" json:"punish_sequencer_address,omitempty"`
	// rewardAddr is bech32 for sdk acc addr
	Rewardee string `protobuf:"bytes,4,opt,name=rewardee,proto3" json:"rewardee,omitempty"`
}

func (m *PunishSequencerProposal) Reset()      { *m = PunishSequencerProposal{} }
func (*PunishSequencerProposal) ProtoMessage() {}
func (*PunishSequencerProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dde1d83e64e296bb, []int{0}
}
func (m *PunishSequencerProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PunishSequencerProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PunishSequencerProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PunishSequencerProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PunishSequencerProposal.Merge(m, src)
}
func (m *PunishSequencerProposal) XXX_Size() int {
	return m.Size()
}
func (m *PunishSequencerProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PunishSequencerProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PunishSequencerProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PunishSequencerProposal)(nil), "dymensionxyz.dymension.sequencer.PunishSequencerProposal")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/sequencer/gov_punish.proto", fileDescriptor_dde1d83e64e296bb)
}

var fileDescriptor_dde1d83e64e296bb = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xab, 0xa8, 0xac, 0xd2, 0x87, 0x73, 0xf4, 0x8b, 0x53, 0x0b, 0x4b,
	0x53, 0xf3, 0x92, 0x53, 0x8b, 0xf4, 0xd3, 0xf3, 0xcb, 0xe2, 0x0b, 0x4a, 0xf3, 0x32, 0x8b, 0x33,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x14, 0x90, 0xb5, 0xe8, 0xc1, 0x39, 0x7a, 0x70, 0x2d,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xc5, 0xfa, 0x20, 0x16, 0x44, 0x9f, 0xd2, 0x56, 0x46,
	0x2e, 0xf1, 0x00, 0xb0, 0x41, 0xc1, 0x30, 0x95, 0x01, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0x89, 0x39,
	0x42, 0x22, 0x5c, 0xac, 0x25, 0x99, 0x25, 0x39, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x10, 0x8e, 0x90, 0x02, 0x17, 0x77, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e,
	0x9e, 0x04, 0x13, 0x58, 0x0e, 0x59, 0x48, 0xc8, 0x82, 0x4b, 0x02, 0xe2, 0xb6, 0x78, 0xb8, 0xed,
	0xf1, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0xcc, 0x60, 0xe5, 0x62, 0x05, 0xa8, 0x56,
	0x3a, 0x42, 0x64, 0x85, 0xa4, 0xb8, 0x38, 0x8a, 0x52, 0xcb, 0x13, 0x8b, 0x52, 0x52, 0x53, 0x25,
	0x58, 0xc0, 0x2a, 0xe1, 0x7c, 0x2b, 0x9e, 0x8e, 0x05, 0xf2, 0x0c, 0x33, 0x16, 0xc8, 0x33, 0xbc,
	0x58, 0x20, 0xcf, 0xe8, 0x14, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0x66, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x38, 0xc2, 0xb1, 0xcc,
	0x58, 0xbf, 0x02, 0x29, 0x30, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x01, 0x62, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x80, 0xf2, 0xdd, 0x8f, 0x7d, 0x01, 0x00, 0x00,
}

func (this *PunishSequencerProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PunishSequencerProposal)
	if !ok {
		that2, ok := that.(PunishSequencerProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.PunishSequencerAddress != that1.PunishSequencerAddress {
		return false
	}
	if this.Rewardee != that1.Rewardee {
		return false
	}
	return true
}
func (m *PunishSequencerProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PunishSequencerProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PunishSequencerProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewardee) > 0 {
		i -= len(m.Rewardee)
		copy(dAtA[i:], m.Rewardee)
		i = encodeVarintGovPunish(dAtA, i, uint64(len(m.Rewardee)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PunishSequencerAddress) > 0 {
		i -= len(m.PunishSequencerAddress)
		copy(dAtA[i:], m.PunishSequencerAddress)
		i = encodeVarintGovPunish(dAtA, i, uint64(len(m.PunishSequencerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGovPunish(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGovPunish(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGovPunish(dAtA []byte, offset int, v uint64) int {
	offset -= sovGovPunish(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PunishSequencerProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGovPunish(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGovPunish(uint64(l))
	}
	l = len(m.PunishSequencerAddress)
	if l > 0 {
		n += 1 + l + sovGovPunish(uint64(l))
	}
	l = len(m.Rewardee)
	if l > 0 {
		n += 1 + l + sovGovPunish(uint64(l))
	}
	return n
}

func sovGovPunish(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGovPunish(x uint64) (n int) {
	return sovGovPunish(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PunishSequencerProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGovPunish
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
			return fmt.Errorf("proto: PunishSequencerProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PunishSequencerProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovPunish
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
				return ErrInvalidLengthGovPunish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovPunish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovPunish
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
				return ErrInvalidLengthGovPunish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovPunish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PunishSequencerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovPunish
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
				return ErrInvalidLengthGovPunish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovPunish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PunishSequencerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewardee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGovPunish
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
				return ErrInvalidLengthGovPunish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGovPunish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewardee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGovPunish(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGovPunish
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
func skipGovPunish(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGovPunish
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
					return 0, ErrIntOverflowGovPunish
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
					return 0, ErrIntOverflowGovPunish
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
				return 0, ErrInvalidLengthGovPunish
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGovPunish
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGovPunish
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGovPunish        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGovPunish          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGovPunish = fmt.Errorf("proto: unexpected end of group")
)
