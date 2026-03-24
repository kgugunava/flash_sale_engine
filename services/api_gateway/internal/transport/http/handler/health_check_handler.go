package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {

}

func (h *HealthCheckHandler) HealthCheckAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}
