package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

var mySigningKey = []byte("mySecretKey")

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"statusText": "failed",
				"statusCode": 401,
				"errorType":  "UnauthorizedException",
				"error":      "No bearer token",
			})
			return
		}

		err := validateToken(bearerToken)
		if err != nil {

			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"statusText": "failed",
				"statusCode": 403,
				"errorType":  "ForbiddenException",
				"error":      "Invalid credentials",
			})
			return
		}
		c.Next()
	}
}

func validateToken(signedToken string) (err error) {
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
