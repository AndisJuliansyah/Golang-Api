package routes

import (
	"golang-api/controllers"
	"golang-api/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello WOrld")
	})

	e.GET("/employee", controllers.FetchAllEmployee, middleware.IsAuthenticated)
	e.POST("/employee", controllers.StoreEmployee, middleware.IsAuthenticated)
	e.PUT("/employee", controllers.UpdateEmployee, middleware.IsAuthenticated)
	e.DELETE("/employee", controllers.DeleteEmployee, middleware.IsAuthenticated)

	e.POST("/register-user", controllers.GenerateHashPassword)
	e.POST("/login", controllers.Checklogin)

	return e
}
