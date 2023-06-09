package services

import (
	"context"
	"errors"
	pb "github.com/qcodelabsllc/qreeket/gateway/gen"
	"github.com/qcodelabsllc/qreeket/gateway/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

// connectionFailedMessage => message to display when a connection to a service fails
const connectionFailedMessage = "something went wrong. Please try again later"

// CreateAnnouncementClient => creates a new announcement client
func CreateAnnouncementClient(ctx context.Context) (pb.AnnouncementServiceClient, context.Context) {
	addr := os.Getenv("ANNOUNCEMENT_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	ctx, _ = utils.CreateOutgoingMetadata(ctx)
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewAnnouncementServiceClient(conn), ctx
}

// CreateAuthServiceClient => creates a new auth client
func CreateAuthServiceClient(ctx context.Context) (pb.AuthServiceClient, context.Context) {
	addr := os.Getenv("AUTH_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	ctx, _ = utils.CreateOutgoingMetadata(ctx)
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewAuthServiceClient(conn), ctx
}

// CreateBillingServiceClient => creates a new billing client
func CreateBillingServiceClient(ctx context.Context) (pb.BillingServiceClient, context.Context) {
	addr := os.Getenv("BILLING_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	ctx, _ = utils.CreateOutgoingMetadata(ctx)
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewBillingServiceClient(conn), ctx
}

// CreateGroupServiceClient => creates a new group client
func CreateGroupServiceClient(ctx context.Context) (pb.GroupChannelServiceClient, context.Context) {
	addr := os.Getenv("GROUPS_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	ctx, _ = utils.CreateOutgoingMetadata(ctx)
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewGroupChannelServiceClient(conn), ctx
}

// CreateMediaServiceClient => creates a new media client
func CreateMediaServiceClient(ctx context.Context) (pb.MediaServiceClient, context.Context) {
	addr := os.Getenv("MEDIA_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	ctx, _ = utils.CreateOutgoingMetadata(ctx)
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewMediaServiceClient(conn), ctx
}

// CreateMessagingServiceClient => creates a new messaging client
func CreateMessagingServiceClient(ctx context.Context) (pb.MessagingServiceClient, context.Context) {
	addr := os.Getenv("MESSAGING_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	ctx, _ = utils.CreateOutgoingMetadata(ctx)
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewMessagingServiceClient(conn), ctx
}

// CreateSmsServiceClient => creates a new sms client
func CreateSmsServiceClient(ctx context.Context) (pb.SmsServiceClient, context.Context) {
	addr := os.Getenv("SMS_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	ctx, _ = utils.CreateOutgoingMetadata(ctx)
	conn, _ := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewSmsServiceClient(conn), ctx
}
