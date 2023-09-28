package services

import (
	"github.com/gin-gonic/gin"
)

type AdminModuleService struct {
}

func NewAdminModuleService() *AdminModuleService {
	return &AdminModuleService{}
}

func (service AdminModuleService) GetAllModules(c *gin.Context) {
}

func (service AdminModuleService) GetOneModule(c *gin.Context) {
	//id := c.Param("id")
	//data, err := u.repository.GetOne(id)
}

func (service AdminModuleService) CreateModules(c *gin.Context) {

}

func (service AdminModuleService) UpdateModules(c *gin.Context) {

}

func (service AdminModuleService) DeleteModules(c *gin.Context) {
	//id := c.Param("id")
	//data, err := u.repository.GetOne(id)
}
