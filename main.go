package main

import (
	"log"

	"github.com/ALPHACOD3RS/Beauty-Salon/internal/api"
	database "github.com/ALPHACOD3RS/Beauty-Salon/internal/db"
	"github.com/gofiber/fiber/v2"
)


func main(){
	db := database.InitDatabase()

	

	app := fiber.New()

	// api

	api.SetupRoutes(app, db)


	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("working")
	})

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}