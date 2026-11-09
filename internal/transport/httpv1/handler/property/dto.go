package property

import (
	"time"
)

type Property struct {
	ID uint `json:"id"`
	// 1 шаг - создание карточки
	Name            string     `json:"name"`             // название площадки
	City            string     `json:"city"`             // ваш город
	Address         *string    `json:"address"`          // адрес площадки
	DistrictAddress *string    `json:"district_address"` // район
	MetroAddress    *string    `json:"metro_address"`    // метро
	WorkTime        *time.Time `json:"work_time"`        // время работы
	Phone           string     `json:"phone"`            // номер телефона
	OnlineScreen    bool       `json:"online_screen"`    // онлайн показ есть/нет

	// 2 шаг - создание карточки
	CategoryID     uint `json:"category_id"`      // категория
	LocationTypeID uint `json:"location_type_id"` // тип места
	FeaturesID     uint `json:"features_id"`      // особенности
	ServicesID     uint `json:"services_id"`      // услуги
	EquipmentID    uint `json:"equipment_idip"`   // оборудование в наличии
	KitchenID      uint `json:"kitchen_id"`       // кухня

	ParkingAvailable bool `json:"parking_available"` // парковка
	ParkingSpaces    int  `json:"parking_spaces"`

	InteriorStyleID *uint `json:"interior_style_id"` // стиль интерьера

	// залы
	HallName          *string `json:"hall_name"`           // название зала
	Area              *string `json:"area"`                // площадь зала кв. м.
	GuestCountBanquet *int    `json:"guest_count_banquet"` // кол-во мест банкет
	GuestCountBuffet  *int    `json:"guest_count_buffet"`  // кол-во мест фуршет
	GuestCountTheater *int    `json:"guest_count_theater"` // кол-во мест рассадка театр
	GuestCountMore    *int    `json:"guest_count_more"`    // кол-во мест/еще
	Description       string  `json:"description"`         // описание площадки

	// цены
	AverageCheckGuest *int    `json:"average_check_guest"` // средний чек на гостя (сумма в рублях)
	BanquetMenuCheck  *int    `json:"banquet_menu_check"`  // банкетное меню на гостя (сумма в рублях)
	BuffetMenuCheck   *int    `json:"buffet_menu_check"`   // фуршетное меню на гостя (сумма в рублях)
	RentalPrice       *int    `json:"rental_price"`        // аренда/час
	RentalWeekDays    *int    `json:"rental_week_days"`    // аренда площадка под мероприятия в будни/выходные и праздничные дни
	DepositMin        *string `json:"deposit_min"`         // депозит (минимальная стоимость закрытия площадки, зала под одного мероприятие)
	DepositInfo       *string `json:"deposit_info"`        // что входит в депозит
	OwnAlcoholAllowed string  `json:"own_alcohol_allowed"` // возможность своего алкоголя
	ServiceSupport    string  `json:"service_support"`     // сервисное обслуживание
	CorkageFee        *string `json:"corkage_fee"`         // пробковый сбор
	ExtraCost         *string `json:"extra_cost"`          // дополнительные затраты
	ExtraInfo         *string `json:"extra_info"`          // дополнительно

	Image string // медиатека

	// номерной фонд
	RoomCategoryID    uint       `json:"room_category_id"`   // категории номеров
	SleepingPlacesID  uint       `json:"sleeping_places_id"` // спальные места
	AmenitiesRoomID   uint       `json:"amenities_room_id"`  // удобство в номерах
	ServicesRoomID    uint       `json:"services_room_id"`   // услуги
	CheckInAfter      *time.Time `json:"check_in_after"`     //время заезда (от)
	CheckInBefore     *time.Time `json:"check_in_before"`    //время заезда (до)
	CheckOutAfter     *time.Time `json:"check_out_after"`    //время выезда (от)
	CheckOutBefore    *time.Time `json:"check_out_before"`   //время выезда (до)
	RoomName          *string    `json:"room_name"`          // название
	AccommodationType string     `json:"accommodation_type"` // тип размещения
	RoomPrice         *int       `json:"room_price"`         // цена
	RoomInfo          *string    `json:"room_info"`          // дополнительно
}

type CreatePropertyInput struct {
	Body Property
}

type CreatePropertyOutput struct {
	Body struct {
		ID uint `json:"id"`
	}
}

type GetPropertyInput struct {
	ID uint `path:"id"`
}

type GetPropertyOutput struct {
	Body Property
}

type UpdatePropertyInput struct {
	ID   uint `path:"id"`
	Body struct {
		Name          *string `json:"name,omitempty" minLength:"1" maxLength:"255" pattern:"^[a-zA-Zа-яА-Я0-9\\s]+$"`
		Description   *string `json:"description,omitempty" maxLength:"10000" pattern:"^[\\p{L}\\p{N}\\p{P}\\p{Z}\\n\\r]+$"`
		Address       *string `json:"address"`
		City          *string `json:"city"`
		Country       *string `json:"country"`
		PricePerNight *int    `json:"price_per_night"`
		PropertyType  *string `json:"property_type"`
	}
}

type UpdatePropertyOutput struct {
	Body Property
}

type DeletePropertyInput struct {
	ID uint `path:"id"`
}

type DeletePropertyOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}
