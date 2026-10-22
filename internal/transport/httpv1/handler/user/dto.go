package user

import "time"

type User struct {
	Avatar    string     `json:"avatar"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	City      string     `json:"city"`
	Birthday  *time.Time `json:"birthday"`
	Gender    string     `json:"gender"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	Email     string     `json:"email"`
}

type RegisterInput struct {
	Body struct {
		Email    string `json:"email"     format:"email"`
		Password string `json:"password"  minLength:"7" maxLength:"64"`
	}
}

type RegisterOutput struct {
	Body struct {
		Message     string `json:"message"`
		AccessToken string `json:"accessToken"`
	}
}

type LoginInput struct {
	Body struct {
		Email    string `json:"email"     format:"email"`
		Password string `json:"password"  minLength:"7" maxLength:"64"`
	}
}

type LoginOutput struct {
	Body struct {
		AccessToken string `json:"accessToken"`
	}
}

type UpdateUserInput struct {
	FirstName string
	LastName  string
}
