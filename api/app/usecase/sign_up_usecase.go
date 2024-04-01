package usecase

import (
	"go-rest-api/gatewayBoundary"
	"go-rest-api/usecase/input"
)

type ISignUpUsecase interface {
	SignUp(input input.SignUpInput) error
}

type signUpUsecase struct {
	companyGateway gatewayBoundary.ICompanyGateway
}

func (signUpUsecase *signUpUsecase) SignUp(input input.SignUpInput) error {
	if err := signUpUsecase.companyGateway.Register(input.Company); err != nil {
		return err
	}
	return nil
}

func NewSignUpUsecase(companyGateway gatewayBoundary.ICompanyGateway) ISignUpUsecase {
	return &signUpUsecase{companyGateway: companyGateway}
}
