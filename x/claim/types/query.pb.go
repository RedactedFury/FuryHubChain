// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/claim/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryAirdropsRequest is request type for the Query/Airdrops RPC method.
type QueryAirdropsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAirdropsRequest) Reset()         { *m = QueryAirdropsRequest{} }
func (m *QueryAirdropsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAirdropsRequest) ProtoMessage()    {}
func (*QueryAirdropsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4d96c11986b085, []int{0}
}
func (m *QueryAirdropsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAirdropsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAirdropsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAirdropsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAirdropsRequest.Merge(m, src)
}
func (m *QueryAirdropsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAirdropsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAirdropsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAirdropsRequest proto.InternalMessageInfo

func (m *QueryAirdropsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAirdropsResponse is response type for the Query/Airdrops RPC method.
type QueryAirdropsResponse struct {
	Airdrops   []Airdrop           `protobuf:"bytes,1,rep,name=airdrops,proto3" json:"airdrops"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAirdropsResponse) Reset()         { *m = QueryAirdropsResponse{} }
func (m *QueryAirdropsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAirdropsResponse) ProtoMessage()    {}
func (*QueryAirdropsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4d96c11986b085, []int{1}
}
func (m *QueryAirdropsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAirdropsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAirdropsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAirdropsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAirdropsResponse.Merge(m, src)
}
func (m *QueryAirdropsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAirdropsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAirdropsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAirdropsResponse proto.InternalMessageInfo

func (m *QueryAirdropsResponse) GetAirdrops() []Airdrop {
	if m != nil {
		return m.Airdrops
	}
	return nil
}

func (m *QueryAirdropsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAirdropRequest is request type for the Query/Airdrop RPC method.
type QueryAirdropRequest struct {
	AirdropId uint64 `protobuf:"varint,1,opt,name=airdrop_id,json=airdropId,proto3" json:"airdrop_id,omitempty"`
}

func (m *QueryAirdropRequest) Reset()         { *m = QueryAirdropRequest{} }
func (m *QueryAirdropRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAirdropRequest) ProtoMessage()    {}
func (*QueryAirdropRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4d96c11986b085, []int{2}
}
func (m *QueryAirdropRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAirdropRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAirdropRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAirdropRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAirdropRequest.Merge(m, src)
}
func (m *QueryAirdropRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAirdropRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAirdropRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAirdropRequest proto.InternalMessageInfo

func (m *QueryAirdropRequest) GetAirdropId() uint64 {
	if m != nil {
		return m.AirdropId
	}
	return 0
}

// QueryAirdropResponse is response type for the Query/Airdrop RPC method.
type QueryAirdropResponse struct {
	Airdrop Airdrop `protobuf:"bytes,1,opt,name=airdrop,proto3" json:"airdrop"`
}

func (m *QueryAirdropResponse) Reset()         { *m = QueryAirdropResponse{} }
func (m *QueryAirdropResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAirdropResponse) ProtoMessage()    {}
func (*QueryAirdropResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4d96c11986b085, []int{3}
}
func (m *QueryAirdropResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAirdropResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAirdropResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAirdropResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAirdropResponse.Merge(m, src)
}
func (m *QueryAirdropResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAirdropResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAirdropResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAirdropResponse proto.InternalMessageInfo

func (m *QueryAirdropResponse) GetAirdrop() Airdrop {
	if m != nil {
		return m.Airdrop
	}
	return Airdrop{}
}

// QueryClaimRecordRequest is request type for the Query/ClaimRecord RPC method.
type QueryClaimRecordRequest struct {
	AirdropId uint64 `protobuf:"varint,1,opt,name=airdrop_id,json=airdropId,proto3" json:"airdrop_id,omitempty"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
}

func (m *QueryClaimRecordRequest) Reset()         { *m = QueryClaimRecordRequest{} }
func (m *QueryClaimRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryClaimRecordRequest) ProtoMessage()    {}
func (*QueryClaimRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4d96c11986b085, []int{4}
}
func (m *QueryClaimRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimRecordRequest.Merge(m, src)
}
func (m *QueryClaimRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimRecordRequest proto.InternalMessageInfo

func (m *QueryClaimRecordRequest) GetAirdropId() uint64 {
	if m != nil {
		return m.AirdropId
	}
	return 0
}

func (m *QueryClaimRecordRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

// QueryClaimRecordResponse is response type for the Query/ClaimRecord RPC method.
type QueryClaimRecordResponse struct {
	ClaimRecord ClaimRecord `protobuf:"bytes,1,opt,name=claim_record,json=claimRecord,proto3" json:"claim_record"`
}

func (m *QueryClaimRecordResponse) Reset()         { *m = QueryClaimRecordResponse{} }
func (m *QueryClaimRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClaimRecordResponse) ProtoMessage()    {}
func (*QueryClaimRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd4d96c11986b085, []int{5}
}
func (m *QueryClaimRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimRecordResponse.Merge(m, src)
}
func (m *QueryClaimRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimRecordResponse proto.InternalMessageInfo

func (m *QueryClaimRecordResponse) GetClaimRecord() ClaimRecord {
	if m != nil {
		return m.ClaimRecord
	}
	return ClaimRecord{}
}

func init() {
	proto.RegisterType((*QueryAirdropsRequest)(nil), "fury.claim.v1beta1.QueryAirdropsRequest")
	proto.RegisterType((*QueryAirdropsResponse)(nil), "fury.claim.v1beta1.QueryAirdropsResponse")
	proto.RegisterType((*QueryAirdropRequest)(nil), "fury.claim.v1beta1.QueryAirdropRequest")
	proto.RegisterType((*QueryAirdropResponse)(nil), "fury.claim.v1beta1.QueryAirdropResponse")
	proto.RegisterType((*QueryClaimRecordRequest)(nil), "fury.claim.v1beta1.QueryClaimRecordRequest")
	proto.RegisterType((*QueryClaimRecordResponse)(nil), "fury.claim.v1beta1.QueryClaimRecordResponse")
}

func init() {
	proto.RegisterFile("fury/claim/v1beta1/query.proto", fileDescriptor_bd4d96c11986b085)
}

var fileDescriptor_bd4d96c11986b085 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6b, 0x13, 0x5d,
	0x14, 0xc6, 0x73, 0xdb, 0xbe, 0x6f, 0x9b, 0x13, 0x57, 0xd7, 0xaa, 0x21, 0xd4, 0x69, 0x18, 0x41,
	0x83, 0x36, 0xf7, 0xda, 0xd4, 0xae, 0xa5, 0x15, 0x2c, 0x05, 0xa1, 0x3a, 0x0b, 0x05, 0x17, 0x96,
	0xc9, 0xe4, 0x32, 0x1d, 0x6c, 0xe6, 0x4e, 0xe7, 0xde, 0x54, 0x4b, 0xe9, 0xc6, 0x4f, 0xa0, 0xb8,
	0x74, 0xe9, 0x27, 0x71, 0xd7, 0x65, 0xc1, 0x8d, 0x2b, 0x91, 0xc4, 0xaf, 0x21, 0x48, 0xee, 0x9c,
	0xf9, 0x13, 0xdb, 0xd8, 0x71, 0x17, 0x4e, 0x9e, 0x73, 0x9e, 0xdf, 0x79, 0xee, 0x61, 0xc0, 0xf6,
	0x62, 0xa1, 0x3c, 0x11, 0x6a, 0xee, 0xed, 0xbb, 0x41, 0x9f, 0x1f, 0xae, 0x76, 0x85, 0x76, 0x57,
	0xf9, 0xc1, 0x40, 0xc4, 0x47, 0x2c, 0x8a, 0xa5, 0x96, 0xf4, 0x7a, 0xaa, 0x61, 0x46, 0xc3, 0x50,
	0xd3, 0x58, 0xf4, 0xa5, 0x2f, 0x8d, 0x84, 0x8f, 0x7f, 0x25, 0xea, 0xc6, 0x92, 0x2f, 0xa5, 0xbf,
	0x2f, 0xb8, 0x1b, 0x05, 0xdc, 0x0d, 0x43, 0xa9, 0x5d, 0x1d, 0xc8, 0x50, 0xe1, 0xbf, 0x77, 0x3d,
	0xa9, 0xfa, 0x52, 0xf1, 0xae, 0xab, 0x44, 0x62, 0x92, 0x59, 0x46, 0xae, 0x1f, 0x84, 0x46, 0x8c,
	0xda, 0x69, 0x6c, 0x09, 0x85, 0xd1, 0xd8, 0xaf, 0x60, 0xf1, 0xd9, 0x78, 0xca, 0x46, 0x10, 0xf7,
	0x62, 0x19, 0x29, 0x47, 0x1c, 0x0c, 0x84, 0xd2, 0xf4, 0x31, 0x40, 0x3e, 0xaf, 0x4e, 0x9a, 0xa4,
	0x55, 0xeb, 0xdc, 0x66, 0x89, 0x39, 0x1b, 0x9b, 0xb3, 0x64, 0x43, 0x9c, 0xc9, 0x9e, 0xba, 0xbe,
	0xc0, 0x5e, 0xa7, 0xd0, 0x69, 0x7f, 0x26, 0x70, 0xed, 0x0f, 0x03, 0x15, 0xc9, 0x50, 0x09, 0xba,
	0x01, 0x0b, 0x2e, 0xd6, 0xea, 0xa4, 0x39, 0xdb, 0xaa, 0x75, 0x96, 0xd9, 0xc5, 0x41, 0x31, 0xec,
	0xdd, 0x9c, 0x3b, 0xfd, 0xbe, 0x5c, 0x71, 0xb2, 0x36, 0xba, 0x35, 0x01, 0x39, 0x63, 0x20, 0xef,
	0x5c, 0x0a, 0x99, 0xf8, 0x4f, 0x50, 0x3e, 0x80, 0xab, 0x45, 0xc8, 0x34, 0x84, 0x9b, 0x00, 0xe8,
	0xb5, 0x1b, 0xf4, 0x4c, 0x08, 0x73, 0x4e, 0x15, 0x2b, 0xdb, 0x3d, 0xfb, 0xc5, 0x64, 0x76, 0xd9,
	0x66, 0x0f, 0x61, 0x1e, 0x45, 0x18, 0x5c, 0xc9, 0xc5, 0xd2, 0x2e, 0xfb, 0x39, 0xdc, 0x30, 0x83,
	0x1f, 0x8d, 0xc5, 0x8e, 0xf0, 0x64, 0xdc, 0x2b, 0x87, 0x44, 0x97, 0xa0, 0x1a, 0x0b, 0x2f, 0x88,
	0x02, 0x11, 0x6a, 0x13, 0x48, 0xd5, 0xc9, 0x0b, 0xf6, 0x1e, 0xd4, 0xcf, 0xcf, 0x45, 0xe8, 0x27,
	0x70, 0xc5, 0xb0, 0xed, 0xc6, 0xa6, 0x8e, 0xe4, 0xb7, 0xa6, 0x91, 0x17, 0x46, 0x20, 0x7d, 0xcd,
	0xcb, 0x4b, 0x9d, 0x5f, 0xb3, 0xf0, 0x9f, 0xb1, 0xa2, 0x1f, 0x08, 0x2c, 0xa4, 0x6f, 0x4f, 0x57,
	0xa6, 0x8d, 0xbb, 0xe8, 0x06, 0x1b, 0xed, 0x92, 0xea, 0x64, 0x03, 0xbb, 0xf5, 0xee, 0xeb, 0xcf,
	0x8f, 0x33, 0x36, 0x6d, 0xf2, 0x29, 0x77, 0x9f, 0xdd, 0xcd, 0x27, 0x02, 0xf3, 0xd8, 0x4e, 0xef,
	0x95, 0x31, 0x49, 0x89, 0x56, 0xca, 0x89, 0x11, 0x68, 0xdd, 0x00, 0x71, 0xda, 0xbe, 0x0c, 0x88,
	0x1f, 0xe7, 0x6f, 0x7a, 0x42, 0xbf, 0x10, 0xa8, 0x15, 0xe2, 0xa5, 0xfc, 0xaf, 0xa6, 0xe7, 0x6f,
	0xa4, 0x71, 0xbf, 0x7c, 0x03, 0x92, 0xee, 0x18, 0xd2, 0x6d, 0xba, 0xf5, 0x4f, 0xa4, 0xbc, 0x78,
	0x30, 0x8a, 0x1f, 0x67, 0x87, 0x76, 0xb2, 0xb9, 0x73, 0x3a, 0xb4, 0xc8, 0xd9, 0xd0, 0x22, 0x3f,
	0x86, 0x16, 0x79, 0x3f, 0xb2, 0x2a, 0x67, 0x23, 0xab, 0xf2, 0x6d, 0x64, 0x55, 0x5e, 0xae, 0xfb,
	0x81, 0xde, 0x1b, 0x74, 0x99, 0x27, 0xfb, 0x99, 0x59, 0x3b, 0x14, 0xfa, 0x8d, 0x8c, 0x5f, 0xe7,
	0xee, 0x87, 0x6b, 0xfc, 0x2d, 0x22, 0xe8, 0xa3, 0x48, 0xa8, 0xee, 0xff, 0xe6, 0x73, 0xb5, 0xf6,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x2f, 0x04, 0xc5, 0x70, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Airdrops returns all airdrops.
	Airdrops(ctx context.Context, in *QueryAirdropsRequest, opts ...grpc.CallOption) (*QueryAirdropsResponse, error)
	// Airdrop returns the specific airdrop.
	Airdrop(ctx context.Context, in *QueryAirdropRequest, opts ...grpc.CallOption) (*QueryAirdropResponse, error)
	// ClaimRecord returns the claim record for the recipient address.
	ClaimRecord(ctx context.Context, in *QueryClaimRecordRequest, opts ...grpc.CallOption) (*QueryClaimRecordResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Airdrops(ctx context.Context, in *QueryAirdropsRequest, opts ...grpc.CallOption) (*QueryAirdropsResponse, error) {
	out := new(QueryAirdropsResponse)
	err := c.cc.Invoke(ctx, "/fury.claim.v1beta1.Query/Airdrops", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Airdrop(ctx context.Context, in *QueryAirdropRequest, opts ...grpc.CallOption) (*QueryAirdropResponse, error) {
	out := new(QueryAirdropResponse)
	err := c.cc.Invoke(ctx, "/fury.claim.v1beta1.Query/Airdrop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ClaimRecord(ctx context.Context, in *QueryClaimRecordRequest, opts ...grpc.CallOption) (*QueryClaimRecordResponse, error) {
	out := new(QueryClaimRecordResponse)
	err := c.cc.Invoke(ctx, "/fury.claim.v1beta1.Query/ClaimRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Airdrops returns all airdrops.
	Airdrops(context.Context, *QueryAirdropsRequest) (*QueryAirdropsResponse, error)
	// Airdrop returns the specific airdrop.
	Airdrop(context.Context, *QueryAirdropRequest) (*QueryAirdropResponse, error)
	// ClaimRecord returns the claim record for the recipient address.
	ClaimRecord(context.Context, *QueryClaimRecordRequest) (*QueryClaimRecordResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Airdrops(ctx context.Context, req *QueryAirdropsRequest) (*QueryAirdropsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Airdrops not implemented")
}
func (*UnimplementedQueryServer) Airdrop(ctx context.Context, req *QueryAirdropRequest) (*QueryAirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Airdrop not implemented")
}
func (*UnimplementedQueryServer) ClaimRecord(ctx context.Context, req *QueryClaimRecordRequest) (*QueryClaimRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimRecord not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Airdrops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAirdropsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Airdrops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fury.claim.v1beta1.Query/Airdrops",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Airdrops(ctx, req.(*QueryAirdropsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Airdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAirdropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Airdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fury.claim.v1beta1.Query/Airdrop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Airdrop(ctx, req.(*QueryAirdropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ClaimRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClaimRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fury.claim.v1beta1.Query/ClaimRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClaimRecord(ctx, req.(*QueryClaimRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fury.claim.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Airdrops",
			Handler:    _Query_Airdrops_Handler,
		},
		{
			MethodName: "Airdrop",
			Handler:    _Query_Airdrop_Handler,
		},
		{
			MethodName: "ClaimRecord",
			Handler:    _Query_ClaimRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fury/claim/v1beta1/query.proto",
}

func (m *QueryAirdropsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAirdropsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAirdropsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAirdropsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAirdropsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAirdropsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Airdrops) > 0 {
		for iNdEx := len(m.Airdrops) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Airdrops[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryAirdropRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAirdropRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAirdropRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AirdropId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.AirdropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryAirdropResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAirdropResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAirdropResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Airdrop.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryClaimRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.AirdropId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.AirdropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryClaimRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ClaimRecord.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAirdropsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAirdropsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Airdrops) > 0 {
		for _, e := range m.Airdrops {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAirdropRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropId != 0 {
		n += 1 + sovQuery(uint64(m.AirdropId))
	}
	return n
}

func (m *QueryAirdropResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Airdrop.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryClaimRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropId != 0 {
		n += 1 + sovQuery(uint64(m.AirdropId))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryClaimRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClaimRecord.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAirdropsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAirdropsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAirdropsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAirdropsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAirdropsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAirdropsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Airdrops", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Airdrops = append(m.Airdrops, Airdrop{})
			if err := m.Airdrops[len(m.Airdrops)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAirdropRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAirdropRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAirdropRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropId", wireType)
			}
			m.AirdropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirdropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAirdropResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAirdropResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAirdropResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Airdrop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Airdrop.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryClaimRecordRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryClaimRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropId", wireType)
			}
			m.AirdropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirdropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryClaimRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryClaimRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClaimRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)