package property

import (
	"context"
	"errors"
	"time"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type PropertyRepo struct {
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) *PropertyRepo {
	return &PropertyRepo{db: db}
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
	Name         *string    // название площадки
	City         *string    // ваш город
	Address      *string    // адрес площадки
	WorkTime     *time.Time // время работы
	Phone        *string    // номер телефона
	OnlineScreen *bool      // онлайн показ есть/нет

	CategoryID     *[]uint // категория
	LocationTypeID *[]uint // тип места
	FeaturesID     *[]uint // особенности
	ServicesID     *[]uint // услуги
	EquipmentID    *[]uint // оборудование в наличии
	KitchenID      *[]uint // кухня

	ParkingAvailable *bool // парковка
	ParkingSpaces    *int

	InteriorStyleID *uint // стиль интерьера

	// залы
	HallName          *string // название зала
	Area              *string // площадь зала кв. м.
	GuestCountBanquet *int    // кол-во мест банкет
	GuestCountBuffet  *int    // кол-во мест фуршет
	GuestCountTheater *int    // кол-во мест рассадка театр
	GuestCountMore    *int    // кол-во мест/еще
	Description       *string // описание площадки

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

	RoomCategoryID    *uint                     // категории номеров
	SleepingPlacesID  *uint                     // спальные места
	AmenitiesRoomID   *uint                     // удобство в номерах
	ServicesRoomID    *uint                     // услуги
	CheckInAfter      *time.Time                //время заезда (от)
	CheckInBefore     *time.Time                //время заезда (до)
	CheckOutAfter     *time.Time                //время выезда (от)
	CheckOutBefore    *time.Time                //время выезда (до)
	RoomName          *string                   // название
	AccommodationType *models.AccommodationType // тип размещения
	RoomPrice         *int                      // цена
	RoomInfo          *string                   // дополнительно

	Image *[]string
}

func (p *PropertyRepo) Create(ctx context.Context, property *models.Property) error {
	query := p.db.WithContext(ctx).Create(property).Error
	return query
}

func (p *PropertyRepo) FindByID(ctx context.Context, id uint) (*models.Property, error) {
	var property models.Property
	err := p.db.WithContext(ctx).Preload("Reviews", nil).Where("id = ?", id).First(&property).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &property, nil
}

func (r *PropertyRepo) FindAll(ctx context.Context, filter PropertyFilter) ([]models.Property, error) {
	query := r.db.WithContext(ctx).Model(&models.Property{})

	if filter.Search != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Search+"%")
	}

	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}

	if filter.LocationTypeID > 0 {
		query = query.Where("location_type_id = ?", filter.LocationTypeID)
	}

	if filter.City != "" {
		query = query.Where("city = ?", filter.City)
	}

	if filter.GuestCountMin != nil {
		query = query.Where("guest_count_banquet >= ?", *filter.GuestCountMin)
	}

	if filter.MinPrice != nil {
		query = query.Where("rental_price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("rental_price <= ?", *filter.MaxPrice)
	}

	var properties []models.Property
	err := query.Find(&properties).Error
	return properties, err
}

func (p *PropertyRepo) FindByOwnerID(ctx context.Context, id uint) ([]models.Property, error) {
	var properties []models.Property
	err := p.db.WithContext(ctx).Where("owner_id = ?", id).Find(&properties).Error
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (p *PropertyRepo) Update(ctx context.Context, id uint, req *UpdatePropertyRequest) error {
	updates := validationUpdates(req)

	if len(updates) == 0 {
		return nil
	}

	return p.db.WithContext(ctx).Model(&models.Property{}).Where("id = ?", id).Updates(updates).Error
}

func (p *PropertyRepo) Delete(ctx context.Context, id uint) error {
	result := p.db.WithContext(ctx).Delete(&models.Property{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("property not found")
	}

	return nil
}
