package api

import (
	user_handler "github.com/ALPHACOD3RS/Beauty-Salon/internal/api/v1"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	api := app.Group("/api")
    v1 := api.Group("/v1")
	v1.Post("/register", func(c *fiber.Ctx) error {
		return user_handler.RegisterUserHandler(c, db)
	})

	
}