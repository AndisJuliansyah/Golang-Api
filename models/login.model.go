package models

import (
	"database/sql"
	"fmt"
	"golang-api/db"
	"golang-api/helper"
)

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func Checklogin(username, password string) (bool, error) {
	var obj Users
	var pwd string

	con := db.CreateCon()

	sqlStatement := "Select * from users where username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Name, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Println("Hash and Password doesn't match")
		return false, err
	}

	return true, nil
}
