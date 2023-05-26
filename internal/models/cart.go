package models

import (
	"github.com/labstack/echo/v4"
)

type (
	Cart struct {
		ID          uint64 `json:"cart_id" query:"cart_id"`
		UID         uint64 `json:"user_id"`
		PID         uint64 `json:"product_id"`
		Amount      int8   `json:"amount"`
		AmountPrice int    `json:"amount_price"`
		Timestamp
	}
)

func (u *Cart) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	return nil
}
