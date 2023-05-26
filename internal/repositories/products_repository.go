package repositories

import (
	"strings"

	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/models/payload/request"
	"github.com/baihakhi/simple-shop/internal/repositories/queries"
)

func (r *repository) GetListProducts(params *request.PaginationRequest) (result []*models.Products, err error) {
	rows, err := r.db.Query(queries.GetListProducts, strings.ToUpper(params.Key), params.Limit, (params.Page-1)*params.Limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var data models.Products
		if rows.Scan(
			&data.ID,
			&data.Code,
			&data.Title,
			&data.Price,
			&data.Weight,
			&data.Stock,
			&data.Address,
			&data.Category,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, &data)
	}

	return result, nil
}

func (r *repository) GetOneProductByID(productID uint64) (*models.Products, error) {
	var result models.Products
	if err := r.db.QueryRow(queries.GetOneProductByID, productID).Scan(
		&result.ID,
		&result.Code,
		&result.Title,
		&result.Price,
		&result.Weight,
		&result.Stock,
		&result.Address,
		&result.Category,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &result, nil
}
