package user_handler

import (
	model "github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"github.com/ALPHACOD3RS/Beauty-Salon/internal/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func RegisterUserHandler(c *fiber.Ctx, db *gorm.DB) error{

	user := new(model.User)

	if err := c.BodyParser(&user); err != nil{

		return c.Status(400).JSON(fiber.Map{
			"msg": "please fill als the neccessary fields!",
		})

	}

	pass, err := utils.HashPassword(string(user.Password))
	if err != nil{
		return c.Status(301).JSON(fiber.Map{
			"msg": "something went wrong",
		})
	}

	user.Password = pass

	if err := db.Create(&user).Error; err != nil{
		return c.Status(400).JSON(fiber.Map{
			"msg": "can not create the user",
		})
	}

	return c.JSON(user)
}