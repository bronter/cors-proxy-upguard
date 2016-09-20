package routes

import (
	proxy "github.com/bronter/cors-proxy-upguard/resources/proxy"
	"github.com/labstack/echo"
)

// Route - Struct containing all the info to initialize a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// Routes - Array of routes
var Routes = []Route{
	Route{
		"ProxyRequest",
		"GET",
		"/api/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"POST",
		"/api/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"PATCH",
		"/api/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"PUT",
		"/api/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"DELETE",
		"/api/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"HEAD",
		"/api/*",
		echo.HandlerFunc(proxy.Proxy),
	},
}
