package schedule

import (
	"eticketing/module/feature/middleware"
	"eticketing/module/feature/schedule/handler"
	"eticketing/module/feature/schedule/repository"
	"eticketing/module/feature/schedule/service"
	"eticketing/utils/token"

	"eticketing/module/feature/schedule/domain"
	user "eticketing/module/feature/user/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	repo domain.ScheduleRepositoryInterface
	serv domain.ScheduleServiceInterface
	hand domain.ScheduleHandlerInterface
)

func InitializeSchedule(db *gorm.DB) {
	repo = repository.NewScheduleRepository(db)
	serv = service.NewScheduleService(repo)
	hand = handler.NewScheduleHandler(serv)
}

func SetupRoutesSchedule(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/schedule")
	api.Get("/list", hand.GetAllSchedules)
	api.Post("/create", middleware.AuthMiddleware(jwt, userService), hand.CreateSchedule)
	api.Put("/update/:id", middleware.AuthMiddleware(jwt, userService), hand.UpdateSchedule)
	api.Delete("/delete/:id", middleware.AuthMiddleware(jwt, userService), hand.DeleteSchedule)
}
