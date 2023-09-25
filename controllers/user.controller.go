package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetOne godoc
// @Summary      Get one user
// @Description  Get user by ID
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  domain.UserDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 404  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users/{id} [get]
func GetOne(g *gin.Context) {

}

// GetAll godoc
// @Summary      List all users
// @Description  Get all users
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   domain.UserDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users [get]
func GetAll(g *gin.Context) {

}

// Create godoc
// @Summary      Add new user
// @Description  Create a new user
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body domain.UserDTO true "New user data"
// @Success      200  {object}  domain.UserDTO
// @Failure		 400  {object}	response.ErrorResponseDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users [post]
func Create(g *gin.Context) {

}

// Delete godoc
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
func Delete(g *gin.Context) {

}

// Update godoc
// @Summary      Update one user
// @Description  Patch user by ID
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        user body domain.UserDTO true "New user data"
// @Success      200  {object}  domain.UserDTO
// @Failure		 400  {object}	response.ErrorResponseDTO
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 404  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users/{id} [patch]
func Update(g *gin.Context) {

}
