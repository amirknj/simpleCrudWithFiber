package main

import (
	"github.com/amirknj/simpleCrud/database"
	"github.com/amirknj/simpleCrud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"amir": "123456",
		}}))
	routes.Setup(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}

}
