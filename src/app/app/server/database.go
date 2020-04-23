package server

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() (err error) {
	connect := "user=postgres password=1111 dbname=db_taskManager sslmode=disable"

	Db, err = sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	return
}
