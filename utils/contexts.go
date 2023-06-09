package utils

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CreateOutgoingMetadata => creates outgoing metadata from the incoming metadata
func CreateOutgoingMetadata(ctx context.Context) (context.Context, error) {
	// get the incoming metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "metadata is not provided")
	}

	// create a new outgoing metadata
	outgoingMD := metadata.New(map[string]string{})

	if len(md) == 0 {
		return metadata.NewOutgoingContext(ctx, outgoingMD), nil
	}

	// attach the incoming metadata to the outgoing metadata
	for k, v := range md {
		outgoingMD[k] = v
	}

	// create a new context with the outgoing metadata
	outgoingMDCtx := metadata.NewOutgoingContext(ctx, outgoingMD)

	return outgoingMDCtx, nil
}
