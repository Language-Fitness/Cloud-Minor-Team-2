package admin

import (
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) GetAllUsers(c *gin.Context) {
}

func (u UserService) GetOneUser(c *gin.Context) {
	//id := c.Param("id")
	//data, err := u.repository.GetOne(id)
}

func (u UserService) CreateUser(c *gin.Context) {

}

func (u UserService) UpdateUser(c *gin.Context) {

}

func (u UserService) DeleteUser(c *gin.Context) {
	//id := c.Param("id")
	//data, err := u.repository.GetOne(id)
}
