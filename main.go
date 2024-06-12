
package main

import (
	"log"

	"github.com/ALPHACOD3RS/Beauty-Salon/internal/api"
	database "github.com/ALPHACOD3RS/Beauty-Salon/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func main(){
	godotenv.Load(".env")
	db := database.InitDatabase()
	app := fiber.New()

	api.SetupRoutes(app, db)


	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(
			fiber.Map{
				"msg": "api is working",
			},
		)
	})

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}

