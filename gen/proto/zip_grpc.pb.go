// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/zip.proto

package proto

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

const (
	ZipProtoService_GetZipInfo_FullMethodName     = "/zippb.ZipProtoService/GetZipInfo"
	ZipProtoService_GetZipAll_FullMethodName      = "/zippb.ZipProtoService/GetZipAll"
	ZipProtoService_GetZipShowYes_FullMethodName  = "/zippb.ZipProtoService/GetZipShowYes"
	ZipProtoService_GetZipShowNo_FullMethodName   = "/zippb.ZipProtoService/GetZipShowNo"
	ZipProtoService_GetZipByAgent_FullMethodName  = "/zippb.ZipProtoService/GetZipByAgent"
	ZipProtoService_GetZipByEstate_FullMethodName = "/zippb.ZipProtoService/GetZipByEstate"
	ZipProtoService_SaveZip_FullMethodName        = "/zippb.ZipProtoService/SaveZip"
	ZipProtoService_UpdateZip_FullMethodName      = "/zippb.ZipProtoService/UpdateZip"
	ZipProtoService_DeleteZip_FullMethodName      = "/zippb.ZipProtoService/DeleteZip"
	ZipProtoService_SearchZip_FullMethodName      = "/zippb.ZipProtoService/SearchZip"
)

// ZipProtoServiceClient is the client API for ZipProtoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZipProtoServiceClient interface {
	GetZipInfo(ctx context.Context, in *ZipIdRequest, opts ...grpc.CallOption) (*ZipIdResponse, error)
	GetZipAll(ctx context.Context, in *ZipAllRequest, opts ...grpc.CallOption) (*ZipAllResponse, error)
	GetZipShowYes(ctx context.Context, in *ZipShowYesRequest, opts ...grpc.CallOption) (*ZipShowYesResponse, error)
	GetZipShowNo(ctx context.Context, in *ZipShowNoRequest, opts ...grpc.CallOption) (*ZipShowNoResponse, error)
	GetZipByAgent(ctx context.Context, in *ZipAgentIdRequest, opts ...grpc.CallOption) (*ZipAgentIdResponse, error)
	GetZipByEstate(ctx context.Context, in *ZipEsateIdRequest, opts ...grpc.CallOption) (*ZipEstateIdResponse, error)
	SaveZip(ctx context.Context, in *ZipInsertRequest, opts ...grpc.CallOption) (*ZipInsertResponse, error)
	UpdateZip(ctx context.Context, in *ZipUpdateRequest, opts ...grpc.CallOption) (*ZipUpdateResponse, error)
	DeleteZip(ctx context.Context, in *ZipDeleteRequest, opts ...grpc.CallOption) (*ZipDeleteResponse, error)
	SearchZip(ctx context.Context, in *ZipSearchRequest, opts ...grpc.CallOption) (*ZipSearchResponse, error)
}

type zipProtoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZipProtoServiceClient(cc grpc.ClientConnInterface) ZipProtoServiceClient {
	return &zipProtoServiceClient{cc}
}

