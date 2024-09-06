package controller

import (
	"github.com/amirknj/simpleCrud/database"
	"github.com/amirknj/simpleCrud/models"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	data := make(map[string]string)

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	post := models.Post{
		Title: data["title"],
		Body:  data["body"],
	}
	database.DB.Create(&post)
	return c.JSON(post)
}
