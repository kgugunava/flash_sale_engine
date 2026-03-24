package main

import (
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/transport/http"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/transport/http/handler"
)

func main() {
	ordersHandler := handler.NewOrdersHandler()
	apiHandlers := http.NewApiHandlers(ordersHandler)
	router := http.NewRouter(*apiHandlers)
	router.Run("0.0.0.0:8080")
}