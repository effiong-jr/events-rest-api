package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("mySignedToken")

func GenerateJWT(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		"email":     email,
		"userId":    userId,
	})

	signedToken, err := token.SignedString(mySigningKey)

	return signedToken, err
}

func VerifyJWT(token string) (int64, error) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return 0, errors.New("unexpected signing method")
		}

		return mySigningKey, nil
	})

	isValidToken := parsedToken.Valid

	if !isValidToken {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	fmt.Println("Claims", claims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	fmt.Println(claims)

	userId := int64(claims["userId"].(float64))

	return userId, err

}
