package main

import (
	"go-sample/cmd/_your_app_/infrastructure"
	"go-sample/pkg/validator"

	"github.com/labstack/echo"
	"gopkg.in/ini.v1"
)

func main() {
	config, err := ini.Load("configs/config.ini")
	if err != nil {
		panic(err.Error())
	}
	port := config.Section("web").Key("port").String()
	dsn := config.Section("db").Key("dsn").String()
	// TODO マイグレーションファイルに分ける
	dbinit()
	e := echo.New()
	e.Validator = validator.Init()
	infrastructure.Init(e, dsn)
	e.Logger.Fatal(e.Start(port))
}

func dbinit() {
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// }
	// db.Migrator().CreateTable(domain.User{})
}
