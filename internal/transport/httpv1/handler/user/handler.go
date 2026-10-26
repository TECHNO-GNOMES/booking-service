package user

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/seaweedfs"
	"github.com/nurgal1ev/booking-service/internal/service/user"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

type UserHandler struct {
	userService *user.UserService
	fileStorage *seaweedfs.FileStorage
}

func NewUserHandler(userService *user.UserService, fileStorage *seaweedfs.FileStorage) *UserHandler {
	return &UserHandler{userService: userService, fileStorage: fileStorage}
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
	if input.Body.Email == nil && input.Body.Phone == nil {
		return nil, huma.Error400BadRequest("email or phone is required")
	}

	if input.Body.Email != nil && input.Body.Phone != nil {
		return nil, huma.Error400BadRequest("provide only email or phone, not both")
	}

	token, err := u.userService.Login(ctx, input.Body.Email, input.Body.Phone, input.Body.Password)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	resp := &LoginOutput{}
	resp.Body.AccessToken = token

	return resp, nil
}

func (u *UserHandler) UpdateProfileHandler(ctx context.Context, input *UpdateUserInput) (*UpdateUserOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	req := &user.UpdateUserRequest{
		FirstName: input.Body.FirstName,
		LastName:  input.Body.LastName,
		Email:     input.Body.Email,
		City:      input.Body.City,
		Birthday:  input.Body.Birthday,
		Gender:    input.Body.Gender,
		Address:   input.Body.Address,
		Phone:     input.Body.Phone,
	}

	updatedProfile, err := u.userService.Update(ctx, userID, req)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	return &UpdateUserOutput{
		Body: User{
			FirstName: updatedProfile.FirstName,
			LastName:  updatedProfile.LastName,
			Email:     updatedProfile.Email,
			City:      updatedProfile.City,
			Gender:    updatedProfile.Gender,
			Address:   updatedProfile.Address,
			Phone:     updatedProfile.Phone,
		},
	}, nil
}

func (u *UserHandler) UpdateAvatarHandler(ctx context.Context, input *UploadAvatarInput) (*UploadAvatarOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	file := input.RawBody.Data()
	avatar := file.Avatar

	url, err := u.fileStorage.Upload(ctx, userID, avatar, avatar.Filename)
	if err != nil {
		return nil, huma.Error400BadRequest("failed to upload avatar: " + err.Error())
	}

	req := &user.UpdateUserRequest{Avatar: &url}
	_, err = u.userService.Update(ctx, userID, req)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	resp := &UploadAvatarOutput{}
	resp.Body.AvatarURL = url
	return resp, nil
}

func (u *UserHandler) GetMeHandler(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	profile, err := u.userService.GetProfile(ctx, userID)
	if err != nil {
		return nil, huma.Error404NotFound(err.Error())
	}

	return &GetUserOutput{
		Body: User{
			Avatar:    profile.Avatar,
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
			Email:     profile.Email,
			City:      profile.City,
			Gender:    profile.Gender,
			Address:   profile.Address,
			Phone:     profile.Phone,
		},
	}, nil
}
