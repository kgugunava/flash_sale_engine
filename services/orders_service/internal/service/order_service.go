package service

import (
	"github.com/kgugunava/flash_sale_engine/orders_service/internal/domain"
)

type OrderServiceInterface interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
}

type OrderService struct {
	
}

func NewOrderService() OrderServiceInterface {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(order *domain.Order) (*domain.Order, error) {
	order.OrderID = "new-order-id"
	return order, nil
}