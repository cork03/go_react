package usecase

import (
	"github.com/google/uuid"
	"go-rest-api/domain"
	"go-rest-api/gatewayBoundary"
	"go-rest-api/gatewayBoundary/mail"
	"go-rest-api/usecase/input"
	"time"
)

type CanNotAuthorized struct {
}

func (e CanNotAuthorized) Error() string {
	return ""
}

type ISignUpUsecase interface {
	MailCertification(input input.MailCertificationInput) error
	SignUp(input input.SignUpInput) error
}

type signUpUsecase struct {
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway
	draftGateway             gatewayBoundary.IDraftGateway
	mailSendGateway          mail.IMailSendGateway
}

func (signUpUsecase *signUpUsecase) MailCertification(input input.MailCertificationInput) error {
	// トークンからメール認証情報を取得
	mailCertification, getByTokenErr := signUpUsecase.mailCertificationGateway.GetByToken(input.Token)
	if getByTokenErr != nil {
		return getByTokenErr
	}
	// 認証情報がなければエラー
	if mailCertification == nil {
		return CanNotAuthorized{}
	}
	// 期限切れならエラー
	location, _ := time.LoadLocation("Asia/Tokyo")
	if time.Now().In(location).After(mailCertification.Expire) {
		return CanNotAuthorized{}
	}
	// 問題なければユーザー情報を登録して本登録
	return nil
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
	// メール送信 @todo queueにいれて非同期でやりたい
	if mailSendErr := signUpUsecase.mailSendGateway.SendMailCertification(user.Email, mailCertification.Token); mailSendErr != nil {
		return mailSendErr
	}
	return nil
}

func NewSignUpUsecase(
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway,
	draftCompanyGateway gatewayBoundary.IDraftGateway,
	mailSendGateway mail.IMailSendGateway,
) ISignUpUsecase {
	return &signUpUsecase{
		mailCertificationGateway: mailCertificationGateway,
		draftGateway:             draftCompanyGateway,
		mailSendGateway:          mailSendGateway,
	}
}
