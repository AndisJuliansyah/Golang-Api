package controllers

import (
	"golang-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllEmployee(c echo.Context) error {
	result, err := models.FetchAllEmployee()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreEmployee(e echo.Context) error {
	name := e.FormValue("name")
	address := e.FormValue("address")
	phone := e.FormValue("phone")

	result, err := models.StoreEmployee(name, address, phone)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	return e.JSON(http.StatusOK, result)
}

func UpdateEmployee(e echo.Context) error {
	id := e.FormValue("id")
	name := e.FormValue("name")
	address := e.FormValue("address")
	phone := e.FormValue("phone")

	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateEmployee(conv_id, name, address, phone)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())

	}

	return e.JSON(http.StatusOK, result)

}

func DeleteEmployee(e echo.Context) error {
	id := e.FormValue("id")

	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteEmployee(conv_id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, result)
}
