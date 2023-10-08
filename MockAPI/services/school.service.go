package services

import (
	"github.com/gin-gonic/gin"
)

type SchoolService struct {
}

func NewSchoolService() *SchoolService {
	return &SchoolService{}
}

func (service SchoolService) GetAllSchools(c *gin.Context) {
}

func (service SchoolService) GetSchoolByID(c *gin.Context) {

}
