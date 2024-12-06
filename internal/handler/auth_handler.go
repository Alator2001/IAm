package handler

import (
	"net/http"
	"user-registration/internal/domain"
	"user-registration/internal/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Передаем JWT-сервис как зависимость
func RegisterHandler(jwtService *service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user domain.User

		// Читаем данные из JSON-запроса
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Хешируем пароль
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing password"})
			return
		}
		user.Password = string(hashedPassword)

		// Сохраняем пользователя в базу данных (пропустим для простоты)
		// На практике здесь нужно вызывать репозиторий

		// Генерируем JWT-токен
		token, err := jwtService.GenerateToken(1) // Здесь вы передаете реальный ID пользователя
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Возвращаем успешный ответ с токеном
		c.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully!",
			"token":   token,
		})
	}
}
