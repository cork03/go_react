package usecase

import (
	"github.com/google/uuid"
	"go-rest-api/domain"
	"go-rest-api/gatewayBoundary"
	"go-rest-api/usecase/input"
	"time"
)

type ISignUpUsecase interface {
	SignUp(input input.SignUpInput) error
}

type signUpUsecase struct {
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway
}

func (signUpUsecase *signUpUsecase) SignUp(input input.SignUpInput) error {
	// メール認証のためのトークンを生成
	uuid := uuid.NewString()
	location, _ := time.LoadLocation("Asia/Tokyo")
	mailCertification := domain.MailCertification{
		Token:  uuid,
		Expire: time.Now().Add(time.Hour * 24).In(location),
	}
	if err := signUpUsecase.mailCertificationGateway.Create(mailCertification); err != nil {
		return err
	}
	// draft会社の登録
	// draftユーザーの登録
	// draftPasswordの登録
	// メール送信

	return nil
}

func NewSignUpUsecase(
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway,
) ISignUpUsecase {
	return &signUpUsecase{
		mailCertificationGateway: mailCertificationGateway,
	}
}
