// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sumtaskapp/sumpb/sum.proto

package sumpb

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

type Sum struct {
	FirstNumber          int32    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int32    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce460da49e8ca7c3, []int{0}
}

func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (m *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(m, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetFirstNumber() int32 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *Sum) GetSecondNumber() int32 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce460da49e8ca7c3, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	SumResult            int32    `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce460da49e8ca7c3, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSumResult() int32 {
	if m != nil {
		return m.SumResult
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "sum.Sum")
	proto.RegisterType((*SumRequest)(nil), "sum.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sum.SumResponse")
}

func init() { proto.RegisterFile("sumtaskapp/sumpb/sum.proto", fileDescriptor_ce460da49e8ca7c3) }

var fileDescriptor_ce460da49e8ca7c3 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xbf, 0x4f, 0x86, 0x30,
	0x10, 0x86, 0x45, 0x82, 0x3f, 0x0e, 0x8c, 0xa6, 0x93, 0x21, 0x31, 0x51, 0x5c, 0x88, 0x31, 0x98,
	0xe0, 0xe2, 0xec, 0xae, 0x03, 0x6c, 0x2e, 0x04, 0xf0, 0x4c, 0x88, 0x94, 0xd6, 0x5e, 0xcf, 0xbf,
	0xdf, 0xf4, 0xea, 0xf7, 0x7d, 0x4b, 0x87, 0xb7, 0x4f, 0xde, 0xe7, 0x5e, 0x28, 0x89, 0xb5, 0x1f,
	0xe9, 0x7b, 0xb4, 0xf6, 0x89, 0x58, 0xdb, 0x29, 0xbc, 0x8d, 0x75, 0xc6, 0x1b, 0x95, 0x12, 0xeb,
	0xea, 0x0d, 0xd2, 0x9e, 0xb5, 0xba, 0x83, 0xe2, 0x6b, 0x71, 0xe4, 0x87, 0x8d, 0xf5, 0x84, 0xee,
	0x3a, 0xb9, 0x4d, 0xea, 0xac, 0xcb, 0x25, 0x7b, 0x97, 0x48, 0xdd, 0xc3, 0x05, 0xe1, 0x6c, 0xb6,
	0xcf, 0x1d, 0x73, 0x2c, 0x4c, 0x11, 0xc3, 0x08, 0x55, 0x35, 0x40, 0xcf, 0xba, 0xc3, 0x1f, 0x46,
	0xf2, 0xaa, 0x84, 0xe0, 0x90, 0xb2, 0xbc, 0x3d, 0x6b, 0x82, 0x3a, 0xfc, 0x8a, 0xf8, 0x11, 0x72,
	0x21, 0xc9, 0x9a, 0x8d, 0x50, 0xdd, 0x00, 0x10, 0xeb, 0xc1, 0x21, 0xf1, 0xea, 0xff, 0xf5, 0xe7,
	0x24, 0x00, 0xaf, 0xbe, 0x7d, 0x91, 0xde, 0x1e, 0xdd, 0xef, 0x32, 0xa3, 0x7a, 0x88, 0x47, 0x5f,
	0xee, 0x1b, 0xa3, 0xaf, 0xbc, 0x3a, 0x04, 0xb1, 0xb6, 0x3a, 0x7a, 0x3d, 0xfd, 0xc8, 0x64, 0xf8,
	0x74, 0x22, 0xab, 0x9f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x4d, 0xcd, 0xdb, 0x13, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sum.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sumtaskapp/sumpb/sum.proto",
}