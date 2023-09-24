package main

import (
	"example/cloud-api/Routes"
	docs "example/cloud-api/docs"
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
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fmt.Println("Starting app")

	r := gin.Default()
	Routes.UserRoutes(r)
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
