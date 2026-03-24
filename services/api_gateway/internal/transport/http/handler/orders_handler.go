package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersHandler struct {

}

func NewOrdersHandler() *OrdersHandler {
	return &OrdersHandler{}
}

func (h *OrdersHandler) CreateOrderPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
	
}

func (h *OrdersHandler) CancelOrderPost(c *gin.Context) {

}

func (h *OrdersHandler) GetOrdersListGet(c *gin.Context) {

}

func (h *OrdersHandler) GetUserOrdersListGet(c *gin.Context) {

}

func (h *OrdersHandler) PayForOrderPost(c *gin.Context) {

}