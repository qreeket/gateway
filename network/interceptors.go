package network

import (
	"context"
	"google.golang.org/grpc"
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
