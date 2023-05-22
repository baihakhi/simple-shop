package middleware

import (
	"context"
	"net/http"

	"github.com/baihakhi/simple-shop/internal/models"
	response "github.com/baihakhi/simple-shop/internal/models/payload/responses"
	"github.com/baihakhi/simple-shop/internal/utils"
	"github.com/labstack/echo/v4"
)

func SetMiddlewareAuthentication(access []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := TokenValid(c.Request())
			if err != nil {
				c.JSON(http.StatusUnauthorized, response.MapResponse{
					Message: response.InvalidToken,
				})
				return err
			}

			isAllowed := utils.IsContain(user.Role, access)
			if !isAllowed {
				c.JSON(http.StatusForbidden, response.MapResponse{
					Message: response.AccessDenied,
				})
				return err
			}
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), models.ACC, user)))
			return next(c)
		}
	}
}
