package libauth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	secretKey = []byte("your-secret-key")
)

type AuthService struct {
	secretKey []byte
}

func NewAuthService(secretKey []byte) *AuthService {
	return &AuthService{secretKey: secretKey}
}

func (s *AuthService) GenerateJWT(userID string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *AuthService) ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(string), nil
	}
	return "", errors.New("invalid token")
}
