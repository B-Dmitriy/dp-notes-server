package db

import (
	"fmt"
	"log"
	"strconv"
)

type User struct {
	id    int
	name  string
	email string
}

const (
	LIMIT = 10
)

func GetUsers(page string) ([]User, error) {
	var offset int
	var users []User

	if pageInt, err := strconv.Atoi(page); err != nil {
		offset = pageInt * LIMIT
	}

	queryRow := fmt.Sprintf("SELECT * FROM %s.users LIMIT %d OFFSET %d", Postgres.schema, LIMIT, offset)

	rows, err := Postgres.DB.Query(queryRow)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.id, &user.name, &user.email); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func GetUserById(id string) (User, error) {
	var user User

	// SQL injection check
	if _, err := strconv.Atoi(id); err != nil {
		return user, fmt.Errorf("id must be a integer")
	}

	queryRow := fmt.Sprintf("SELECT * FROM %s.users WHERE id = %s;", Postgres.schema, id)

	row := Postgres.DB.QueryRow(queryRow)

	err := row.Scan(&user.id, &user.name, &user.email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUserById(id string) error {
	// SQL injection check
	if _, err := strconv.Atoi(id); err != nil {
		return fmt.Errorf("id must be a integer")
	}

	execRow := fmt.Sprintf("DELETE FROM %s.users WHERE id = %s;", Postgres.schema, id)

	_, err := Postgres.DB.Exec(execRow)
	if err != nil {
		return err
	}

	return nil
}
