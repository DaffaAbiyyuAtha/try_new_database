package models

import (
	"context"
	"fazztrack/backend/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" form:"email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
	Username string `json:"username" form:"username" db:"username"`
}

func FindAllUsers() []User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select * from "users"`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}
	return users
}

func FindOneUser(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select * from "users"`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for _, v := range users {
		if v.Id == id {
			user = v
		}
	}

	return user
}

func DeleteUser(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`DELETE FROM "users" WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found")
	}

	return nil
}

func CreateUser(user User) error {
	db := lib.DB()
	defer db.Close(context.Background())

	_, err := db.Exec(
		context.Background(),
		`INSERT INTO "users" (email, password, username) VALUES ($1, $2, $3)`,
		user.Email, user.Password, user.Username,
	)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}

func EditUser(email string, username string, password string, id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4`

	db.Exec(context.Background(), dataSql, email, username, password, id)

}
