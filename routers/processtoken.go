package routers

import (
	"errors"
	"strings"

	"github.com/GerardCod/twittorGoBackend/bd"

	"github.com/GerardCod/twittorGoBackend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string

var UserID string

//ProccessToken is a function that verifies the token's validity
func ProccessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("Tomate.admin.1234")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("token invalid")
	}

	token = strings.TrimSpace(splitToken[1])
	validToken, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, found, ID := bd.UserAlreadyExists(claims.Email)
		if found {

		}
	}
}
