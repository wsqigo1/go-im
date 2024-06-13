// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.24.4
// source: apps/social/rpc/social.proto

package social

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Social_FriendPutIn_FullMethodName       = "/user.social/FriendPutIn"
	Social_FriendPutInHandle_FullMethodName = "/user.social/FriendPutInHandle"
	Social_FriendPutInList_FullMethodName   = "/user.social/FriendPutInList"
	Social_FriendList_FullMethodName        = "/user.social/FriendList"
	Social_GroupCreate_FullMethodName       = "/user.social/GroupCreate"
	Social_GroupPutin_FullMethodName        = "/user.social/GroupPutin"
	Social_GroupPutinList_FullMethodName    = "/user.social/GroupPutinList"
	Social_GroupPutInHandle_FullMethodName  = "/user.social/GroupPutInHandle"
	Social_GroupList_FullMethodName         = "/user.social/GroupList"
	Social_GroupUsers_FullMethodName        = "/user.social/GroupUsers"
)

// SocialClient is the client API for Social service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// svc
type SocialClient interface {
	FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error)
	FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error)
	FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error)
	FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error)
	GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error)
	GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error)
	GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error)
	GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error)
	GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error)
	GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error)
}

type socialClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialClient(cc grpc.ClientConnInterface) SocialClient {
	return &socialClient{cc}
}

