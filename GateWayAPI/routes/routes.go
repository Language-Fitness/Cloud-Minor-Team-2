package routes

import (
	AuthHandler "example/cloud-api/handlers/auth"
	ModuleHandler "example/cloud-api/handlers/module"
	OpenHandler "example/cloud-api/handlers/open"
	SchoolHandler "example/cloud-api/handlers/school"
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

		userRouter.GET("/", userHandler.GetAllUsers)
		userRouter.GET("/leaderboard", userHandler.GetAllUsers)
		userRouter.GET("/:id", userHandler.GetUserByID)
	}
}

func SchoolRoutes(r *gin.Engine) {

	schoolService := services.NewSchoolService()
	schoolHandler := SchoolHandler.NewSchoolHandler(schoolService)

	schoolRouter := r.Group("/school")
	{
		schoolRouter.Use(middlewares.Auth())

		schoolRouter.GET("/", schoolHandler.GetAllSchools)
		schoolRouter.GET("/:id", schoolHandler.GetSchoolByID)
	}
}

func ModuleRoutes(r *gin.Engine) {
	moduleService := services.NewModuleService()
	moduleHandler := ModuleHandler.NewModuleHandler(moduleService)

	moduleRouter := r.Group("/module/:module_id")
	{
		moduleRouter.Use(middlewares.Auth())

		// /module/:module_id/
		moduleRouter.GET("/", moduleHandler.GetModuleByID)
		// /module/:module_id/courses/
		moduleRouter.GET("/courses", moduleHandler.GetAllCourses)
		// Define a route for "module/:module_id/course/:course_id"
		courseRouter := moduleRouter.Group("/course/:course_id")
		{
			// Middleware for the /module/:module_id/course/:course_id group
			// courseRouter.Use(middlewares.SomeMiddleware())

			// /module/:module_id/course/:course_id/
			courseRouter.GET("/", moduleHandler.GetAllCoursesById)
			// /module/:module_id/course/:course_id/exercises
			courseRouter.GET("/exercises", moduleHandler.GetAllExercises)

			// Define a route for "module/:module_id/course/:course_id/exercise/:exercise_id"
			exerciseRouter := courseRouter.Group("/exercise/:exercise_id")
			{
				// /module/:module_id/course/:course_id/exercise/:exercise_id
				exerciseRouter.GET("/", moduleHandler.GetExerciseById)

				resultRouter := exerciseRouter.Group("/result")
				{
					resultRouter.GET("/", moduleHandler.GetResultForExercise)
					resultRouter.POST("/", moduleHandler.CreateResultForExercise)
				}
			}
		}
	}

	r.GET("/modules", moduleHandler.GetAllModules)
}

func OpenAiRoutes(r *gin.Engine) {

	openService := services.NewOpenService()
	openHandler := OpenHandler.NewOpenHandler(openService)

	openRouter := r.Group("/help")
	{
		openRouter.GET("/", openHandler.GetFeedback)
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
