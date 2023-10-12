package controllers

import (
	"Gin-API/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUsers(context *gin.Context) {

	allUsers, err := services.GetAllUsers()

	if err != nil {
		HandleException(context, err, 500)
		return
	}

	HandleOkResponse(context, "All Users", allUsers)
}

func GetUserByID(context *gin.Context) {

	// Retrieve the `id` parameter from the URL
	id := context.Param("id")

	// Convert the `id` parameter to an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		HandleException(context, err, 400)
		return
	}

	user, err := services.GetUserByID(idInt)

	if err != nil {
		HandleException(context, err, 404)
		return
	}
	response := "User with ID: " + strconv.Itoa(idInt)
	HandleOkResponse(context, response, user)
}

func HandleOkResponse(context *gin.Context, message string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"statusText": "success",
		"statusCode": 200,
		"message":    message,
		"data":       data,
	})
}

func HandleException(context *gin.Context, err error, code int) {
	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"statusText": "failure",
		"statusCode": code,
		"errorType":  "BadRequestException",
		"error":      err.Error(),
	})
}
