package infrastructure

import (
	"github.com/santduv/gyma-api/internal/database"
	app "github.com/santduv/gyma-api/internal/modules/users/app/use-cases"
	"github.com/santduv/gyma-api/internal/modules/users/domain/ports"
	"github.com/santduv/gyma-api/internal/modules/users/infrastructure/handlers"
	"github.com/santduv/gyma-api/internal/modules/users/infrastructure/repositories"
)

type UserContext struct {
	UserRepository    ports.UserRepository
	CreateUserUseCase *app.CreateUserUseCase
	UserHandler       *handlers.UserHandler
}

func NewUserContext() *UserContext {
	userRepository := repositories.NewUserMongoRepository(database.GetCollection("users"))
	createUserUseCase := app.NewCreateUserUseCase(userRepository)

	return &UserContext{
		UserRepository:    userRepository,
		CreateUserUseCase: createUserUseCase,
		UserHandler:       handlers.NewUserHandler(createUserUseCase),
	}
}
