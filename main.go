package main

import (
	"github.com/bronter/cors-proxy-upguard/resources"
	"github.com/labstack/echo/engine/fasthttp"
)

func main() {
	// File path for TLS Cert
	tlsCertFile := ""
	// File path for TLS key
	tlsKeyFile := ""
	// Start the server
	println("Listening at :80")
	resources.NewRouter().Run(fasthttp.WithTLS(":80", tlsCertFile, tlsKeyFile))
}
