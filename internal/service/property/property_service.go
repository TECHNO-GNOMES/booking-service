package property

import (
	"context"
	"errors"

	"github.com/nurgal1ev/booking-service/internal/models"
	"github.com/nurgal1ev/booking-service/internal/repository/property"
	"github.com/nurgal1ev/booking-service/internal/repository/user"
)

type PropertyService struct {
	propertyRepo *property.PropertyRepo
	userRepo     *user.UserRepo
}

func NewPropertyService(propertyRepo *property.PropertyRepo, userRepo *user.UserRepo) *PropertyService {
	return &PropertyService{
		propertyRepo: propertyRepo,
		userRepo:     userRepo,
	}
}

func (s *PropertyService) Create(ctx context.Context, p *Property) (*models.Property, error) {
	user, err := s.userRepo.FindByID(ctx, p.OwnerID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	var property = models.Property{
		Name:              p.Name,
		City:              p.City,
		Address:           p.Address,
		DistrictAddress:   p.DistrictAddress,
		MetroAddress:      p.MetroAddress,
		WorkTime:          p.WorkTime,
		Phone:             p.Phone,
		OnlineScreen:      p.OnlineScreen,
		CategoryID:        p.CategoryID,
		LocationTypeID:    p.LocationTypeID,
		FeaturesID:        p.FeaturesID,
		ServicesID:        p.ServicesID,
		EquipmentID:       p.EquipmentID,
		KitchenID:         p.KitchenID,
		ParkingAvailable:  p.ParkingAvailable,
		ParkingSpaces:     p.ParkingSpaces,
		InteriorStyleID:   p.InteriorStyleID,
		HallName:          p.HallName,
		Area:              p.Area,
		GuestCountBanquet: p.GuestCountBanquet,
		GuestCountBuffet:  p.GuestCountBuffet,
		GuestCountTheater: p.GuestCountTheater,
		GuestCountMore:    p.GuestCountMore,
		Description:       p.Description,
		AverageCheckGuest: p.AverageCheckGuest,
		BanquetMenuCheck:  p.BanquetMenuCheck,
		BuffetMenuCheck:   p.BuffetMenuCheck,
		RentalPrice:       p.RentalPrice,
		RentalWeekDays:    p.RentalWeekDays,
		DepositMin:        p.DepositMin,
		DepositInfo:       p.DepositInfo,
		OwnAlcoholAllowed: p.OwnAlcoholAllowed,
		ServiceSupport:    p.ServiceSupport,
		CorkageFee:        p.CorkageFee,
		ExtraCost:         p.ExtraCost,
		ExtraInfo:         p.ExtraInfo,
		Image:             p.Image,
		RoomCategoryID:    p.RoomCategoryID,
		SleepingPlacesID:  p.SleepingPlacesID,
		AmenitiesRoomID:   p.AmenitiesRoomID,
		ServicesRoomID:    p.ServicesRoomID,
		CheckInAfter:      p.CheckInAfter,
		CheckInBefore:     p.CheckInBefore,
		CheckOutAfter:     p.CheckOutAfter,
		CheckOutBefore:    p.CheckOutBefore,
		RoomName:          p.RoomName,
		AccommodationType: p.AccommodationType,
		RoomPrice:         p.RoomPrice,
		RoomInfo:          p.RoomInfo,
		OwnerID:           p.OwnerID,
	}

	err = s.propertyRepo.Create(ctx, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (s *PropertyService) GetById(ctx context.Context, id uint) (*models.Property, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}

	property, err := s.propertyRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if property == nil {
		return nil, err
	}

	return property, nil
}

func (s *PropertyService) GetAll(ctx context.Context, filter PropertyFilter) ([]models.Property, error) {
	if filter.MinPrice != nil && filter.MaxPrice != nil {
		if *filter.MinPrice > *filter.MaxPrice {
			return nil, errors.New("minPrice cannot be greater than maxPrice")
		}
	}

	if filter.CategoryID > 0 {
		if _, ok := models.Category[filter.CategoryID]; !ok {
			return nil, errors.New("invalid category ID")
		}
	}

	if filter.LocationTypeID > 0 {
		if _, ok := models.LocationType[filter.LocationTypeID]; !ok {
			return nil, errors.New("invalid location type ID")
		}
	}

	repoFilter := property.PropertyFilter{
		Search:         filter.Search,
		CategoryID:     filter.CategoryID,
		LocationTypeID: filter.LocationTypeID,
		GuestCountMin:  filter.GuestCountMin,
		MinPrice:       filter.MinPrice,
		MaxPrice:       filter.MaxPrice,
		City:           filter.City,
	}

	properties, err := s.propertyRepo.FindAll(ctx, repoFilter)
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (s *PropertyService) Update(ctx context.Context, id uint, userID uint, userRole string, req *UpdatePropertyRequest) (*models.Property, error) {
	propertyId, err := s.propertyRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if propertyId == nil {
		return nil, errors.New("property not found")
	}

	if propertyId.OwnerID != userID && userRole != "admin" {
		return nil, errors.New("you don't have permission to update this property")
	}

	request := &property.UpdatePropertyRequest{
		Name:         req.Name,
		City:         req.City,
		Address:      req.Address,
		WorkTime:     req.WorkTime,
		Phone:        req.Phone,
		OnlineScreen: req.OnlineScreen,

		CategoryID:     req.CategoryID,
		LocationTypeID: req.LocationTypeID,
		FeaturesID:     req.FeaturesID,
		ServicesID:     req.ServicesID,
		KitchenID:      req.KitchenID,

		ParkingAvailable:  req.ParkingAvailable,
		ParkingSpaces:     req.ParkingSpaces,
		InteriorStyleID:   req.InteriorStyleID,
		HallName:          req.HallName,
		Area:              req.Area,
		GuestCountBanquet: req.GuestCountBanquet,
		GuestCountBuffet:  req.GuestCountBuffet,
		GuestCountTheater: req.GuestCountTheater,
		GuestCountMore:    req.GuestCountMore,
		Description:       req.Description,
		AverageCheckGuest: req.AverageCheckGuest,
		BanquetMenuCheck:  req.BanquetMenuCheck,
		BuffetMenuCheck:   req.BuffetMenuCheck,
		RentalPrice:       req.RentalPrice,
		RentalWeekDays:    req.RentalWeekDays,
		DepositMin:        req.DepositMin,
		DepositInfo:       req.DepositInfo,
		OwnAlcoholAllowed: req.OwnAlcoholAllowed,
		ServiceSupport:    req.ServiceSupport,
		CorkageFee:        req.CorkageFee,
		ExtraCost:         req.ExtraCost,
		ExtraInfo:         req.ExtraInfo,
		RoomCategoryID:    req.RoomCategoryID,
		SleepingPlacesID:  req.SleepingPlacesID,
		AmenitiesRoomID:   req.AmenitiesRoomID,
		ServicesRoomID:    req.ServicesRoomID,
		CheckInAfter:      req.CheckInAfter,
		CheckInBefore:     req.CheckInBefore,
		CheckOutAfter:     req.CheckOutAfter,
		CheckOutBefore:    req.CheckOutBefore,
		RoomName:          req.RoomName,
		AccommodationType: req.AccommodationType,
		RoomPrice:         req.RoomPrice,
		RoomInfo:          req.RoomInfo,
		Image:             req.Image,
	}

	err = s.propertyRepo.Update(ctx, id, request)
	if err != nil {
		return nil, err
	}

	return s.propertyRepo.FindByID(ctx, id)
}

func (s *PropertyService) Delete(ctx context.Context, id uint, userID uint, userRole string) error {
	propertyId, err := s.propertyRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if propertyId == nil {
		return errors.New("property not found")
	}

	if propertyId.OwnerID != userID && userRole != "admin" {
		return errors.New("you don't have permission to delete this property")
	}

	err = s.propertyRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
