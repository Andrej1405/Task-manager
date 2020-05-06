package mappers

import (
	config "app/app/config"
	"app/app/model/entities"
	"database/sql"
	"fmt"
)

type UserMapper struct {
	User entities.User
}

func (u *UserMapper) GetUserByEmail(email string) (user entities.User, err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `SELECT * FROM users WHERE email = $1`
	row := db.QueryRow(query, email)

	user = entities.User{}

	err = row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func (u *UserMapper) AddNewUser(user *entities.User) (err error) {
	db, err := sql.Open("postgres", config.InitConnectionString())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := `INSERT INTO users (email, password) VALUES ($1, $2) returning id`
	db.QueryRow(query, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
