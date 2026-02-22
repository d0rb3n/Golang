package repository

import (
	_mysql "practice4/internal/repository/mysql"
	"practice4/internal/repository/mysql/users"
	"practice4/pkg/modules"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
	GetUserByID(id int) (*modules.User, error)
	CreateUser(user modules.User) (int64, error)
	UpdateUser(id int, user modules.User) error
	DeleteUser(id int) (int64, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *_mysql.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}
