package v1

import (
	"fmt"
	"os"
	"time"

	"github.com/ALPHACOD3RS/Beauty-Salon/internal/database"
	"github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"github.com/ALPHACOD3RS/Beauty-Salon/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RegisterUserHandler(c *fiber.Ctx) error {
	newUser := new(models.User)

	db := database.InitDatabase()

	userID := uuid.New().String()

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Please fill all the necessary fields!",
			"err": err.Error(),
		})
	}

	// Set default role if not provided
	if newUser.Role == "" {
		newUser.Role = models.CustomerRole
	}

	pass, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"msg": "Something went wrong",
			"err": err.Error(),
		})
	}

	newUser.Password = pass
	newUser.UserID = userID


	if err := db.Create(&newUser).Error; err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		fmt.Printf("User data: %+v\n", newUser)

		return c.Status(400).JSON(fiber.Map{
			"msg": "Cannot create the user",
			"err": err.Error(),
		})
	}

	return c.JSON(newUser)
}

func LoginHandler(c *fiber.Ctx, db *gorm.DB) error{

	inputUser := new(models.User)

	if err := c.BodyParser(&inputUser); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"msg": "something wen wrong",
		})
	}

	var user models.User

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
		"id": user.UserID,
		"email": user.Email,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil{
		return err
	}
	
	response := struct {
        User  models.User `json:"user"`
        Token string     `json:"token"`
    }{
        User:  user,
        Token: signedToken,
    }

	return c.Status(200).JSON(response)
	
}