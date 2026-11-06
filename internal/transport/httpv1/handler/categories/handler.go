package categories

import (
	"context"

	"github.com/nurgal1ev/booking-service/internal/models"
)

type CategoriesHandler struct{}

func NewCategoriesHandler() *CategoriesHandler {
	return &CategoriesHandler{}
}

func (h *CategoriesHandler) GetAllCategoriesHandler(ctx context.Context, input *GetCategoriesInput) (*GetCategoriesOutput, error) {
	result := map[string]any{
		"category":        models.Category,
		"location_type":   models.LocationType,
		"features":        models.Features,
		"services":        models.Services,
		"equipment":       models.Equipment,
		"kitchen":         models.Kitchen,
		"interior_style":  models.InteriorStyle,
		"room_category":   models.RoomCategory,
		"sleeping_places": models.SleepingPlaces,
		"services_room":   models.ServicesRoom,
		"amenities_room":  models.AmenitiesRoom,
	}

	return &GetCategoriesOutput{Body: result}, nil
}
