package usecase

import (
	"fmt"
	"practice4/internal/repository"
	"practice4/pkg/modules"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUsers() ([]modules.User, error) {
	return u.repo.GetUsers()
}

func (u *UserUsecase) GetUserByID(id int) (*modules.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserUsecase) CreateUser(user modules.User) (int64, error) {
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) UpdateUser(id int, user modules.User) error {
	return u.repo.UpdateUser(id, user)
}

func (u *UserUsecase) DeleteUser(id int) (int64, error) {
	return u.repo.DeleteUser(id)
}

// Usecase знает бизнес-логику — например форматирует ответ
func (u *UserUsecase) CreateUserFormatted(user modules.User) string {
	id, err := u.repo.CreateUser(user)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("created user with id=%d", id)
}
