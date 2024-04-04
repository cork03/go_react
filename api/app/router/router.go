package router

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/driver"
	mailDriver "go-rest-api/driver/mail"
	"go-rest-api/gateway"
	"go-rest-api/gateway/mail"
	"go-rest-api/usecase"
	"gorm.io/gorm"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	db := db.Main()
	e.POST("/signup", createSignUpController(db).SignUp)

	return e
}

func createSignUpController(db *gorm.DB) *controller.SignUpController {
	MailCertificationDriver := driver.NewMailCertificationDriver(db)
	MailCertificationGateway := gateway.NewMailCertificationGateway(MailCertificationDriver)
	DraftCompanyDriver := driver.NewDraftDriver(db)
	DraftCompanyGateway := gateway.NewDraftGateway(DraftCompanyDriver)
	MailSendDriver := mailDriver.NewMailSendDriver()
	MailSendGateway := mail.NewMailSendGateway(MailSendDriver)
	SignUpUsecase := usecase.NewSignUpUsecase(MailCertificationGateway, DraftCompanyGateway, MailSendGateway)
	return controller.NewUserController(SignUpUsecase)
}
