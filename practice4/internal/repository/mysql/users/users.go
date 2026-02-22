package users

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_mysql "practice4/internal/repository/mysql"
	"practice4/pkg/modules"
)

type Repository struct {
	db               *_mysql.Dialect
	executionTimeout time.Duration
}

func NewUserRepository(db *_mysql.Dialect) *Repository {
	return &Repository{
		db:               db,
		executionTimeout: time.Second * 5,
	}
}

func (r *Repository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.DB.Select(&users, "SELECT id, name, email, age, phone FROM users")
	if err != nil {
		return nil, fmt.Errorf("GetUsers error: %w", err)
	}
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*modules.User, error) {
	var user modules.User
	err := r.db.DB.Get(&user, "SELECT id, name, email, age, phone FROM users WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with id=%d not found", id)
		}
		return nil, fmt.Errorf("GetUserByID error: %w", err)
	}
	return &user, nil
}

func (r *Repository) CreateUser(user modules.User) (int64, error) {
	result, err := r.db.DB.Exec(
		"INSERT INTO users (name, email, age, phone) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Age, user.Phone,
	)
	if err != nil {
		return 0, fmt.Errorf("CreateUser error: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("LastInsertId error: %w", err)
	}

	return id, nil
}

func (r *Repository) UpdateUser(id int, user modules.User) error {
	result, err := r.db.DB.Exec(
		"UPDATE users SET name=?, email=?, age=?, phone=? WHERE id=?",
		user.Name, user.Email, user.Age, user.Phone, id,
	)
	if err != nil {
		return fmt.Errorf("UpdateUser error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected error: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with id=%d does not exist", id)
	}

	return nil
}

func (r *Repository) DeleteUser(id int) (int64, error) {
	result, err := r.db.DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return 0, fmt.Errorf("DeleteUser error: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("RowsAffected error: %w", err)
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("user with id=%d does not exist", id)
	}

	return rowsAffected, nil
}
