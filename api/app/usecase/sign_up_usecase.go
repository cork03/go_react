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
	draftGateway             gatewayBoundary.IDraftGateway
}

func (signUpUsecase *signUpUsecase) SignUp(input input.SignUpInput) error {
	// @todo 既に登録されているメールアドレスかどうかを確認
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
	// draft情報の登録
	user, draftErr := signUpUsecase.draftGateway.Create(
		input.Company,
		input.User,
		input.UserPassword,
		resMailCertification.ID,
	)
	if draftErr != nil {
		return draftErr
	}
	// メール送信
	println(user.Email)
	return nil
}

func NewSignUpUsecase(
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway,
	draftCompanyGateway gatewayBoundary.IDraftGateway,
) ISignUpUsecase {
	return &signUpUsecase{
		mailCertificationGateway: mailCertificationGateway,
		draftGateway:             draftCompanyGateway,
	}
}
