package main

import (
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/app"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/config"
)

func main() {
	cfg := config.Load()
	apiGatewayApp := app.NewApiGatewayApp(*cfg)
	apiGatewayApp.Router.Run(app.GetServerURL(apiGatewayApp))
}