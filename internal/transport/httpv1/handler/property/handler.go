package property

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/models"
	"github.com/nurgal1ev/booking-service/internal/service/property"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

type PropertyHandler struct {
	propertyService *property.PropertyService
}

func NewPropertyHandler(propertyService *property.PropertyService) *PropertyHandler {
	return &PropertyHandler{
		propertyService: propertyService,
	}
}

func (h *PropertyHandler) CreatePropertyHandler(ctx context.Context, input *CreatePropertyInput) (*CreatePropertyOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	/*
		if input.Body.Name == "" {
			return nil, huma.Error400BadRequest("name is required")
		}

		if input.Body.PricePerNight <= 0 {
			return nil, huma.Error400BadRequest("price per night must be greater than 0")
		}

		if input.Body.PropertyType == "" {
			return nil, huma.Error400BadRequest("property type is required")
		}*/

	property := &property.Property{
		Name:              input.Body.Name,
		City:              input.Body.City,
		Address:           input.Body.Address,
		DistrictAddress:   input.Body.DistrictAddress,
		MetroAddress:      input.Body.MetroAddress,
		WorkTime:          input.Body.WorkTime,
		Phone:             input.Body.Phone,
		OnlineScreen:      input.Body.OnlineScreen,
		CategoryID:        input.Body.CategoryID,
		LocationTypeID:    input.Body.LocationTypeID,
		FeaturesID:        input.Body.FeaturesID,
		ServicesID:        input.Body.ServicesID,
		EquipmentID:       input.Body.EquipmentID,
		KitchenID:         input.Body.KitchenID,
		ParkingAvailable:  input.Body.ParkingAvailable,
		ParkingSpaces:     input.Body.ParkingSpaces,
		InteriorStyleID:   input.Body.InteriorStyleID,
		HallName:          input.Body.HallName,
		Area:              input.Body.Area,
		GuestCountBanquet: input.Body.GuestCountBanquet,
		GuestCountBuffet:  input.Body.GuestCountBuffet,
		GuestCountTheater: input.Body.GuestCountTheater,
		GuestCountMore:    input.Body.GuestCountMore,
		Description:       input.Body.Description,
		AverageCheckGuest: input.Body.AverageCheckGuest,
		BanquetMenuCheck:  input.Body.BanquetMenuCheck,
		BuffetMenuCheck:   input.Body.BuffetMenuCheck,
		RentalPrice:       input.Body.RentalPrice,
		RentalWeekDays:    input.Body.RentalWeekDays,
		DepositMin:        input.Body.DepositMin,
		DepositInfo:       input.Body.DepositInfo,
		OwnAlcoholAllowed: (*models.OwnAlcoholAllowed)(&input.Body.OwnAlcoholAllowed),
		ServiceSupport:    (*models.ServiceSupport)(&input.Body.ServiceSupport),
		CorkageFee:        input.Body.CorkageFee,
		ExtraCost:         input.Body.ExtraCost,
		ExtraInfo:         input.Body.ExtraInfo,

		Image: input.Body.Image,

		RoomCategoryID:    input.Body.RoomCategoryID,
		SleepingPlacesID:  input.Body.SleepingPlacesID,
		AmenitiesRoomID:   input.Body.AmenitiesRoomID,
		ServicesRoomID:    input.Body.ServicesRoomID,
		CheckInAfter:      input.Body.CheckInAfter,
		CheckInBefore:     input.Body.CheckInBefore,
		CheckOutAfter:     input.Body.CheckOutAfter,
		CheckOutBefore:    input.Body.CheckOutBefore,
		RoomName:          input.Body.RoomName,
		AccommodationType: (models.AccommodationType)(input.Body.AccommodationType),
		RoomPrice:         input.Body.RoomPrice,
		RoomInfo:          input.Body.RoomInfo,
		OwnerID:           userID,
	}

	result, err := h.propertyService.Create(ctx, property)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	return &CreatePropertyOutput{
		Body: struct {
			ID uint `json:"id"`
		}{ID: result.ID},
	}, nil
}

