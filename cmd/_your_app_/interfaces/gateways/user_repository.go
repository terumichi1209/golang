package gateways

import (
	"go-sample/cmd/_your_app_/domain"
)

// UserRepository test
type UserRepository struct {
	SQLHandler
}

// Store test
func (db *UserRepository) Store(u domain.User) {
	db.Create(&u)
}

// Select test
func (db *UserRepository) Select() []domain.User {
	user := []domain.User{}
	db.FindAll(&user)
	return user
}

// Delete test
func (db *UserRepository) Delete(id string) {
	user := []domain.User{}
	db.DeleteByID(&user, id)
}
