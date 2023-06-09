package services

import (
	"context"
	pb "github.com/qcodelabsllc/qreeket/gateway/gen"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *GatewayService) Login(ctx context.Context, req *pb.LoginRequest) (*wrapperspb.StringValue, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.Login(outgoingCtx, req)
}

func (s *GatewayService) Register(ctx context.Context, req *pb.RegisterRequest) (*wrapperspb.StringValue, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.Register(outgoingCtx, req)
}

func (s *GatewayService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*wrapperspb.StringValue, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.ResetPassword(outgoingCtx, req)
}

func (s *GatewayService) RequestPasswordReset(ctx context.Context, req *pb.RequestPasswordResetRequest) (*wrapperspb.StringValue, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.RequestPasswordReset(outgoingCtx, req)
}

func (s *GatewayService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.Logout(outgoingCtx, req)
}

func (s *GatewayService) VerifyPassword(ctx context.Context, req *wrapperspb.StringValue) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.VerifyPassword(outgoingCtx, req)
}

func (s *GatewayService) UpgradeToPremium(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.UpgradeToPremium(outgoingCtx, req)
}

func (s *GatewayService) DowngradeToStandard(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.DowngradeToStandard(outgoingCtx, req)
}

func (s *GatewayService) RequestPublicAccessToken(ctx context.Context, _ *emptypb.Empty) (*wrapperspb.StringValue, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.RequestPublicAccessToken(outgoingCtx, &emptypb.Empty{})
}

func (s *GatewayService) ValidateAccessToken(ctx context.Context, req *emptypb.Empty) (*pb.ValidateAccessTokenResponse, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.ValidateAccessToken(outgoingCtx, req)
}

func (s *GatewayService) GetAccount(ctx context.Context, req *emptypb.Empty) (*pb.Account, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.GetAccount(outgoingCtx, req)
}

func (s *GatewayService) GetAccountByPhoneNumber(ctx context.Context, req *wrapperspb.StringValue) (*pb.Account, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.GetAccountByPhoneNumber(outgoingCtx, req)
}

func (s *GatewayService) GetAccountById(ctx context.Context, req *wrapperspb.StringValue) (*pb.Account, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.GetAccountById(outgoingCtx, req)
}

func (s *GatewayService) UpdateAccount(ctx context.Context, req *pb.Account) (*pb.Account, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.UpdateAccount(outgoingCtx, req)
}

func (s *GatewayService) DeleteAccount(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.DeleteAccount(outgoingCtx, req)
}

func (s *GatewayService) AuthenticateAccount(ctx context.Context, req *pb.AuthenticateWithSocialAccountRequest) (*wrapperspb.StringValue, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.AuthenticateAccount(outgoingCtx, req)
}

func (s *GatewayService) CheckEmail(ctx context.Context, req *wrapperspb.StringValue) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.CheckEmail(outgoingCtx, req)
}

func (s *GatewayService) CheckPhoneNumber(ctx context.Context, req *wrapperspb.StringValue) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.CheckPhoneNumber(outgoingCtx, req)
}

func (s *GatewayService) GetCountries(ctx context.Context, req *emptypb.Empty) (*pb.GetCountriesResponse, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.GetCountries(outgoingCtx, req)
}

func (s *GatewayService) GetCountryById(ctx context.Context, req *wrapperspb.StringValue) (*pb.Country, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.GetCountryById(outgoingCtx, req)
}

func (s *GatewayService) AddCountry(ctx context.Context, req *pb.Country) (*pb.Country, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.AddCountry(outgoingCtx, req)
}

func (s *GatewayService) DeleteCountry(ctx context.Context, req *wrapperspb.StringValue) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.DeleteCountry(outgoingCtx, req)
}

func (s *GatewayService) GetCollegesForCountry(ctx context.Context, req *wrapperspb.StringValue) (*pb.GetCollegesResponse, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.GetCollegesForCountry(outgoingCtx, req)
}

func (s *GatewayService) GetCollegeById(ctx context.Context, req *wrapperspb.StringValue) (*pb.College, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.GetCollegeById(outgoingCtx, req)
}

func (s *GatewayService) AddCollege(ctx context.Context, req *pb.College) (*pb.College, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.AddCollege(outgoingCtx, req)
}

func (s *GatewayService) DeleteCollege(ctx context.Context, req *wrapperspb.StringValue) (*emptypb.Empty, error) {
	client, outgoingCtx := CreateAuthServiceClient(ctx)
	return client.DeleteCollege(outgoingCtx, req)
}
