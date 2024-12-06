package domain

// User представляет сущность пользователя
type User struct {
	ID       int    `json:"id"`       // Уникальный идентификатор
	Name     string `json:"name"`     // Имя пользователя
	Email    string `json:"email"`    // Email пользователя
	Password string `json:"password"` // Пароль (хранится в виде хэша)
}
