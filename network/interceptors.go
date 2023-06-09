package network

import (
	"context"
	"github.com/qcodelabsllc/qreeket/gateway/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

const permissionDeniedMessage = "You don't have permission to access this resource"

// AuthUnaryInterceptor => intercepts all incoming requests and checks if the user is authenticated and authorized to access the resource
// Also, it attaches tracing to the context
func AuthUnaryInterceptor(parentCtx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("unary interceptor: %v\n", info.FullMethod)
	// @todo check if the user is authenticated and authorized to access the resource

	// configure outgoing metadata
	//ctx, err := utils.CreateOutgoingMetadata(parentCtx)
	//if err != nil {
	//	return nil, err
	//}
	return handler(parentCtx, req)
}

// AuthStreamInterceptor => intercepts all incoming requests and checks if the user is authenticated and authorized to access the resource
// Also, it attaches tracing to the context
func AuthStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("stream interceptor: %v\n", info.FullMethod)
	switch info.FullMethod {
	case "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo":
		return handler(srv, ss)
	default:
		// @todo check if the user is authenticated and authorized to access the resource

		// configure outgoing metadata
		ctx, err := utils.CreateOutgoingMetadata(ss.Context())
		if err != nil {
			return err
		}
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			return status.Errorf(codes.PermissionDenied, permissionDeniedMessage)
		}
		_ = ss.SetHeader(md)
		return handler(srv, ss)
	}

}
