package user_handler

import (
	"os"
	"time"

	model "github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"github.com/ALPHACOD3RS/Beauty-Salon/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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


func LoginHandler(c *fiber.Ctx, db *gorm.DB) error{

	inputUser := new(model.User)

	if err := c.BodyParser(&inputUser); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"msg": "something wen wrong",
		})
	}

	var user model.User

	if err := db.Where("email = ?", inputUser.Email).First(&user).Error; err != nil{
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid credential ",
		})
	}

	if !utils.VerifyHashedPassword(inputUser.Password, user.Password){
		return c.Status(400).JSON(fiber.Map{
			"msg": "password or email is incorrect",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"email": user.Email,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil{
		return err
	}
	
	response := struct {
        User  model.User `json:"user"`
        Token string     `json:"token"`
    }{
        User:  user,
        Token: signedToken,
    }

	return c.Status(200).JSON(response)
	
}