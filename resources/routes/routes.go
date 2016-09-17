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
		"/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"POST",
		"/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"PATCH",
		"/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"PUT",
		"/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"DELETE",
		"/*",
		echo.HandlerFunc(proxy.Proxy),
	},
	Route{
		"ProxyRequest",
		"HEAD",
		"/*",
		echo.HandlerFunc(proxy.Proxy),
	},
}
