package repositories

import (
	"fmt"

	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/models/payload/request"
	"github.com/baihakhi/simple-shop/internal/repositories/queries"
)

func (r *repository) CreateCart(data *models.Cart) (int64, error) {
	var result int64
	if err := r.db.QueryRow(queries.CreateCart,
		data.UID,
		data.PID,
		data.Amount,
		data.AmountPrice).Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}

func (r *repository) GetListCart(params *request.PaginationRequest, userID uint64) (result []*models.Cart, err error) {
	rows, err := r.db.Query(queries.GetListCart, userID, params.Limit, (params.Page-1)*params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println(userID)
	for rows.Next() {
		var data models.Cart
		if rows.Scan(
			&data.ID,
			&data.UID,
			&data.PID,
			&data.Amount,
			&data.AmountPrice,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			return nil, err
		}
		fmt.Println(data)
		result = append(result, &data)
	}

	return result, nil
}

func (r *repository) GetOneCartsById(cartID uint64) (*models.Cart, error) {
	var result models.Cart

	if err := r.db.QueryRow(queries.GetOneCartsById, cartID).
		Scan(
			&result.ID,
			&result.UID,
			&result.PID,
			&result.Amount,
			&result.AmountPrice,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) DeleteCart(cartID uint64) error {
	_, err := r.db.Exec(queries.DeleteCart, cartID)

	return err
}
