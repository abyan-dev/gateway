package main

import (
	"log"

	"github.com/abyan-dev/gateway/pkg/info"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.All("/api/auth/*", func(c *fiber.Ctx) error {
		targetURL := "http://auth:8080/api" + c.OriginalURL()[len("/api/auth"):]
		if err := proxy.Do(c, targetURL); err != nil {
			return err
		}
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	info.IsRunning()

	log.Fatal(app.Listen(":8080"))
}
