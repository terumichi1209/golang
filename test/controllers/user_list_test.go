package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-sample/cmd/_your_app_/domain"
	"go-sample/cmd/_your_app_/infrastructure"
	"go-sample/cmd/_your_app_/interfaces/controllers"
	"go-sample/test"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestListUser(t *testing.T) {
	db, dsn := test.InitDB()

	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(
		httptest.NewRequest(http.MethodGet, "/users", nil),
		rec,
	)

	// create test data
	db.Exec("truncate users")
	user := domain.User{Name: "名前", Email: "teru@gmail.com", Password: "password"}
	db.Create(&user)

	controller := controllers.NewUserController(infrastructure.NewSQLHandler(dsn))
	if assert.NoError(t, controller.List(c)) {
		assert.Exactly(t, http.StatusOK, rec.Code)
		users := []domain.User{user}
		expected, _ := json.Marshal(users)
		assert.JSONEq(t, string(expected), rec.Body.String())
	}
}
