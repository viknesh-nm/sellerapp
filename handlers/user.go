package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/viknesh-nm/sellerapp/backend"
	"github.com/viknesh-nm/sellerapp/domain"
)

// User holds for the user handlers method from routes
type User struct{}

// List returns the userList
func (u *User) List(c echo.Context) error {
	req := &domain.UserRequest{}

	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	list, err := backend.User.UserList(*req)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(200, list)
}

// ListMongo list the user list from mongodb
func (u *User) ListMongo(c echo.Context) error {
	req := &domain.UserRequest{}

	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	list, err := backend.User.UserListMongo(*req)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(200, list)
}

// MongoConvertions transfers the data from mysql to moongoDB
func (u *User) MongoConvertions(c echo.Context) error {

	err := backend.User.MongoConvertions()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(200, `{"status": "success"}`)
}
