// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: users-service/src/gRPC/proto/users.proto

// Go Only  Headers

// desc:

package users

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UsersServices_Ping_FullMethodName             = "/users.UsersServices/Ping"
	UsersServices_Create_FullMethodName           = "/users.UsersServices/Create"
	UsersServices_Update_FullMethodName           = "/users.UsersServices/Update"
	UsersServices_Delete_FullMethodName           = "/users.UsersServices/Delete"
	UsersServices_GetAll_FullMethodName           = "/users.UsersServices/GetAll"
	UsersServices_GetById_FullMethodName          = "/users.UsersServices/GetById"
	UsersServices_Follow_FullMethodName           = "/users.UsersServices/Follow"
	UsersServices_Unfollow_FullMethodName         = "/users.UsersServices/Unfollow"
	UsersServices_GetFollowers_FullMethodName     = "/users.UsersServices/GetFollowers"
	UsersServices_GetFollowings_FullMethodName    = "/users.UsersServices/GetFollowings"
	UsersServices_RedifinePassword_FullMethodName = "/users.UsersServices/RedifinePassword"
	UsersServices_Login_FullMethodName            = "/users.UsersServices/Login"
)

// UsersServicesClient is the client API for UsersServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// desc:
type UsersServicesClient interface {
	Ping(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*PongMsg, error)
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetAll(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserResponse], error)
	GetById(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Follow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*OK, error)
	Unfollow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*OK, error)
	GetFollowers(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserResponse], error)
	GetFollowings(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserResponse], error)
	RedifinePassword(ctx context.Context, in *RedifinePasswordMessage, opts ...grpc.CallOption) (*OK, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenMessage, error)
}

type usersServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServicesClient(cc grpc.ClientConnInterface) UsersServicesClient {
	return &usersServicesClient{cc}
}

