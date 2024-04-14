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
	signUpController := createSignUpController(db)
	e.POST("/signup", signUpController.SignUp)
	e.POST("/mail-certification", signUpController.MailCertification)

	return e
}

func createSignUpController(db *gorm.DB) *controller.SignUpController {
	MailCertificationDriver := driver.NewMailCertificationDriver(db)
	MailCertificationGateway := gateway.NewMailCertificationGateway(MailCertificationDriver)
	DraftCompanyDriver := driver.NewDraftDriver(db)
	CompanyDriver := driver.NewCompanyDriver(db)
	DraftCompanyGateway := gateway.NewDraftGateway(DraftCompanyDriver, CompanyDriver)
	MailSendDriver := mailDriver.NewMailSendDriver()
	MailSendGateway := mail.NewMailSendGateway(MailSendDriver)
	userDriver := driver.NewUserDriver(db)
	UserGateway := gateway.NewUserGateway(userDriver)

	SignUpUsecase := usecase.NewSignUpUsecase(
		MailCertificationGateway,
		DraftCompanyGateway,
		MailSendGateway,
		UserGateway,
	)
	return controller.NewUserController(SignUpUsecase)
}
