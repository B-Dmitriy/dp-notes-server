package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB     *sql.DB
	schema string
}

var Postgres DataBase

func InitDB(user, password, schema string) error {

	connStr := "user=" + user + " password=" + password + "  dbname=webservice sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	Postgres.DB = db
	Postgres.schema = schema

	return nil
}
