package main

import (
	"github.com/bronter/cors-proxy-upguard/resources"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	// Start the server
	println("Listening at :9000")
	resources.NewRouter().Run(standard.New(":9000"))
}
