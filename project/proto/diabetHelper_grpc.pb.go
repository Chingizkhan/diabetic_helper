// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: diabetHelper.proto

package diabetHelper

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
	DiabetHelper_AddSL_FullMethodName    = "/diabetHelper.DiabetHelper/AddSL"
	DiabetHelper_UpdateSL_FullMethodName = "/diabetHelper.DiabetHelper/UpdateSL"
	DiabetHelper_FindSL_FullMethodName   = "/diabetHelper.DiabetHelper/FindSL"
)

// DiabetHelperClient is the client API for DiabetHelper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiabetHelperClient interface {
	AddSL(ctx context.Context, in *AddSLRequest, opts ...grpc.CallOption) (*SugarLevel, error)
	UpdateSL(ctx context.Context, in *UpdateSLRequest, opts ...grpc.CallOption) (*SugarLevel, error)
	FindSL(ctx context.Context, in *FindSLRequest, opts ...grpc.CallOption) (*SugarLevels, error)
}

type diabetHelperClient struct {
	cc grpc.ClientConnInterface
}

func NewDiabetHelperClient(cc grpc.ClientConnInterface) DiabetHelperClient {
	return &diabetHelperClient{cc}
}

func (c *diabetHelperClient) AddSL(ctx context.Context, in *AddSLRequest, opts ...grpc.CallOption) (*SugarLevel, error) {
	out := new(SugarLevel)
	err := c.cc.Invoke(ctx, DiabetHelper_AddSL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diabetHelperClient) UpdateSL(ctx context.Context, in *UpdateSLRequest, opts ...grpc.CallOption) (*SugarLevel, error) {
	out := new(SugarLevel)
	err := c.cc.Invoke(ctx, DiabetHelper_UpdateSL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diabetHelperClient) FindSL(ctx context.Context, in *FindSLRequest, opts ...grpc.CallOption) (*SugarLevels, error) {
	out := new(SugarLevels)
	err := c.cc.Invoke(ctx, DiabetHelper_FindSL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiabetHelperServer is the server API for DiabetHelper service.
// All implementations must embed UnimplementedDiabetHelperServer
// for forward compatibility
type DiabetHelperServer interface {
	AddSL(context.Context, *AddSLRequest) (*SugarLevel, error)
	UpdateSL(context.Context, *UpdateSLRequest) (*SugarLevel, error)
	FindSL(context.Context, *FindSLRequest) (*SugarLevels, error)
	mustEmbedUnimplementedDiabetHelperServer()
}

// UnimplementedDiabetHelperServer must be embedded to have forward compatible implementations.
type UnimplementedDiabetHelperServer struct {
}

func (UnimplementedDiabetHelperServer) AddSL(context.Context, *AddSLRequest) (*SugarLevel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSL not implemented")
}
func (UnimplementedDiabetHelperServer) UpdateSL(context.Context, *UpdateSLRequest) (*SugarLevel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSL not implemented")
}
func (UnimplementedDiabetHelperServer) FindSL(context.Context, *FindSLRequest) (*SugarLevels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSL not implemented")
}
func (UnimplementedDiabetHelperServer) mustEmbedUnimplementedDiabetHelperServer() {}

// UnsafeDiabetHelperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiabetHelperServer will
// result in compilation errors.
type UnsafeDiabetHelperServer interface {
	mustEmbedUnimplementedDiabetHelperServer()
}

func RegisterDiabetHelperServer(s grpc.ServiceRegistrar, srv DiabetHelperServer) {
	s.RegisterService(&DiabetHelper_ServiceDesc, srv)
}

func _DiabetHelper_AddSL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiabetHelperServer).AddSL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiabetHelper_AddSL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiabetHelperServer).AddSL(ctx, req.(*AddSLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiabetHelper_UpdateSL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiabetHelperServer).UpdateSL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiabetHelper_UpdateSL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiabetHelperServer).UpdateSL(ctx, req.(*UpdateSLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiabetHelper_FindSL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiabetHelperServer).FindSL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiabetHelper_FindSL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiabetHelperServer).FindSL(ctx, req.(*FindSLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiabetHelper_ServiceDesc is the grpc.ServiceDesc for DiabetHelper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiabetHelper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "diabetHelper.DiabetHelper",
	HandlerType: (*DiabetHelperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSL",
			Handler:    _DiabetHelper_AddSL_Handler,
		},
		{
			MethodName: "UpdateSL",
			Handler:    _DiabetHelper_UpdateSL_Handler,
		},
		{
			MethodName: "FindSL",
			Handler:    _DiabetHelper_FindSL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diabetHelper.proto",
}
