package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

type Route struct {
	Name string
	Method string
	URLPattern string
	HandlerFunc gin.HandlerFunc
}

func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
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

type ApiHandleFunctions struct {

}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		
	}
}