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
		Phone    *string `json:"phone" minLength:"11" maxLength:"12"`
		Email    *string `json:"email"     format:"email"`
		Password string  `json:"password"  minLength:"7" maxLength:"64"`
	}
}

type LoginOutput struct {
	Body struct {
		AccessToken string `json:"accessToken"`
	}
}

type UpdateUserInput struct {
	Body struct {
		Avatar    *string    `json:"avatar"`
		FirstName *string    `json:"firstname" minLength:"1" maxLength:"50" pattern:"^[a-zA-Zа-яА-Я\\s-]+$"`
		LastName  *string    `json:"lastname" minLength:"1" maxLength:"50" pattern:"^[a-zA-Zа-яА-Я\\s-]+$"`
		City      *string    `json:"city" minLength:"1" maxLength:"100"`
		Birthday  *time.Time `json:"birthday" format:"date-time"`
		Gender    *string    `json:"gender" enum:"male,female"`
		Phone     *string    `json:"phone" minLength:"11" maxLength:"12"`
		Address   *string    `json:"address" maxLength:"255"`
		Email     *string    `json:"email" format:"email"`
	}
}

type UpdateUserOutput struct {
	Body User
}

type GetUserInput struct{}

type GetUserOutput struct {
	Body User
}
