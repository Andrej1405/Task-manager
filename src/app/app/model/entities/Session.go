package entities

import (
	"time"
)

// Структура сессии пользователя
type Session struct {
	Email      string
	Created    time.Time
	Expiration int64
}
