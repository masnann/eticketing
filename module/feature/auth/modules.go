package auth

import (
	"eticketing/module/feature/auth/domain"
	utils "eticketing/utils/hash"
	"eticketing/utils/token"
	"os"

	"eticketing/module/feature/auth/handler"
	"eticketing/module/feature/auth/repository"
	"eticketing/module/feature/auth/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	repo domain.AuthRepositoryInterface
	serv domain.AuthServiceInterface
	hand domain.AuthHandlerInterface
	hash utils.HashInterface
	jwt  token.JWTInterface
)

func InitializeAuth(db *gorm.DB) {
	secret := os.Getenv("SECRET")
	hash = utils.NewHash()
	jwt = token.NewJWT(secret)

	repo = repository.NewAuthRepository(db)
	serv = service.NewAuthService(repo, hash, jwt)
	hand = handler.NewAuthHandler(serv)
}

func SetupRoutesAuth(app *fiber.App) {
	api := app.Group("/api/v1/auth")
	api.Post("/login", hand.Login)
	api.Post("/register", hand.Register)
}
