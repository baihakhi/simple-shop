package services

import (
	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/models/payload/request"
)

func (s *service) GetListProducts(params *request.PaginationRequest) ([]*models.Products, error) {
	return s.repositories.GetListProducts(params)
}
