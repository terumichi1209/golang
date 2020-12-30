package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-sample/cmd/_your_app_/interfaces/gateways"
)

// SQLHandler test
type SQLHandler struct {
	db *gorm.DB
}

// NewSQLHandler test
func NewSQLHandler() gateways.SQLHandler {
	dsn := "root:password@tcp(go_sample)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.db = db
	return sqlHandler
}

// Create test
func (handler *SQLHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

// FindAll test
func (handler *SQLHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

// DeleteByID test
func (handler *SQLHandler) DeleteByID(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}
