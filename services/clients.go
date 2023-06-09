package services

import (
	"errors"
	pb "github.com/qcodelabsllc/qreeket/gateway/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

// connectionFailedMessage => message to display when a connection to a service fails
const connectionFailedMessage = "something went wrong. Please try again later"

// CreateAnnouncementClient => creates a new announcement client
func CreateAnnouncementClient() pb.AnnouncementServiceClient {
	addr := os.Getenv("ANNOUNCEMENT_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewAnnouncementServiceClient(conn)
}

// CreateAuthServiceClient => creates a new auth client
func CreateAuthServiceClient() pb.AuthServiceClient {
	addr := os.Getenv("AUTH_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewAuthServiceClient(conn)
}

// CreateBillingServiceClient => creates a new billing client
func CreateBillingServiceClient() pb.BillingServiceClient {
	addr := os.Getenv("BILLING_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewBillingServiceClient(conn)
}

// CreateGroupServiceClient => creates a new group client
func CreateGroupServiceClient() pb.GroupChannelServiceClient {
	addr := os.Getenv("GROUPS_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewGroupChannelServiceClient(conn)
}

// CreateMediaServiceClient => creates a new media client
func CreateMediaServiceClient() pb.MediaServiceClient {
	addr := os.Getenv("MEDIA_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewMediaServiceClient(conn)
}

// CreateMessagingServiceClient => creates a new messaging client
func CreateMessagingServiceClient() pb.MessagingServiceClient {
	addr := os.Getenv("MESSAGING_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewMessagingServiceClient(conn)
}

// CreateSmsServiceClient => creates a new sms client
func CreateSmsServiceClient() pb.SmsServiceClient {
	addr := os.Getenv("SMS_SERVICE_URL")
	if len(addr) == 0 {
		panic(errors.New(connectionFailedMessage))
	}
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return pb.NewSmsServiceClient(conn)
}
