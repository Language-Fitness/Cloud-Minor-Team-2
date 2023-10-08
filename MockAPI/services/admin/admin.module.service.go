package admin

import (
	"github.com/gin-gonic/gin"
)

type ModuleService struct {
}

func NewAdminModuleService() *ModuleService {
	return &ModuleService{}
}

func (service ModuleService) GetAllModules(c *gin.Context) {
}

func (service ModuleService) GetOneModule(c *gin.Context) {
	//id := c.Param("id")
	//data, err := u.repository.GetOne(id)
}

func (service ModuleService) CreateModules(c *gin.Context) {

}

func (service ModuleService) UpdateModules(c *gin.Context) {

}

func (service ModuleService) DeleteModules(c *gin.Context) {
	//id := c.Param("id")
	//data, err := u.repository.GetOne(id)
}
