package user

import (
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *Handler {
	return &Handler{userService: userService}
}

// GetAllUsers godoc
// @Summary      List all users
// @Description  Get all users
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   response.UserResponseDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	h.userService.GetAllUsers(c)
}

// GetUserByID godoc
// @Summary      Get one user
// @Description  Get user by ID
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  response.UserResponseDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 404  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	h.userService.GetOneUser(c)
}

// CreateUser godoc
// @Summary      Add new user
// @Description  Create a new user
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body dto.UserDTO true "New user data"
// @Success      200  {object}  response.UserResponseDTO
// @Failure		 400  {object}	response.ErrorResponseDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	h.userService.CreateUser(c)
}

// UpdateUser godoc
// @Summary      Update one user
// @Description  Patch user by ID
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        user body dto.UserDTO true "New user data"
// @Success      200  {object}  response.UserResponseDTO
// @Failure		 400  {object}	response.ErrorResponseDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 404  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users/{id} [patch]
func (h *Handler) UpdateUser(c *gin.Context) {
	h.userService.UpdateUser(c)
}

// DeleteUser godoc
// @Summary      Remove one user
// @Description  Delete a user by ID
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 404  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	h.userService.DeleteUser(c)
}
