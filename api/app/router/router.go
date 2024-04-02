package router

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/driver"
	"go-rest-api/gateway"
	"go-rest-api/usecase"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	db := db.Main()
	CompanyDriver := driver.NewCompanyDriver(db)
	CompanyGateway := gateway.NewCompanyGateway(CompanyDriver)
	MailCertificationDriver := driver.NewMailCertificationDriver(db)
	MailCertificationGateway := gateway.NewMailCertificationGateway(MailCertificationDriver)
	SignUpUsecase := usecase.NewSignUpUsecase(CompanyGateway, MailCertificationGateway)
	SignUpController := controller.NewUserController(SignUpUsecase)
	e.POST("/signup", SignUpController.SignUp)

	return e
}
