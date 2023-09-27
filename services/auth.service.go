package services

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySigningKey = []byte("mySecretKey")

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a AuthService) CreateToken(userId string, role string) (tokenString string, err error) {
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
