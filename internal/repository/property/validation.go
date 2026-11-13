package property

func validationUpdates(req *UpdatePropertyRequest) map[string]any {
	updates := make(map[string]any)

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.City != nil {
		updates["city"] = *req.City
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}

	if req.WorkTime != nil {
		updates["work_time"] = *req.WorkTime
	}

	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}

	if req.OnlineScreen != nil {
		updates["online_screen"] = *req.OnlineScreen
	}

	if req.CategoryID != nil {
		updates["category_id"] = req.CategoryID // GORM сам сериализует в jsonb
	}

	if req.LocationTypeID != nil {
		updates["location_type_id"] = req.LocationTypeID
	}

	if req.FeaturesID != nil {
		updates["features_id"] = req.FeaturesID
	}

	if req.ServicesID != nil {
		updates["services_id"] = req.ServicesID
	}

	if req.EquipmentID != nil {
		updates["equipment_id"] = req.EquipmentID
	}

	if req.KitchenID != nil {
		updates["kitchen_id"] = req.KitchenID
	}

	if req.ParkingAvailable != nil {
		updates["parking_available"] = *req.ParkingAvailable
	}

	if req.ParkingSpaces != nil {
		updates["parking_spaces"] = *req.ParkingSpaces
	}

	if req.InteriorStyleID != nil {
		updates["interior_style_id"] = req.InteriorStyleID
	}

	if req.HallName != nil {
		updates["hall_name"] = *req.HallName
	}
	if req.Area != nil {
		updates["area"] = *req.Area
	}

	if req.GuestCountBanquet != nil {
		updates["guest_count_banquet"] = *req.GuestCountBanquet
	}

	if req.GuestCountBuffet != nil {
		updates["guest_count_buffet"] = *req.GuestCountBuffet
	}

	if req.GuestCountTheater != nil {
		updates["guest_count_theater"] = *req.GuestCountTheater
	}

	if req.GuestCountMore != nil {
		updates["guest_count_more"] = *req.GuestCountMore
	}

	if req.Description != nil {
		updates["description"] = *req.Description
	}

	if req.AverageCheckGuest != nil {
		updates["average_check_guest"] = *req.AverageCheckGuest
	}

	if req.BanquetMenuCheck != nil {
		updates["banquet_menu_check"] = *req.BanquetMenuCheck
	}

	if req.BuffetMenuCheck != nil {
		updates["buffet_menu_check"] = *req.BuffetMenuCheck
	}

	if req.RentalPrice != nil {
		updates["rental_price"] = *req.RentalPrice
	}

	if req.RentalWeekDays != nil {
		updates["rental_week_days"] = *req.RentalWeekDays
	}

	if req.DepositMin != nil {
		updates["deposit_min"] = *req.DepositMin
	}

	if req.DepositInfo != nil {
		updates["deposit_info"] = *req.DepositInfo
	}

	if req.OwnAlcoholAllowed != nil {
		updates["own_alcohol_allowed"] = *req.OwnAlcoholAllowed
	}

	if req.ServiceSupport != nil {
		updates["service_support"] = *req.ServiceSupport
	}

	if req.CorkageFee != nil {
		updates["corkage_fee"] = *req.CorkageFee
	}

	if req.ExtraCost != nil {
		updates["extra_cost"] = *req.ExtraCost
	}

	if req.ExtraInfo != nil {
		updates["extra_info"] = *req.ExtraInfo
	}

	if req.RoomCategoryID != nil {
		updates["room_category_id"] = *req.RoomCategoryID
	}

	if req.SleepingPlacesID != nil {
		updates["sleeping_places_id"] = *req.SleepingPlacesID
	}

	if req.AmenitiesRoomID != nil {
		updates["amenities_room_id"] = *req.AmenitiesRoomID
	}

	if req.ServicesRoomID != nil {
		updates["service_support"] = *req.ServiceSupport
	}

	if req.CheckInAfter != nil {
		updates["check_in_after"] = *req.CheckInAfter
	}

	if req.CheckInBefore != nil {
		updates["check_in_before"] = *req.CheckInBefore
	}

	if req.CheckOutAfter != nil {
		updates["check_out_after"] = *req.CheckOutAfter
	}

	if req.CheckOutBefore != nil {
		updates["check_out_before"] = *req.CheckOutBefore
	}

	if req.RoomName != nil {
		updates["room_name"] = *req.RoomName
	}

	if req.AccommodationType != nil {
		updates["accommodation_type"] = *req.AccommodationType
	}

	if req.RoomPrice != nil {
		updates["room_price"] = *req.RoomPrice
	}

	if req.RoomInfo != nil {
		updates["room_info"] = *req.RoomInfo
	}

	if req.Image != nil {
		updates["image"] = *req.Image
	}

	return updates
}
