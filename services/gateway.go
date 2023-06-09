package services

import (
	pb "github.com/qcodelabsllc/qreeket/gateway/gen"
)

type GatewayService struct {
	pb.UnimplementedQreeketGatewayServiceServer
}

func NewGatewayService() *GatewayService {
	return &GatewayService{}
}
