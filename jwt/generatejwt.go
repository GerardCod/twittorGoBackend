package jwt

import (
	"time"

	"github.com/GerardCod/twittorGoBackend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT converts an user model to JWT format
func GenerateJWT(user models.User) (string, error) {
	key := []byte("Tomate.admin.1234")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"last_name": user.LastName,
		"biografy":  user.Biografy,
		"location":  user.Location,
		"web_site":  user.WebSite,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
