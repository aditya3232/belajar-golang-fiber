package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second, // diatas 5 detik write  maka timeout, stop requestnya
		ReadTimeout:  5 * time.Second,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
