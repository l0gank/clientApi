// Code generated by protoc-gen-go. DO NOT EDIT.
// source: referential.proto

package referential

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

type ReferentialReq struct {
	TblName              string   `protobuf:"bytes,1,opt,name=TblName,proto3" json:"TblName,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReferentialReq) Reset()         { *m = ReferentialReq{} }
func (m *ReferentialReq) String() string { return proto.CompactTextString(m) }
func (*ReferentialReq) ProtoMessage()    {}
func (*ReferentialReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e6f978450c34954, []int{0}
}

func (m *ReferentialReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReferentialReq.Unmarshal(m, b)
}
func (m *ReferentialReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReferentialReq.Marshal(b, m, deterministic)
}
func (m *ReferentialReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferentialReq.Merge(m, src)
}
func (m *ReferentialReq) XXX_Size() int {
	return xxx_messageInfo_ReferentialReq.Size(m)
}
func (m *ReferentialReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferentialReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReferentialReq proto.InternalMessageInfo

func (m *ReferentialReq) GetTblName() string {
	if m != nil {
		return m.TblName
	}
	return ""
}

func (m *ReferentialReq) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReferentialReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReferentialReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReferentialListResp struct {
	ReferentialList      []*ReferentialResp `protobuf:"bytes,1,rep,name=referentialList,proto3" json:"referentialList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ReferentialListResp) Reset()         { *m = ReferentialListResp{} }
func (m *ReferentialListResp) String() string { return proto.CompactTextString(m) }
func (*ReferentialListResp) ProtoMessage()    {}
func (*ReferentialListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e6f978450c34954, []int{1}
}

func (m *ReferentialListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReferentialListResp.Unmarshal(m, b)
}
func (m *ReferentialListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReferentialListResp.Marshal(b, m, deterministic)
}
func (m *ReferentialListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferentialListResp.Merge(m, src)
}
func (m *ReferentialListResp) XXX_Size() int {
	return xxx_messageInfo_ReferentialListResp.Size(m)
}
func (m *ReferentialListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferentialListResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReferentialListResp proto.InternalMessageInfo

func (m *ReferentialListResp) GetReferentialList() []*ReferentialResp {
	if m != nil {
		return m.ReferentialList
	}
	return nil
}

// The response message containing the greetings
type ReferentialResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	SubCode              int32    `protobuf:"varint,2,opt,name=SubCode,proto3" json:"SubCode,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReferentialResp) Reset()         { *m = ReferentialResp{} }
func (m *ReferentialResp) String() string { return proto.CompactTextString(m) }
func (*ReferentialResp) ProtoMessage()    {}
func (*ReferentialResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e6f978450c34954, []int{2}
}

func (m *ReferentialResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReferentialResp.Unmarshal(m, b)
}
func (m *ReferentialResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReferentialResp.Marshal(b, m, deterministic)
}
func (m *ReferentialResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferentialResp.Merge(m, src)
}
func (m *ReferentialResp) XXX_Size() int {
	return xxx_messageInfo_ReferentialResp.Size(m)
}
func (m *ReferentialResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferentialResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReferentialResp proto.InternalMessageInfo

func (m *ReferentialResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReferentialResp) GetSubCode() int32 {
	if m != nil {
		return m.SubCode
	}
	return 0
}

func (m *ReferentialResp) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReferentialResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ReferentialReq)(nil), "omsapi.pkg.utl.grpc.referential.ReferentialReq")
	proto.RegisterType((*ReferentialListResp)(nil), "omsapi.pkg.utl.grpc.referential.ReferentialListResp")
	proto.RegisterType((*ReferentialResp)(nil), "omsapi.pkg.utl.grpc.referential.ReferentialResp")
}

func init() { proto.RegisterFile("referential.proto", fileDescriptor_1e6f978450c34954) }

var fileDescriptor_1e6f978450c34954 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0xad, 0xbf, 0xe4, 0xa3, 0xe2, 0x2a, 0x28, 0x1c, 0x8b, 0xc5, 0x42, 0xe4, 0x29, 0x93, 0xa9,
	0x0a, 0xbf, 0x00, 0x46, 0x10, 0x83, 0xc3, 0xd4, 0x2d, 0x69, 0x8f, 0x2a, 0x22, 0xc1, 0x8e, 0xed,
	0x22, 0xf1, 0xd7, 0x99, 0x50, 0x2c, 0x05, 0xdc, 0x2e, 0x90, 0xed, 0xbd, 0xd3, 0x3b, 0xbf, 0xf7,
	0x4e, 0x86, 0x73, 0x4b, 0x2f, 0x64, 0xe9, 0xcd, 0xd7, 0x65, 0x23, 0x8d, 0xd5, 0x5e, 0xe3, 0x95,
	0x6e, 0x5d, 0x69, 0x6a, 0x69, 0x5e, 0xb7, 0x72, 0xe7, 0x1b, 0xb9, 0xb5, 0x66, 0x2d, 0x23, 0x99,
	0xd8, 0xc0, 0xa9, 0xfa, 0xa1, 0x8a, 0x3a, 0xe4, 0x30, 0x7d, 0xae, 0x9a, 0xa7, 0xb2, 0x25, 0xce,
	0x32, 0x96, 0x1f, 0xab, 0x81, 0x22, 0x42, 0x7a, 0xaf, 0x37, 0xc4, 0xff, 0x65, 0x2c, 0xff, 0xaf,
	0x02, 0xc6, 0x33, 0x48, 0x1e, 0xe8, 0x83, 0x27, 0x41, 0xd9, 0xc3, 0x5e, 0x15, 0x96, 0xd3, 0x30,
	0x0a, 0x58, 0x74, 0x70, 0x11, 0xb9, 0x3c, 0xd6, 0xce, 0x2b, 0x72, 0x06, 0x57, 0x30, 0xb7, 0xfb,
	0x63, 0xce, 0xb2, 0x24, 0x9f, 0x2d, 0x17, 0xf2, 0x97, 0xdc, 0x72, 0x2f, 0xb4, 0x33, 0xea, 0xf0,
	0x21, 0x41, 0x30, 0x3f, 0xd0, 0x7c, 0xe7, 0x67, 0x51, 0x7e, 0x0e, 0xd3, 0x62, 0x57, 0x45, 0xb5,
	0x06, 0xfa, 0xb7, 0x66, 0xcb, 0x4f, 0x06, 0x18, 0xf9, 0x14, 0x64, 0xdf, 0xeb, 0x35, 0xa1, 0x86,
	0xb4, 0x4f, 0x81, 0xd7, 0xe3, 0x8a, 0x74, 0x97, 0xb7, 0x63, 0x16, 0x86, 0x43, 0x8a, 0x09, 0x3a,
	0x80, 0x9e, 0x15, 0xde, 0x52, 0xd9, 0x8e, 0xb7, 0x1d, 0x7d, 0x70, 0x31, 0x59, 0xb0, 0xbb, 0x93,
	0xd5, 0x2c, 0x92, 0x54, 0x47, 0xe1, 0xcf, 0xdd, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x1d,
	0xce, 0xde, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReferentialServiceClient is the client API for ReferentialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReferentialServiceClient interface {
	List(ctx context.Context, in *ReferentialReq, opts ...grpc.CallOption) (*ReferentialListResp, error)
	ListStream(ctx context.Context, in *ReferentialReq, opts ...grpc.CallOption) (ReferentialService_ListStreamClient, error)
}

type referentialServiceClient struct {
	cc *grpc.ClientConn
}

func NewReferentialServiceClient(cc *grpc.ClientConn) ReferentialServiceClient {
	return &referentialServiceClient{cc}
}

func (c *referentialServiceClient) List(ctx context.Context, in *ReferentialReq, opts ...grpc.CallOption) (*ReferentialListResp, error) {
	out := new(ReferentialListResp)
	err := c.cc.Invoke(ctx, "/omsapi.pkg.utl.grpc.referential.ReferentialService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referentialServiceClient) ListStream(ctx context.Context, in *ReferentialReq, opts ...grpc.CallOption) (ReferentialService_ListStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ReferentialService_serviceDesc.Streams[0], "/omsapi.pkg.utl.grpc.referential.ReferentialService/ListStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &referentialServiceListStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReferentialService_ListStreamClient interface {
	Recv() (*ReferentialResp, error)
	grpc.ClientStream
}

type referentialServiceListStreamClient struct {
	grpc.ClientStream
}

func (x *referentialServiceListStreamClient) Recv() (*ReferentialResp, error) {
	m := new(ReferentialResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReferentialServiceServer is the server API for ReferentialService service.
type ReferentialServiceServer interface {
	List(context.Context, *ReferentialReq) (*ReferentialListResp, error)
	ListStream(*ReferentialReq, ReferentialService_ListStreamServer) error
}

// UnimplementedReferentialServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReferentialServiceServer struct {
}

func (*UnimplementedReferentialServiceServer) List(ctx context.Context, req *ReferentialReq) (*ReferentialListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedReferentialServiceServer) ListStream(req *ReferentialReq, srv ReferentialService_ListStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStream not implemented")
}

func RegisterReferentialServiceServer(s *grpc.Server, srv ReferentialServiceServer) {
	s.RegisterService(&_ReferentialService_serviceDesc, srv)
}

func _ReferentialService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferentialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferentialServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omsapi.pkg.utl.grpc.referential.ReferentialService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferentialServiceServer).List(ctx, req.(*ReferentialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferentialService_ListStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReferentialReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReferentialServiceServer).ListStream(m, &referentialServiceListStreamServer{stream})
}

type ReferentialService_ListStreamServer interface {
	Send(*ReferentialResp) error
	grpc.ServerStream
}

type referentialServiceListStreamServer struct {
	grpc.ServerStream
}

func (x *referentialServiceListStreamServer) Send(m *ReferentialResp) error {
	return x.ServerStream.SendMsg(m)
}

var _ReferentialService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "omsapi.pkg.utl.grpc.referential.ReferentialService",
	HandlerType: (*ReferentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ReferentialService_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStream",
			Handler:       _ReferentialService_ListStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "referential.proto",
}
