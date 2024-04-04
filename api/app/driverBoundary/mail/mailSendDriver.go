package mail

type IMailSendDriver interface {
	SendMailCertification(email string) error
}
