package route

import (
	"eticketing/module/feature/auth"
	"eticketing/module/feature/schedule"
	users "eticketing/module/feature/user"
	user "eticketing/module/feature/user/domain"
	"eticketing/utils/token"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwt token.JWTInterface, userService user.UserServiceInterface) {
	users.InitializeUser(db)
	users.SetupRoutesUser(app)
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
	schedule.InitializeSchedule(db)
	schedule.SetupRoutesSchedule(app, jwt, userService)
}
