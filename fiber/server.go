package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, World")
		return c.JSON(fiber.Map{"name": "Hello", "age": 20})
	})

	app.Listen(":3000")
}
