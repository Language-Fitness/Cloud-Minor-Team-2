package auth

import (
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *Handler {
	return &Handler{authService: authService}
}

// Login godoc
// @Summary      Log in to get a bearer token
// @Description  Login with credentials
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body dto.LoginDTO true "Login data"
// @Success      200  {object}  response.LoginResponseDTO
// @Failure		 400  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /auth/login [post]
func (h *Handler) Login(g *gin.Context) {
	token, err := h.authService.CreateToken("aaaa-bbbb-cccc-dddd", "admin")

	if err == nil {
		g.IndentedJSON(http.StatusOK, gin.H{
			"bearer_token": token,
		})
	}
}
