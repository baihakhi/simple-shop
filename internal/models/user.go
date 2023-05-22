package models

import (
	response "github.com/baihakhi/simple-shop/internal/models/payload/responses"
	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID       uint64 `json:"user_id"`
		Username string `form:"username"`
		Fullname string `json:"fullname"`
		Address  string `json:"address"`
		Role     string `json:"role"`
		Password string `json:"password,omitempty"`
		Balance  int64  `json:"balance"`
		Timestamp
	}
	AccStr string
)

const (
	ACC    AccStr = "account"
	URole  string = "User Role"
	UUname string = "Username"
	UPass  string = "Password"

	// Role For User
	RoleAdmin    string = "ADMIN"
	RoleCustomer string = "CUSTOMER"
)

var (
	IsValidUserGroup = map[string]bool{
		RoleAdmin:    true,
		RoleCustomer: true,
	}
)

func (u *User) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}

	binder := &echo.DefaultBinder{}
	binder.BindHeaders(c, u)

	return nil
}

func (u *User) Validate() error {
	if u.Username == "" {
		return response.ErrorBuilder(UUname, response.MDTR)
	}

	if u.Password == "" {
		return response.ErrorBuilder(UPass, response.MDTR)
	}

	return nil
}
