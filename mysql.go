package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "database:3306"
	user   = "root"
	pass   = "qwerty"
	dbname = "shortlinks"
	driver = "mysql"
)

func connect() *sql.DB {

	db, err := sql.Open(driver, user+":"+pass+"@tcp("+host+")/"+dbname)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
