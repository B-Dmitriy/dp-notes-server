package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
}

var Postgres DataBase

func InitDB(user, password string) error {

	connStr := "user=" + user + " password=" + password + "  dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Postgres.DB = db

	return nil
}
