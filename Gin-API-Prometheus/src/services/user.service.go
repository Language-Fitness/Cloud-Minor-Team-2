package services

import (
	"Gin-API-Prometeus/src/dtos"
	"errors"
)

var mockUsers = []dtos.UserDto{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	// Add more mock users as needed
}

func GetAllUsers() ([]dtos.UserDto, error) {
	if len(mockUsers) == 0 {
		return nil, errors.New("No users found")
	}
	return mockUsers, nil
}

func GetUserByID(id int) (dtos.UserDto, error) {
	for _, user := range mockUsers {
		if user.ID == id {
			return user, nil
		}
	}

	return dtos.UserDto{}, errors.New("User not found")
}
