package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySigningKey = []byte("mySecretKey")

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return mySigningKey, nil
		},
	)

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Role, claims.RegisteredClaims.Issuer)
		return nil
	} else {
		err = errors.New("token expired")
		return
	}
}

func CreateToken(userId string, role string) (tokenString string, err error) {

	claims := CustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   userId,
			Audience:  []string{"movie-api"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(mySigningKey)
	if err != nil {
		// Handle error
	}
	return tokenString, err
}
