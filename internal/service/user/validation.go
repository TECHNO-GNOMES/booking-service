package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/nyaruka/phonenumbers/v2"
)

func (u *User) Validate() error {
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email is required")
	}

	if !govalidator.IsEmail(u.Email) {
		return errors.New("invalid email")
	}

	if strings.TrimSpace(u.Phone) != "" {
		num, err := phonenumbers.Parse(u.Phone, "RU")
		if err != nil {
			return fmt.Errorf("invalid phone format: %w", err)
		}

		if !phonenumbers.IsValidNumber(num) {
			return errors.New("phone number is not valid")
		}
	}

	if strings.TrimSpace(u.Password) == "" {
		return errors.New("password is required")
	}
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if u.Role == "" {
		u.Role = "guest"
	}
	if u.Role != "guest" && u.Role != "admin" {
		return errors.New("role must be 'guest' or 'admin'")
	}

	return nil
}
