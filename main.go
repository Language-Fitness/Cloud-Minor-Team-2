package main

import (
	docs "example/cloud-api/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Cloud Minor - Go - Gin API
// @version         1.0
// @description     Api used for the Cloud minor project - Language Fitness App..
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	fmt.Println("Starting app")

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET(":id", GetOne)
			users.GET("", GetAll)
			users.POST("", Create)
			users.DELETE(":id", Delete)
			users.PATCH(":id", Update)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}

// GetOne godoc
// @Summary      Get one user
// @Description  Get user by ID
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
