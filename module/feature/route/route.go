package route

import (
	"eticketing/module/feature/auth"
	users "eticketing/module/feature/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	users.InitializeUser(db)
	users.SetupRoutesUser(app)
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
}
