package usecase

import "go-sample/cmd/_your_app_/domain"

// UserInteractor test
type UserInteractor struct {
	UserRepository UserRepository
}

// Add test
func (interactor *UserInteractor) Add(u domain.User) {
	interactor.UserRepository.Store(u)
}

// FindAll test
func (interactor *UserInteractor) FindAll() []domain.User {
	return interactor.UserRepository.Select()
}

// Delete test
func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}
