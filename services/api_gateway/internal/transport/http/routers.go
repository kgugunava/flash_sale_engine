package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/transport/http/handler"

)

type Route struct {
	Name string
	Method string
	URLPattern string
	HandlerFunc gin.HandlerFunc
}

func NewRouter(handlers ApiHandlers) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handlers)
}

func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandlers) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.URLPattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.URLPattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.URLPattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.URLPattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.URLPattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandlers struct {
	OrdersHandler handler.OrdersHandler
}

func NewApiHandlers(ordersHandler *handler.OrdersHandler) *ApiHandlers {
	return &ApiHandlers{OrdersHandler: *ordersHandler}
}

func getRoutes(handlers ApiHandlers) []Route {
	return []Route{ 
		{
			Name: "CreateOrderPost",
			Method: http.MethodPost,
			URLPattern: "orders/create",
			HandlerFunc: handlers.OrdersHandler.CreateOrderPost,
		},
		{
			Name: "CancelOrderPost",
			Method: http.MethodPost,
			URLPattern: "orders/cancel/:order_id",
			HandlerFunc: handlers.OrdersHandler.CancelOrderPost,
		},
		{
			Name: "GetOrdersListGet",
			Method: http.MethodGet,
			URLPattern: "orders/list",
			HandlerFunc: handlers.OrdersHandler.GetOrdersListGet,
		},
		{
			Name: "GetUserOrdersListGet",
			Method: http.MethodGet,
			URLPattern: "orders/users/:user_id",
			HandlerFunc: handlers.OrdersHandler.GetUserOrdersListGet,
		},
		{
			Name: "PayForOrderPost",
			Method: http.MethodPost,
			URLPattern: "orders/pay/:order_id",
			HandlerFunc: handlers.OrdersHandler.PayForOrderPost,
		},
	}
}