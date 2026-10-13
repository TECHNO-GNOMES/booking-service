package booking

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/service/booking"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

type BookingHandler struct {
	bookingService *booking.BookingService
}

func NewBookingHandler(bookingService *booking.BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
	}
}

func (h *BookingHandler) Create(ctx context.Context, input *CreateBookingInput) (*CreateBookingOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	booking := &booking.Booking{
		UnitID:     input.Body.UnitID,
		CheckIn:    input.Body.CheckIn,
		CheckOut:   input.Body.CheckOut,
		GuestCount: input.Body.GuestCount,
	}

	result, err := h.bookingService.Create(ctx, userID, booking)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	return &CreateBookingOutput{Body: Booking{
		ID:         result.ID,
		UserID:     result.UserID,
		UnitID:     result.UnitID,
		CheckIn:    result.CheckIn,
		CheckOut:   result.CheckOut,
		TotalPrice: result.TotalPrice,
		Status:     result.Status,
		GuestCount: result.GuestCount,
	}}, nil
}

func (h *BookingHandler) GetAllBookings(ctx context.Context, input *GetBookingInput) (*GetBookingOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	bookings, err := h.bookingService.GetUserBookings(ctx, userID)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	result := make([]Booking, len(bookings))
	for i, booking := range bookings {
		result[i] = Booking{
			ID:         booking.ID,
			UserID:     booking.UserID,
			UnitID:     booking.UnitID,
			CheckIn:    booking.CheckIn,
			CheckOut:   booking.CheckOut,
			TotalPrice: booking.TotalPrice,
			Status:     booking.Status,
			GuestCount: booking.GuestCount,
		}
	}

	return &GetBookingOutput{Body: result}, nil
}
