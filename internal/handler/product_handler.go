package handler

import (
	"net/http"

	"github.com/baihakhi/simple-shop/internal/models/payload/request"
	response "github.com/baihakhi/simple-shop/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetListProducts(c echo.Context) error {
	params := new(request.PaginationRequest)

	if err := params.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}

	result, err := h.service.GetListProducts(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	pagination := &response.Pagination{
		Page:      params.Page,
		Limit:     params.Limit,
		TotalData: len(result),
	}
	pagination.CountTotalPage()

	return c.JSON(http.StatusOK, response.MapResponse{
		Message:  response.SUCCESS,
		Data:     result,
		Metadata: pagination,
	})
}
