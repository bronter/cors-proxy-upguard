package main

import (
	"github.com/bronter/cors-proxy-upguard/resources"
	"github.com/labstack/echo/engine/fasthttp"
)

func main() {
	// Start the server
	println("Listening at :9000")
	resources.NewRouter().Run(fasthttp.New(":9000"))
}
