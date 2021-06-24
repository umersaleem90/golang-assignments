package utilities

import (
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool{
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func ComparePasswordHash(password, hash string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        return errors.New("incorrect password")
    }
    return nil
}
