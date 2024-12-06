package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTService предоставляет методы для работы с JWT
type JWTService struct {
	secretKey string
}

// NewJWTService создает новый экземпляр JWTService
func NewJWTService(secretKey string) *JWTService {
	return &JWTService{secretKey: secretKey}
}

// GenerateToken создает JWT-токен для пользователя
func (s *JWTService) GenerateToken(userID int) (string, error) {
	// Определяем время истечения токена
	expirationTime := time.Now().Add(24 * time.Hour)

	// Создаем claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	return token.SignedString([]byte(s.secretKey))
}

// ParseToken проверяет и парсит JWT-токен
func (s *JWTService) ParseToken(tokenString string) (jwt.MapClaims, error) {
	// Парсим токен с функцией проверки подписи
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Убеждаемся, что используется правильный метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})

	// Проверяем, не было ли ошибок при парсинге
	if err != nil {
		return nil, err
	}

	// Проверяем валидность токена и извлекаем claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
