package property

import (
	"time"

	"github.com/nurgal1ev/booking-service/internal/models"
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
	CategoryID     []uint // категория
	LocationTypeID []uint // тип места
	FeaturesID     []uint // особенности
	ServicesID     []uint // услуги
	EquipmentID    []uint // оборудование в наличии
	KitchenID      []uint // кухня

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

	Image []string // медиатека

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
