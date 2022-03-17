// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: registry/tx.proto

package registry

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	RegisterService(ctx context.Context, in *MsgRegisterService, opts ...grpc.CallOption) (*MsgRegisterServiceResponse, error)
	RegisterName(ctx context.Context, in *MsgRegisterName, opts ...grpc.CallOption) (*MsgRegisterNameResponse, error)
	CheckName(ctx context.Context, in *MsgCheckName, opts ...grpc.CallOption) (*MsgCheckNameResponse, error)
	AccessName(ctx context.Context, in *MsgAccessName, opts ...grpc.CallOption) (*MsgAccessNameResponse, error)
	UpdateName(ctx context.Context, in *MsgUpdateName, opts ...grpc.CallOption) (*MsgUpdateNameResponse, error)
	AccessService(ctx context.Context, in *MsgAccessService, opts ...grpc.CallOption) (*MsgAccessServiceResponse, error)
	UpdateService(ctx context.Context, in *MsgUpdateService, opts ...grpc.CallOption) (*MsgUpdateServiceResponse, error)
	CreateWhoIs(ctx context.Context, in *MsgCreateWhoIs, opts ...grpc.CallOption) (*MsgCreateWhoIsResponse, error)
	UpdateWhoIs(ctx context.Context, in *MsgUpdateWhoIs, opts ...grpc.CallOption) (*MsgUpdateWhoIsResponse, error)
	DeleteWhoIs(ctx context.Context, in *MsgDeleteWhoIs, opts ...grpc.CallOption) (*MsgDeleteWhoIsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterService(ctx context.Context, in *MsgRegisterService, opts ...grpc.CallOption) (*MsgRegisterServiceResponse, error) {
	out := new(MsgRegisterServiceResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterName(ctx context.Context, in *MsgRegisterName, opts ...grpc.CallOption) (*MsgRegisterNameResponse, error) {
	out := new(MsgRegisterNameResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/RegisterName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CheckName(ctx context.Context, in *MsgCheckName, opts ...grpc.CallOption) (*MsgCheckNameResponse, error) {
	out := new(MsgCheckNameResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/CheckName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AccessName(ctx context.Context, in *MsgAccessName, opts ...grpc.CallOption) (*MsgAccessNameResponse, error) {
	out := new(MsgAccessNameResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/AccessName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateName(ctx context.Context, in *MsgUpdateName, opts ...grpc.CallOption) (*MsgUpdateNameResponse, error) {
	out := new(MsgUpdateNameResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AccessService(ctx context.Context, in *MsgAccessService, opts ...grpc.CallOption) (*MsgAccessServiceResponse, error) {
	out := new(MsgAccessServiceResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/AccessService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateService(ctx context.Context, in *MsgUpdateService, opts ...grpc.CallOption) (*MsgUpdateServiceResponse, error) {
	out := new(MsgUpdateServiceResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/UpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateWhoIs(ctx context.Context, in *MsgCreateWhoIs, opts ...grpc.CallOption) (*MsgCreateWhoIsResponse, error) {
	out := new(MsgCreateWhoIsResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/CreateWhoIs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateWhoIs(ctx context.Context, in *MsgUpdateWhoIs, opts ...grpc.CallOption) (*MsgUpdateWhoIsResponse, error) {
	out := new(MsgUpdateWhoIsResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/UpdateWhoIs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteWhoIs(ctx context.Context, in *MsgDeleteWhoIs, opts ...grpc.CallOption) (*MsgDeleteWhoIsResponse, error) {
	out := new(MsgDeleteWhoIsResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.registry.Msg/DeleteWhoIs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations should embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	RegisterService(context.Context, *MsgRegisterService) (*MsgRegisterServiceResponse, error)
	RegisterName(context.Context, *MsgRegisterName) (*MsgRegisterNameResponse, error)
	CheckName(context.Context, *MsgCheckName) (*MsgCheckNameResponse, error)
	AccessName(context.Context, *MsgAccessName) (*MsgAccessNameResponse, error)
	UpdateName(context.Context, *MsgUpdateName) (*MsgUpdateNameResponse, error)
	AccessService(context.Context, *MsgAccessService) (*MsgAccessServiceResponse, error)
	UpdateService(context.Context, *MsgUpdateService) (*MsgUpdateServiceResponse, error)
	CreateWhoIs(context.Context, *MsgCreateWhoIs) (*MsgCreateWhoIsResponse, error)
	UpdateWhoIs(context.Context, *MsgUpdateWhoIs) (*MsgUpdateWhoIsResponse, error)
	DeleteWhoIs(context.Context, *MsgDeleteWhoIs) (*MsgDeleteWhoIsResponse, error)
}

// UnimplementedMsgServer should be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RegisterService(context.Context, *MsgRegisterService) (*MsgRegisterServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedMsgServer) RegisterName(context.Context, *MsgRegisterName) (*MsgRegisterNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterName not implemented")
}
func (UnimplementedMsgServer) CheckName(context.Context, *MsgCheckName) (*MsgCheckNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckName not implemented")
}
func (UnimplementedMsgServer) AccessName(context.Context, *MsgAccessName) (*MsgAccessNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessName not implemented")
}
func (UnimplementedMsgServer) UpdateName(context.Context, *MsgUpdateName) (*MsgUpdateNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedMsgServer) AccessService(context.Context, *MsgAccessService) (*MsgAccessServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessService not implemented")
}
func (UnimplementedMsgServer) UpdateService(context.Context, *MsgUpdateService) (*MsgUpdateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedMsgServer) CreateWhoIs(context.Context, *MsgCreateWhoIs) (*MsgCreateWhoIsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWhoIs not implemented")
}
func (UnimplementedMsgServer) UpdateWhoIs(context.Context, *MsgUpdateWhoIs) (*MsgUpdateWhoIsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWhoIs not implemented")
}
func (UnimplementedMsgServer) DeleteWhoIs(context.Context, *MsgDeleteWhoIs) (*MsgDeleteWhoIsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWhoIs not implemented")
}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterService(ctx, req.(*MsgRegisterService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/RegisterName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterName(ctx, req.(*MsgRegisterName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CheckName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCheckName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CheckName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/CheckName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CheckName(ctx, req.(*MsgCheckName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AccessName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAccessName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AccessName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/AccessName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AccessName(ctx, req.(*MsgAccessName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateName(ctx, req.(*MsgUpdateName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AccessService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAccessService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AccessService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/AccessService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AccessService(ctx, req.(*MsgAccessService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateService(ctx, req.(*MsgUpdateService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateWhoIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateWhoIs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateWhoIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/CreateWhoIs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateWhoIs(ctx, req.(*MsgCreateWhoIs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateWhoIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateWhoIs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateWhoIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/UpdateWhoIs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateWhoIs(ctx, req.(*MsgUpdateWhoIs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteWhoIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteWhoIs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteWhoIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.registry.Msg/DeleteWhoIs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteWhoIs(ctx, req.(*MsgDeleteWhoIs))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sonrio.sonr.registry.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _Msg_RegisterService_Handler,
		},
		{
			MethodName: "RegisterName",
			Handler:    _Msg_RegisterName_Handler,
		},
		{
			MethodName: "CheckName",
			Handler:    _Msg_CheckName_Handler,
		},
		{
			MethodName: "AccessName",
			Handler:    _Msg_AccessName_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _Msg_UpdateName_Handler,
		},
		{
			MethodName: "AccessService",
			Handler:    _Msg_AccessService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _Msg_UpdateService_Handler,
		},
		{
			MethodName: "CreateWhoIs",
			Handler:    _Msg_CreateWhoIs_Handler,
		},
		{
			MethodName: "UpdateWhoIs",
			Handler:    _Msg_UpdateWhoIs_Handler,
		},
		{
			MethodName: "DeleteWhoIs",
			Handler:    _Msg_DeleteWhoIs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/tx.proto",
}