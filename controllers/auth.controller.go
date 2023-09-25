package controllers

import (
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary      Log in to get a bearer token
// @Description  Login with credentials
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body domain.LoginDTO true "Login data"
// @Success      200  {object}  response.LoginResponseDTO
// @Failure		 400  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /auth/login [post]
func Login(g *gin.Context) {
	token, err := services.CreateToken("aaaa-bbbb-cccc-dddd", "admin")

	if err == nil {
		g.IndentedJSON(http.StatusOK, gin.H{
			"bearer_token": token,
		})
	}
}
