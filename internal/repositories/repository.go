package repositories

import (
	"database/sql"

	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/models/payload/request"
)

type repository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) Repositories {
	return &repository{
		db: db,
	}
}

type Repositories interface {
	// User repository
	CreateUser(data *models.User) (string, error)
	GetPasswordByUsername(username string) (string, error)
	GetOneUsersByUsername(username string) (*models.User, error)

	// Product Repository
	GetListProducts(params *request.PaginationRequest) ([]*models.Products, error)
}
