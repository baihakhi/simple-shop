package services

import (
	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/models/payload/request"
	"github.com/baihakhi/simple-shop/internal/repositories"
)

type service struct {
	repositories repositories.Repositories
}

func InitService(repo repositories.Repositories) Services {
	return &service{
		repositories: repo,
	}
}

type Services interface {
	// User Services
	CreateUser(data *models.User) (string, error)
	Login(data *models.User) (string, error)
	GetOneUsersByUsername(username string) (*models.User, error)

	// Product Service
	GetListProducts(params *request.PaginationRequest) ([]*models.Products, error)

	// Cart Services
	CreateCart(data *models.Cart) (int64, error)
	GetListCart(params *request.PaginationRequest, userID uint64) ([]*models.Cart, error)
	GetOneCartsById(cardID uint64) (*models.Cart, error)
	DeleteCart(cartID uint64) error
}
