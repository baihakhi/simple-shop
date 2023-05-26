package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/baihakhi/simple-shop/internal/models/payload/request"
	response "github.com/baihakhi/simple-shop/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

func (h *Handler) AddToCart(c echo.Context) error {
	data := new(models.Cart)
	account := c.Request().Context().Value(models.ACC).(*models.User)

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}

	user, err := h.service.GetOneUsersByUsername(account.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}
	data.UID = user.ID
	result, err := h.service.CreateCart(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.MapResponse{
		Message: response.CREATED,
		Data: map[string]int64{
			"cart ID": result,
		},
	})
}

func (h *Handler) GetListCart(c echo.Context) error {
	data := new(request.PaginationRequest)
	account := c.Request().Context().Value(models.ACC).(*models.User)

	if err := data.GetDataFromHTTPRequest(c); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: response.BADREQUEST,
		})
	}
	user, err := h.service.GetOneUsersByUsername(account.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	fmt.Println(user)
	result, err := h.service.GetListCart(data, user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.MapResponse{
		Message: response.SUCCESS,
		Data:    result,
	})
}

func (h *Handler) DeleteCart(c echo.Context) error {
	account := c.Request().Context().Value(models.ACC).(*models.User)

	cartId, _ := strconv.Atoi(c.Param("cart_id"))

	cart, err := h.service.GetOneCartsById(uint64(cartId))
	fmt.Println(cart, cartId, err)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.MapResponse{
			Message: err.Error(),
		})
	}

	user, err := h.service.GetOneUsersByUsername(account.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	if cart.UID != user.ID {
		return c.JSON(http.StatusForbidden, response.MapResponse{
			Message: response.AccessDenied,
		})
	}
	fmt.Println(cart, user)
	if err := h.service.DeleteCart(uint64(cartId)); err != nil {
		return c.JSON(http.StatusBadRequest, response.MapResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.MapResponse{
		Message: response.SUCCESS,
		Data:    response.SUCCESS,
	})
}
