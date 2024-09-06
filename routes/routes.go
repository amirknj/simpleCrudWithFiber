package routes

import (
	"github.com/amirknj/simpleCrud/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/v1/post", controller.Create)
	//app.Get("api/v1/post/:id")
}
