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
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		accessToken := strings.Split(bearerToken, "Bearer ")[1]
		err := services.ValidateToken(accessToken)
		if err != nil {

			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
