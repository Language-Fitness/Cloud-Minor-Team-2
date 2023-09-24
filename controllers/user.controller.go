package controllers

import "github.com/gin-gonic/gin"

// GetOne godoc
// @Summary      Get one user
// @Description  Get user by ID
// @Security BearerAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  Domain.User
// @Failure		 401  {object}	Domain.ErrorResponseDTO
// @Failure		 403  {object}	Domain.ErrorResponseDTO
// @Failure		 404  {object}	Domain.ErrorResponseDTO
// @Failure		 500  {object}	Domain.ErrorResponseDTO
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
// @Success      200  {array}   Domain.User
// @Failure		 401  {object}	Domain.ErrorResponseDTO
// @Failure		 403  {object}	Domain.ErrorResponseDTO
// @Failure		 500  {object}	Domain.ErrorResponseDTO
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
// @Param        user body Domain.User true "New user data"
// @Success      200  {object}  Domain.User
// @Failure		 400  {object}	Domain.ErrorResponseDTO
// @Failure		 401  {object}	Domain.ErrorResponseDTO
// @Failure		 403  {object}	Domain.ErrorResponseDTO
// @Failure		 500  {object}	Domain.ErrorResponseDTO
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
// @Failure		 401  {object}	Domain.ErrorResponseDTO
// @Failure		 403  {object}	Domain.ErrorResponseDTO
// @Failure		 404  {object}	Domain.ErrorResponseDTO
// @Failure		 500  {object}	Domain.ErrorResponseDTO
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
// @Param        user body Domain.User true "New user data"
// @Success      200  {object}  Domain.User
// @Failure		 400  {object}	Domain.ErrorResponseDTO
// @Failure		 401  {object}	Domain.ErrorResponseDTO
// @Failure		 403  {object}	Domain.ErrorResponseDTO
// @Failure		 404  {object}	Domain.ErrorResponseDTO
// @Failure		 500  {object}	Domain.ErrorResponseDTO
// @Router       /users/{id} [patch]
func Update(g *gin.Context) {

}
