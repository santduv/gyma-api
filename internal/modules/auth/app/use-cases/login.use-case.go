package app

import (
	"context"
	"errors"
	"regexp"

	"github.com/golang-jwt/jwt/v5"
	"github.com/santduv/gyma-api/internal/config"
	"github.com/santduv/gyma-api/internal/modules/auth/app/dto"
	"github.com/santduv/gyma-api/internal/modules/shared/app/constants"
	httpErrors "github.com/santduv/gyma-api/internal/modules/shared/app/http-errors"
	"github.com/santduv/gyma-api/internal/modules/shared/app/types"
	"github.com/santduv/gyma-api/internal/modules/users/domain/entities"
	"github.com/santduv/gyma-api/internal/modules/users/domain/ports"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepository ports.UserRepository
	ctx            context.Context
}

func NewLoginUseCase(userRepository ports.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepository: userRepository,
	}
}

func (u *LoginUseCase) Execute(ctx context.Context, loginDto dto.LoginDto) (*types.ApiResponse, *httpErrors.HttpError) {
	u.ctx = ctx

	err := u.validateDto(loginDto)

	if err != nil {
		return nil, httpErrors.NewBadRequestError("invalid request body", &types.JsonMap{
			"error": err.Error(),
		})
	}

	user, errGetUser := u.getUser(loginDto.Email)

	if errGetUser != nil {
		return nil, errGetUser
	}

	err = u.validatePassword(loginDto.Password, user.Password)

	if err != nil {
		return nil, httpErrors.NewUnauthorizedError("invalid credentials", nil)
	}

	token, err := u.generateJwtToken(user)

	if err != nil {
		return nil, httpErrors.NewInternalServerError("failed to generate token", &types.JsonMap{
			"error": err.Error(),
		})
	}

	res := types.ApiResponse{
		Status:  constants.HTTP_STATUS_OK,
		Message: "login successful",
		Data: &dto.AuthResponseDto{
			AccessToken: token,
		},
	}

	return &res, nil
}

func (u *LoginUseCase) validateDto(loginDto dto.LoginDto) error {
	if loginDto.Email == "" {
		return errors.New("email is required")
	}

	if loginDto.Password == "" {
		return errors.New("password is required")
	}

	if len(loginDto.Password) < 8 {
		return errors.New("password should be at least 8 characters")
	}

	// Regex for email validation
	regex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

	match, err := regexp.MatchString(regex, loginDto.Email)

	if err != nil {
		return errors.New("error validating email: " + err.Error())
	}

	if !match {
		return errors.New("invalid email format")
	}

	return nil

}

func (u *LoginUseCase) getUser(email string) (*entities.User, *httpErrors.HttpError) {
	user, err := u.userRepository.FindOne(u.ctx, &ports.FindUser{
		Email: &email,
	})

	if err != nil {
		return nil, httpErrors.NewInternalServerError("failed to get user", &types.JsonMap{
			"error": err.Error(),
		})
	}

	if user == nil {
		return nil, httpErrors.NewUnauthorizedError("invalid credentials", nil)
	}

	return user, nil
}

func (u *LoginUseCase) validatePassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return errors.New("invalid password")
	}

	return nil
}

func (u *LoginUseCase) generateJwtToken(user *entities.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID.Hex()
	claims["nickname"] = user.Nickname
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte(config.Envs.JWTSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
