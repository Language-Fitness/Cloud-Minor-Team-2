package services

import (
	"github.com/gin-gonic/gin"
)

type ModuleService struct {
}

func NewModuleService() *ModuleService {
	return &ModuleService{}
}

func (service ModuleService) GetAllModules(c *gin.Context) {
}

func (service ModuleService) GetResultForExercise(context *gin.Context) {

}

func (service ModuleService) CreateResultForExercise(context *gin.Context) {

}

func (service ModuleService) GetExerciseById(context *gin.Context) {

}

func (service ModuleService) GetAllExercises(context *gin.Context) {

}

func (service ModuleService) GetAllCourses(context *gin.Context) {

}

func (service ModuleService) GetAllCoursesById(context *gin.Context) {

}

func (service ModuleService) GetModuleByID(context *gin.Context) {

}
