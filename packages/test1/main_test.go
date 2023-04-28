package main_test

import (
	"github.com/gofiber/fiber/v2"
	"redesigned-pancake/packages/test1"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	t.Setenv("SERVER_ADDR", "http://localhost:3000")
	app := fiber.New()
	app.Get("/api", func(c *fiber.Ctx) error {
		defer func(app *fiber.App) {
			_ = app.Shutdown()
		}(app)
		return c.SendString("Hello")
	})
	go func() {
		time.Sleep(time.Second)
		main.Run()
	}()
	if err := app.Listen(":3000"); err != nil {
		t.Fatal(err)
	}
}
