package handler

import (
	"net/http"

	"github.com/baihakhi/simple-shop/internal/models"
	response "github.com/baihakhi/simple-shop/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {
	data := new(models.User)

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}
	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}
	result, err := h.service.CreateUser(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.SUCCESS,
		Data: map[string]string{
			"username": result,
		},
	})
}

func (h *Handler) Login(c echo.Context) error {
	data := new(models.User)

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}
	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	result, err := h.service.Login(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.SUCCESS,
		Data: map[string]string{
			"token": result,
		},
	})
}
