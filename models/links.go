package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

type Link struct {
	id             int
	Ownerid        int
	Short          string
	Full           string
	ExpirationTime time.Time
}

func LinkAdd(link Link, db *sql.DB) error {

	linkShort := link.Short

	if linkShort == "" {
		linkShort = linkGenerate()
	}

	q := `INSERT INTO links (ownerid, short, full, expiration_date) VALUES (?, ?, ?, ?)`

	_, err := db.Query(q, link.Ownerid, linkShort, link.Full, sql.NullTime{Time: link.ExpirationTime, Valid: !link.ExpirationTime.IsZero()})

	if err != nil {
		fmt.Println(err)
	}

	return err

}

func LinkGetFull(short string, db *sql.DB) (string, error) {
	var fullLink string

	query := `SELECT full FROM links WHERE short=? AND (expiration_date > NOW() || expiration_date IS NULL) ORDER BY id DESC`

	results, err := db.Query(query, short)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for results.Next() {
		err := results.Scan(&fullLink)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		break
	}

	return fullLink, nil
}

func LinkGetShort(full string, db *sql.DB) (string, error) {

	var shortLink string

	query := `SELECT short FROM links WHERE full=? AND (expiration_date > NOW() || expiration_date IS NULL) ORDER BY id DESC`

	results, err := db.Query(query, full)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for results.Next() {
		err := results.Scan(&shortLink)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		break
	}

	return shortLink, nil

}

func linkGenerate() string {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	linkLenght := 6

	b := make([]rune, linkLenght)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}
