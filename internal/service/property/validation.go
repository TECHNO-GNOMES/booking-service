package property

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nyaruka/phonenumbers/v2"
)

func (p *Property) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("property name cannot be empty")
	}

	if strings.TrimSpace(p.City) == "" {
		return errors.New("property city cannot be empty")
	}

	if strings.TrimSpace(p.WorkTime.String()) == "" {
		return errors.New("property work time cannot be empty")
	}

	if strings.TrimSpace(p.Phone) != "" {
		num, err := phonenumbers.Parse(p.Phone, "RU")
		if err != nil {
			return fmt.Errorf("invalid phone format: %w", err)
		}

		if !phonenumbers.IsValidNumber(num) {
			return errors.New("phone number is not valid")
		}
	}

	if p.CategoryID == 0 {
		return errors.New("property category ID cannot be empty")
	}

	if p.LocationTypeID == 0 {
		return errors.New("property location type ID cannot be empty")
	}

	if p.FeaturesID == 0 {
		return errors.New("property features ID cannot be empty")
	}

	if p.ServicesID == 0 {
		return errors.New("property services ID cannot be empty")
	}

	if p.EquipmentID == 0 {
		return errors.New("property equipment ID cannot be empty")
	}

	if p.KitchenID == 0 {
		return errors.New("property kitchen ID cannot be empty")
	}

	if p.RoomCategoryID == 0 {
		return errors.New("property room category ID cannot be empty")
	}

	if p.SleepingPlacesID == 0 {
		return errors.New("property sleeping places ID cannot be empty")
	}

	if p.AmenitiesRoomID == 0 {
		return errors.New("property amenities room ID cannot be empty")
	}

	if p.ServicesRoomID == 0 {
		return errors.New("property services room ID cannot be empty")
	}

	if p.OwnerID == 0 {
		return errors.New("owner id is required")
	}

	return errors.New("property type must be: hotel, apartment, house, or villa")
}
