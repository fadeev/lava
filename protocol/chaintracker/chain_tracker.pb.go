// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chainTracker.proto

package chaintracker

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type LatestBlockData struct {
	FromBlock     int64 `protobuf:"varint,1,opt,name=fromBlock,proto3" json:"fromBlock,omitempty"`
	ToBlock       int64 `protobuf:"varint,2,opt,name=toBlock,proto3" json:"toBlock,omitempty"`
	SpecificBlock int64 `protobuf:"varint,3,opt,name=specificBlock,proto3" json:"specificBlock,omitempty"`
}

func (m *LatestBlockData) Reset()         { *m = LatestBlockData{} }
func (m *LatestBlockData) String() string { return proto.CompactTextString(m) }
func (*LatestBlockData) ProtoMessage()    {}
func (*LatestBlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f7d15fc8a35cee, []int{0}
}
func (m *LatestBlockData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LatestBlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LatestBlockData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LatestBlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestBlockData.Merge(m, src)
}
func (m *LatestBlockData) XXX_Size() int {
	return m.Size()
}
func (m *LatestBlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestBlockData.DiscardUnknown(m)
}

var xxx_messageInfo_LatestBlockData proto.InternalMessageInfo

func (m *LatestBlockData) GetFromBlock() int64 {
	if m != nil {
		return m.FromBlock
	}
	return 0
}

func (m *LatestBlockData) GetToBlock() int64 {
	if m != nil {
		return m.ToBlock
	}
	return 0
}

func (m *LatestBlockData) GetSpecificBlock() int64 {
	if m != nil {
		return m.SpecificBlock
	}
	return 0
}

type LatestBlockDataResponse struct {
	LatestBlock     int64         `protobuf:"varint,1,opt,name=latestBlock,proto3" json:"latestBlock,omitempty"`
	RequestedHashes []*BlockStore `protobuf:"bytes,2,rep,name=requestedHashes,proto3" json:"requestedHashes,omitempty"`
}

