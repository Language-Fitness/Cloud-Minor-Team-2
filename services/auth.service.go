package services

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func ValidateToken(signedToken string) (err error) {
	return
}

//func createToken(userId string, role string) (token string, err error) {
//	claims := CustomClaims{
//		Role: "admin",
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
//			IssuedAt:  jwt.NewNumericDate(time.Now()),
//			NotBefore: jwt.NewNumericDate(time.Now()),
//			Issuer:    "test",
//			Subject:   userId,
//			Audience:  []string{"movie-api"},
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	tokenString, err := token.SignedString(mySigningKey)
//	if err != nil {
//		// Handle error
//	}
//	return
//}
