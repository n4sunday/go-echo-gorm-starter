package routers

import (
	"main/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Echo")
	})
	api := e.Group("/api")

	users := api.Group("/users")

	// Users Controller
	users.GET("", controllers.GetUsers)

	return e
}
