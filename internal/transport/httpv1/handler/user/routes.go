package user

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

func RegisterRoutes(api huma.API, handler *UserHandler, authSecret string) {
	authMiddleware := huma.Middlewares{
		middleware.AuthMiddleware(api, authSecret),
	}

	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/api/v1/auth/register",
		Summary:       "Регистрация",
		Tags:          []string{"Users"},
		DefaultStatus: http.StatusCreated,
	}, handler.RegisterHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodPost,
		Path:    "/api/v1/auth/login",
		Summary: "Авторизация",
		Tags:    []string{"Users"},
	}, handler.LoginHandler)

	huma.Register(api, huma.Operation{
		Method:      http.MethodPatch,
		Path:        "/api/v1/users/me",
		Summary:     "Обновить профиль",
		Tags:        []string{"Users"},
		Middlewares: authMiddleware,
		Security:    []map[string][]string{{"jwt": {}}},
	}, handler.UpdateProfileHandler)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/api/v1/users/me",
		Summary:     "Получить текущий профиль",
		Tags:        []string{"Users"},
		Middlewares: authMiddleware,
		Security:    []map[string][]string{{"jwt": {}}},
	}, handler.GetMeHandler)
}
