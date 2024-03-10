package main

import (
	"log"
	"time"

	"github.com/jackboy2fly/password-manager/internal/config"
	"github.com/jackboy2fly/password-manager/internal/database"
	"github.com/jackboy2fly/password-manager/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Init()

	app := fiber.New(fiber.Config{
		AppName:            config.Config("APP_NAME") + " " + config.Config("APP_VERSION"),
		EnableIPValidation: true,
		IdleTimeout:        time.Second * 5,
		ReadTimeout:        time.Second * 2,
	})

	router.Init(app)
	log.Fatal(app.Listen(":" + config.Config("PWM_PORT")))
}
