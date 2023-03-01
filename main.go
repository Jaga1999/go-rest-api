package main

import (
	"github.com/Jaga1999/go-rest-api/user"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")

}

func Routers(app *fiber.App) {
	app.Get("/user", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/User/:id", user.UpdateUser)

}

func main() {

	user.InitialMigration()

	app := fiber.New()

	app.Get("/", hello)
	Routers(app)

	app.Listen(":3000")
}
