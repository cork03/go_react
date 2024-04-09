package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-rest-api/domain"
	"go-rest-api/usecase"
	"go-rest-api/usecase/input"
	"log/slog"
	"net/http"
	"os"
)

type SignUpController struct {
	signUpUsecase usecase.ISignUpUsecase
}

func (signUpController *SignUpController) SignUp(c echo.Context) error {
	signUpRequest := signUpRequest{}
	if err := c.Bind(&signUpRequest); err != nil {
		return err
	}
	userPassword, passwordErr := domain.NewUserPassword(signUpRequest.User.Password)
	if passwordErr != nil {
		return passwordErr
	}
	signUpUsecaseInput := input.SignUpInput{
		Company: domain.Company{
			Name:       signUpRequest.Company.Name,
			PostalCode: signUpRequest.Company.PostalCode,
			Prefecture: signUpRequest.Company.Prefecture,
			Town:       signUpRequest.Company.Town,
			Area:       signUpRequest.Company.Area,
			Tel:        signUpRequest.Company.Tel,
		},
		User: domain.User{
			Name:  signUpRequest.User.Name,
			Email: signUpRequest.User.Email,
		},
		UserPassword: userPassword,
	}
	err := signUpController.signUpUsecase.SignUp(signUpUsecaseInput)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")
}

type CompanyRequest struct {
	Name       string `json:"name"`
	PostalCode string `json:"postal_code"`
	Prefecture string `json:"prefecture"`
	Town       string `json:"town"`
	Area       string `json:"area"`
	Tel        string `json:"tel"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpRequest struct {
	Company CompanyRequest `json:"company"`
	User    UserRequest    `json:"user"`
}

func (signUpController *SignUpController) MailCertification(c echo.Context) error {
	mailCertificationRequest := mailCertificationRequest{}
	// @todo slogをdiする
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	if err := c.Bind(&mailCertificationRequest); err != nil {
		slog.Error(err.Error())
		return err
	}
	err := signUpController.signUpUsecase.MailCertification(
		input.MailCertificationInput{
			Token: mailCertificationRequest.Token,
		},
	)
	if errors.Is(err, usecase.ExpiredError{}) {
		return c.JSON(http.StatusBadRequest, "認証期限が切れています")
	}
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, "success")
}

type mailCertificationRequest struct {
	Token string `json:"token"`
}

func NewUserController(signUpUsecase usecase.ISignUpUsecase) *SignUpController {
	return &SignUpController{signUpUsecase: signUpUsecase}
}
