package routes

import (
	"Gin-API/src/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/", controllers.GetAllUsers)
		userRouter.GET("/:id", controllers.GetUserByID)
	}

}