func (h *PropertyHandler) GetPropertyHandler(ctx context.Context, input *GetPropertyInput) (*GetPropertyOutput, error) {
	property, err := h.propertyService.GetById(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("property not found")
	}

	if property == nil {
		return nil, huma.Error404NotFound("property not found")
	}

	var ownAlcoholAllowed string
	if property.OwnAlcoholAllowed != nil {
		ownAlcoholAllowed = string(*property.OwnAlcoholAllowed)
	}

	var serviceSupport string
	if property.ServiceSupport != nil {
		serviceSupport = string(*property.ServiceSupport)
	}

	return &GetPropertyOutput{
		Body: Property{
			Name:              property.Name,
			City:              property.City,
			Address:           property.Address,
			DistrictAddress:   property.DistrictAddress,
			MetroAddress:      property.MetroAddress,
			WorkTime:          property.WorkTime,
			Phone:             property.Phone,
			OnlineScreen:      property.OnlineScreen,
			CategoryID:        property.CategoryID,
			LocationTypeID:    property.LocationTypeID,
			FeaturesID:        property.FeaturesID,
			ServicesID:        property.ServicesID,
			EquipmentID:       property.EquipmentID,
			KitchenID:         property.KitchenID,
			ParkingAvailable:  property.ParkingAvailable,
			ParkingSpaces:     property.ParkingSpaces,
			InteriorStyleID:   property.InteriorStyleID,
			HallName:          property.HallName,
			Area:              property.Area,
			GuestCountBanquet: property.GuestCountBanquet,
			GuestCountBuffet:  property.GuestCountBuffet,
			GuestCountTheater: property.GuestCountTheater,
			GuestCountMore:    property.GuestCountMore,
			Description:       property.Description,
			AverageCheckGuest: property.AverageCheckGuest,
			BanquetMenuCheck:  property.BanquetMenuCheck,
			BuffetMenuCheck:   property.BuffetMenuCheck,
			RentalPrice:       property.RentalPrice,
			RentalWeekDays:    property.RentalWeekDays,
			DepositMin:        property.DepositMin,
			DepositInfo:       property.DepositInfo,
			OwnAlcoholAllowed: ownAlcoholAllowed,
			ServiceSupport:    serviceSupport,
			CorkageFee:        property.CorkageFee,
			ExtraCost:         property.ExtraCost,
			ExtraInfo:         property.ExtraInfo,
			Image:             property.Image,
			RoomCategoryID:    property.RoomCategoryID,
			SleepingPlacesID:  property.SleepingPlacesID,
			AmenitiesRoomID:   property.AmenitiesRoomID,
			ServicesRoomID:    property.ServicesRoomID,
			CheckInAfter:      property.CheckInAfter,
			CheckInBefore:     property.CheckInBefore,
			CheckOutAfter:     property.CheckOutAfter,
			CheckOutBefore:    property.CheckOutBefore,
			RoomName:          property.RoomName,
			AccommodationType: string(property.AccommodationType),
			RoomPrice:         property.RoomPrice,
			RoomInfo:          property.RoomInfo,
		},
	}, nil
}

