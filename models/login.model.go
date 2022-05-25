package models

import (
	"database/sql"
	"fmt"

	"github.com/oxwazz/l210526-twitter-app-be/helpers"

	"github.com/oxwazz/l210526-twitter-app-be/db"
)

type Userr struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func Login(username, password string) (bool, error) {
	var obj Userr
	var pwd string

	fmt.Println(username)

	con := db.CreateCon()

	sqlStatement := "SELECT id, username, password FROM users WHERE username = $1"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error", err)
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
