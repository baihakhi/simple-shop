package services

import (
	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/models/payload/request"
)

func (s *service) CreateCart(data *models.Cart) (int64, error) {
	product, err := s.repositories.GetOneProductByID(uint64(data.PID))
	if err != nil {
		return 0, err
	}
	data.AmountPrice = product.Price * int(data.Amount)

	return s.repositories.CreateCart(data)
}

func (s *service) GetListCart(params *request.PaginationRequest, userID uint64) ([]*models.Cart, error) {
	return s.repositories.GetListCart(params, userID)
}

func (s *service) GetOneCartsById(cardID uint64) (*models.Cart, error) {
	return s.repositories.GetOneCartsById(cardID)
}

func (s *service) DeleteCart(cartID uint64) error {
	return s.repositories.DeleteCart(cartID)
}
