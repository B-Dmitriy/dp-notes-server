package db

type User struct {
	id    int
	name  string
	email string
}

func GetUserById(id string) (User, error) {
	var user User

	row := Postgres.DB.QueryRow("SELECT * FROM test.users WHERE id = $1;", id)

	err := row.Scan(&user.id, &user.name, &user.email)
	if err != nil {
		return user, err
	}

	return user, nil
}
