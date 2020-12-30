package infrastructure

import (
	controllers "go-sample/cmd/_your_app_/interfaces/controllers"

	"github.com/labstack/echo"
)

// Init test
func Init(e *echo.Echo) {
	userController := controllers.NewUserController(NewSQLHandler())

	e.GET("/users", userController.List)
	e.POST("/users", userController.Create)
	e.DELETE("/users/:id", userController.Delete)
}