func (c *socialClient) FriendPutIn(ctx context.Context, in *FriendPutInReq, opts ...grpc.CallOption) (*FriendPutInResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FriendPutInResp)
	err := c.cc.Invoke(ctx, Social_FriendPutIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) FriendPutInHandle(ctx context.Context, in *FriendPutInHandleReq, opts ...grpc.CallOption) (*FriendPutInHandleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FriendPutInHandleResp)
	err := c.cc.Invoke(ctx, Social_FriendPutInHandle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) FriendPutInList(ctx context.Context, in *FriendPutInListReq, opts ...grpc.CallOption) (*FriendPutInListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FriendPutInListResp)
	err := c.cc.Invoke(ctx, Social_FriendPutInList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) FriendList(ctx context.Context, in *FriendListReq, opts ...grpc.CallOption) (*FriendListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FriendListResp)
	err := c.cc.Invoke(ctx, Social_FriendList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) GroupCreate(ctx context.Context, in *GroupCreateReq, opts ...grpc.CallOption) (*GroupCreateResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupCreateResp)
	err := c.cc.Invoke(ctx, Social_GroupCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) GroupPutin(ctx context.Context, in *GroupPutinReq, opts ...grpc.CallOption) (*GroupPutinResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupPutinResp)
	err := c.cc.Invoke(ctx, Social_GroupPutin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) GroupPutinList(ctx context.Context, in *GroupPutinListReq, opts ...grpc.CallOption) (*GroupPutinListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupPutinListResp)
	err := c.cc.Invoke(ctx, Social_GroupPutinList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) GroupPutInHandle(ctx context.Context, in *GroupPutInHandleReq, opts ...grpc.CallOption) (*GroupPutInHandleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupPutInHandleResp)
	err := c.cc.Invoke(ctx, Social_GroupPutInHandle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) GroupList(ctx context.Context, in *GroupListReq, opts ...grpc.CallOption) (*GroupListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupListResp)
	err := c.cc.Invoke(ctx, Social_GroupList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialClient) GroupUsers(ctx context.Context, in *GroupUsersReq, opts ...grpc.CallOption) (*GroupUsersResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupUsersResp)
	err := c.cc.Invoke(ctx, Social_GroupUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialServer is the server API for Social service.
// All implementations must embed UnimplementedSocialServer
// for forward compatibility
//
// svc
type SocialServer interface {
	FriendPutIn(context.Context, *FriendPutInReq) (*FriendPutInResp, error)
	FriendPutInHandle(context.Context, *FriendPutInHandleReq) (*FriendPutInHandleResp, error)
	FriendPutInList(context.Context, *FriendPutInListReq) (*FriendPutInListResp, error)
	FriendList(context.Context, *FriendListReq) (*FriendListResp, error)
	GroupCreate(context.Context, *GroupCreateReq) (*GroupCreateResp, error)
	GroupPutin(context.Context, *GroupPutinReq) (*GroupPutinResp, error)
	GroupPutinList(context.Context, *GroupPutinListReq) (*GroupPutinListResp, error)
	GroupPutInHandle(context.Context, *GroupPutInHandleReq) (*GroupPutInHandleResp, error)
	GroupList(context.Context, *GroupListReq) (*GroupListResp, error)
	GroupUsers(context.Context, *GroupUsersReq) (*GroupUsersResp, error)
	mustEmbedUnimplementedSocialServer()
}

// UnimplementedSocialServer must be embedded to have forward compatible implementations.
type UnimplementedSocialServer struct {
}

func (UnimplementedSocialServer) FriendPutIn(context.Context, *FriendPutInReq) (*FriendPutInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendPutIn not implemented")
}
func (UnimplementedSocialServer) FriendPutInHandle(context.Context, *FriendPutInHandleReq) (*FriendPutInHandleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendPutInHandle not implemented")
}
func (UnimplementedSocialServer) FriendPutInList(context.Context, *FriendPutInListReq) (*FriendPutInListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendPutInList not implemented")
}
func (UnimplementedSocialServer) FriendList(context.Context, *FriendListReq) (*FriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendList not implemented")
}
func (UnimplementedSocialServer) GroupCreate(context.Context, *GroupCreateReq) (*GroupCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupCreate not implemented")
}
func (UnimplementedSocialServer) GroupPutin(context.Context, *GroupPutinReq) (*GroupPutinResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPutin not implemented")
}
func (UnimplementedSocialServer) GroupPutinList(context.Context, *GroupPutinListReq) (*GroupPutinListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPutinList not implemented")
}
func (UnimplementedSocialServer) GroupPutInHandle(context.Context, *GroupPutInHandleReq) (*GroupPutInHandleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPutInHandle not implemented")
}
func (UnimplementedSocialServer) GroupList(context.Context, *GroupListReq) (*GroupListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupList not implemented")
}
func (UnimplementedSocialServer) GroupUsers(context.Context, *GroupUsersReq) (*GroupUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupUsers not implemented")
}
func (UnimplementedSocialServer) mustEmbedUnimplementedSocialServer() {}

// UnsafeSocialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialServer will
// result in compilation errors.
type UnsafeSocialServer interface {
	mustEmbedUnimplementedSocialServer()
}

func RegisterSocialServer(s grpc.ServiceRegistrar, srv SocialServer) {
	s.RegisterService(&Social_ServiceDesc, srv)
}

func _Social_FriendPutIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendPutInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).FriendPutIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_FriendPutIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).FriendPutIn(ctx, req.(*FriendPutInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_FriendPutInHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendPutInHandleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).FriendPutInHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_FriendPutInHandle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).FriendPutInHandle(ctx, req.(*FriendPutInHandleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_FriendPutInList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendPutInListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).FriendPutInList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_FriendPutInList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).FriendPutInList(ctx, req.(*FriendPutInListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_FriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).FriendList(ctx, req.(*FriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_GroupCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).GroupCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_GroupCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).GroupCreate(ctx, req.(*GroupCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_GroupPutin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPutinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).GroupPutin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_GroupPutin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).GroupPutin(ctx, req.(*GroupPutinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_GroupPutinList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPutinListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).GroupPutinList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_GroupPutinList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).GroupPutinList(ctx, req.(*GroupPutinListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_GroupPutInHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPutInHandleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).GroupPutInHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_GroupPutInHandle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).GroupPutInHandle(ctx, req.(*GroupPutInHandleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_GroupList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).GroupList(ctx, req.(*GroupListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Social_GroupUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServer).GroupUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Social_GroupUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServer).GroupUsers(ctx, req.(*GroupUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Social_ServiceDesc is the grpc.ServiceDesc for Social service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Social_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.social",
	HandlerType: (*SocialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FriendPutIn",
			Handler:    _Social_FriendPutIn_Handler,
		},
		{
			MethodName: "FriendPutInHandle",
			Handler:    _Social_FriendPutInHandle_Handler,
		},
		{
			MethodName: "FriendPutInList",
			Handler:    _Social_FriendPutInList_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _Social_FriendList_Handler,
		},
		{
			MethodName: "GroupCreate",
			Handler:    _Social_GroupCreate_Handler,
		},
		{
			MethodName: "GroupPutin",
			Handler:    _Social_GroupPutin_Handler,
		},
		{
			MethodName: "GroupPutinList",
			Handler:    _Social_GroupPutinList_Handler,
		},
		{
			MethodName: "GroupPutInHandle",
			Handler:    _Social_GroupPutInHandle_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Social_GroupList_Handler,
		},
		{
			MethodName: "GroupUsers",
			Handler:    _Social_GroupUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/social/rpc/social.proto",
}