func (m *LatestBlockDataResponse) Reset()         { *m = LatestBlockDataResponse{} }
func (m *LatestBlockDataResponse) String() string { return proto.CompactTextString(m) }
func (*LatestBlockDataResponse) ProtoMessage()    {}
func (*LatestBlockDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f7d15fc8a35cee, []int{1}
}
func (m *LatestBlockDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LatestBlockDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LatestBlockDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LatestBlockDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestBlockDataResponse.Merge(m, src)
}
func (m *LatestBlockDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *LatestBlockDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestBlockDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LatestBlockDataResponse proto.InternalMessageInfo

func (m *LatestBlockDataResponse) GetLatestBlock() int64 {
	if m != nil {
		return m.LatestBlock
	}
	return 0
}

func (m *LatestBlockDataResponse) GetRequestedHashes() []*BlockStore {
	if m != nil {
		return m.RequestedHashes
	}
	return nil
}

type BlockStore struct {
	Block int64  `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	Hash  string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *BlockStore) Reset()         { *m = BlockStore{} }
func (m *BlockStore) String() string { return proto.CompactTextString(m) }
func (*BlockStore) ProtoMessage()    {}
func (*BlockStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_90f7d15fc8a35cee, []int{2}
}
func (m *BlockStore) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockStore.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStore.Merge(m, src)
}
func (m *BlockStore) XXX_Size() int {
	return m.Size()
}
func (m *BlockStore) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStore.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStore proto.InternalMessageInfo

func (m *BlockStore) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *BlockStore) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*LatestBlockData)(nil), "chainTracker.LatestBlockData")
	proto.RegisterType((*LatestBlockDataResponse)(nil), "chainTracker.LatestBlockDataResponse")
	proto.RegisterType((*BlockStore)(nil), "chainTracker.BlockStore")
}

func init() { proto.RegisterFile("chainTracker.proto", fileDescriptor_90f7d15fc8a35cee) }

var fileDescriptor_90f7d15fc8a35cee = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x5b, 0xf0, 0x4f, 0x18, 0x54, 0xe2, 0x6a, 0xb4, 0xa9, 0xd8, 0x90, 0x46, 0x13, 0x4e,
	0x25, 0x41, 0xc3, 0x03, 0xa0, 0x46, 0x8d, 0xc6, 0x43, 0x51, 0x0f, 0xc6, 0xcb, 0x52, 0x07, 0xda,
	0x50, 0xd8, 0xba, 0xbb, 0xd5, 0x78, 0xf2, 0x15, 0x7c, 0x24, 0x8f, 0x1e, 0x39, 0x7a, 0x34, 0xf0,
	0x22, 0xc6, 0xad, 0x84, 0xb6, 0x26, 0xde, 0x76, 0xbe, 0xdf, 0xec, 0x7c, 0x3b, 0xdf, 0x02, 0xf1,
	0x7c, 0x1a, 0x8c, 0xae, 0x39, 0xf5, 0x06, 0xc8, 0x9d, 0x88, 0x33, 0xc9, 0xc8, 0x4a, 0x5a, 0x33,
	0xad, 0x3e, 0x63, 0xfd, 0x10, 0x1b, 0x8a, 0x75, 0xe3, 0x5e, 0xe3, 0x99, 0xd3, 0x28, 0x42, 0x2e,
	0x92, 0x6e, 0x73, 0x27, 0xcf, 0x71, 0x18, 0xc9, 0x97, 0x04, 0xda, 0x0c, 0x2a, 0x97, 0x54, 0xa2,
	0x90, 0xed, 0x90, 0x79, 0x83, 0x63, 0x2a, 0x29, 0xa9, 0x42, 0xa9, 0xc7, 0xd9, 0x50, 0x09, 0x86,
	0x5e, 0xd3, 0xeb, 0x45, 0x77, 0x2e, 0x10, 0x03, 0x96, 0x25, 0x4b, 0x58, 0x41, 0xb1, 0x59, 0x49,
	0xf6, 0x60, 0x55, 0x44, 0xe8, 0x05, 0xbd, 0xc0, 0x4b, 0x78, 0x51, 0xf1, 0xac, 0x68, 0xbf, 0xc2,
	0x76, 0xce, 0xd0, 0x45, 0x11, 0xb1, 0x91, 0x40, 0x52, 0x83, 0x72, 0x38, 0x47, 0xbf, 0xd6, 0x69,
	0x89, 0xb4, 0xa1, 0xc2, 0xf1, 0x31, 0x46, 0x21, 0xf1, 0xe1, 0x8c, 0x0a, 0x1f, 0x85, 0x51, 0xa8,
	0x15, 0xeb, 0xe5, 0xa6, 0xe1, 0x64, 0x62, 0x52, 0xdd, 0x1d, 0xc9, 0x38, 0xba, 0xf9, 0x0b, 0x76,
	0x0b, 0x60, 0x8e, 0xc9, 0x26, 0x2c, 0x76, 0x53, 0x6e, 0x49, 0x41, 0x08, 0x2c, 0xf8, 0x54, 0xf8,
	0x6a, 0xc3, 0x92, 0xab, 0xce, 0xcd, 0x77, 0x1d, 0x36, 0x8e, 0x52, 0x26, 0x1d, 0xe4, 0x4f, 0x81,
	0x87, 0xe4, 0x02, 0xd6, 0x4f, 0x51, 0xa6, 0x76, 0xba, 0x8a, 0x87, 0x64, 0xcb, 0x49, 0x42, 0x77,
	0x66, 0xa1, 0x3b, 0x27, 0x3f, 0xa1, 0x9b, 0xd5, 0x3f, 0xfa, 0xcd, 0xf9, 0x48, 0xb6, 0x0e, 0x6f,
	0x69, 0x18, 0xa3, 0xad, 0x91, 0x7b, 0x20, 0xd9, 0x61, 0xea, 0x47, 0x76, 0xb3, 0xdb, 0xe5, 0xb0,
	0xb9, 0xff, 0x2f, 0x9e, 0xc5, 0x6b, 0x6b, 0xed, 0xfa, 0xc7, 0xc4, 0xd2, 0xc7, 0x13, 0x4b, 0xff,
	0x9a, 0x58, 0xfa, 0xdb, 0xd4, 0xd2, 0xc6, 0x53, 0x4b, 0xfb, 0x9c, 0x5a, 0xda, 0xdd, 0x9a, 0xd3,
	0x50, 0x33, 0x64, 0x32, 0xa3, 0xbb, 0xa4, 0xde, 0x77, 0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x72,
	0xe1, 0x3c, 0xae, 0x7e, 0x02, 0x00, 0x00,
}

func (m *LatestBlockData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LatestBlockData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LatestBlockData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SpecificBlock != 0 {
		i = encodeVarintChainTracker(dAtA, i, uint64(m.SpecificBlock))
		i--
		dAtA[i] = 0x18
	}
	if m.ToBlock != 0 {
		i = encodeVarintChainTracker(dAtA, i, uint64(m.ToBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.FromBlock != 0 {
		i = encodeVarintChainTracker(dAtA, i, uint64(m.FromBlock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LatestBlockDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LatestBlockDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LatestBlockDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequestedHashes) > 0 {
		for iNdEx := len(m.RequestedHashes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestedHashes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintChainTracker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.LatestBlock != 0 {
		i = encodeVarintChainTracker(dAtA, i, uint64(m.LatestBlock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockStore) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockStore) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockStore) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintChainTracker(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Block != 0 {
		i = encodeVarintChainTracker(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChainTracker(dAtA []byte, offset int, v uint64) int {
	offset -= sovChainTracker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LatestBlockData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FromBlock != 0 {
		n += 1 + sovChainTracker(uint64(m.FromBlock))
	}
	if m.ToBlock != 0 {
		n += 1 + sovChainTracker(uint64(m.ToBlock))
	}
	if m.SpecificBlock != 0 {
		n += 1 + sovChainTracker(uint64(m.SpecificBlock))
	}
	return n
}

func (m *LatestBlockDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LatestBlock != 0 {
		n += 1 + sovChainTracker(uint64(m.LatestBlock))
	}
	if len(m.RequestedHashes) > 0 {
		for _, e := range m.RequestedHashes {
			l = e.Size()
			n += 1 + l + sovChainTracker(uint64(l))
		}
	}
	return n
}

func (m *BlockStore) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovChainTracker(uint64(m.Block))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovChainTracker(uint64(l))
	}
	return n
}

func sovChainTracker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChainTracker(x uint64) (n int) {
	return sovChainTracker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LatestBlockData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainTracker
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
			return fmt.Errorf("proto: LatestBlockData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LatestBlockData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromBlock", wireType)
			}
			m.FromBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToBlock", wireType)
			}
			m.ToBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpecificBlock", wireType)
			}
			m.SpecificBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpecificBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChainTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainTracker
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
func (m *LatestBlockDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainTracker
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
			return fmt.Errorf("proto: LatestBlockDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LatestBlockDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlock", wireType)
			}
			m.LatestBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestedHashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainTracker
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
				return ErrInvalidLengthChainTracker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChainTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestedHashes = append(m.RequestedHashes, &BlockStore{})
			if err := m.RequestedHashes[len(m.RequestedHashes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChainTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainTracker
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
func (m *BlockStore) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainTracker
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
			return fmt.Errorf("proto: BlockStore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockStore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainTracker
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
				return ErrInvalidLengthChainTracker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChainTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainTracker
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
func skipChainTracker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChainTracker
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
					return 0, ErrIntOverflowChainTracker
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
					return 0, ErrIntOverflowChainTracker
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
				return 0, ErrInvalidLengthChainTracker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChainTracker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChainTracker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChainTracker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChainTracker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChainTracker = fmt.Errorf("proto: unexpected end of group")
)
