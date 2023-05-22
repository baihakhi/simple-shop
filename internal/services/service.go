package services

import (
	"github.com/baihakhi/simple-shop/internal/models"
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
}
