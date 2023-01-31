package jwt

import (
	// "errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"persons/model"
	"time"
)

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func CreateToken(person model.PersonAuthReq) (string, error) {
	expirationTime := time.Now().Add(1 * time.Minute)

	claims := Claims{
		UserName: person.UserName,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenstring := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenstring.SignedString([]byte("secretkey_dhbfkjhbf3o41nd"))
	fmt.Println(token)
	if err != nil {
		return "", err
	}
	return token, nil
}
