package security

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHandler struct{}

func (p *PasswordHandler) Hash(password string) (string, error) {
	pwdBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if the plain password matches the hashed password
func (p *PasswordHandler) VerifyPassword(hashedPassword string, plainPassword string) bool {
	fmt.Println(hashedPassword, plainPassword)
	passwordByteEnc := []byte(plainPassword)
	hashedPasswordBytes := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordByteEnc)
	if err != nil {
		log.Printf("PasswordHandler Error verifying password: %s", err)
		return false
	}
	return true
}
