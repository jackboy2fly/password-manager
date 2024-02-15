package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Init(router *fiber.App) {
	router.Use(healthcheck.New())
	router.Use(idempotency.New())
	router.Get("/monitor", monitor.New(monitor.Config{Title: "Brady PWM Monitor"}))

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Brady PWM is up and running!")
	})

}
