package booking

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

func RegisterRoutes(api huma.API, handler *BookingHandler, authSecret string) {
	authMiddleware := huma.Middlewares{
		middleware.AuthMiddleware(api, authSecret),
	}

	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/api/v1/bookings",
		Summary:       "Создать бронь",
		Tags:          []string{"Bookings"},
		Description:   "Создаёт новую бронь для текущего пользователя",
		Security:      []map[string][]string{{"jwt": {}}},
		Middlewares:   authMiddleware,
		DefaultStatus: http.StatusCreated,
	}, handler.CreateBookingHandler)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/api/v1/bookings",
		Summary:     "Мои брони",
		Tags:        []string{"Bookings"},
		Description: "Возвращает список броней текущего пользователя",
		Security:    []map[string][]string{{"jwt": {}}},
		Middlewares: authMiddleware,
	}, handler.GetAllBookingHandler)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/api/v1/bookings/{id}",
		Summary:     "Получить бронь",
		Tags:        []string{"Bookings"},
		Description: "Возвращает бронь по ID",
		Security:    []map[string][]string{{"jwt": {}}},
		Middlewares: authMiddleware,
	}, handler.GetBookingHandler)

	huma.Register(api, huma.Operation{
		Method:      http.MethodPatch,
		Path:        "/api/v1/bookings/{id}/cancel",
		Summary:     "Отменить бронь",
		Tags:        []string{"Bookings"},
		Description: "Отменяет бронь. Доступно только владельцу брони.",
		Security:    []map[string][]string{{"jwt": {}}},
		Middlewares: authMiddleware,
	}, handler.CancelBookingHandler)

	huma.Register(api, huma.Operation{
		Method:      http.MethodPatch,
		Path:        "/api/v1/bookings/{id}/confirm",
		Summary:     "Подтвердить бронь",
		Tags:        []string{"Bookings"},
		Description: "Подтверждает бронь. Доступно только администратору.",
		Security:    []map[string][]string{{"jwt": {}}},
		Middlewares: authMiddleware,
	}, handler.ConfirmBookingHandler)
}
