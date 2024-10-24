package internals

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// ComparePasswords will create password using bcrypt
func ComparePasswords(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("The '" + password + "' and '" + hashedPassword + "' strings don't match")
	}
	return nil
}
