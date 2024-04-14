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

type ExistUser struct{}

func (e ExistUser) Error() string {
	return ""
}

type ISignUpUsecase interface {
	MailCertification(input input.MailCertificationInput) error
	SignUp(input input.SignUpInput) error
}

type signUpUsecase struct {
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway
	draftGateway             gatewayBoundary.IDraftGateway
	userGateway              gatewayBoundary.IUserGateway
	mailSendGateway          mail.IMailSendGateway
}

func (signUpUsecase *signUpUsecase) MailCertification(input input.MailCertificationInput) error {
	// @todo 既に登録されているメールアドレスかどうかを確認

	// トークンからメール認証情報を取得
	drafts, getByTokenErr := signUpUsecase.mailCertificationGateway.GetDraftsByToken(input.Token)
	if getByTokenErr != nil {
		return getByTokenErr
	}
	// 認証情報がなければエラー
	if drafts == nil {
		return CanNotAuthorized{}
	}
	// 期限切れならエラー
	location, _ := time.LoadLocation("Asia/Tokyo")
	if time.Now().In(location).After(drafts.MailCertification.Expire) {
		return CanNotAuthorized{}
	}
	// 問題なければユーザー情報を登録して本登録
	if signUpErr := signUpUsecase.draftGateway.BookRegistration(*drafts); signUpErr != nil {
		return signUpErr
	}
	return nil
}

func (signUpUsecase *signUpUsecase) SignUp(input input.SignUpInput) error {
	// 既に登録されているかのチェック
	existUser, existUserErr := signUpUsecase.userGateway.ExistByEmail(input.User.Email)
	if existUserErr != nil {
		return existUserErr
	}
	if existUser {
		return ExistUser{}
	}
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
	if mailSendErr := signUpUsecase.mailSendGateway.SendMailCertification(user.Email, mailCertification.Token, mailCertification.Expire); mailSendErr != nil {
		return mailSendErr
	}
	return nil
}

func NewSignUpUsecase(
	mailCertificationGateway gatewayBoundary.IMailCertificationGateway,
	draftCompanyGateway gatewayBoundary.IDraftGateway,
	mailSendGateway mail.IMailSendGateway,
	userGateway gatewayBoundary.IUserGateway,
) ISignUpUsecase {
	return &signUpUsecase{
		mailCertificationGateway: mailCertificationGateway,
		draftGateway:             draftCompanyGateway,
		mailSendGateway:          mailSendGateway,
		userGateway:              userGateway,
	}
}
