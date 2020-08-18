package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var authkey string

func InitAuth(signKey string) {
	authkey = signKey
}

func CreateToken(posID int) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["pos_id"] = posID
	atClaims["exp"] = time.Now().Add(3 * time.Hour).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(authkey))
	if err != nil {
		return "", err
	}
	return token, nil
}
