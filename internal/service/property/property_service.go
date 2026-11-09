package property

import (
	"context"
	"errors"
	"time"

	"github.com/nurgal1ev/booking-service/internal/models"
	"github.com/nurgal1ev/booking-service/internal/repository/property"
	"github.com/nurgal1ev/booking-service/internal/repository/user"
)

type Property struct {
	ID uint
	// 1 шаг - создание карточки
	Name            string     // название площадки
	City            string     // ваш город
	Address         *string    // адрес площадки
	DistrictAddress *string    // район
	MetroAddress    *string    // метро
	WorkTime        *time.Time // время работы
	Phone           string     // номер телефона
	OnlineScreen    bool       // онлайн показ есть/нет

	// 2 шаг - создание карточки
	CategoryID     uint // категория
	LocationTypeID uint // тип места
	FeaturesID     uint // особенности
	ServicesID     uint // услуги
	EquipmentID    uint // оборудование в наличии
	KitchenID      uint // кухня

	ParkingAvailable bool // парковка
	ParkingSpaces    int

	InteriorStyleID *uint // стиль интерьера

	// залы
	HallName          *string // название зала
	Area              *string // площадь зала кв. м.
	GuestCountBanquet *int    // кол-во мест банкет
	GuestCountBuffet  *int    // кол-во мест фуршет
	GuestCountTheater *int    // кол-во мест рассадка театр
	GuestCountMore    *int    // кол-во мест/еще
	Description       string  // описание площадки

	// цены
	AverageCheckGuest *int                      // средний чек на гостя (сумма в рублях)
	BanquetMenuCheck  *int                      // банкетное меню на гостя (сумма в рублях)
	BuffetMenuCheck   *int                      // фуршетное меню на гостя (сумма в рублях)
	RentalPrice       *int                      // аренда/час
	RentalWeekDays    *int                      // аренда площадка под мероприятия в будни/выходные и праздничные дни
	DepositMin        *string                   // депозит (минимальная стоимость закрытия площадки, зала под одного мероприятие)
	DepositInfo       *string                   // что входит в депозит
	OwnAlcoholAllowed *models.OwnAlcoholAllowed // возможность своего алкоголя
	ServiceSupport    *models.ServiceSupport    // сервисное обслуживание
	CorkageFee        *string                   // пробковый сбор
	ExtraCost         *string                   // дополнительные затраты
	ExtraInfo         *string                   // дополнительно

	Image string // медиатека

	// номерной фонд
	RoomCategoryID    uint                     // категории номеров
	SleepingPlacesID  uint                     // спальные места
	AmenitiesRoomID   uint                     // удобство в номерах
	ServicesRoomID    uint                     // услуги
	CheckInAfter      *time.Time               //время заезда (от)
	CheckInBefore     *time.Time               //время заезда (до)
	CheckOutAfter     *time.Time               //время выезда (от)
	CheckOutBefore    *time.Time               //время выезда (до)
	RoomName          *string                  // название
	AccommodationType models.AccommodationType // тип размещения
	RoomPrice         *int                     // цена
	RoomInfo          *string                  // дополнительно
	OwnerID           uint
}

type PropertyFilter struct {
	Search         string
	CategoryID     uint
	LocationTypeID uint
	GuestCountMin  *int
	MinPrice       *int
	MaxPrice       *int
	City           string
}

type UpdatePropertyRequest struct {
	Name          *string
	Description   *string
	Address       *string
	City          *string
	Country       *string
	PricePerNight *int
	PropertyType  *string
}

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
		Name:          req.Name,
		Description:   req.Description,
		Address:       req.Address,
		City:          req.City,
		Country:       req.Country,
		PricePerNight: req.PricePerNight,
		PropertyType:  req.PropertyType,
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
