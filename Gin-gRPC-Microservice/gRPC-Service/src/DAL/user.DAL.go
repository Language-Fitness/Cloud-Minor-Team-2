package DAL

import (
	"fmt"
	"gRPC-Service/proto/pb"
)

var users = []pb.User{
	{Id: 1, Name: "Alice", Email: "alice@example.com"},
	{Id: 2, Name: "Bob", Email: "bob@example.com"},
	{Id: 3, Name: "Charlie", Email: "charlie@example.com"},
	{Id: 4, Name: "David", Email: "david@example.com"},
}

func GetAllUsers() ([]*pb.User, error) {
	// Create a slice of pointers to pb.User.
	var userPointers []*pb.User

	for _, u := range users {
		// Create a pointer to the user and add it to the slice.
		userPointer := &pb.User{
			Id:    u.Id,
			Name:  u.Name,
			Email: u.Email,
		}
		userPointers = append(userPointers, userPointer)
	}
	return userPointers, nil
}

func GetUserByID(id int32) (*pb.User, error) {
	for _, u := range users {
		if u.Id == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}
