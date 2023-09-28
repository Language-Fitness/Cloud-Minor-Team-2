package school

import (
	"cloud-api/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	schoolService *services.SchoolService
}

func NewSchoolHandler(schoolService *services.SchoolService) *Handler {
	return &Handler{schoolService: schoolService}
}

// GetAllSchools godoc
// @Summary      List all schools
// @Description  Get all schools
// @Security BearerAuth
// @Tags         schools
// @Accept       json
// @Produce      json
// @Success      200  {array}   response
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users [get]
func (h *Handler) GetAllSchools(c *gin.Context) {
	h.schoolService.GetAllSchools(c)
}

// GetSchoolByID godoc
// @Summary      Get one schools
// @Description  Get schools by ID
// @Security BearerAuth
// @Tags         schools
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  response
// @Failure		 401  {object}	response.ErrorResponseDTO
// @Failure		 403  {object}	response.ErrorResponseDTO
// @Failure		 404  {object}	response.ErrorResponseDTO
// @Failure		 500  {object}	response.ErrorResponseDTO
// @Router       /users/{id} [get]
func (h *Handler) GetSchoolByID(c *gin.Context) {
	h.schoolService.GetSchoolByID(c)
}
