package main

import (
	"Gin-API/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run(":8080")
}
