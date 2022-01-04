// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product-info.proto

package product

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_info_2607e15ce5d0eb36, []int{0}
}
func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (dst *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(dst, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ProductId struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductId) Reset()         { *m = ProductId{} }
func (m *ProductId) String() string { return proto.CompactTextString(m) }
func (*ProductId) ProtoMessage()    {}
func (*ProductId) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_info_2607e15ce5d0eb36, []int{1}
}
func (m *ProductId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductId.Unmarshal(m, b)
}
func (m *ProductId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductId.Marshal(b, m, deterministic)
}
func (dst *ProductId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductId.Merge(dst, src)
}
func (m *ProductId) XXX_Size() int {
	return xxx_messageInfo_ProductId.Size(m)
}
func (m *ProductId) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductId.DiscardUnknown(m)
}

var xxx_messageInfo_ProductId proto.InternalMessageInfo

func (m *ProductId) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Product)(nil), "product.Product")
	proto.RegisterType((*ProductId)(nil), "product.ProductId")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductInfoClient is the client API for ProductInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductInfoClient interface {
	// 添加商品
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductId, error)
	// 获取商品
	GetProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error)
}

type productInfoClient struct {
	cc *grpc.ClientConn
}

func NewProductInfoClient(cc *grpc.ClientConn) ProductInfoClient {
	return &productInfoClient{cc}
}

func (c *productInfoClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductId, error) {
	out := new(ProductId)
	err := c.cc.Invoke(ctx, "/product.ProductInfo/addProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productInfoClient) GetProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/product.ProductInfo/getProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInfoServer is the server API for ProductInfo service.
type ProductInfoServer interface {
	// 添加商品
	AddProduct(context.Context, *Product) (*ProductId, error)
	// 获取商品
	GetProduct(context.Context, *ProductId) (*Product, error)
}

func RegisterProductInfoServer(s *grpc.Server, srv ProductInfoServer) {
	s.RegisterService(&_ProductInfo_serviceDesc, srv)
}

func _ProductInfo_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductInfo/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductInfo_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductInfo/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).GetProduct(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductInfo",
	HandlerType: (*ProductInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addProduct",
			Handler:    _ProductInfo_AddProduct_Handler,
		},
		{
			MethodName: "getProduct",
			Handler:    _ProductInfo_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product-info.proto",
}

func init() { proto.RegisterFile("product-info.proto", fileDescriptor_product_info_2607e15ce5d0eb36) }

var fileDescriptor_product_info_2607e15ce5d0eb36 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0xcd, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x87, 0x8a, 0x29, 0xf9, 0x73, 0xb1, 0x07, 0x40, 0x98, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x4c, 0x60, 0x11, 0x30, 0x5b, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5, 0x38, 0xb9, 0x28,
	0xb3, 0xa0, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x19, 0x2c, 0x85, 0x2c, 0xa4, 0xa4, 0xc8, 0xc5, 0x09,
	0x35, 0xd0, 0x33, 0x45, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0x6a, 0x2a, 0x84,
	0x63, 0x54, 0xca, 0xc5, 0x0d, 0x53, 0x92, 0x97, 0x96, 0x2f, 0x64, 0xc4, 0xc5, 0x95, 0x98, 0x92,
	0x02, 0x73, 0x85, 0x80, 0x1e, 0xd4, 0x69, 0x7a, 0x50, 0x11, 0x29, 0x21, 0x74, 0x11, 0xcf, 0x14,
	0x90, 0x9e, 0xf4, 0xd4, 0x12, 0x98, 0x1e, 0x2c, 0x2a, 0xa4, 0x30, 0xcc, 0x49, 0x62, 0x03, 0x7b,
	0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xe7, 0xc1, 0xad, 0x10, 0x01, 0x00, 0x00,
}