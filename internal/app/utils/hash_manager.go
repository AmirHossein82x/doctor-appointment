package utils

import (

	"golang.org/x/crypto/bcrypt"
)

type UserHashManager interface {
	HashPassword(string) (string, error)
	VerifyPassword(string, string) bool
}

func HashPassword(RawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(RawPassword), bcrypt.DefaultCost)
	return string(bytes), err

}

func VerifyPassword(RawPassword string, HashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(RawPassword))
	return err == nil
}