package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name  string
	Email string
}

type UserPassword struct {
	password string
}

func (userPassword UserPassword) GetPassword() string {
	return userPassword.password
}

func NewUserPassword(rawPassword string) (UserPassword, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return UserPassword{}, err
	}
	return UserPassword{password: string(hashed)}, nil
}
