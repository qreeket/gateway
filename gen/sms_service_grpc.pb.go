// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: sms_service.proto

package qreeket

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SmsService_SendPhoneVerificationCode_FullMethodName   = "/qreeket.SmsService/send_phone_verification_code"
	SmsService_VerifyPhoneVerificationCode_FullMethodName = "/qreeket.SmsService/verify_phone_verification_code"
)

// SmsServiceClient is the client API for SmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsServiceClient interface {
	// phone verification
	SendPhoneVerificationCode(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VerifyPhoneVerificationCode(ctx context.Context, in *VerifyPhoneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type smsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServiceClient(cc grpc.ClientConnInterface) SmsServiceClient {
	return &smsServiceClient{cc}
}

func (c *smsServiceClient) SendPhoneVerificationCode(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SmsService_SendPhoneVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsServiceClient) VerifyPhoneVerificationCode(ctx context.Context, in *VerifyPhoneRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SmsService_VerifyPhoneVerificationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServiceServer is the server API for SmsService service.
// All implementations must embed UnimplementedSmsServiceServer
// for forward compatibility
type SmsServiceServer interface {
	// phone verification
	SendPhoneVerificationCode(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error)
	VerifyPhoneVerificationCode(context.Context, *VerifyPhoneRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSmsServiceServer()
}

// UnimplementedSmsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServiceServer struct {
}

func (UnimplementedSmsServiceServer) SendPhoneVerificationCode(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneVerificationCode not implemented")
}
func (UnimplementedSmsServiceServer) VerifyPhoneVerificationCode(context.Context, *VerifyPhoneRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPhoneVerificationCode not implemented")
}
func (UnimplementedSmsServiceServer) mustEmbedUnimplementedSmsServiceServer() {}

// UnsafeSmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServiceServer will
// result in compilation errors.
type UnsafeSmsServiceServer interface {
	mustEmbedUnimplementedSmsServiceServer()
}

func RegisterSmsServiceServer(s grpc.ServiceRegistrar, srv SmsServiceServer) {
	s.RegisterService(&SmsService_ServiceDesc, srv)
}

func _SmsService_SendPhoneVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).SendPhoneVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsService_SendPhoneVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).SendPhoneVerificationCode(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsService_VerifyPhoneVerificationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).VerifyPhoneVerificationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsService_VerifyPhoneVerificationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).VerifyPhoneVerificationCode(ctx, req.(*VerifyPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsService_ServiceDesc is the grpc.ServiceDesc for SmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qreeket.SmsService",
	HandlerType: (*SmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send_phone_verification_code",
			Handler:    _SmsService_SendPhoneVerificationCode_Handler,
		},
		{
			MethodName: "verify_phone_verification_code",
			Handler:    _SmsService_VerifyPhoneVerificationCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms_service.proto",
}
