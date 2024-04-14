package driverBoundary

type IUserDriver interface {
	ExistByEmail(email string) (bool, error)
}
