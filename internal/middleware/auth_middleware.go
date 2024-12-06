package middleware

import (
	"net/http"
	"user-registration/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware проверяет JWT токен
func AuthMiddleware(jwtService *service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Проверяем и парсим токен
		claims, err := jwtService.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		// Сохраняем user_id из токена в контекст Gin (для последующего использования)
		c.Set("user_id", claims["user_id"])

		// Если всё валидно, продолжаем обработку запроса
		c.Next()
	}
}
