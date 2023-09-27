package routes

import (
	AuthHandler "example/cloud-api/handlers/auth"
	UserHandler "example/cloud-api/handlers/user"
	"example/cloud-api/middlewares"
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := UserHandler.NewUserHandler(userService)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/", middlewares.Auth(), userHandler.CreateUser)
		userRouter.GET("/", middlewares.Auth(), userHandler.GetAllUsers)
		userRouter.GET("/:id", middlewares.Auth(), userHandler.GetUserByID)
		userRouter.PUT("/:id", middlewares.Auth(), userHandler.UpdateUser)
		userRouter.DELETE("/:id", middlewares.Auth(), userHandler.DeleteUser)
	}
}

func AuthRoutes(r *gin.Engine) {

	authService := services.NewAuthService()
	authHandler := AuthHandler.NewAuthHandler(authService)

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/login", authHandler.Login)
		//authRouter.POST("/register", authHandler.Register)
	}
}
