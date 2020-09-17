package bd

import (
	"github.com/GerardCod/twittorGoBackend/models"
	"golang.org/x/crypto/bcrypt"
)

//Login does the signin in the DB
func Login(email string, password string) (models.User, bool) {
	user, found, _ := UserAlreadyExists(email)

	if !found {
		return user, found
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, found
}
