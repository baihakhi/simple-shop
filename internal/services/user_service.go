package services

import (
	"strings"

	"github.com/baihakhi/simple-shop/internal/middleware"
	"github.com/baihakhi/simple-shop/internal/models"
	hash "github.com/baihakhi/simple-shop/internal/utils/bcrypt"
)

func (s *service) CreateUser(data *models.User) (string, error) {
	hashedPass, err := hash.Encrypt(data.Password)
	if err != nil {
		return "", err
	}
	data.Password = hashedPass
	data.Role = strings.ToUpper(models.RoleCustomer)
	return s.repositories.CreateUser(data)
}

func (s *service) Login(data *models.User) (string, error) {
	pass, err := s.repositories.GetPasswordByUsername(strings.ToLower(data.Username))
	if err != nil {
		return "", err
	}

	if err := hash.VerifyPassword(pass, data.Password); err != nil {
		return "", err
	}

	acc, err := s.repositories.GetOneUsersByUsername(strings.ToLower(data.Username))
	if err != nil {
		return "", err
	}

	return middleware.CreateToken(*acc)
}
