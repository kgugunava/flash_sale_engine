package app

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/client/grpc"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/config"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/service"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/transport/http"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/transport/http/handler"
	"github.com/kgugunava/flash_sale_engine/pkg/logger"
)

type ApiGatewayApp struct {
	Cfg *config.Config
	Router *gin.Engine
	Logger *logger.Logger
}

func NewApiGatewayApp(cfg config.Config) *ApiGatewayApp {
	apiGatewayApp := &ApiGatewayApp{
		Cfg: &cfg,
		Logger: logger.NewLogger(),
	}
	orderClient := grpc.NewOrderClient("localhost:8090", apiGatewayApp.Logger, 10 * time.Second)
	if orderClient == nil {
		apiGatewayApp.Logger.Error("Failed to initialize OrderClient")
		return nil
	}
	
	ordersService := service.NewOrderService(orderClient)
	ordersHandler := handler.NewOrdersHandler(ordersService)
	apiHandlers := http.NewApiHandlers(ordersHandler)
	apiGatewayApp.Router = http.NewRouter(*apiHandlers)

	return apiGatewayApp
}

func GetServerURL(apiGatewayApp *ApiGatewayApp) string {
	return fmt.Sprintf("%s%s", apiGatewayApp.Cfg.ServerAddress, apiGatewayApp.Cfg.HTTPPort)
}