func (h *PropertyHandler) UpdatePropertyHandler(ctx context.Context, input *UpdatePropertyInput) (*UpdatePropertyOutput, error) {
	propertyID, err := h.propertyService.GetById(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("property not found")
	}

	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	userRole, _ := middleware.GetUserRoleFromContext(ctx)

	if propertyID.OwnerID != userID {
		return nil, huma.Error403Forbidden("forbidden")
	}

	req := &property.UpdatePropertyRequest{
		Name:              input.Body.Name,
		City:              input.Body.City,
		Address:           input.Body.Address,
		WorkTime:          input.Body.WorkTime,
		Phone:             input.Body.Phone,
		OnlineScreen:      input.Body.OnlineScreen,
		CategoryID:        input.Body.CategoryID,
		LocationTypeID:    input.Body.LocationTypeID,
		FeaturesID:        input.Body.FeaturesID,
		ServicesID:        input.Body.ServicesID,
		EquipmentID:       input.Body.EquipmentID,
		KitchenID:         input.Body.KitchenID,
		ParkingAvailable:  input.Body.ParkingAvailable,
		ParkingSpaces:     input.Body.ParkingSpaces,
		InteriorStyleID:   input.Body.InteriorStyleID,
		HallName:          input.Body.HallName,
		Area:              input.Body.Area,
		GuestCountBanquet: input.Body.GuestCountBanquet,
		GuestCountBuffet:  input.Body.GuestCountBuffet,
		GuestCountTheater: input.Body.GuestCountTheater,
		GuestCountMore:    input.Body.GuestCountMore,
		Description:       input.Body.Description,
		AverageCheckGuest: input.Body.AverageCheckGuest,
		BanquetMenuCheck:  input.Body.BanquetMenuCheck,
		BuffetMenuCheck:   input.Body.BuffetMenuCheck,
		RentalPrice:       input.Body.RentalPrice,
		RentalWeekDays:    input.Body.RentalWeekDays,
		DepositMin:        input.Body.DepositMin,
		DepositInfo:       input.Body.DepositInfo,
		OwnAlcoholAllowed: (*models.OwnAlcoholAllowed)(input.Body.OwnAlcoholAllowed),
		ServiceSupport:    (*models.ServiceSupport)(input.Body.ServiceSupport),
		CorkageFee:        input.Body.CorkageFee,
		ExtraCost:         input.Body.ExtraCost,
		ExtraInfo:         input.Body.ExtraInfo,
		RoomCategoryID:    input.Body.RoomCategoryID,
		SleepingPlacesID:  input.Body.SleepingPlacesID,
		AmenitiesRoomID:   input.Body.AmenitiesRoomID,
		ServicesRoomID:    input.Body.ServicesRoomID,
		CheckInAfter:      input.Body.CheckInAfter,
		CheckInBefore:     input.Body.CheckInBefore,
		CheckOutAfter:     input.Body.CheckOutAfter,
		CheckOutBefore:    input.Body.CheckOutBefore,
		RoomName:          input.Body.RoomName,
		Image:             (*models.Image)(input.Body.Image),
	}

	updatedProperty, err := h.propertyService.Update(ctx, input.ID, userID, userRole, req)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	var ownAlcohol string
	if updatedProperty.OwnAlcoholAllowed != nil {
		ownAlcohol = string(*updatedProperty.OwnAlcoholAllowed)
	}

	var serviceSupport string
	if updatedProperty.ServiceSupport != nil {
		serviceSupport = string(*updatedProperty.ServiceSupport)
	}

	return &UpdatePropertyOutput{
		Body: Property{
			Name:              updatedProperty.Name,
			City:              updatedProperty.City,
			Address:           updatedProperty.Address,
			DistrictAddress:   updatedProperty.DistrictAddress,
			MetroAddress:      updatedProperty.MetroAddress,
			WorkTime:          updatedProperty.WorkTime,
			Phone:             updatedProperty.Phone,
			OnlineScreen:      updatedProperty.OnlineScreen,
			CategoryID:        updatedProperty.CategoryID,
			LocationTypeID:    updatedProperty.LocationTypeID,
			FeaturesID:        updatedProperty.FeaturesID,
			ServicesID:        updatedProperty.ServicesID,
			EquipmentID:       updatedProperty.EquipmentID,
			KitchenID:         updatedProperty.KitchenID,
			ParkingAvailable:  updatedProperty.ParkingAvailable,
			ParkingSpaces:     updatedProperty.ParkingSpaces,
			InteriorStyleID:   updatedProperty.InteriorStyleID,
			HallName:          updatedProperty.HallName,
			Area:              updatedProperty.Area,
			GuestCountBanquet: updatedProperty.GuestCountBanquet,
			GuestCountBuffet:  updatedProperty.GuestCountBuffet,
			GuestCountTheater: updatedProperty.GuestCountTheater,
			GuestCountMore:    updatedProperty.GuestCountMore,
			Description:       updatedProperty.Description,
			AverageCheckGuest: updatedProperty.AverageCheckGuest,
			BanquetMenuCheck:  updatedProperty.BanquetMenuCheck,
			BuffetMenuCheck:   updatedProperty.BuffetMenuCheck,
			RentalPrice:       updatedProperty.RentalPrice,
			RentalWeekDays:    updatedProperty.RentalWeekDays,
			DepositMin:        updatedProperty.DepositMin,
			DepositInfo:       updatedProperty.DepositInfo,
			OwnAlcoholAllowed: ownAlcohol,
			ServiceSupport:    serviceSupport,
			CorkageFee:        updatedProperty.CorkageFee,
			ExtraCost:         updatedProperty.ExtraCost,
			ExtraInfo:         updatedProperty.ExtraInfo,
			Image:             updatedProperty.Image,
			RoomCategoryID:    updatedProperty.RoomCategoryID,
			SleepingPlacesID:  updatedProperty.SleepingPlacesID,
			AmenitiesRoomID:   updatedProperty.AmenitiesRoomID,
			ServicesRoomID:    updatedProperty.ServicesRoomID,
			CheckInAfter:      updatedProperty.CheckInAfter,
			CheckInBefore:     updatedProperty.CheckInBefore,
			CheckOutAfter:     updatedProperty.CheckOutAfter,
			CheckOutBefore:    updatedProperty.CheckOutBefore,
			RoomName:          updatedProperty.RoomName,
			AccommodationType: string(updatedProperty.AccommodationType),
			RoomPrice:         updatedProperty.RoomPrice,
			RoomInfo:          updatedProperty.RoomInfo,
		},
	}, nil
}

func (h *PropertyHandler) DeletePropertyHandler(ctx context.Context, input *DeletePropertyInput) (*DeletePropertyOutput, error) {
	propertyID, err := h.propertyService.GetById(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("property not found")
	}

	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	userRole, _ := middleware.GetUserRoleFromContext(ctx)

	err = h.propertyService.Delete(ctx, propertyID.ID, userID, userRole)
	if err != nil {
		return nil, err
	}

	return &DeletePropertyOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "property deleted successfully",
		},
	}, nil
}
