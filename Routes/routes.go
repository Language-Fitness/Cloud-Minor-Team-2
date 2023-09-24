package Routes

import (
	"example/cloud-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	userRouter := router.Group("/users")

	{
		userRouter.POST("/", controllers.Create)
		userRouter.GET("/", controllers.GetAll)
		userRouter.GET("/:id", controllers.GetOne)
		userRouter.PUT("/:id", controllers.Update)
		userRouter.DELETE("/:id", controllers.Delete)
	}
}
