// protoc --go_out=../grpc --go-grpc_out=../grpc node.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: node.proto

package node

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
	ManuscriptNodeService_ReceiveNewTask_FullMethodName  = "/node.ManuscriptNodeService/ReceiveNewTask"
	ManuscriptNodeService_GetOperatorInfo_FullMethodName = "/node.ManuscriptNodeService/GetOperatorInfo"
)

// ManuscriptNodeServiceClient is the client API for ManuscriptNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManuscriptNodeServiceClient interface {
	ReceiveNewTask(ctx context.Context, in *NewTaskRequest, opts ...grpc.CallOption) (*NewTaskResponse, error)
	GetOperatorInfo(ctx context.Context, in *GetOperatorInfoRequest, opts ...grpc.CallOption) (*GetOperatorInfoResponse, error)
}

type manuscriptNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManuscriptNodeServiceClient(cc grpc.ClientConnInterface) ManuscriptNodeServiceClient {
	return &manuscriptNodeServiceClient{cc}
}

func (c *manuscriptNodeServiceClient) ReceiveNewTask(ctx context.Context, in *NewTaskRequest, opts ...grpc.CallOption) (*NewTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewTaskResponse)
	err := c.cc.Invoke(ctx, ManuscriptNodeService_ReceiveNewTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manuscriptNodeServiceClient) GetOperatorInfo(ctx context.Context, in *GetOperatorInfoRequest, opts ...grpc.CallOption) (*GetOperatorInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOperatorInfoResponse)
	err := c.cc.Invoke(ctx, ManuscriptNodeService_GetOperatorInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManuscriptNodeServiceServer is the server API for ManuscriptNodeService service.
// All implementations must embed UnimplementedManuscriptNodeServiceServer
// for forward compatibility.
type ManuscriptNodeServiceServer interface {
	ReceiveNewTask(context.Context, *NewTaskRequest) (*NewTaskResponse, error)
	GetOperatorInfo(context.Context, *GetOperatorInfoRequest) (*GetOperatorInfoResponse, error)
	mustEmbedUnimplementedManuscriptNodeServiceServer()
}

// UnimplementedManuscriptNodeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedManuscriptNodeServiceServer struct{}

func (UnimplementedManuscriptNodeServiceServer) ReceiveNewTask(context.Context, *NewTaskRequest) (*NewTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveNewTask not implemented")
}
func (UnimplementedManuscriptNodeServiceServer) GetOperatorInfo(context.Context, *GetOperatorInfoRequest) (*GetOperatorInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatorInfo not implemented")
}
func (UnimplementedManuscriptNodeServiceServer) mustEmbedUnimplementedManuscriptNodeServiceServer() {}
func (UnimplementedManuscriptNodeServiceServer) testEmbeddedByValue()                               {}

// UnsafeManuscriptNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManuscriptNodeServiceServer will
// result in compilation errors.
type UnsafeManuscriptNodeServiceServer interface {
	mustEmbedUnimplementedManuscriptNodeServiceServer()
}

func RegisterManuscriptNodeServiceServer(s grpc.ServiceRegistrar, srv ManuscriptNodeServiceServer) {
	// If the following call pancis, it indicates UnimplementedManuscriptNodeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ManuscriptNodeService_ServiceDesc, srv)
}

func _ManuscriptNodeService_ReceiveNewTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManuscriptNodeServiceServer).ReceiveNewTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManuscriptNodeService_ReceiveNewTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManuscriptNodeServiceServer).ReceiveNewTask(ctx, req.(*NewTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManuscriptNodeService_GetOperatorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatorInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManuscriptNodeServiceServer).GetOperatorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ManuscriptNodeService_GetOperatorInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManuscriptNodeServiceServer).GetOperatorInfo(ctx, req.(*GetOperatorInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManuscriptNodeService_ServiceDesc is the grpc.ServiceDesc for ManuscriptNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManuscriptNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "node.ManuscriptNodeService",
	HandlerType: (*ManuscriptNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveNewTask",
			Handler:    _ManuscriptNodeService_ReceiveNewTask_Handler,
		},
		{
			MethodName: "GetOperatorInfo",
			Handler:    _ManuscriptNodeService_GetOperatorInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
