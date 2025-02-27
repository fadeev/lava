// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subscription/subscription.proto

package types

import (
	fmt "fmt"
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

type Subscription struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Consumer        string `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Block           uint64 `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	PlanIndex       string `protobuf:"bytes,4,opt,name=plan_index,json=planIndex,proto3" json:"plan_index,omitempty"`
	PlanBlock       uint64 `protobuf:"varint,5,opt,name=plan_block,json=planBlock,proto3" json:"plan_block,omitempty"`
	DurationTotal   uint64 `protobuf:"varint,6,opt,name=duration_total,json=durationTotal,proto3" json:"duration_total,omitempty"`
	DurationLeft    uint64 `protobuf:"varint,7,opt,name=duration_left,json=durationLeft,proto3" json:"duration_left,omitempty"`
	MonthExpiryTime uint64 `protobuf:"varint,8,opt,name=month_expiry_time,json=monthExpiryTime,proto3" json:"month_expiry_time,omitempty"`
	PrevExpiryBlock uint64 `protobuf:"varint,9,opt,name=prev_expiry_block,json=prevExpiryBlock,proto3" json:"prev_expiry_block,omitempty"`
	MonthCuTotal    uint64 `protobuf:"varint,10,opt,name=month_cu_total,json=monthCuTotal,proto3" json:"month_cu_total,omitempty"`
	MonthCuLeft     uint64 `protobuf:"varint,11,opt,name=month_cu_left,json=monthCuLeft,proto3" json:"month_cu_left,omitempty"`
	PrevCuLeft      uint64 `protobuf:"varint,12,opt,name=prev_cu_left,json=prevCuLeft,proto3" json:"prev_cu_left,omitempty"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac47bc0f89224537, []int{0}
}
func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return m.Size()
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Subscription) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *Subscription) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Subscription) GetPlanIndex() string {
	if m != nil {
		return m.PlanIndex
	}
	return ""
}

func (m *Subscription) GetPlanBlock() uint64 {
	if m != nil {
		return m.PlanBlock
	}
	return 0
}

func (m *Subscription) GetDurationTotal() uint64 {
	if m != nil {
		return m.DurationTotal
	}
	return 0
}

func (m *Subscription) GetDurationLeft() uint64 {
	if m != nil {
		return m.DurationLeft
	}
	return 0
}

func (m *Subscription) GetMonthExpiryTime() uint64 {
	if m != nil {
		return m.MonthExpiryTime
	}
	return 0
}

func (m *Subscription) GetPrevExpiryBlock() uint64 {
	if m != nil {
		return m.PrevExpiryBlock
	}
	return 0
}

func (m *Subscription) GetMonthCuTotal() uint64 {
	if m != nil {
		return m.MonthCuTotal
	}
	return 0
}

func (m *Subscription) GetMonthCuLeft() uint64 {
	if m != nil {
		return m.MonthCuLeft
	}
	return 0
}

func (m *Subscription) GetPrevCuLeft() uint64 {
	if m != nil {
		return m.PrevCuLeft
	}
	return 0
}

func init() {
	proto.RegisterType((*Subscription)(nil), "lavanet.lava.subscription.Subscription")
}

func init() { proto.RegisterFile("subscription/subscription.proto", fileDescriptor_ac47bc0f89224537) }

var fileDescriptor_ac47bc0f89224537 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x9b, 0xdb, 0xff, 0xa7, 0x69, 0x2f, 0x77, 0xb8, 0x8b, 0x51, 0x30, 0x96, 0xaa, 0x50,
	0x44, 0x92, 0x85, 0x6f, 0x50, 0x51, 0x10, 0x5c, 0xd5, 0xae, 0xdc, 0x84, 0x24, 0x9d, 0xda, 0xc1,
	0x24, 0x13, 0x26, 0x93, 0xd2, 0xbe, 0x85, 0x8f, 0xe1, 0xa3, 0xb8, 0xec, 0xd2, 0xa5, 0xb4, 0x2f,
	0x22, 0x73, 0xa6, 0x09, 0x75, 0x35, 0x9c, 0xdf, 0xf7, 0x3b, 0x99, 0x0f, 0x32, 0x70, 0x9e, 0x17,
	0x61, 0x1e, 0x49, 0x9e, 0x29, 0x2e, 0x52, 0xef, 0x78, 0x70, 0x33, 0x29, 0x94, 0x20, 0x27, 0x71,
	0xb0, 0x0a, 0x52, 0xa6, 0x5c, 0x7d, 0xba, 0xc7, 0xc2, 0xe8, 0xa3, 0x0e, 0xf6, 0xf3, 0x11, 0x20,
	0x14, 0xda, 0x91, 0x64, 0x81, 0x12, 0x92, 0x5a, 0x43, 0x6b, 0xdc, 0x9d, 0x96, 0x23, 0x39, 0x85,
	0x4e, 0x24, 0xd2, 0xbc, 0x48, 0x98, 0xa4, 0x7f, 0x30, 0xaa, 0x66, 0xf2, 0x1f, 0x9a, 0x61, 0x2c,
	0xa2, 0x37, 0x5a, 0x1f, 0x5a, 0xe3, 0xc6, 0xd4, 0x0c, 0xe4, 0x0c, 0x20, 0x8b, 0x83, 0xd4, 0xe7,
	0xe9, 0x9c, 0xad, 0x69, 0x03, 0x77, 0xba, 0x9a, 0x3c, 0x6a, 0x50, 0xc5, 0x66, 0xb3, 0x89, 0x9b,
	0x18, 0x4f, 0x70, 0xfb, 0x0a, 0x06, 0xf3, 0x42, 0x06, 0xba, 0x95, 0xaf, 0x84, 0x0a, 0x62, 0xda,
	0x42, 0xa5, 0x5f, 0xd2, 0x99, 0x86, 0xe4, 0x02, 0x2a, 0xe0, 0xc7, 0x6c, 0xa1, 0x68, 0x1b, 0x2d,
	0xbb, 0x84, 0x4f, 0x6c, 0xa1, 0xc8, 0x35, 0xfc, 0x4b, 0x44, 0xaa, 0x96, 0x3e, 0x5b, 0x67, 0x5c,
	0x6e, 0x7c, 0xc5, 0x13, 0x46, 0x3b, 0x28, 0xfe, 0xc5, 0xe0, 0x1e, 0xf9, 0x8c, 0x27, 0x4c, 0xbb,
	0x99, 0x64, 0xab, 0x52, 0x35, 0xed, 0xba, 0xc6, 0xd5, 0x81, 0x51, 0x4d, 0xc7, 0x4b, 0x18, 0x98,
	0xef, 0x46, 0xc5, 0xa1, 0x23, 0x98, 0xdb, 0x91, 0xde, 0x15, 0xa6, 0xe2, 0x08, 0xfa, 0x95, 0x85,
	0x15, 0x7b, 0x28, 0xf5, 0x0e, 0x12, 0x36, 0x1c, 0x82, 0x8d, 0xb7, 0x96, 0x8a, 0x8d, 0x0a, 0x68,
	0x66, 0x8c, 0xc9, 0xc3, 0xe7, 0xce, 0xb1, 0xb6, 0x3b, 0xc7, 0xfa, 0xde, 0x39, 0xd6, 0xfb, 0xde,
	0xa9, 0x6d, 0xf7, 0x4e, 0xed, 0x6b, 0xef, 0xd4, 0x5e, 0x6e, 0x5e, 0xb9, 0x5a, 0x16, 0xa1, 0x1b,
	0x89, 0xc4, 0x3b, 0xfc, 0x6a, 0x3c, 0xbd, 0xf5, 0xaf, 0xd7, 0xe0, 0xa9, 0x4d, 0xc6, 0xf2, 0xb0,
	0x85, 0x8f, 0xe2, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x6e, 0x85, 0xd9, 0x37, 0x02, 0x00,
	0x00,
}

func (m *Subscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PrevCuLeft != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PrevCuLeft))
		i--
		dAtA[i] = 0x60
	}
	if m.MonthCuLeft != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.MonthCuLeft))
		i--
		dAtA[i] = 0x58
	}
	if m.MonthCuTotal != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.MonthCuTotal))
		i--
		dAtA[i] = 0x50
	}
	if m.PrevExpiryBlock != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PrevExpiryBlock))
		i--
		dAtA[i] = 0x48
	}
	if m.MonthExpiryTime != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.MonthExpiryTime))
		i--
		dAtA[i] = 0x40
	}
	if m.DurationLeft != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.DurationLeft))
		i--
		dAtA[i] = 0x38
	}
	if m.DurationTotal != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.DurationTotal))
		i--
		dAtA[i] = 0x30
	}
	if m.PlanBlock != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.PlanBlock))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PlanIndex) > 0 {
		i -= len(m.PlanIndex)
		copy(dAtA[i:], m.PlanIndex)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.PlanIndex)))
		i--
		dAtA[i] = 0x22
	}
	if m.Block != 0 {
		i = encodeVarintSubscription(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSubscription(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubscription(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubscription(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Subscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovSubscription(uint64(m.Block))
	}
	l = len(m.PlanIndex)
	if l > 0 {
		n += 1 + l + sovSubscription(uint64(l))
	}
	if m.PlanBlock != 0 {
		n += 1 + sovSubscription(uint64(m.PlanBlock))
	}
	if m.DurationTotal != 0 {
		n += 1 + sovSubscription(uint64(m.DurationTotal))
	}
	if m.DurationLeft != 0 {
		n += 1 + sovSubscription(uint64(m.DurationLeft))
	}
	if m.MonthExpiryTime != 0 {
		n += 1 + sovSubscription(uint64(m.MonthExpiryTime))
	}
	if m.PrevExpiryBlock != 0 {
		n += 1 + sovSubscription(uint64(m.PrevExpiryBlock))
	}
	if m.MonthCuTotal != 0 {
		n += 1 + sovSubscription(uint64(m.MonthCuTotal))
	}
	if m.MonthCuLeft != 0 {
		n += 1 + sovSubscription(uint64(m.MonthCuLeft))
	}
	if m.PrevCuLeft != 0 {
		n += 1 + sovSubscription(uint64(m.PrevCuLeft))
	}
	return n
}

func sovSubscription(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubscription(x uint64) (n int) {
	return sovSubscription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Subscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubscription
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
			return fmt.Errorf("proto: Subscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
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
				return ErrInvalidLengthSubscription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubscription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlanIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanBlock", wireType)
			}
			m.PlanBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationTotal", wireType)
			}
			m.DurationTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationTotal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationLeft", wireType)
			}
			m.DurationLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthExpiryTime", wireType)
			}
			m.MonthExpiryTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MonthExpiryTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevExpiryBlock", wireType)
			}
			m.PrevExpiryBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevExpiryBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthCuTotal", wireType)
			}
			m.MonthCuTotal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MonthCuTotal |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthCuLeft", wireType)
			}
			m.MonthCuLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MonthCuLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevCuLeft", wireType)
			}
			m.PrevCuLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubscription
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevCuLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSubscription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubscription
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
func skipSubscription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
					return 0, ErrIntOverflowSubscription
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
				return 0, ErrInvalidLengthSubscription
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubscription
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubscription
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubscription        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubscription          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubscription = fmt.Errorf("proto: unexpected end of group")
)
