package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/domain"
	model_errors "github.com/kgugunava/flash_sale_engine/api_gateway/internal/model/errors"
	model_orders_requests "github.com/kgugunava/flash_sale_engine/api_gateway/internal/model/orders/requests"
	"github.com/kgugunava/flash_sale_engine/api_gateway/internal/service"
)

type OrdersHandler struct {
	ordersService *service.OrdersService
}

func NewOrdersHandler(ordersService *service.OrdersService) *OrdersHandler {
	return &OrdersHandler{ordersService: ordersService}
}

func (h *OrdersHandler) CreateOrderPost(c *gin.Context) {
	var req model_orders_requests.OrdersCreateOrderPostRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model_errors.ErrorResponse{
			Error: model_errors.ErrorResponseError{
				Code: "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	domainReq := domain.JsonCreateOrderRequestToDomain(&req)
	if domainReq == nil {
		c.JSON(http.StatusInternalServerError, model_errors.ErrorResponse{
			Error: model_errors.ErrorResponseError{
				Code: "INTERNAL_ERROR",
				Message: "Failed to convert request to domain model",
			},
		})
		return
	}

	resp, err := h.ordersService.CreateOrder(c.Request.Context(), domainReq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, model_errors.ErrorResponse{
			Error: model_errors.ErrorResponseError{
				Code: "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	jsonResp := domain.DomainCreateOrderResponseToJson201Response(resp)

	c.JSON(201, jsonResp)
}

func (h *OrdersHandler) CancelOrderPost(c *gin.Context) {

}

func (h *OrdersHandler) GetOrdersListGet(c *gin.Context) {

}

func (h *OrdersHandler) GetUserOrdersListGet(c *gin.Context) {

}

func (h *OrdersHandler) PayForOrderPost(c *gin.Context) {

}