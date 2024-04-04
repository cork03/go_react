package mail

type IMailSendGateway interface {
	SendMailCertification(email string) error
}
