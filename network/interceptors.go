package network

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthUnaryInterceptor => intercepts all incoming requests and checks if the user is authenticated and authorized to access the resource
// Also, it attaches tracing to the context
func AuthUnaryInterceptor(parentCtx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// TODO - create a new auth client
	// TODO - attach tracing to the context
	return handler(parentCtx, req)
}

// AuthStreamInterceptor => intercepts all incoming requests and checks if the user is authenticated and authorized to access the resource
// Also, it attaches tracing to the context
func AuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// TODO - create a new auth client
	// TODO - attach tracing to the context
	return handler(srv, ss)
}

// convertErrToStatus => converts an error to a gRPC status
func convertErrToStatus(err error) error {
	// check if the error is a gRPC status
	if s, ok := status.FromError(err); ok {
		return status.Errorf(s.Code(), s.Message())
	}

	// return an internal server error
	return status.Errorf(codes.Internal, err.Error())
}
