package middleware

import (
    "github.com/gin-gonic/gin"
	"github.com/google/uuid"
    "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/trace"
)

func TracingMiddleware(serviceName string) gin.HandlerFunc {
    tracer := otel.Tracer(serviceName)
    
    return func(c *gin.Context) {
        // === ДО ОБРАБОТКИ ЗАПРОСА ===

        traceID := c.GetHeader("X-Trace-ID")
        if traceID == "" {
			traceID = uuid.New().String()
        }
        
        c.Set("trace_id", traceID)
        c.Header("X-Trace-ID", traceID)
        
        ctx, span := tracer.Start(
            c.Request.Context(),
            c.Request.URL.Path, 
            trace.WithSpanKind(trace.SpanKindServer), 
            trace.WithAttributes(
                attribute.String("http.method", c.Request.Method),
                attribute.String("http.url", c.Request.URL.String()),
                attribute.String("http.client_ip", c.ClientIP()),
                attribute.String("http.user_agent", c.Request.UserAgent()),
                
                attribute.String("trace_id", traceID),
                attribute.String("service.name", serviceName),
            ),
        )
        
        c.Request = c.Request.WithContext(ctx)
        
        c.Next()
        
        // === ПОСЛЕ ОБРАБОТКИ ЗАПРОСА ===

        span.SetAttributes(
            attribute.Int("http.status_code", c.Writer.Status()),
            attribute.Int("http.response_size", c.Writer.Size()),
        )
        
        if c.Writer.Status() >= 500 {
            span.SetStatus(codes.Code(c.Writer.Status()), "Internal server error")
            for _, err := range c.Errors {
                span.RecordError(err)
            }
        }
        
        span.End()
    }
}