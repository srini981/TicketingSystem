// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: rbac.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RbacService_CreateRole_FullMethodName      = "/rbacService.rbacService/createRole"
	RbacService_GetRole_FullMethodName         = "/rbacService.rbacService/GetRole"
	RbacService_UpdateRole_FullMethodName      = "/rbacService.rbacService/updateRole"
	RbacService_DeleteRole_FullMethodName      = "/rbacService.rbacService/deleteRole"
	RbacService_AssignRole_FullMethodName      = "/rbacService.rbacService/assignRole"
	RbacService_UnAssignRole_FullMethodName    = "/rbacService.rbacService/unAssignRole"
	RbacService_GetAllRoles_FullMethodName     = "/rbacService.rbacService/GetAllRoles"
	RbacService_GetAllUserRoles_FullMethodName = "/rbacService.rbacService/GetAllUserRoles"
)

// RbacServiceClient is the client API for RbacService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RbacServiceClient interface {
	CreateRole(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	GetRole(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleResponse, error)
	UpdateRole(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	DeleteRole(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AssignRole(ctx context.Context, in *AssignUser, opts ...grpc.CallOption) (*RoleID, error)
	UnAssignRole(ctx context.Context, in *AssignUser, opts ...grpc.CallOption) (*RoleID, error)
	GetAllRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Roles, error)
	GetAllUserRoles(ctx context.Context, in *AssignUser, opts ...grpc.CallOption) (*Roles, error)
}

type rbacServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRbacServiceClient(cc grpc.ClientConnInterface) RbacServiceClient {
	return &rbacServiceClient{cc}
}

func (c *rbacServiceClient) CreateRole(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RbacService_CreateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetRole(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*RoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RbacService_GetRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) UpdateRole(ctx context.Context, in *RoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, RbacService_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) DeleteRole(ctx context.Context, in *RoleID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RbacService_DeleteRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) AssignRole(ctx context.Context, in *AssignUser, opts ...grpc.CallOption) (*RoleID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleID)
	err := c.cc.Invoke(ctx, RbacService_AssignRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) UnAssignRole(ctx context.Context, in *AssignUser, opts ...grpc.CallOption) (*RoleID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleID)
	err := c.cc.Invoke(ctx, RbacService_UnAssignRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetAllRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Roles, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Roles)
	err := c.cc.Invoke(ctx, RbacService_GetAllRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacServiceClient) GetAllUserRoles(ctx context.Context, in *AssignUser, opts ...grpc.CallOption) (*Roles, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Roles)
	err := c.cc.Invoke(ctx, RbacService_GetAllUserRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RbacServiceServer is the server API for RbacService service.
// All implementations must embed UnimplementedRbacServiceServer
// for forward compatibility.
type RbacServiceServer interface {
	CreateRole(context.Context, *RoleRequest) (*RoleResponse, error)
	GetRole(context.Context, *RoleID) (*RoleResponse, error)
	UpdateRole(context.Context, *RoleRequest) (*RoleResponse, error)
	DeleteRole(context.Context, *RoleID) (*emptypb.Empty, error)
	AssignRole(context.Context, *AssignUser) (*RoleID, error)
	UnAssignRole(context.Context, *AssignUser) (*RoleID, error)
	GetAllRoles(context.Context, *emptypb.Empty) (*Roles, error)
	GetAllUserRoles(context.Context, *AssignUser) (*Roles, error)
	mustEmbedUnimplementedRbacServiceServer()
}

// UnimplementedRbacServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRbacServiceServer struct{}

func (UnimplementedRbacServiceServer) CreateRole(context.Context, *RoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRbacServiceServer) GetRole(context.Context, *RoleID) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRbacServiceServer) UpdateRole(context.Context, *RoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRbacServiceServer) DeleteRole(context.Context, *RoleID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRbacServiceServer) AssignRole(context.Context, *AssignUser) (*RoleID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignRole not implemented")
}
func (UnimplementedRbacServiceServer) UnAssignRole(context.Context, *AssignUser) (*RoleID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnAssignRole not implemented")
}
func (UnimplementedRbacServiceServer) GetAllRoles(context.Context, *emptypb.Empty) (*Roles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoles not implemented")
}
func (UnimplementedRbacServiceServer) GetAllUserRoles(context.Context, *AssignUser) (*Roles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserRoles not implemented")
}
func (UnimplementedRbacServiceServer) mustEmbedUnimplementedRbacServiceServer() {}
func (UnimplementedRbacServiceServer) testEmbeddedByValue()                     {}

// UnsafeRbacServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RbacServiceServer will
// result in compilation errors.
type UnsafeRbacServiceServer interface {
	mustEmbedUnimplementedRbacServiceServer()
}

func RegisterRbacServiceServer(s grpc.ServiceRegistrar, srv RbacServiceServer) {
	// If the following call pancis, it indicates UnimplementedRbacServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RbacService_ServiceDesc, srv)
}

func _RbacService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).CreateRole(ctx, req.(*RoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetRole(ctx, req.(*RoleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).UpdateRole(ctx, req.(*RoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).DeleteRole(ctx, req.(*RoleID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_AssignRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).AssignRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_AssignRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).AssignRole(ctx, req.(*AssignUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_UnAssignRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).UnAssignRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_UnAssignRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).UnAssignRole(ctx, req.(*AssignUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetAllRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetAllRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_GetAllRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetAllRoles(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacService_GetAllUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacServiceServer).GetAllUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RbacService_GetAllUserRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacServiceServer).GetAllUserRoles(ctx, req.(*AssignUser))
	}
	return interceptor(ctx, in, info, handler)
}

// RbacService_ServiceDesc is the grpc.ServiceDesc for RbacService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RbacService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rbacService.rbacService",
	HandlerType: (*RbacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createRole",
			Handler:    _RbacService_CreateRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _RbacService_GetRole_Handler,
		},
		{
			MethodName: "updateRole",
			Handler:    _RbacService_UpdateRole_Handler,
		},
		{
			MethodName: "deleteRole",
			Handler:    _RbacService_DeleteRole_Handler,
		},
		{
			MethodName: "assignRole",
			Handler:    _RbacService_AssignRole_Handler,
		},
		{
			MethodName: "unAssignRole",
			Handler:    _RbacService_UnAssignRole_Handler,
		},
		{
			MethodName: "GetAllRoles",
			Handler:    _RbacService_GetAllRoles_Handler,
		},
		{
			MethodName: "GetAllUserRoles",
			Handler:    _RbacService_GetAllUserRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rbac.proto",
}
