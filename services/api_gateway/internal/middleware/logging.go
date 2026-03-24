package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/kgugunava/flash_sale_engine/pkg/logger"
)

func LoggingMiddleware(log *logger.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()

		traceID := c.GetHeader("X-Trace-ID")
		if traceID == "" {
            traceID = uuid.New().String()
        }

		c.Set("trace_id", traceID)
		c.Header("X-Trace-ID", traceID)

		log.Info("HTTP request started",
            logger.String("trace_id", traceID),
            logger.String("method", c.Request.Method),
            logger.String("path", c.Request.URL.Path),
            logger.String("client_ip", c.ClientIP()),
            logger.String("user_agent", c.Request.UserAgent()),
            logger.Time("start_time", start),
        )

		c.Next()

		latency := time.Since(start)
        
        status := c.Writer.Status()
        
        bodySize := c.Writer.Size()
        
        log.Info("HTTP request completed",
            logger.String("trace_id", traceID),
            logger.String("method", c.Request.Method),
            logger.String("path", c.Request.URL.Path),
            logger.Int("status", status),
            logger.Int("body_size", bodySize),
            logger.Duration("latency", latency),
            logger.String("client_ip", c.ClientIP()),
            logger.Time("start_time", start),
            logger.Time("end_time", time.Now()),
        )
	}
	
}