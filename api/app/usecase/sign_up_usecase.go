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
	draftCompanyGateway      gatewayBoundary.IDraftCompanyGateway
}

func (signUpUsecase *signUpUsecase) SignUp(input input.SignUpInput) error {
	// メール認証のためのトークンを生成
	uuid := uuid.NewString()
	location, _ := time.LoadLocation("Asia/Tokyo")
	mailCertification := domain.MailCertification{
		Token:  uuid,
		Expire: time.Now().Add(time.Hour * 24).In(location),
	}
	resMailCertification, mailCertificationErr := signUpUsecase.mailCertificationGateway.Create(mailCertification)
	if mailCertificationErr != nil {
		return mailCertificationErr
	}
	// draft会社の登録
	if draftCompanyErr := signUpUsecase.draftCompanyGateway.Create(input.Company, resMailCertification.ID); draftCompanyErr != nil {
		return draftCompanyErr
	}
	// draftユーザーの登録
	// draftPasswordの登録
	// メール送信

	return nil
}

func NewSignUpUsecase(
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway,
	draftCompanyGateway gatewayBoundary.IDraftCompanyGateway,
) ISignUpUsecase {
	return &signUpUsecase{
		mailCertificationGateway: mailCertificationGateway,
		draftCompanyGateway:      draftCompanyGateway,
	}
}
