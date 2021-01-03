package test

import (
	"go/build"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB test
func InitDB() (*gorm.DB, string) {
	// TODO どこかに移す
	appPath := build.Default.GOROOT + "/src/go-sample"
	config, err := ini.Load(appPath + "/configs/config.ini")
	if err != nil {
		panic(err.Error())
	}
	dsn := config.Section("testdb").Key("dsn").String()

	// create test data
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	return db, dsn
}
