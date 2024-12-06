package main

import (
	"log"
	"os"

	"user-registration/internal/handler"
	"user-registration/internal/middleware"
	"user-registration/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set in .env")
	}

	jwtService := service.NewJWTService(secretKey) // Создаем JWT-сервис с секретным ключом

	// Маршрут для регистрации
	r.POST("/register", handler.RegisterHandler(jwtService))

	// Защищенный маршрут с middleware
	r.GET("/protected", middleware.AuthMiddleware(jwtService), func(c *gin.Context) {
		// Извлекаем user_id из контекста
		userID, _ := c.Get("user_id")
		c.JSON(200, gin.H{
			"message": "You are authorized",
			"user_id": userID,
		})
	})

	r.Run(":8080")
}
