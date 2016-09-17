package resources

import (
	routes "github.com/bronter/cors-proxy-upguard/resources/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter - Returns a router with iniialized routes
func NewRouter() *echo.Echo {
	router := echo.New()

	// Deal with CORS
	router.Use(middleware.CORS())

	// Set up logger middleware
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Method: ${method}\tURI: ${uri}\tStatus: ${status}\tLatency: ${latency_human}\n",
	}))

	handlerMap := map[string]func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) {
		"GET":     router.GET,
		"POST":    router.POST,
		"PUT":     router.PUT,
		"DELETE":  router.DELETE,
		"PATCH":   router.PATCH,
		"HEAD":    router.HEAD,
		"OPTIONS": router.OPTIONS,
	}

	for _, route := range routes.Routes {
		// Set up logging for each request
		handler := route.HandlerFunc

		handlerMap[route.Method](route.Pattern, handler)
	}

	return router
}
