package app

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/config"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/transport/http/handler"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/transport/http"
)

type ApiGatewayApp struct {
	Cfg *config.Config
	Router *gin.Engine
}

func NewApiGatewayApp(cfg config.Config) *ApiGatewayApp {
	apiGatewayApp := &ApiGatewayApp{
		Cfg: &cfg,
	}
	ordersHandler := handler.NewOrdersHandler()
	apiHandlers := http.NewApiHandlers(ordersHandler)
	apiGatewayApp.Router = http.NewRouter(*apiHandlers)

	return apiGatewayApp
}

func GetServerURL(apiGatewayApp *ApiGatewayApp) string {
	return fmt.Sprintf("%s%s", apiGatewayApp.Cfg.ServerAddress, apiGatewayApp.Cfg.HTTPPort)
}