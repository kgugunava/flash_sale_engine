package main

import (
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/app"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/config"
)

func main() {
	cfg := config.Load()
	apiGatewayApp := app.NewApiGatewayApp(*cfg)
	if apiGatewayApp == nil {
		panic("Failed to initialize API Gateway App")
	}
	apiGatewayApp.Router.Run(app.GetServerURL(apiGatewayApp))
}