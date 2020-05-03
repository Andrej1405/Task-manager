package config

import (
	"fmt"
)

// Строка подключения к БД
func InitConnectionString() string {
	connect := "user=postgres password=1111 dbname=db_taskManager sslmode=disable"

	return fmt.Sprintf(connect)
}
