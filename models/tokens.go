package models

import (
	"database/sql"
)

type Token struct {
	id              int
	Ownerid         int
	Token           string
	expiration_date string
}

func TokenCheck(token string, db *sql.DB) (Token, error) {

	var t Token

	query := `SELECT * FROM tokens WHERE token=? AND expiration_date > NOW() ORDER BY id DESC;`

	results, err := db.Query(query, token)

	if err != nil {
		return t, err
	}

	for results.Next() {
		err := results.Scan(&t.id, &t.Ownerid, &t.Token, &t.expiration_date)
		if err != nil {
			return t, err
		}
		break
	}

	return t, err

}

func TokenGetNew(user User, db *sql.DB) (Token, error) {

	var t Token

	err := tokenCreate(user.id, db)

	if err != nil {
		return t, err
	}

	t, err = tokenGet(user.id, db)

	return t, err
}

func tokenGet(ownerid int, db *sql.DB) (Token, error) {

	var t Token

	query := `SELECT * FROM tokens WHERE ownerid=? AND expiration_date > NOW() ORDER BY id DESC;`

	results, err := db.Query(query, ownerid)

	if err != nil {
		return t, err
	}

	for results.Next() {
		err := results.Scan(&t.id, &t.Ownerid, &t.Token, &t.expiration_date)
		if err != nil {
			return t, err
		}
		break
	}

	return t, err

}

func tokenCreate(ownerid int, db *sql.DB) error {

	newtoken := "testtoken4"

	query := `INSERT INTO tokens (ownerid, token, expiration_date) VALUES (?, ?, DATE_ADD(NOW(), INTERVAL 30 MINUTE));`
	_, err := db.Query(query, ownerid, newtoken)

	return err

}
