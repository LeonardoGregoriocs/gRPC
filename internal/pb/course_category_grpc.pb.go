// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	CreateCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateCategoryStreamClient, error)
	ListCategories(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*CategoryList, error)
	GetCategory(ctx context.Context, in *CategoryGetRequest, opts ...grpc.CallOption) (*Category, error)
	GetCategoryByName(ctx context.Context, in *CategoryGetRequestName, opts ...grpc.CallOption) (*Category, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) CreateCategoryStream(ctx context.Context, opts ...grpc.CallOption) (CategoryService_CreateCategoryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[0], "/pb.CategoryService/CreateCategoryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceCreateCategoryStreamClient{stream}
	return x, nil
}

type CategoryService_CreateCategoryStreamClient interface {
	Send(*CreateCategoryRequest) error
	CloseAndRecv() (*CategoryList, error)
	grpc.ClientStream
}

type categoryServiceCreateCategoryStreamClient struct {
	grpc.ClientStream
}

func (x *categoryServiceCreateCategoryStreamClient) Send(m *CreateCategoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceCreateCategoryStreamClient) CloseAndRecv() (*CategoryList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CategoryList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *categoryServiceClient) ListCategories(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*CategoryList, error) {
	out := new(CategoryList)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/ListCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategory(ctx context.Context, in *CategoryGetRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategoryByName(ctx context.Context, in *CategoryGetRequestName, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategoryByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
	CreateCategoryStream(CategoryService_CreateCategoryStreamServer) error
	ListCategories(context.Context, *Blank) (*CategoryList, error)
	GetCategory(context.Context, *CategoryGetRequest) (*Category, error)
	GetCategoryByName(context.Context, *CategoryGetRequestName) (*Category, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) CreateCategoryStream(CategoryService_CreateCategoryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCategoryStream not implemented")
}
func (UnimplementedCategoryServiceServer) ListCategories(context.Context, *Blank) (*CategoryList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategory(context.Context, *CategoryGetRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategoryByName(context.Context, *CategoryGetRequestName) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryByName not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_CreateCategoryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).CreateCategoryStream(&categoryServiceCreateCategoryStreamServer{stream})
}

type CategoryService_CreateCategoryStreamServer interface {
	SendAndClose(*CategoryList) error
	Recv() (*CreateCategoryRequest, error)
	grpc.ServerStream
}

type categoryServiceCreateCategoryStreamServer struct {
	grpc.ServerStream
}

func (x *categoryServiceCreateCategoryStreamServer) SendAndClose(m *CategoryList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceCreateCategoryStreamServer) Recv() (*CreateCategoryRequest, error) {
	m := new(CreateCategoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CategoryService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/ListCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).ListCategories(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategory(ctx, req.(*CategoryGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategoryByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryGetRequestName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategoryByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategoryByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategoryByName(ctx, req.(*CategoryGetRequestName))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryService_CreateCategory_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _CategoryService_ListCategories_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CategoryService_GetCategory_Handler,
		},
		{
			MethodName: "GetCategoryByName",
			Handler:    _CategoryService_GetCategoryByName_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCategoryStream",
			Handler:       _CategoryService_CreateCategoryStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/course_category.proto",
}
