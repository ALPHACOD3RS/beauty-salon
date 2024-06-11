package main

import (
	database "github.com/ALPHACOD3RS/Beauty-Salon/internal/db"
	"github.com/gofiber/fiber/v2"
)


func main(){
	database.InitDatabase()

	

	app := fiber.New()


	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("working")
	})

	app.Listen(":8000")

}