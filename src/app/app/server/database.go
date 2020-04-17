package app

import (
	"fmt"

	_ "github.com/lib/pq"
)

type ConfigDb struct {
}

const (
	server   = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1111"
	database = "db_taskManager"
	sslmode  = "disable"
)

// GetConnectionString возвращает string сконфигурированную строку подключения к Postgre
func GetConnectionString() string {
	return fmt.Sprintf("user=%s password=%s database=%s sslmode=disable", user, password, database)
}

// GetNewConnectionString получение новой строки подключения
// func GetNewConnectionString(username string, pass string) (string, error) {
// 	return fmt.Sprintf("user=%s password=%s database=%s sslmode=disable", username, pass, database), nil
// }
