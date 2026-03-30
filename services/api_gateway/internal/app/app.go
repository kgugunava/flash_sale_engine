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
	
	// Пытаемся подключиться 5 раз с интервалом в 1 секунду
	var orderClient *grpc.OrderClient
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		orderClient = grpc.NewOrderClient("localhost:8020", apiGatewayApp.Logger, 10 * time.Second)
		if orderClient != nil {
			apiGatewayApp.Logger.Info("Successfully connected to OrderService",
				logger.Int("attempt", i+1),
			)
			break
		}
		
		if i < maxRetries-1 {
			apiGatewayApp.Logger.Warn("Failed to connect to OrderService, retrying...",
				logger.Int("attempt", i+1),
				logger.Int("max_retries", maxRetries),
			)
			time.Sleep(time.Second)
		}
	}
	
	if orderClient == nil {
		apiGatewayApp.Logger.Panic("Failed to initialize OrderClient after all retries",
			logger.Int("max_retries", maxRetries),
		)
		return nil
	}

	if orderClient != nil {
		apiGatewayApp.Logger.Info("order client not nil")
	}

	fmt.Print("orderClient: ", orderClient)
	
	ordersService := service.NewOrderService(orderClient)
	ordersHandler := handler.NewOrdersHandler(ordersService)
	apiHandlers := http.NewApiHandlers(ordersHandler)
	apiGatewayApp.Router = http.NewRouter(*apiHandlers)

	return apiGatewayApp
}

func GetServerURL(apiGatewayApp *ApiGatewayApp) string {
	return fmt.Sprintf("%s%s", apiGatewayApp.Cfg.ServerAddress, apiGatewayApp.Cfg.HTTPPort)
}