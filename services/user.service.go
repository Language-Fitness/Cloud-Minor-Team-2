package services

import (
	"example/cloud-api/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
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
