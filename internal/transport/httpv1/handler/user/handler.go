package user

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/service/user"
)

type UserHandler struct {
	userService *user.UserService
}

func NewUserHandler(userService *user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) RegisterHandler(ctx context.Context, input *RegisterInput) (*RegisterOutput, error) {
	user := &user.User{
		Email:    input.Body.Email,
		Password: input.Body.Password,
	}

	createdUser, err := u.userService.Create(ctx, user)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	token, err := u.userService.GenerateToken(createdUser.ID, createdUser.Email, createdUser.Role)
	if err != nil {
		return nil, err
	}

	resp := &RegisterOutput{}
	resp.Body.AccessToken = token
	resp.Body.Message = "successful registration"

	return resp, nil
}

func (u *UserHandler) LoginHandler(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	token, err := u.userService.Login(ctx, input.Body.Email, input.Body.Password)
	if err != nil {
		return nil, err
	}

	resp := &LoginOutput{}
	resp.Body.AccessToken = token

	return resp, nil
}

func (u *UserHandler)