func (c *usersServicesClient) Ping(ctx context.Context, in *PingMsg, opts ...grpc.CallOption) (*PongMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PongMsg)
	err := c.cc.Invoke(ctx, UsersServices_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UsersServices_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UsersServices_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) Delete(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UsersServices_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) GetAll(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &UsersServices_ServiceDesc.Streams[0], UsersServices_GetAll_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetAllUsersRequest, UserResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UsersServices_GetAllClient = grpc.ServerStreamingClient[UserResponse]

func (c *usersServicesClient) GetById(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UsersServices_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) Follow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*OK, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OK)
	err := c.cc.Invoke(ctx, UsersServices_Follow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) Unfollow(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*OK, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OK)
	err := c.cc.Invoke(ctx, UsersServices_Unfollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) GetFollowers(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &UsersServices_ServiceDesc.Streams[1], UsersServices_GetFollowers_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UserIdRequest, UserResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UsersServices_GetFollowersClient = grpc.ServerStreamingClient[UserResponse]

func (c *usersServicesClient) GetFollowings(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UserResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &UsersServices_ServiceDesc.Streams[2], UsersServices_GetFollowings_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UserIdRequest, UserResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UsersServices_GetFollowingsClient = grpc.ServerStreamingClient[UserResponse]

func (c *usersServicesClient) RedifinePassword(ctx context.Context, in *RedifinePasswordMessage, opts ...grpc.CallOption) (*OK, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OK)
	err := c.cc.Invoke(ctx, UsersServices_RedifinePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServicesClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*TokenMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenMessage)
	err := c.cc.Invoke(ctx, UsersServices_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServicesServer is the server API for UsersServices service.
// All implementations must embed UnimplementedUsersServicesServer
// for forward compatibility.
//
// desc:
type UsersServicesServer interface {
	Ping(context.Context, *PingMsg) (*PongMsg, error)
	Create(context.Context, *CreateUserRequest) (*UserResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UserResponse, error)
	Delete(context.Context, *UserIdRequest) (*UserResponse, error)
	GetAll(*GetAllUsersRequest, grpc.ServerStreamingServer[UserResponse]) error
	GetById(context.Context, *UserIdRequest) (*UserResponse, error)
	Follow(context.Context, *UserIdRequest) (*OK, error)
	Unfollow(context.Context, *UserIdRequest) (*OK, error)
	GetFollowers(*UserIdRequest, grpc.ServerStreamingServer[UserResponse]) error
	GetFollowings(*UserIdRequest, grpc.ServerStreamingServer[UserResponse]) error
	RedifinePassword(context.Context, *RedifinePasswordMessage) (*OK, error)
	Login(context.Context, *LoginRequest) (*TokenMessage, error)
	mustEmbedUnimplementedUsersServicesServer()
}

// UnimplementedUsersServicesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServicesServer struct{}

func (UnimplementedUsersServicesServer) Ping(context.Context, *PingMsg) (*PongMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUsersServicesServer) Create(context.Context, *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUsersServicesServer) Update(context.Context, *UpdateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUsersServicesServer) Delete(context.Context, *UserIdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUsersServicesServer) GetAll(*GetAllUsersRequest, grpc.ServerStreamingServer[UserResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUsersServicesServer) GetById(context.Context, *UserIdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUsersServicesServer) Follow(context.Context, *UserIdRequest) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedUsersServicesServer) Unfollow(context.Context, *UserIdRequest) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedUsersServicesServer) GetFollowers(*UserIdRequest, grpc.ServerStreamingServer[UserResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedUsersServicesServer) GetFollowings(*UserIdRequest, grpc.ServerStreamingServer[UserResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetFollowings not implemented")
}
func (UnimplementedUsersServicesServer) RedifinePassword(context.Context, *RedifinePasswordMessage) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedifinePassword not implemented")
}
func (UnimplementedUsersServicesServer) Login(context.Context, *LoginRequest) (*TokenMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsersServicesServer) mustEmbedUnimplementedUsersServicesServer() {}
func (UnimplementedUsersServicesServer) testEmbeddedByValue()                       {}

// UnsafeUsersServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServicesServer will
// result in compilation errors.
type UnsafeUsersServicesServer interface {
	mustEmbedUnimplementedUsersServicesServer()
}

func RegisterUsersServicesServer(s grpc.ServiceRegistrar, srv UsersServicesServer) {
	// If the following call pancis, it indicates UnimplementedUsersServicesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UsersServices_ServiceDesc, srv)
}

func _UsersServices_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).Ping(ctx, req.(*PingMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).Delete(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServicesServer).GetAll(m, &grpc.GenericServerStream[GetAllUsersRequest, UserResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UsersServices_GetAllServer = grpc.ServerStreamingServer[UserResponse]

func _UsersServices_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).GetById(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_Follow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).Follow(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_Unfollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).Unfollow(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_GetFollowers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServicesServer).GetFollowers(m, &grpc.GenericServerStream[UserIdRequest, UserResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UsersServices_GetFollowersServer = grpc.ServerStreamingServer[UserResponse]

func _UsersServices_GetFollowings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServicesServer).GetFollowings(m, &grpc.GenericServerStream[UserIdRequest, UserResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UsersServices_GetFollowingsServer = grpc.ServerStreamingServer[UserResponse]

func _UsersServices_RedifinePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedifinePasswordMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).RedifinePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_RedifinePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).RedifinePassword(ctx, req.(*RedifinePasswordMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersServices_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersServices_ServiceDesc is the grpc.ServiceDesc for UsersServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UsersServices",
	HandlerType: (*UsersServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UsersServices_Ping_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UsersServices_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UsersServices_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UsersServices_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _UsersServices_GetById_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _UsersServices_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _UsersServices_Unfollow_Handler,
		},
		{
			MethodName: "RedifinePassword",
			Handler:    _UsersServices_RedifinePassword_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UsersServices_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _UsersServices_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFollowers",
			Handler:       _UsersServices_GetFollowers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFollowings",
			Handler:       _UsersServices_GetFollowings_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users-service/src/gRPC/proto/users.proto",
}
