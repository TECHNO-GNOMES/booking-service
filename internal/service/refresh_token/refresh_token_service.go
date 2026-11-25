package refresh_token

import "github.com/nurgal1ev/booking-service/internal/repository/refresh_token"

type RefreshTokenService struct {
	refreshTokenRepo *refresh_token.RefreshTokenRepo
}

func NewRefreshTokenService(refreshTokenRepo *refresh_token.RefreshTokenRepo) *RefreshTokenService {
	return &RefreshTokenService{refreshTokenRepo}
}
