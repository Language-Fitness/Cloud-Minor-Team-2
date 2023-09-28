package routes

import (
	AuthHandler "example/cloud-api/handlers/auth"
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
)

func AdminUserRoutes(r *gin.Engine) {

}

func AdminSchoolRoutes(r *gin.Engine) {

}

func AdminModuleRoutes(r *gin.Engine) {

}

func AdminResultRoutes(r *gin.Engine) {

}

func AdminOpenAiRoutes(r *gin.Engine) {

}

func AdminAuthRoutes(r *gin.Engine) {

	authService := services.NewAuthService()
	authHandler := AuthHandler.NewAuthHandler(authService)

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/login", authHandler.Login)
		//authRouter.POST("/register", authHandler.Register)
	}
}
