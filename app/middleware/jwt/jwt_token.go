package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"persons/config"
	"persons/model"
	"time"
)

type Claims struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func CreateToken(person model.PersonAuthReq) (string, error) {
	expirationTime := time.Now().Add(120 * time.Minute)

	claims := Claims{
		UserName: person.UserName,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenstring := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenstring.SignedString([]byte(config.Jwt.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateJwtAuthToken(token string) error {
	tkn, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})
	if err != nil {
		return errors.New("Unauthorized")
	}
	if !tkn.Valid {
		return errors.New("Unauthorized")
	}
	return nil
}
