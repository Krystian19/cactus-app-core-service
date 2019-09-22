// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/release.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Release struct {
	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Poster       string `protobuf:"bytes,3,opt,name=poster,proto3" json:"poster,omitempty"`
	Background   string `protobuf:"bytes,4,opt,name=background,proto3" json:"background,omitempty"`
	ReleaseOrder string `protobuf:"bytes,5,opt,name=release_order,json=releaseOrder,proto3" json:"release_order,omitempty"`
	CreatedAt    string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relations
	AnimeId              int64    `protobuf:"varint,8,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{0}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Release) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Release) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func (m *Release) GetBackground() string {
	if m != nil {
		return m.Background
	}
	return ""
}

func (m *Release) GetReleaseOrder() string {
	if m != nil {
		return m.ReleaseOrder
	}
	return ""
}

func (m *Release) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Release) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Release) GetAnimeId() int64 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

type ReleaseQuery struct {
	AnimeId int64 `protobuf:"varint,3,opt,name=anime_id,json=animeId,proto3" json:"anime_id,omitempty"`
	// Pagination
	Offset               int64    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseQuery) Reset()         { *m = ReleaseQuery{} }
func (m *ReleaseQuery) String() string { return proto.CompactTextString(m) }
func (*ReleaseQuery) ProtoMessage()    {}
func (*ReleaseQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{1}
}

func (m *ReleaseQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseQuery.Unmarshal(m, b)
}
func (m *ReleaseQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseQuery.Marshal(b, m, deterministic)
}
func (m *ReleaseQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseQuery.Merge(m, src)
}
func (m *ReleaseQuery) XXX_Size() int {
	return xxx_messageInfo_ReleaseQuery.Size(m)
}
func (m *ReleaseQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseQuery proto.InternalMessageInfo

func (m *ReleaseQuery) GetAnimeId() int64 {
	if m != nil {
		return m.AnimeId
	}
	return 0
}

func (m *ReleaseQuery) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReleaseQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReleaseRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseRequest) Reset()         { *m = ReleaseRequest{} }
func (m *ReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseRequest) ProtoMessage()    {}
func (*ReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{2}
}

func (m *ReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseRequest.Unmarshal(m, b)
}
func (m *ReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseRequest.Merge(m, src)
}
func (m *ReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseRequest.Size(m)
}
func (m *ReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseRequest proto.InternalMessageInfo

func (m *ReleaseRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReleasesRequest struct {
	Query                *ReleaseQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReleasesRequest) Reset()         { *m = ReleasesRequest{} }
func (m *ReleasesRequest) String() string { return proto.CompactTextString(m) }
func (*ReleasesRequest) ProtoMessage()    {}
func (*ReleasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{3}
}

func (m *ReleasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleasesRequest.Unmarshal(m, b)
}
func (m *ReleasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleasesRequest.Marshal(b, m, deterministic)
}
func (m *ReleasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasesRequest.Merge(m, src)
}
func (m *ReleasesRequest) XXX_Size() int {
	return xxx_messageInfo_ReleasesRequest.Size(m)
}
func (m *ReleasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasesRequest proto.InternalMessageInfo

func (m *ReleasesRequest) GetQuery() *ReleaseQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type ReleaseResponse struct {
	Release              *Release `protobuf:"bytes,1,opt,name=release,proto3" json:"release,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseResponse) Reset()         { *m = ReleaseResponse{} }
func (m *ReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseResponse) ProtoMessage()    {}
func (*ReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{4}
}

