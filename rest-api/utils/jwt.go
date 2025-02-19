package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretSignature string = "!$&!d[Z5Uzi,;!0Ve(-C#2+^9}2N6:uf"

func GenerateToken(email string, id int64) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"id":     id,
		"expire": time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretSignature))
}

func VerifyToken(strToken string) (int64, error) {
	token, err := jwt.Parse(strToken, func(tkn *jwt.Token) (interface{}, error) {
		_, ok := tkn.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretSignature), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token provided")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("wrong token payload provided")
	}

	id := claims["id"].(float64)

	return int64(id), nil
}
