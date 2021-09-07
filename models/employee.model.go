package models

import (
	"golang-api/db"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type Employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required,number"`
}

func FetchAllEmployee() (Response, error) {
	var obj Employee
	var arrobj []Employee
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM employee"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Phone)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Massage = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreEmployee(name string, address string, phone string) (Response, error) {
	var res Response

	valid := validator.New()

	data := Employee{
		Name:    name,
		Address: address,
		Phone:   phone,
	}

	err := valid.Struct(data)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "Insert employee (name, address, phone) Values (?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, phone)

	if err != nil {
		return res, err
	}

	lastinsertedId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Massage = "Success"
	res.Data = map[string]int64{
		"lastinsertId": lastinsertedId,
	}

	return res, nil
}

func UpdateEmployee(id int, name string, address string, phone string) (Response, error) {
	var res Response

	valid := validator.New()

	data := Employee{
		Name:    name,
		Address: address,
		Phone:   phone,
	}

	err := valid.Struct(data)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "Update employee SET name = ?, address = ?, phone = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, phone, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Massage = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

func DeleteEmployee(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "Delete from employee WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, nil
	}

	res.Status = http.StatusOK
	res.Massage = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}
