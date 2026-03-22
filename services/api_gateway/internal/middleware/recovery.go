package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/kgugunava/flash_sale_engine/pkg/logger"
	models_errors "github.com/kgugunava/flash_sale_engine/api_gateway/internal/model/errors"
)

func RecoveryMiddleware(log logger.Logger) gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {

			var stack []byte
			err := recover()

			if err != nil {
				stack = debug.Stack()
			}

			log.Error("Panic recovered", 
				logger.Any("error", err),
				logger.String("stack", string(stack)),
				logger.String("path", c.Request.URL.Path),
				logger.String("method", c.Request.Method),
				logger.String("client_ip", c.ClientIP()),
			)

			c.JSON(http.StatusInternalServerError, models_errors.ErrorResponse{
				Error: models_errors.ErrorResponseError{
					Code: "INTERNAL_ERROR",
					Message: "Internal server error",
				},
			})

			c.Abort()
		}()
		
		c.Next()
	}
}