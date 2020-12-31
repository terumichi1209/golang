package main

import (
	"go-sample/cmd/_your_app_/infrastructure"
	"go-sample/pkg/validator"

	"github.com/labstack/echo"
)

const (
	// TODO env
	port = ":8080"
	dsn  = "root:password@tcp(go_sample)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	// TODO マイグレーションファイルに分ける
	dbinit()
	e := echo.New()
	e.Validator = validator.Init()
	infrastructure.Init(e)
	e.Logger.Fatal(e.Start(port))
}

func dbinit() {
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// }
	// db.Migrator().CreateTable(domain.User{})
}
