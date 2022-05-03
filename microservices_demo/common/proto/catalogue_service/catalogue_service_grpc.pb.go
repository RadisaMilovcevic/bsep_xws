// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: catalogue_service.proto

package catalogue

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

// CatalogueServiceClient is the client API for CatalogueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogueServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type catalogueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogueServiceClient(cc grpc.ClientConnInterface) CatalogueServiceClient {
	return &catalogueServiceClient{cc}
}

func (c *catalogueServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/catalogue.CatalogueService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogueServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/catalogue.CatalogueService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogueServiceServer is the server API for CatalogueService service.
// All implementations must embed UnimplementedCatalogueServiceServer
// for forward compatibility
type CatalogueServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedCatalogueServiceServer()
}

// UnimplementedCatalogueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogueServiceServer struct {
}

func (UnimplementedCatalogueServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCatalogueServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCatalogueServiceServer) mustEmbedUnimplementedCatalogueServiceServer() {}

// UnsafeCatalogueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogueServiceServer will
// result in compilation errors.
type UnsafeCatalogueServiceServer interface {
	mustEmbedUnimplementedCatalogueServiceServer()
}

func RegisterCatalogueServiceServer(s grpc.ServiceRegistrar, srv CatalogueServiceServer) {
	s.RegisterService(&CatalogueService_ServiceDesc, srv)
}

func _CatalogueService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogue.CatalogueService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogueService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogueServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalogue.CatalogueService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogueServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogueService_ServiceDesc is the grpc.ServiceDesc for CatalogueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalogue.CatalogueService",
	HandlerType: (*CatalogueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CatalogueService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CatalogueService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalogue_service.proto",
}
