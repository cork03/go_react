package gatewayBoundary

type IUserGateway interface {
	ExistByEmail(email string) (bool, error)
}
