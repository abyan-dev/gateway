package main

import (
	"crypto/tls"
	"log"

	"github.com/abyan-dev/gateway/pkg/info"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/valyala/fasthttp"
)

func main() {
	app := fiber.New()

	proxy.WithTlsConfig(&tls.Config{
		InsecureSkipVerify: true,
	})

	proxy.WithClient(&fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		DisablePathNormalizing:   true,
	})

	// HTTP security middleware
	app.Use(helmet.New())

	// Logger
	app.Use(logger.New())

	// Proxies
	app.Use("/api/auth", proxy.Forward("http://auth/api"))

	info.IsRunning()

	log.Fatal(app.Listen(":8080"))
}
