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
	SingleBed AccommodationType = "single-bed"
	DoubleBed AccommodationType = "double-bed"
	TripleBed AccommodationType = "triple-bed"
	OtherBed  AccommodationType = "other-bed"
)

type Property struct {
	gorm.Model
	// 1 шаг - создание карточки
	Name            string
	City            string
	Address         string
	DistrictAddress string
	MetroAddress    string
	WorkTime        *time.Time
	OnlineScreen    string

	// 2 шаг - создание карточки
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`

	LocationTypeID uint
	LocationType   LocationType `gorm:"foreignKey:LocationTypeID"`

	FeaturesID uint
	Features   Features `gorm:"foreignKey:FeaturesID"`

	ServicesID uint
	Services   Services `gorm:"foreignKey:ServicesID"`

	EquipmentID uint
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`

	KitchenID uint
	Kitchen   Kitchen `gorm:"foreignKey:KitchenID"`

	ParkingID uint
	Parking   Parking `gorm:"foreignKey:ParkingID"`

	ParkingSpacesID uint
	ParkingSpaces   ParkingSpaces `gorm:"foreignKey:ParkingSpacesID"`

	InteriorStyleID uint
	InteriorStyle   InteriorStyle `gorm:"foreignKey:InteriorStyleID"`

	// залы
	HallName          string
	Area              string
	GuestCountBanquet int
	GuestCountBuffet  int
	GuestCountTheater int
	GuestCountMore    int
	Description       string

	// цены
	AverageCheckGuest int
	BanquetMenuCheck  int
	BuffetMenuCheck   int
	RentalPrice       int
	RentalWeekDays    int
	DepositMin        string
	DepositInfo       string
	OwnAlcoholAllowed OwnAlcoholAllowed
	ServiceSupport    ServiceSupport
	CorkageFee        string
	ExtraCost         string
	ExtraInfo         string

	Image string

	// номерной фонд
	RoomCategoryID uint
	RoomCategory   RoomCategory `gorm:"foreignKey:RoomCategoryID"`

	SleepingPlacesID uint
	SleepingPlaces   SleepingPlaces `gorm:"foreignKey:SleepingPlacesID"`

	CheckInAfterID uint
	CheckInAfter   CheckInAfter `gorm:"foreignKey:CheckInAfterID"`

	CheckInBeforeID uint
	CheckInBefore   CheckInBefore `gorm:"foreignKey:CheckInBeforeID"`

	CheckOutAfterID uint
	CheckOutAfter   CheckOutAfter `gorm:"foreignKey:CheckOutAfterID"`

	CheckOutBeforeID uint
	CheckOutBefore   CheckOutBefore `gorm:"foreignKey:CheckOutBeforeID"`

	RoomName          string
	AccommodationType AccommodationType
	RoomPrice         int
	RoomInfo          string

	OwnerID uint
	Owner   User `gorm:"foreignKey:OwnerID"`
}