func (m *ReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseResponse.Unmarshal(m, b)
}
func (m *ReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseResponse.Merge(m, src)
}
func (m *ReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseResponse.Size(m)
}
func (m *ReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseResponse proto.InternalMessageInfo

func (m *ReleaseResponse) GetRelease() *Release {
	if m != nil {
		return m.Release
	}
	return nil
}

type ReleasesResponse struct {
	Releases             []*Release `protobuf:"bytes,1,rep,name=releases,proto3" json:"releases,omitempty"`
	Count                uint64     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReleasesResponse) Reset()         { *m = ReleasesResponse{} }
func (m *ReleasesResponse) String() string { return proto.CompactTextString(m) }
func (*ReleasesResponse) ProtoMessage()    {}
func (*ReleasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_987cf5695bafb180, []int{5}
}

func (m *ReleasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleasesResponse.Unmarshal(m, b)
}
func (m *ReleasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleasesResponse.Marshal(b, m, deterministic)
}
func (m *ReleasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasesResponse.Merge(m, src)
}
func (m *ReleasesResponse) XXX_Size() int {
	return xxx_messageInfo_ReleasesResponse.Size(m)
}
func (m *ReleasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasesResponse proto.InternalMessageInfo

func (m *ReleasesResponse) GetReleases() []*Release {
	if m != nil {
		return m.Releases
	}
	return nil
}

func (m *ReleasesResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Release)(nil), "proto.Release")
	proto.RegisterType((*ReleaseQuery)(nil), "proto.ReleaseQuery")
	proto.RegisterType((*ReleaseRequest)(nil), "proto.ReleaseRequest")
	proto.RegisterType((*ReleasesRequest)(nil), "proto.ReleasesRequest")
	proto.RegisterType((*ReleaseResponse)(nil), "proto.ReleaseResponse")
	proto.RegisterType((*ReleasesResponse)(nil), "proto.ReleasesResponse")
}

func init() { proto.RegisterFile("proto/release.proto", fileDescriptor_987cf5695bafb180) }

var fileDescriptor_987cf5695bafb180 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x55, 0x9a, 0xa6, 0x49, 0x2f, 0xa5, 0x20, 0x17, 0x8a, 0xa9, 0x04, 0x8a, 0xc2, 0xa6, 0xb0,
	0x28, 0x52, 0xd9, 0x20, 0x95, 0x4d, 0x97, 0xac, 0x10, 0x06, 0x89, 0x65, 0x95, 0xc6, 0xb7, 0xc8,
	0x9a, 0x36, 0x4e, 0x6d, 0x67, 0xa4, 0xf9, 0x82, 0xf9, 0xd2, 0xf9, 0x8f, 0x51, 0x6c, 0xa7, 0x8f,
	0xcc, 0xac, 0xa2, 0xf3, 0xb8, 0xce, 0x39, 0xd7, 0x86, 0x49, 0xa5, 0xa4, 0x91, 0x5f, 0x15, 0xee,
	0x31, 0xd7, 0xb8, 0xb0, 0x88, 0x44, 0xf6, 0x93, 0x3d, 0x04, 0x10, 0x33, 0x27, 0x90, 0x31, 0xf4,
	0x04, 0xa7, 0x41, 0x1a, 0xcc, 0x43, 0xd6, 0x13, 0x9c, 0xbc, 0x81, 0xc8, 0x08, 0xb3, 0x47, 0xda,
	0x4b, 0x83, 0xf9, 0x90, 0x39, 0x40, 0xa6, 0x30, 0xa8, 0xa4, 0x36, 0xa8, 0x68, 0x68, 0x69, 0x8f,
	0xc8, 0x47, 0x80, 0x6d, 0x5e, 0xdc, 0xfc, 0x57, 0xb2, 0x2e, 0x39, 0xed, 0x5b, 0xed, 0x82, 0x21,
	0x9f, 0xe0, 0xa5, 0x4f, 0xb0, 0x91, 0x8a, 0xa3, 0xa2, 0x91, 0xb5, 0x8c, 0x3c, 0xf9, 0xab, 0xe1,
	0xc8, 0x07, 0x80, 0x42, 0x61, 0x6e, 0x90, 0x6f, 0x72, 0x43, 0x07, 0xd6, 0x31, 0xf4, 0xcc, 0xda,
	0x34, 0x72, 0x5d, 0xf1, 0x56, 0x8e, 0x9d, 0xec, 0x99, 0xb5, 0x21, 0xef, 0x21, 0xc9, 0x4b, 0x71,
	0xc0, 0x8d, 0xe0, 0x34, 0xb1, 0x35, 0x62, 0x8b, 0x7f, 0xf2, 0xec, 0x1f, 0x8c, 0x7c, 0xcd, 0xdf,
	0x35, 0xaa, 0xbb, 0x2b, 0x6b, 0x78, 0x65, 0x6d, 0x0a, 0xca, 0xdd, 0x4e, 0xa3, 0xb1, 0x25, 0x42,
	0xe6, 0x51, 0xb3, 0x8e, 0xbd, 0x38, 0x08, 0x63, 0x83, 0x87, 0xcc, 0x81, 0x2c, 0x85, 0xb1, 0x3f,
	0x98, 0xe1, 0xb1, 0x46, 0x6d, 0xba, 0x6b, 0xcc, 0x7e, 0xc0, 0x2b, 0xef, 0xd0, 0xad, 0xe5, 0x33,
	0x44, 0xc7, 0x26, 0x86, 0x75, 0xbd, 0x58, 0x4e, 0xdc, 0x9d, 0x2c, 0x2e, 0x13, 0x32, 0xe7, 0xc8,
	0x56, 0xa7, 0x69, 0x86, 0xba, 0x92, 0xa5, 0x46, 0x32, 0x87, 0xd8, 0x2f, 0xcd, 0xcf, 0x8f, 0xaf,
	0xe7, 0x59, 0x2b, 0x67, 0x7f, 0xe1, 0xf5, 0xf9, 0xd7, 0x7e, 0xfa, 0x0b, 0x24, 0x5e, 0xd6, 0x34,
	0x48, 0xc3, 0x67, 0xc6, 0x4f, 0x7a, 0x53, 0xb9, 0x90, 0x75, 0x69, 0xec, 0x0b, 0xe8, 0x33, 0x07,
	0x96, 0xf7, 0xc1, 0xa9, 0xf3, 0x1f, 0x54, 0xb7, 0xa2, 0x40, 0xf2, 0xfd, 0xfc, 0x8a, 0xde, 0x76,
	0x4e, 0x73, 0x95, 0x67, 0xd3, 0x2e, 0xed, 0xe3, 0xac, 0x20, 0x69, 0x23, 0x92, 0x8e, 0xa7, 0x5d,
	0xd7, 0xec, 0xdd, 0x13, 0xde, 0x0d, 0x6f, 0x07, 0x96, 0xff, 0xf6, 0x18, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0xb2, 0x31, 0xe9, 0xe2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReleaseServiceClient is the client API for ReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReleaseServiceClient interface {
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
	Releases(ctx context.Context, in *ReleasesRequest, opts ...grpc.CallOption) (*ReleasesResponse, error)
}

