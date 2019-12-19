package models

import "database/sql"

import "fmt"

type tableRow struct {
	short string
	full  string
	count int
}

type StatTable []tableRow

func (r tableRow) GetString() string {

	return fmt.Sprintf("<tr><td>%v</td><td>%v</td><td>%v</td></tr>", r.short, r.full, r.count)

}

func RecordFollowing(short string, full string, db *sql.DB) error {

	q := `INSERT INTO statistics (short, full) VALUES (?, ?);`

	_, err := db.Query(q, short, full)

	if err != nil {
		return err
	}

	return nil
}

func GetPopularTable(db *sql.DB) (StatTable, error) {

	limit := 10
	var table StatTable
	var row tableRow

	q := `SELECT short, full, count(*) as count FROM statistics GROUP BY short, full ORDER BY count DESC LIMIT ?;`
	results, err := db.Query(q, limit)

	if err != nil {
		fmt.Println(err)
		return table, err
	}

	for results.Next() {

		err := results.Scan(&row.short, &row.full, &row.count)
		if err != nil {
			fmt.Println(err)
			return table, err
		}

		table = append(table, row)
	}

	return table, nil

}
