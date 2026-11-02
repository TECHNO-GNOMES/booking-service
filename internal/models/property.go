package models

import (
	"time"

	"gorm.io/gorm"
)

type OwnAlcoholAllowed string
type ServiceSupport string
type AccommodationType string

const (
	YesAlcohol OwnAlcoholAllowed = "yes"
	NoAlcohol  OwnAlcoholAllowed = "no"
	Half       OwnAlcoholAllowed = "half"
	Agreement  OwnAlcoholAllowed = "agreement"
)

const (
	TenPercent     ServiceSupport = "ten_percent"
	FifteenPercent ServiceSupport = "fifteen_percent"
	NoService      ServiceSupport = "no"
	OtherService   ServiceSupport = "other"
)

const (
	SingleBed AccommodationType = "single_bed"
	DoubleBed AccommodationType = "double_bed"
	TripleBed AccommodationType = "triple_bed"
	OtherBed  AccommodationType = "other_bed"
)

type Property struct {
	gorm.Model
	// 1 шаг - создание карточки
	Name            string     // название площадки
	City            string     // ваш город
	Address         *string    // адрес площадки
	DistrictAddress *string    // район
	MetroAddress    *string    // метро
	WorkTime        *time.Time // время работы
	OnlineScreen    bool       // онлайн показ есть/нет

	// 2 шаг - создание карточки
	CategoryID uint     // категория
	Category   Category `gorm:"foreignKey:CategoryID"`

	LocationTypeID uint         // тип места
	LocationType   LocationType `gorm:"foreignKey:LocationTypeID"`

	FeaturesID uint     // особенности
	Features   Features `gorm:"foreignKey:FeaturesID"`

	ServicesID uint     // услуги
	Services   Services `gorm:"foreignKey:ServicesID"`

	EquipmentID uint      // оборудование в наличии
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`

	KitchenID uint    // кухня
	Kitchen   Kitchen `gorm:"foreignKey:KitchenID"`

	ParkingID uint    // парковка
	Parking   Parking `gorm:"foreignKey:ParkingID"`

	ParkingSpacesID *uint         // количество парковочных мест
	ParkingSpaces   ParkingSpaces `gorm:"foreignKey:ParkingSpacesID"`

	InteriorStyleID *uint         // стиль интерьера
	InteriorStyle   InteriorStyle `gorm:"foreignKey:InteriorStyleID"`

	// залы
	HallName          *string // название зала
	Area              *string // площадь зала кв. м.
	GuestCountBanquet *int    // кол-во мест банкет
	GuestCountBuffet  *int    // кол-во мест фуршет
	GuestCountTheater *int    // кол-во мест рассадка театр
	GuestCountMore    *int    // кол-во мест/еще
	Description       string  // описание площадки

	// цены
	AverageCheckGuest *int               // средний чек на гостя (сумма в рублях)
	BanquetMenuCheck  *int               // банкетное меню на гостя (сумма в рублях)
	BuffetMenuCheck   *int               // фуршетное меню на гостя (сумма в рублях)
	RentalPrice       *int               // аренда/час
	RentalWeekDays    *int               // аренда площадка под мероприятия в будни/выходные и праздничные дни
	DepositMin        *string            // депозит (минимальная стоимость закрытия площадки, зала под одного мероприятие)
	DepositInfo       *string            // что входит в депозит
	OwnAlcoholAllowed *OwnAlcoholAllowed // возможность своего алкоголя
	ServiceSupport    *ServiceSupport    // сервисное обслуживание
	CorkageFee        *string            // пробковый сбор
	ExtraCost         *string            // дополнительные затраты
	ExtraInfo         *string            // дополнительно

	Image string // медиатека

	// номерной фонд
	RoomCategoryID    uint              // категории номеров
	RoomCategory      RoomCategory      `gorm:"foreignKey:RoomCategoryID"`
	SleepingPlacesID  uint              // спальные места
	SleepingPlaces    *SleepingPlaces   `gorm:"foreignKey:SleepingPlacesID"`
	AmenitiesRoomID   uint              // удобство в номерах
	AmenitiesRoom     AmenitiesRoom     `gorm:"foreignKey:AmenitiesRoomID"`
	ServicesRoomID    uint              // услуги
	ServicesRoom      ServicesRoom      `gorm:"foreignKey:ServicesRoomID"`
	CheckInAfter      *time.Time        //время заезда (от)
	CheckInBefore     *time.Time        //время заезда (до)
	CheckOutAfter     *time.Time        //время выезда (от)
	CheckOutBefore    *time.Time        //время выезда (до)
	RoomName          *string           // название
	AccommodationType AccommodationType // тип размещения
	RoomPrice         *int              // цена
	RoomInfo          *string           // дополнительно

	OwnerID uint
	Owner   User `gorm:"foreignKey:OwnerID"`
}
