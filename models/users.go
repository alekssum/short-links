package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	id       int
	login    string
	password string
}

type Users []User

func UserCheckAuth(login string, password string, db *sql.DB) bool {

	u := UserGetByLogin(login, db)

	if u.password == password {
		return true
	}

	return false

}

func UserGetByLogin(login string, db *sql.DB) User {

	var user User

	query := `SELECT * FROM users WHERE login = ?`

	results, err := db.Query(query, login)

	if err != nil {
		fmt.Println(err)
	}

	for results.Next() {
		err := results.Scan(&user.id, &user.login, &user.password)
		if err != nil {
			fmt.Println(err)
			return user
		}
	}

	return user

}
