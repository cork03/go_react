package usecase

import "go-rest-api/usecase/input"

type ISignUpUsecase interface {
	SignUp(input input.SignUpInput) error
}

type signUpUsecase struct {
}

func (signUpUsecase *signUpUsecase) SignUp(input input.SignUpInput) error {
	// companyの登録
	return nil
}

func NewSignUpUsecase() ISignUpUsecase {
	return &signUpUsecase{}
}
