package controllers

import (
	"go-sample/cmd/_your_app_/domain"
	"go-sample/cmd/_your_app_/interfaces/gateways"
	"go-sample/cmd/_your_app_/usecase"
	"net/http"

	"github.com/labstack/echo"
)

// UserController test
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController test
func NewUserController(sqlHandler gateways.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &gateways.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// List ユーザ取得
func (controller *UserController) List(c echo.Context) error {
	users := controller.Interactor.FindAll()
	return c.JSON(http.StatusOK, users)
}

// Create ユーザ作成
func (controller *UserController) Create(c echo.Context) error {
	user := domain.User{}
	c.Bind(&user)
	controller.Interactor.Add(user)
	return c.NoContent(http.StatusCreated)
}

// Delete ユーザ削除
func (controller *UserController) Delete(c echo.Context) error {
	id := c.Param("id")
	controller.Interactor.Delete(id)
	return c.NoContent(http.StatusNoContent)
}
