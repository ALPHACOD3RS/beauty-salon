package v1

import (
	model "github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func CreateAppoinemtn(c *fiber.Ctx, db *gorm.DB) error{
	appointment := new(model.Appointment)

	if err := c.BodyParser(&appointment); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"msg":"Invalid data",
		})
	}

	return c.JSON("")
}