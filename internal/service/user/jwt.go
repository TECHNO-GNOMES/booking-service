package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateJWT(userID uint, email string, role string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func (s *UserService) GenerateToken(userID uint, email string, role string) (string, error) {
	return generateJWT(userID, email, role, s.jwtSecret)
}
