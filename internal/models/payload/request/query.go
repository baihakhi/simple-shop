package request

import "github.com/labstack/echo/v4"

type (
	PaginationRequest struct {
		Page  int64  `json:"page" query:"page"`
		Limit int64  `json:"per_page" query:"per_page"`
		Key   string `json:"key" query:"q"`
	}
)

func (p *PaginationRequest) ValidatePagination() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10
	}
}

func (u *PaginationRequest) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, u)

	u.ValidatePagination()

	return nil
}
