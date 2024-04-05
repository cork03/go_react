package mail

type IMailSendDriver interface {
	SendMailCertification(email string, token string) error
}
