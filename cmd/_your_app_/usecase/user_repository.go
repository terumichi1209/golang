package usecase

import (
	"go-sample/cmd/_your_app_/domain"
)

// UserRepository test
type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
