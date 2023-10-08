package services

import "github.com/gin-gonic/gin"

type OpenService struct {
}

func (s OpenService) GetFeedback(context *gin.Context) {

}

func NewOpenService() *OpenService {
	return &OpenService{}
}
