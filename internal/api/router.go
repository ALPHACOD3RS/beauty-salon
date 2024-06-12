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
	v1.Get("/service", func(c *fiber.Ctx) error {
		return user_handler.GetAllServiceHandler(c, db)
	})
	v1.Get("/service/:id", func(c *fiber.Ctx) error {
		return user_handler.GetServiceByIdHandler(c, db)
	})
	
	v1.Post("/appointment", func(c *fiber.Ctx) error {
		return user_handler.CreateAppointmentHandler(c, db)
	})
	v1.Get("/appointment", func(c *fiber.Ctx) error {
		return user_handler.GetAllAppointmentsHandler(c, db)
	})


	

	v1.Use(midleware.AuthMidleware)
	v1.Post("/service", func(c *fiber.Ctx) error {
		return user_handler.CreateServiceHandler(c, db)
	})
	v1.Put("/service/:id", func(c *fiber.Ctx) error {
		return user_handler.UpdateServiceHandler(c, db)
	})
	v1.Delete("/service/:id", func(c *fiber.Ctx) error {
		return user_handler.DeleteServiceHandler(c, db)
	})
	
	
	


	

	
}