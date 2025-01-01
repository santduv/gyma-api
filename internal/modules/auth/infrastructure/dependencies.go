package infrastructure

import (
	"github.com/santduv/gyma-api/internal/database"
	app "github.com/santduv/gyma-api/internal/modules/auth/app/use-cases"
	"github.com/santduv/gyma-api/internal/modules/auth/infrastructure/handlers"
	"github.com/santduv/gyma-api/internal/modules/users/domain/ports"
	"github.com/santduv/gyma-api/internal/modules/users/infrastructure/repositories"
)

type AuthContext struct {
	AuthHandler    *handlers.AuthHandler
	LoginUseCase   *app.LoginUseCase
	UserRepository ports.UserRepository
}

func NewAuthContext() *AuthContext {
	userRepository := repositories.NewUserMongoRepository(database.GetCollection("users"))
	loginUseCase := app.NewLoginUseCase(userRepository)

	authHandler := handlers.NewAuthHandler(loginUseCase)

	return &AuthContext{
		AuthHandler:    authHandler,
		LoginUseCase:   loginUseCase,
		UserRepository: userRepository,
	}
}
