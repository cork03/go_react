package controller

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/domain"
	"go-rest-api/usecase"
	"go-rest-api/usecase/input"
	"net/http"
)

type UserController struct {
	signUpUsecase usecase.ISignUpUsecase
}

func (userController *UserController) SignUp(c echo.Context) error {
	signUpRequest := signUpRequest{}
	if err := c.Bind(&signUpRequest); err != nil {
		return err
	}
	signUpUsecaseInput := input.SignUpInput{Company: domain.Company{
		Name:       signUpRequest.Company.Name,
		PostalCode: signUpRequest.Company.PostalCode,
		Prefecture: signUpRequest.Company.Prefecture,
		Town:       signUpRequest.Company.Town,
		Area:       signUpRequest.Company.Area,
		Tel:        signUpRequest.Company.Tel,
	}}
	err := userController.signUpUsecase.SignUp(signUpUsecaseInput)
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

func NewUserController(signUpUsecase usecase.ISignUpUsecase) UserController {
	return UserController{signUpUsecase: signUpUsecase}
}
