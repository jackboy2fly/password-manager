package main

import (
	"log"
	"time"

	"github.com/jackboy2fly/password-manager/internal/config"
	"github.com/jackboy2fly/password-manager/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:            "Brady PWM",
		EnableIPValidation: true,
		IdleTimeout:        time.Second * 5,
		ReadTimeout:        time.Second * 2,
	})

	router.Init(app)
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
