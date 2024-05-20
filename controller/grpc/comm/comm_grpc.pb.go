// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/comm.proto

package comm

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

// CommServiceClient is the client API for CommService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommServiceClient interface {
	Ping(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error)
	Stop(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error)
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error)
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error)
	StrlnValue(ctx context.Context, in *StrlnValueRequest, opts ...grpc.CallOption) (*StrlnValueResponse, error)
	AppendValue(ctx context.Context, in *AppendValueRequest, opts ...grpc.CallOption) (*AppendValueResponse, error)
	// Raft
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
}

type commServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommServiceClient(cc grpc.ClientConnInterface) CommServiceClient {
	return &commServiceClient{cc}
}

func (c *commServiceClient) Ping(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commServiceClient) Stop(ctx context.Context, in *BasicRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	out := new(BasicResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commServiceClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueResponse, error) {
	out := new(GetValueResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commServiceClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/SetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commServiceClient) StrlnValue(ctx context.Context, in *StrlnValueRequest, opts ...grpc.CallOption) (*StrlnValueResponse, error) {
	out := new(StrlnValueResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/StrlnValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commServiceClient) AppendValue(ctx context.Context, in *AppendValueRequest, opts ...grpc.CallOption) (*AppendValueResponse, error) {
	out := new(AppendValueResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/AppendValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commServiceClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commServiceClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/comm.CommService/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommServiceServer is the server API for CommService service.
// All implementations must embed UnimplementedCommServiceServer
// for forward compatibility
type CommServiceServer interface {
	Ping(context.Context, *BasicRequest) (*BasicResponse, error)
	Stop(context.Context, *BasicRequest) (*BasicResponse, error)
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
	StrlnValue(context.Context, *StrlnValueRequest) (*StrlnValueResponse, error)
	AppendValue(context.Context, *AppendValueRequest) (*AppendValueResponse, error)
	// Raft
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	mustEmbedUnimplementedCommServiceServer()
}

// UnimplementedCommServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommServiceServer struct {
}

func (UnimplementedCommServiceServer) Ping(context.Context, *BasicRequest) (*BasicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCommServiceServer) Stop(context.Context, *BasicRequest) (*BasicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedCommServiceServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (UnimplementedCommServiceServer) SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (UnimplementedCommServiceServer) StrlnValue(context.Context, *StrlnValueRequest) (*StrlnValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StrlnValue not implemented")
}
func (UnimplementedCommServiceServer) AppendValue(context.Context, *AppendValueRequest) (*AppendValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendValue not implemented")
}
func (UnimplementedCommServiceServer) AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (UnimplementedCommServiceServer) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedCommServiceServer) mustEmbedUnimplementedCommServiceServer() {}

// UnsafeCommServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommServiceServer will
// result in compilation errors.
type UnsafeCommServiceServer interface {
	mustEmbedUnimplementedCommServiceServer()
}

func RegisterCommServiceServer(s grpc.ServiceRegistrar, srv CommServiceServer) {
	s.RegisterService(&CommService_ServiceDesc, srv)
}

func _CommService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).Ping(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).Stop(ctx, req.(*BasicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommService_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommService_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommService_StrlnValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrlnValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).StrlnValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/StrlnValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).StrlnValue(ctx, req.(*StrlnValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommService_AppendValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).AppendValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/AppendValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).AppendValue(ctx, req.(*AppendValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommService_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommService_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.CommService/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommService_ServiceDesc is the grpc.ServiceDesc for CommService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comm.CommService",
	HandlerType: (*CommServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CommService_Ping_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _CommService_Stop_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _CommService_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _CommService_SetValue_Handler,
		},
		{
			MethodName: "StrlnValue",
			Handler:    _CommService_StrlnValue_Handler,
		},
		{
			MethodName: "AppendValue",
			Handler:    _CommService_AppendValue_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _CommService_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _CommService_RequestVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/comm.proto",
}
