package router

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/controller"
	"go-rest-api/usecase"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	controller := controller.NewUserController(usecase.NewSignUpUsecase())
	e.POST("/signup", controller.SignUp)

	return e
}
