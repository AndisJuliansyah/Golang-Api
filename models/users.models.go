package models

import (
	"golang-api/db"
	"golang-api/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func StoreUsers(name, username, password string) (Response, error) {
	var res Response

	valid := validator.New()

	data := User{
		Name:     name,
		Username: username,
		Password: password,
	}

	err := valid.Struct(data)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "Insert users (name, username, password) Values (?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	hash, _ := helper.HashPassword(password)

	result, err := stmt.Exec(name, username, hash)
	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Massage = "Success"
	res.Data = map[string]int64{
		"lastInsertId": lastInsertId,
	}

	return res, nil
}
