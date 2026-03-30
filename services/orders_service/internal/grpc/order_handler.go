package grpc

import (
	"context"

	"github.com/kgugunava/flash_sale_engine/orders_service/internal/service"
	"github.com/kgugunava/flash_sale_engine/orders_service/internal/domain"
	"github.com/kgugunava/flash_sale_engine/pkg/logger"
	pb_order "github.com/kgugunava/flash_sale_engine/shared/proto/order"
	pb_common "github.com/kgugunava/flash_sale_engine/shared/proto/common"
)

var TraceIDPlaceholder string = "placeholder_trace_id"

type OrderHandler struct {
	pb_order.UnimplementedOrderServiceServer

	ordersService service.OrderServiceInterface
	log *logger.Logger
}

func NewOrderHandler(log *logger.Logger, ordersService service.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		ordersService: ordersService,
		log: log,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb_order.CreateOrderRequest) (*pb_order.CreateOrderResponse, error) {
	h.log.Info("CreateOrder request received",
        logger.String("trace_id", TraceIDPlaceholder),
        logger.String("user_id", req.UserId),
        logger.String("item_name", req.ItemName),
    )

	order := domain.ProtoCreateOrderRequestToDomainOrder(req)
	
	order, err := h.ordersService.CreateOrder(order)

	if err != nil {
		h.log.Error("Error in CreateOrder from Orders Service",
			logger.Any("error", err.Error()),
		)
		return nil, err
	}

	return &pb_order.CreateOrderResponse{
		Order: domain.DomainOrderToProto(order),
		Status: &pb_common.ResponseStatus{
			Success: true,
			Code: "NO ERROR",
			Message: "New order created",
		},
		Error: nil,
	}, nil
}

