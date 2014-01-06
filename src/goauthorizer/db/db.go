package db

import (
	_ "github.com/lib/pq"
	"database/sql"
)

var Connect *sql.DB

// Get connection to postgres
func ConnectToDb() (*sql.DB, error) {
	if Connect != nil {
		return Connect, nil
	}
	// Use global env vars
	dbi, _ := sql.Open("postgres", "")
    err := dbi.Ping()
	if err != nil {
		return nil, err
	}
	Connect = dbi
	return Connect, nil
}
