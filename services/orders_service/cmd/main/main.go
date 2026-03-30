package main

import (
	"context"

	"github.com/kgugunava/flash_sale_engine/orders_service/internal/grpc"
	"github.com/kgugunava/flash_sale_engine/orders_service/internal/service"
	"github.com/kgugunava/flash_sale_engine/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	ordersService := service.NewOrderService()
	server := grpc.NewGRPCServer(8020, log, ordersService)
	server.Start(context.Background())
}