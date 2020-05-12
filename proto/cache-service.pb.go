// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache-service.proto

package cacheService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Item is what is stored in the cache
type Item struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expiration           string   `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Item) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Item) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

type GetKey struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKey) Reset()         { *m = GetKey{} }
func (m *GetKey) String() string { return proto.CompactTextString(m) }
func (*GetKey) ProtoMessage()    {}
func (*GetKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{1}
}

func (m *GetKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKey.Unmarshal(m, b)
}
func (m *GetKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKey.Marshal(b, m, deterministic)
}
func (m *GetKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKey.Merge(m, src)
}
func (m *GetKey) XXX_Size() int {
	return xxx_messageInfo_GetKey.Size(m)
}
func (m *GetKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKey.DiscardUnknown(m)
}

var xxx_messageInfo_GetKey proto.InternalMessageInfo

func (m *GetKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type AllItems struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllItems) Reset()         { *m = AllItems{} }
func (m *AllItems) String() string { return proto.CompactTextString(m) }
func (*AllItems) ProtoMessage()    {}
func (*AllItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{2}
}

func (m *AllItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllItems.Unmarshal(m, b)
}
func (m *AllItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllItems.Marshal(b, m, deterministic)
}
func (m *AllItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllItems.Merge(m, src)
}
func (m *AllItems) XXX_Size() int {
	return xxx_messageInfo_AllItems.Size(m)
}
func (m *AllItems) XXX_DiscardUnknown() {
	xxx_messageInfo_AllItems.DiscardUnknown(m)
}

var xxx_messageInfo_AllItems proto.InternalMessageInfo

func (m *AllItems) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Success struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Success) Reset()         { *m = Success{} }
func (m *Success) String() string { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()    {}
func (*Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b0ee6ef8b54c4e4, []int{3}
}

func (m *Success) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Success.Unmarshal(m, b)
}
func (m *Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Success.Marshal(b, m, deterministic)
}
func (m *Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Success.Merge(m, src)
}
func (m *Success) XXX_Size() int {
	return xxx_messageInfo_Success.Size(m)
}
func (m *Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Success proto.InternalMessageInfo

func (m *Success) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Item)(nil), "cacheService.Item")
	proto.RegisterType((*GetKey)(nil), "cacheService.GetKey")
	proto.RegisterType((*AllItems)(nil), "cacheService.AllItems")
	proto.RegisterType((*Success)(nil), "cacheService.Success")
}

func init() { proto.RegisterFile("cache-service.proto", fileDescriptor_9b0ee6ef8b54c4e4) }

var fileDescriptor_9b0ee6ef8b54c4e4 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0x8d, 0xfd, 0x9b, 0xf6, 0x20, 0x63, 0x2d, 0x21, 0x82, 0x94, 0x78, 0xc9, 0xa5, 0x1b,
	0xa8, 0x1e, 0x44, 0xf0, 0x10, 0x54, 0x82, 0x08, 0x1e, 0xd2, 0x27, 0x48, 0xd3, 0xb1, 0x06, 0xb7,
	0x26, 0x24, 0x9b, 0x62, 0x9e, 0xc7, 0x17, 0x95, 0xdd, 0x25, 0x25, 0xc1, 0xf4, 0x36, 0xfb, 0xcd,
	0xf7, 0xb3, 0xf3, 0xc1, 0x45, 0x1c, 0xc5, 0x9f, 0xb4, 0x2c, 0x28, 0x3f, 0x24, 0x31, 0xb1, 0x2c,
	0x4f, 0x45, 0x8a, 0x53, 0x05, 0xae, 0x35, 0x66, 0x5f, 0xed, 0xd2, 0x74, 0xc7, 0xc9, 0x53, 0xbb,
	0x4d, 0xf9, 0xe1, 0xd1, 0x3e, 0x13, 0x95, 0xa6, 0x3a, 0xef, 0x70, 0xf6, 0x2a, 0x68, 0x8f, 0xe7,
	0x60, 0x7e, 0x51, 0x65, 0x19, 0x0b, 0xc3, 0x1d, 0x87, 0x72, 0xc4, 0x19, 0xf4, 0x0f, 0x11, 0x2f,
	0xc9, 0xea, 0x29, 0x4c, 0x3f, 0xf0, 0x1a, 0x80, 0x7e, 0xb2, 0x24, 0x8f, 0x44, 0x92, 0x7e, 0x5b,
	0xa6, 0x5a, 0x35, 0x10, 0xc7, 0x86, 0x41, 0x40, 0xe2, 0x8d, 0xaa, 0xff, 0x8e, 0xce, 0x1d, 0x8c,
	0x7c, 0xce, 0x65, 0x5c, 0x81, 0x2e, 0xf4, 0x13, 0x39, 0x58, 0xc6, 0xc2, 0x74, 0x27, 0x2b, 0x64,
	0xcd, 0x2f, 0x33, 0xc9, 0x09, 0x35, 0xc1, 0xb9, 0x81, 0xe1, 0xba, 0x8c, 0x63, 0x2a, 0x0a, 0xb4,
	0x60, 0x58, 0xe8, 0x51, 0xd9, 0x8e, 0xc2, 0xfa, 0xb9, 0xfa, 0xed, 0xc1, 0xf4, 0xa9, 0xe1, 0x80,
	0x4b, 0x30, 0xfd, 0xed, 0x16, 0x3b, 0x7c, 0xed, 0x0e, 0x0c, 0x3d, 0x30, 0x03, 0x12, 0x38, 0x6b,
	0xaf, 0xf4, 0x25, 0x9d, 0x82, 0x47, 0x98, 0x04, 0x24, 0x8e, 0xe7, 0xcc, 0x99, 0x2e, 0x99, 0xd5,
	0x25, 0xb3, 0x17, 0x59, 0xb2, 0x3d, 0x6f, 0x4b, 0x8f, 0xfc, 0x7b, 0x18, 0x3f, 0x13, 0x27, 0x41,
	0xb2, 0xa9, 0xee, 0xd4, 0xcb, 0x36, 0x5a, 0x77, 0xf0, 0x50, 0x2b, 0x7d, 0xce, 0x4f, 0xc6, 0x76,
	0x6b, 0x37, 0x03, 0x45, 0xbb, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x16, 0xaf, 0x1e, 0x35,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheServiceClient interface {
	Add(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	Get(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Item, error)
	GetAllItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllItems, error)
	DeleteKey(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Success, error)
	DeleteAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Success, error)
}

type cacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewCacheServiceClient(cc *grpc.ClientConn) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Add(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Get(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetAllItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllItems, error) {
	out := new(AllItems)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteKey(ctx context.Context, in *GetKey, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/DeleteKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/cacheService.CacheService/DeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
type CacheServiceServer interface {
	Add(context.Context, *Item) (*Item, error)
	Get(context.Context, *GetKey) (*Item, error)
	GetAllItems(context.Context, *empty.Empty) (*AllItems, error)
	DeleteKey(context.Context, *GetKey) (*Success, error)
	DeleteAll(context.Context, *empty.Empty) (*Success, error)
}

// UnimplementedCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (*UnimplementedCacheServiceServer) Add(ctx context.Context, req *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCacheServiceServer) Get(ctx context.Context, req *GetKey) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCacheServiceServer) GetAllItems(ctx context.Context, req *empty.Empty) (*AllItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}
func (*UnimplementedCacheServiceServer) DeleteKey(ctx context.Context, req *GetKey) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (*UnimplementedCacheServiceServer) DeleteAll(ctx context.Context, req *empty.Empty) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAll not implemented")
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Add(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*GetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetAllItems(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/DeleteKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).DeleteKey(ctx, req.(*GetKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).DeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacheService.CacheService/DeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).DeleteAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cacheService.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CacheService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
		{
			MethodName: "GetAllItems",
			Handler:    _CacheService_GetAllItems_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _CacheService_DeleteKey_Handler,
		},
		{
			MethodName: "DeleteAll",
			Handler:    _CacheService_DeleteAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache-service.proto",
}