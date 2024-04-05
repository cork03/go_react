package mail

type IMailSendGateway interface {
	SendMailCertification(email string, token string) error
}
