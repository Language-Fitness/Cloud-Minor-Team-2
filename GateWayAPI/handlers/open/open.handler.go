package open

import (
	"example/cloud-api/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	openService *services.OpenService
}

func NewOpenHandler(openService *services.OpenService) *Handler {
	return &Handler{openService: openService}
}

func (h Handler) GetFeedback(context *gin.Context) {
	h.openService.GetFeedback(context)
}
