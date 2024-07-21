package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	webApp := fiber.New()
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	webApp.Post("/tasks", CreateTask)
	webApp.Get("/tasks/:id", GetTask)
	webApp.Patch("/tasks/:id", UpdateTask)
	webApp.Delete("/tasks/:id", DeleteTask)

	logrus.Fatal(webApp.Listen(":8080"))
}
