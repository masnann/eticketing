package main

import (
	"eticketing/config"
	"eticketing/module/feature/route"
	"eticketing/utils/database"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	var initConfig = config.InitConfig()
	db := database.InitPGSDatabase(*initConfig)
	database.Migrate(db)
	route.SetupRoutes(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Ruti Store")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	err := app.Listen(":" + port)
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