type releaseServiceClient struct {
	cc *grpc.ClientConn
}

func NewReleaseServiceClient(cc *grpc.ClientConn) ReleaseServiceClient {
	return &releaseServiceClient{cc}
}

func (c *releaseServiceClient) Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *releaseServiceClient) Releases(ctx context.Context, in *ReleasesRequest, opts ...grpc.CallOption) (*ReleasesResponse, error) {
	out := new(ReleasesResponse)
	err := c.cc.Invoke(ctx, "/proto.ReleaseService/Releases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReleaseServiceServer is the server API for ReleaseService service.
type ReleaseServiceServer interface {
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	Releases(context.Context, *ReleasesRequest) (*ReleasesResponse, error)
}

// UnimplementedReleaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReleaseServiceServer struct {
}

func (*UnimplementedReleaseServiceServer) Release(ctx context.Context, req *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (*UnimplementedReleaseServiceServer) Releases(ctx context.Context, req *ReleasesRequest) (*ReleasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Releases not implemented")
}

func RegisterReleaseServiceServer(s *grpc.Server, srv ReleaseServiceServer) {
	s.RegisterService(&_ReleaseService_serviceDesc, srv)
}

func _ReleaseService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Release(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReleaseService_Releases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReleaseServiceServer).Releases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReleaseService/Releases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReleaseServiceServer).Releases(ctx, req.(*ReleasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReleaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ReleaseService",
	HandlerType: (*ReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Release",
			Handler:    _ReleaseService_Release_Handler,
		},
		{
			MethodName: "Releases",
			Handler:    _ReleaseService_Releases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/release.proto",
}
