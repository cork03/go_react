package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-rest-api/domain"
	"go-rest-api/usecase"
	"go-rest-api/usecase/input"
	"log/slog"
	"net/http"
)

type SignUpController struct {
	signUpUsecase usecase.ISignUpUsecase
}

func (signUpController *SignUpController) SignUp(c echo.Context) error {
	signUpRequest := signUpRequest{}
	if err := c.Bind(&signUpRequest); err != nil {
		slog.Error(err.Error())
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
	if errors.Is(err, usecase.ExistUser{}) {
		return c.JSON(http.StatusBadRequest, errorResponse{Message: "既に登録されているメールアドレスです。"})
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")
}

type companyRequest struct {
	Name       string `json:"name"`
	PostalCode string `json:"postal_code"`
	Prefecture string `json:"prefecture"`
	Town       string `json:"town"`
	Area       string `json:"area"`
	Tel        string `json:"tel"`
}

type userRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpRequest struct {
	Company companyRequest `json:"company"`
	User    userRequest    `json:"user"`
}

func (signUpController *SignUpController) MailCertification(c echo.Context) error {
	mailCertificationRequest := mailCertificationRequest{}
	if err := c.Bind(&mailCertificationRequest); err != nil {
		slog.Error(err.Error())
		return err
	}
	err := signUpController.signUpUsecase.MailCertification(
		input.MailCertificationInput{
			Token: mailCertificationRequest.Token,
		},
	)
	if errors.Is(err, usecase.ExistUser{}) {
		return c.JSON(http.StatusBadRequest, errorResponse{Message: "既に登録されているメールアドレスです。"})
	}
	if errors.Is(err, usecase.CanNotAuthorized{}) {
		return c.JSON(http.StatusUnauthorized, errorResponse{Message: "認証期限が切れている。または無効なURLです。"})
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

type errorResponse struct {
	Message string `json:"message"`
}

func NewUserController(signUpUsecase usecase.ISignUpUsecase) *SignUpController {
	return &SignUpController{signUpUsecase: signUpUsecase}
}
