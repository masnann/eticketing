package main

import (
	"eticketing/config"
	"eticketing/module/feature/middleware"
	"eticketing/module/feature/route"
	"eticketing/module/feature/user/repository"
	"eticketing/module/feature/user/service"
	"eticketing/utils/database"
	"eticketing/utils/token"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	var initConfig = config.InitConfig()
	db := database.InitPGSDatabase(*initConfig)
	jwtService := token.NewJWT(initConfig.Secret)

	middleware.SetupMiddlewares(app)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	database.Migrate(db)
	route.SetupRoutes(app, db, jwtService, userService)

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
