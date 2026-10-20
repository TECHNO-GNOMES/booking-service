package user

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
