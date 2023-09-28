package main

import (
	docs "example/cloud-api/docs"
	"example/cloud-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Cloud Minor - Go - Gin API
// @version         1.0
// @description     Api used for the Cloud minor project - Language Fitness controllers..
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fmt.Println("Starting app")

	r := gin.Default()
	routes.UserRoutes(r)
	routes.AuthRoutes(r)
	routes.OpenAiRoutes(r)
	routes.ModuleRoutes(r)
	routes.SchoolRoutes(r)

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
