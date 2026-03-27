package service

import (
	"context"

	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/client/grpc"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/domain"
	// pb "github.com/kgugunava/flash_sale_engine/shared/proto/order"
)

type OrdersService struct {
	orderClient *grpc.OrderClient
}

func NewOrderService(orderClient *grpc.OrderClient) *OrdersService {
	return &OrdersService{
		orderClient: orderClient,
	}
}

func (s *OrdersService) CreateOrder(ctx context.Context, domainReq *domain.CreateOrderRequest) (*domain.CreateOrderResponse, error) {
	protoReq := domain.DomainCreateOrderRequestToProto(domainReq)

	resp, err := s.orderClient.CreateOrder(ctx, protoReq)

	if err != nil {
		return nil, err
	}

	return domain.ProtoCreateOrderResponseToDomain(resp), nil
}