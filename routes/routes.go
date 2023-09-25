package routes

import (
	"example/cloud-api/controllers"
	"example/cloud-api/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	userRouter := router.Group("/users")
	{
		userRouter.POST("/", middlewares.Auth(), controllers.Create)
		userRouter.GET("/", middlewares.Auth(), controllers.GetAll)
		userRouter.GET("/:id", middlewares.Auth(), controllers.GetOne)
		userRouter.PUT("/:id", middlewares.Auth(), controllers.Update)
		userRouter.DELETE("/:id", middlewares.Auth(), controllers.Delete)
	}
}
