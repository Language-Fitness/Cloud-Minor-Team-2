package routes

import (
	AuthHandler "example/cloud-api/handlers/auth"
	UserHandler "example/cloud-api/handlers/user"
	"example/cloud-api/middlewares"
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	userService := services.NewUserService()
	userHandler := UserHandler.NewUserHandler(userService)

	userRouter := r.Group("/users")
	{
		userRouter.Use(middlewares.Auth())

		userRouter.POST("/", userHandler.CreateUser)
		userRouter.GET("/", userHandler.GetAllUsers)
		userRouter.GET("/:id", userHandler.GetUserByID)
		userRouter.PUT("/:id", userHandler.UpdateUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
	}
}

func SchoolRoutes(r *gin.Engine) {

	userService := services.NewUserService()
	userHandler := UserHandler.NewUserHandler(userService)

	userRouter := r.Group("/school")
	{
		userRouter.Use(middlewares.Auth())

		userRouter.GET("/", userHandler.GetAllUsers)
		userRouter.GET("/:id", userHandler.GetUserByID)
	}
}

func ModuleRoutes(r *gin.Engine) {
	userService := services.NewUserService()
	userHandler := UserHandler.NewUserHandler(userService)

	moduleRouter := r.Group("/module/:module_id")
	{
		moduleRouter.Use(middlewares.Auth())

		// /module/:module_id/
		moduleRouter.GET("/", userHandler.GetModuleByID)
		// /module/:module_id/courses/
		moduleRouter.GET("/courses", userHandler.GetCoursesForModule)
		// Define a route for "module/:module_id/course/:course_id"
		courseRouter := moduleRouter.Group("/course/:course_id")
		{
			// Middleware for the /module/:module_id/course/:course_id group
			// courseRouter.Use(middlewares.SomeMiddleware())

			// /module/:module_id/course/:course_id/
			courseRouter.GET("/", someHandler)
			// /module/:module_id/course/:course_id/exercises
			courseRouter.GET("/exercises", someHandler)

			// Define a route for "module/:module_id/course/:course_id/exercise/:exercise_id"
			exerciseRouter := courseRouter.Group("/exercise/:exercise_id")
			{
				// /module/:module_id/course/:course_id/exercise/:exercise_id
				exerciseRouter.GET("/", someHandler)
			}
		}
	}

	r.GET("/modules", userHandler.GetAllModules)
}

func ResultRoutes(r *gin.Engine) {

}

func OpenAiRoutes(r *gin.Engine) {

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
