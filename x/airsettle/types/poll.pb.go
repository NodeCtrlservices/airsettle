// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: airsettle/airsettle/poll.proto

package types

import (
	fmt "fmt"
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

type Poll struct {
	PollId          string   `protobuf:"bytes,1,opt,name=pollId,proto3" json:"pollId,omitempty"`
	ChainId         string   `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	NewValidator    string   `protobuf:"bytes,3,opt,name=newValidator,proto3" json:"newValidator,omitempty"`
	VotesDoneBy     []string `protobuf:"bytes,4,rep,name=votesDoneBy,proto3" json:"votesDoneBy,omitempty"`
	Votes           []string `protobuf:"bytes,5,rep,name=votes,proto3" json:"votes,omitempty"`
	TotalValidators uint64   `protobuf:"varint,6,opt,name=totalValidators,proto3" json:"totalValidators,omitempty"`
	IsComplete      bool     `protobuf:"varint,7,opt,name=isComplete,proto3" json:"isComplete,omitempty"`
	StartDate       string   `protobuf:"bytes,8,opt,name=startDate,proto3" json:"startDate,omitempty"`
	PollCreator     string   `protobuf:"bytes,9,opt,name=pollCreator,proto3" json:"pollCreator,omitempty"`
}

func (m *Poll) Reset()         { *m = Poll{} }
func (m *Poll) String() string { return proto.CompactTextString(m) }
func (*Poll) ProtoMessage()    {}
func (*Poll) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5c8e1057afeac62, []int{0}
}
func (m *Poll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Poll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Poll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Poll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Poll.Merge(m, src)
}
func (m *Poll) XXX_Size() int {
	return m.Size()
}
func (m *Poll) XXX_DiscardUnknown() {
	xxx_messageInfo_Poll.DiscardUnknown(m)
}

var xxx_messageInfo_Poll proto.InternalMessageInfo

func (m *Poll) GetPollId() string {
	if m != nil {
		return m.PollId
	}
	return ""
}

func (m *Poll) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *Poll) GetNewValidator() string {
	if m != nil {
		return m.NewValidator
	}
	return ""
}

func (m *Poll) GetVotesDoneBy() []string {
	if m != nil {
		return m.VotesDoneBy
	}
	return nil
}

func (m *Poll) GetVotes() []string {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *Poll) GetTotalValidators() uint64 {
	if m != nil {
		return m.TotalValidators
	}
	return 0
}

func (m *Poll) GetIsComplete() bool {
	if m != nil {
		return m.IsComplete
	}
	return false
}

func (m *Poll) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Poll) GetPollCreator() string {
	if m != nil {
		return m.PollCreator
	}
	return ""
}

func init() {
	proto.RegisterType((*Poll)(nil), "airsettle.airsettle.Poll")
}

func init() { proto.RegisterFile("airsettle/airsettle/poll.proto", fileDescriptor_a5c8e1057afeac62) }

var fileDescriptor_a5c8e1057afeac62 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xeb, 0x36, 0x4d, 0x9b, 0x03, 0x09, 0xc9, 0x20, 0x64, 0x09, 0x64, 0x45, 0x9d, 0x3c,
	0x95, 0x01, 0xf1, 0x02, 0x6d, 0x97, 0x6e, 0x28, 0x03, 0x03, 0x9b, 0x21, 0x27, 0x11, 0xc9, 0xc4,
	0x91, 0x7d, 0x02, 0xfa, 0x16, 0x3c, 0x01, 0xcf, 0xc3, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0x62,
	0x20, 0x09, 0x6c, 0xf7, 0x7f, 0x9f, 0x65, 0xdd, 0xfd, 0x20, 0x75, 0xe1, 0x3c, 0x12, 0x19, 0xbc,
	0xe8, 0xa7, 0xca, 0x1a, 0xb3, 0xac, 0x9c, 0x25, 0xcb, 0x8f, 0x3b, 0xba, 0xec, 0xa6, 0xc5, 0xdb,
	0x18, 0xa2, 0x6b, 0x6b, 0x0c, 0x3f, 0x85, 0xb8, 0x7d, 0xbb, 0xcd, 0x05, 0x4b, 0x99, 0x4a, 0xb2,
	0x9f, 0xc4, 0x05, 0xcc, 0xee, 0x1f, 0x74, 0x51, 0x6e, 0x73, 0x31, 0x0e, 0xe2, 0x37, 0xf2, 0x05,
	0x1c, 0x96, 0xf8, 0x7c, 0xa3, 0x4d, 0x91, 0x6b, 0xb2, 0x4e, 0x4c, 0x82, 0xfe, 0xc3, 0x78, 0x0a,
	0x07, 0x4f, 0x96, 0xd0, 0x6f, 0x6c, 0x89, 0xab, 0x9d, 0x88, 0xd2, 0x89, 0x4a, 0xb2, 0x21, 0xe2,
	0x27, 0x30, 0x0d, 0x51, 0x4c, 0x83, 0xfb, 0x0e, 0x5c, 0xc1, 0x11, 0x59, 0xd2, 0xa6, 0xfb, 0xc9,
	0x8b, 0x38, 0x65, 0x2a, 0xca, 0xfe, 0x63, 0x2e, 0x01, 0x0a, 0xbf, 0xb6, 0x8f, 0x95, 0x41, 0x42,
	0x31, 0x4b, 0x99, 0x9a, 0x67, 0x03, 0xc2, 0xcf, 0x21, 0xf1, 0xa4, 0x1d, 0x6d, 0x34, 0xa1, 0x98,
	0x87, 0x15, 0x7b, 0xd0, 0xee, 0xd7, 0xde, 0xb9, 0x76, 0x18, 0x4e, 0x48, 0x82, 0x1f, 0xa2, 0xd5,
	0xd5, 0x7b, 0x2d, 0xd9, 0xbe, 0x96, 0xec, 0xb3, 0x96, 0xec, 0xb5, 0x91, 0xa3, 0x7d, 0x23, 0x47,
	0x1f, 0x8d, 0x1c, 0xdd, 0x9e, 0xf5, 0x2d, 0xbf, 0x0c, 0x1a, 0xa7, 0x5d, 0x85, 0xfe, 0x2e, 0x0e,
	0x9d, 0x5f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x15, 0x9d, 0x60, 0x26, 0x95, 0x01, 0x00, 0x00,
}

func (m *Poll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Poll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Poll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PollCreator) > 0 {
		i -= len(m.PollCreator)
		copy(dAtA[i:], m.PollCreator)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.PollCreator)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.StartDate) > 0 {
		i -= len(m.StartDate)
		copy(dAtA[i:], m.StartDate)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.StartDate)))
		i--
		dAtA[i] = 0x42
	}
	if m.IsComplete {
		i--
		if m.IsComplete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.TotalValidators != 0 {
		i = encodeVarintPoll(dAtA, i, uint64(m.TotalValidators))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Votes[iNdEx])
			copy(dAtA[i:], m.Votes[iNdEx])
			i = encodeVarintPoll(dAtA, i, uint64(len(m.Votes[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VotesDoneBy) > 0 {
		for iNdEx := len(m.VotesDoneBy) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VotesDoneBy[iNdEx])
			copy(dAtA[i:], m.VotesDoneBy[iNdEx])
			i = encodeVarintPoll(dAtA, i, uint64(len(m.VotesDoneBy[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.NewValidator) > 0 {
		i -= len(m.NewValidator)
		copy(dAtA[i:], m.NewValidator)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.NewValidator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PollId) > 0 {
		i -= len(m.PollId)
		copy(dAtA[i:], m.PollId)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.PollId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoll(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoll(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Poll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PollId)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.NewValidator)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	if len(m.VotesDoneBy) > 0 {
		for _, s := range m.VotesDoneBy {
			l = len(s)
			n += 1 + l + sovPoll(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, s := range m.Votes {
			l = len(s)
			n += 1 + l + sovPoll(uint64(l))
		}
	}
	if m.TotalValidators != 0 {
		n += 1 + sovPoll(uint64(m.TotalValidators))
	}
	if m.IsComplete {
		n += 2
	}
	l = len(m.StartDate)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.PollCreator)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	return n
}

func sovPoll(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoll(x uint64) (n int) {
	return sovPoll(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Poll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoll
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
			return fmt.Errorf("proto: Poll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Poll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PollId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewValidator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewValidator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotesDoneBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VotesDoneBy = append(m.VotesDoneBy, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalValidators", wireType)
			}
			m.TotalValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalValidators |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsComplete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsComplete = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollCreator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
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
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PollCreator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoll
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
func skipPoll(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoll
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
					return 0, ErrIntOverflowPoll
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
					return 0, ErrIntOverflowPoll
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
				return 0, ErrInvalidLengthPoll
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoll
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoll
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoll        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoll          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoll = fmt.Errorf("proto: unexpected end of group")
)