func (c *zipProtoServiceClient) GetZipInfo(ctx context.Context, in *ZipIdRequest, opts ...grpc.CallOption) (*ZipIdResponse, error) {
	out := new(ZipIdResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_GetZipInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) GetZipAll(ctx context.Context, in *ZipAllRequest, opts ...grpc.CallOption) (*ZipAllResponse, error) {
	out := new(ZipAllResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_GetZipAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) GetZipShowYes(ctx context.Context, in *ZipShowYesRequest, opts ...grpc.CallOption) (*ZipShowYesResponse, error) {
	out := new(ZipShowYesResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_GetZipShowYes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) GetZipShowNo(ctx context.Context, in *ZipShowNoRequest, opts ...grpc.CallOption) (*ZipShowNoResponse, error) {
	out := new(ZipShowNoResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_GetZipShowNo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) GetZipByAgent(ctx context.Context, in *ZipAgentIdRequest, opts ...grpc.CallOption) (*ZipAgentIdResponse, error) {
	out := new(ZipAgentIdResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_GetZipByAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) GetZipByEstate(ctx context.Context, in *ZipEsateIdRequest, opts ...grpc.CallOption) (*ZipEstateIdResponse, error) {
	out := new(ZipEstateIdResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_GetZipByEstate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) SaveZip(ctx context.Context, in *ZipInsertRequest, opts ...grpc.CallOption) (*ZipInsertResponse, error) {
	out := new(ZipInsertResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_SaveZip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) UpdateZip(ctx context.Context, in *ZipUpdateRequest, opts ...grpc.CallOption) (*ZipUpdateResponse, error) {
	out := new(ZipUpdateResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_UpdateZip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) DeleteZip(ctx context.Context, in *ZipDeleteRequest, opts ...grpc.CallOption) (*ZipDeleteResponse, error) {
	out := new(ZipDeleteResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_DeleteZip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zipProtoServiceClient) SearchZip(ctx context.Context, in *ZipSearchRequest, opts ...grpc.CallOption) (*ZipSearchResponse, error) {
	out := new(ZipSearchResponse)
	err := c.cc.Invoke(ctx, ZipProtoService_SearchZip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZipProtoServiceServer is the server API for ZipProtoService service.
// All implementations must embed UnimplementedZipProtoServiceServer
// for forward compatibility
type ZipProtoServiceServer interface {
	GetZipInfo(context.Context, *ZipIdRequest) (*ZipIdResponse, error)
	GetZipAll(context.Context, *ZipAllRequest) (*ZipAllResponse, error)
	GetZipShowYes(context.Context, *ZipShowYesRequest) (*ZipShowYesResponse, error)
	GetZipShowNo(context.Context, *ZipShowNoRequest) (*ZipShowNoResponse, error)
	GetZipByAgent(context.Context, *ZipAgentIdRequest) (*ZipAgentIdResponse, error)
	GetZipByEstate(context.Context, *ZipEsateIdRequest) (*ZipEstateIdResponse, error)
	SaveZip(context.Context, *ZipInsertRequest) (*ZipInsertResponse, error)
	UpdateZip(context.Context, *ZipUpdateRequest) (*ZipUpdateResponse, error)
	DeleteZip(context.Context, *ZipDeleteRequest) (*ZipDeleteResponse, error)
	SearchZip(context.Context, *ZipSearchRequest) (*ZipSearchResponse, error)
	mustEmbedUnimplementedZipProtoServiceServer()
}

// UnimplementedZipProtoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedZipProtoServiceServer struct {
}

func (UnimplementedZipProtoServiceServer) GetZipInfo(context.Context, *ZipIdRequest) (*ZipIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZipInfo not implemented")
}
func (UnimplementedZipProtoServiceServer) GetZipAll(context.Context, *ZipAllRequest) (*ZipAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZipAll not implemented")
}
func (UnimplementedZipProtoServiceServer) GetZipShowYes(context.Context, *ZipShowYesRequest) (*ZipShowYesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZipShowYes not implemented")
}
func (UnimplementedZipProtoServiceServer) GetZipShowNo(context.Context, *ZipShowNoRequest) (*ZipShowNoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZipShowNo not implemented")
}
func (UnimplementedZipProtoServiceServer) GetZipByAgent(context.Context, *ZipAgentIdRequest) (*ZipAgentIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZipByAgent not implemented")
}
func (UnimplementedZipProtoServiceServer) GetZipByEstate(context.Context, *ZipEsateIdRequest) (*ZipEstateIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetZipByEstate not implemented")
}
func (UnimplementedZipProtoServiceServer) SaveZip(context.Context, *ZipInsertRequest) (*ZipInsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveZip not implemented")
}
func (UnimplementedZipProtoServiceServer) UpdateZip(context.Context, *ZipUpdateRequest) (*ZipUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateZip not implemented")
}
func (UnimplementedZipProtoServiceServer) DeleteZip(context.Context, *ZipDeleteRequest) (*ZipDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteZip not implemented")
}
func (UnimplementedZipProtoServiceServer) SearchZip(context.Context, *ZipSearchRequest) (*ZipSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchZip not implemented")
}
func (UnimplementedZipProtoServiceServer) mustEmbedUnimplementedZipProtoServiceServer() {}

// UnsafeZipProtoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZipProtoServiceServer will
// result in compilation errors.
type UnsafeZipProtoServiceServer interface {
	mustEmbedUnimplementedZipProtoServiceServer()
}

func RegisterZipProtoServiceServer(s grpc.ServiceRegistrar, srv ZipProtoServiceServer) {
	s.RegisterService(&ZipProtoService_ServiceDesc, srv)
}

func _ZipProtoService_GetZipInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).GetZipInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_GetZipInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).GetZipInfo(ctx, req.(*ZipIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_GetZipAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).GetZipAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_GetZipAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).GetZipAll(ctx, req.(*ZipAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_GetZipShowYes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipShowYesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).GetZipShowYes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_GetZipShowYes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).GetZipShowYes(ctx, req.(*ZipShowYesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_GetZipShowNo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipShowNoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).GetZipShowNo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_GetZipShowNo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).GetZipShowNo(ctx, req.(*ZipShowNoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_GetZipByAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipAgentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).GetZipByAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_GetZipByAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).GetZipByAgent(ctx, req.(*ZipAgentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_GetZipByEstate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipEsateIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).GetZipByEstate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_GetZipByEstate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).GetZipByEstate(ctx, req.(*ZipEsateIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_SaveZip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipInsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).SaveZip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_SaveZip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).SaveZip(ctx, req.(*ZipInsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_UpdateZip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).UpdateZip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_UpdateZip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).UpdateZip(ctx, req.(*ZipUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_DeleteZip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).DeleteZip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_DeleteZip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).DeleteZip(ctx, req.(*ZipDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZipProtoService_SearchZip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZipSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZipProtoServiceServer).SearchZip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZipProtoService_SearchZip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZipProtoServiceServer).SearchZip(ctx, req.(*ZipSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZipProtoService_ServiceDesc is the grpc.ServiceDesc for ZipProtoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZipProtoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zippb.ZipProtoService",
	HandlerType: (*ZipProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetZipInfo",
			Handler:    _ZipProtoService_GetZipInfo_Handler,
		},
		{
			MethodName: "GetZipAll",
			Handler:    _ZipProtoService_GetZipAll_Handler,
		},
		{
			MethodName: "GetZipShowYes",
			Handler:    _ZipProtoService_GetZipShowYes_Handler,
		},
		{
			MethodName: "GetZipShowNo",
			Handler:    _ZipProtoService_GetZipShowNo_Handler,
		},
		{
			MethodName: "GetZipByAgent",
			Handler:    _ZipProtoService_GetZipByAgent_Handler,
		},
		{
			MethodName: "GetZipByEstate",
			Handler:    _ZipProtoService_GetZipByEstate_Handler,
		},
		{
			MethodName: "SaveZip",
			Handler:    _ZipProtoService_SaveZip_Handler,
		},
		{
			MethodName: "UpdateZip",
			Handler:    _ZipProtoService_UpdateZip_Handler,
		},
		{
			MethodName: "DeleteZip",
			Handler:    _ZipProtoService_DeleteZip_Handler,
		},
		{
			MethodName: "SearchZip",
			Handler:    _ZipProtoService_SearchZip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/zip.proto",
}
