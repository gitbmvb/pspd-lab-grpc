// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LLMServiceClient is the client API for LLMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LLMServiceClient interface {
	// Chamada unária para health check
	HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	// Stream de resposta da LLM
	GenerateText(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (LLMService_GenerateTextClient, error)
	// Chamada para carregar modelos
	LoadModel(ctx context.Context, in *ModelRequest, opts ...grpc.CallOption) (*ModelResponse, error)
}

type lLMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLLMServiceClient(cc grpc.ClientConnInterface) LLMServiceClient {
	return &lLMServiceClient{cc}
}

func (c *lLMServiceClient) HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/LLMService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lLMServiceClient) GenerateText(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (LLMService_GenerateTextClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LLMService_serviceDesc.Streams[0], "/LLMService/GenerateText", opts...)
	if err != nil {
		return nil, err
	}
	x := &lLMServiceGenerateTextClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LLMService_GenerateTextClient interface {
	Recv() (*TextResponse, error)
	grpc.ClientStream
}

type lLMServiceGenerateTextClient struct {
	grpc.ClientStream
}

func (x *lLMServiceGenerateTextClient) Recv() (*TextResponse, error) {
	m := new(TextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lLMServiceClient) LoadModel(ctx context.Context, in *ModelRequest, opts ...grpc.CallOption) (*ModelResponse, error) {
	out := new(ModelResponse)
	err := c.cc.Invoke(ctx, "/LLMService/LoadModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LLMServiceServer is the server API for LLMService service.
// All implementations must embed UnimplementedLLMServiceServer
// for forward compatibility
type LLMServiceServer interface {
	// Chamada unária para health check
	HealthCheck(context.Context, *HealthRequest) (*HealthResponse, error)
	// Stream de resposta da LLM
	GenerateText(*PromptRequest, LLMService_GenerateTextServer) error
	// Chamada para carregar modelos
	LoadModel(context.Context, *ModelRequest) (*ModelResponse, error)
	mustEmbedUnimplementedLLMServiceServer()
}

// UnimplementedLLMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLLMServiceServer struct {
}

func (UnimplementedLLMServiceServer) HealthCheck(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedLLMServiceServer) GenerateText(*PromptRequest, LLMService_GenerateTextServer) error {
	return status.Errorf(codes.Unimplemented, "method GenerateText not implemented")
}
func (UnimplementedLLMServiceServer) LoadModel(context.Context, *ModelRequest) (*ModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadModel not implemented")
}
func (UnimplementedLLMServiceServer) mustEmbedUnimplementedLLMServiceServer() {}

// UnsafeLLMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LLMServiceServer will
// result in compilation errors.
type UnsafeLLMServiceServer interface {
	mustEmbedUnimplementedLLMServiceServer()
}

func RegisterLLMServiceServer(s *grpc.Server, srv LLMServiceServer) {
	s.RegisterService(&_LLMService_serviceDesc, srv)
}

func _LLMService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LLMService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).HealthCheck(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LLMService_GenerateText_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PromptRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LLMServiceServer).GenerateText(m, &lLMServiceGenerateTextServer{stream})
}

type LLMService_GenerateTextServer interface {
	Send(*TextResponse) error
	grpc.ServerStream
}

type lLMServiceGenerateTextServer struct {
	grpc.ServerStream
}

func (x *lLMServiceGenerateTextServer) Send(m *TextResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LLMService_LoadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).LoadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LLMService/LoadModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).LoadModel(ctx, req.(*ModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LLMService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LLMService",
	HandlerType: (*LLMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _LLMService_HealthCheck_Handler,
		},
		{
			MethodName: "LoadModel",
			Handler:    _LLMService_LoadModel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GenerateText",
			Handler:       _LLMService_GenerateText_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/llm-service.proto",
}
