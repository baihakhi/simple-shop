package repositories

import (
	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/repositories/queries"
)

func (r *repository) CreateUser(data *models.User) (string, error) {
	var username string
	if err := r.db.QueryRow(queries.CreateUsers,
		data.Username,
		data.Fullname,
		data.Address,
		data.Role,
		data.Password).
		Scan(&username); err != nil {
		return "", err
	}

	return username, nil
}

func (r *repository) GetOneUsersByUsername(username string) (*models.User, error) {
	var result models.User

	if err := r.db.QueryRow(queries.GetOneUsersByUsername, username).
		Scan(
			&result.ID,
			&result.Username,
			&result.Fullname,
			&result.Address,
			&result.Role,
			&result.Balance,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) GetPasswordByUsername(username string) (string, error) {
	var result string

	if err := r.db.QueryRow(queries.GetPasswordByUsername, username).
		Scan(&result); err != nil {
		return "", err
	}
	return result, nil
}
