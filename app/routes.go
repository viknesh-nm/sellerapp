package app

import (
	"github.com/labstack/echo"
	"github.com/viknesh-nm/sellerapp/handlers"
)

//Routes creates the handlers, initializes the  muxes & sub-muxes & returns the final mux
func Routes() *echo.Echo {

	e := echo.New()

	user := handlers.User{}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, "success")
	})

	e.GET("/users", user.List)
	e.GET("/users_mongo", user.ListMongo)
	e.GET("/conversion", user.MongoConvertions)

	return e

}
