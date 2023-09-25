package middlewares

import (
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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

		accessToken := strings.Split(bearerToken, "Bearer ")[1]
		err := services.ValidateToken(accessToken)
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
