package school

import (
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	moduleService *services.ModuleService
}

func (h Handler) GetAllModules(context *gin.Context) {
	h.moduleService.GetAllModules(context)
}

func (h Handler) GetResultForExercise(context *gin.Context) {
	h.moduleService.GetResultForExercise(context)
}

func (h Handler) CreateResultForExercise(context *gin.Context) {
	h.moduleService.CreateResultForExercise(context)
}

func (h Handler) GetExerciseById(context *gin.Context) {
	h.moduleService.GetExerciseById(context)
}

func (h Handler) GetAllExercises(context *gin.Context) {
	h.moduleService.GetAllExercises(context)
}

func (h Handler) GetAllCourses(context *gin.Context) {
	h.moduleService.GetAllCourses(context)
}

func (h Handler) GetAllCoursesById(context *gin.Context) {
	h.moduleService.GetAllCoursesById(context)
}

func (h Handler) GetModuleByID(context *gin.Context) {
	h.moduleService.GetModuleByID(context)
}

func NewModuleHandler(moduleService *services.ModuleService) *Handler {
	return &Handler{moduleService: moduleService}
}
