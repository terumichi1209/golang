package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/users", createUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func createUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	return c.String(http.StatusOK, "name:"+name+" email:"+email+" password:"+password)
}
