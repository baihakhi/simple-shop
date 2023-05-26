package router

import (
	"net/http"

	"github.com/baihakhi/simple-shop/internal/handler"
	"github.com/baihakhi/simple-shop/internal/middleware"
	"github.com/baihakhi/simple-shop/internal/models"
	"github.com/labstack/echo/v4"
)

func InitRouter(server *echo.Echo, handler *handler.Handler) {
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Internal API Simple Shop App!")
	})

	v1 := server.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", handler.CreateUser)
			user.POST("/login", handler.Login)
		}

		product := v1.Group("/product")
		{
			product.GET("", handler.GetListProducts, middleware.SetMiddlewareAuthentication([]string{models.RoleAdmin, models.RoleCustomer}))
		}

		cart := v1.Group("/cart")
		{
			cart.POST("", handler.AddToCart, middleware.SetMiddlewareAuthentication([]string{models.RoleCustomer}))
			cart.GET("", handler.GetListCart, middleware.SetMiddlewareAuthentication([]string{models.RoleCustomer, models.RoleAdmin}))
			cart.DELETE("/:cart_id", handler.DeleteCart, middleware.SetMiddlewareAuthentication([]string{models.RoleCustomer}))
		}
	}
}
