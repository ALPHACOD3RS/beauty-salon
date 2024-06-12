package api

import (
	"github.com/ALPHACOD3RS/Beauty-Salon/internal/api/midleware"
	user_handler "github.com/ALPHACOD3RS/Beauty-Salon/internal/api/v1"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	api := app.Group("/api")
    v1 := api.Group("/v1")
	//auth api
	v1.Post("/register", func(c *fiber.Ctx) error {
		return user_handler.RegisterUserHandler(c, db)
	})

	v1.Post("/login", func(c *fiber.Ctx) error {
		return user_handler.LoginHandler(c, db)
	})

	v1.Use(midleware.AuthMidleware)
	


	

	
}