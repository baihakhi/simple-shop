package router

import (
	"net/http"

	"github.com/baihakhi/simple-shop/internal/handler"
	"github.com/labstack/echo/v4"
)

func InitRouter(server *echo.Echo, handler *handler.Handler) {
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Internal API Simple Shop App!")
	})
}
