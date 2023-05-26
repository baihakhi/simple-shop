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
	GetOneProductByID(productID uint64) (result *models.Products, err error)

	// Cart Repository
	CreateCart(data *models.Cart) (int64, error)
	GetListCart(params *request.PaginationRequest, userID uint64) (result []*models.Cart, err error)
	GetOneCartsById(cartID uint64) (*models.Cart, error)
	DeleteCart(cartID uint64) error
}
