package services

import (
	"Gin-API/src/dtos"
	"context"
	"fmt"
	"log"
	"time"

	"Gin-API/proto/pb"
	"google.golang.org/grpc"
)

func GetAllUsers() ([]dtos.UserDto, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Example 1: Get all users
	allUsersReq := &pb.GetAllUsersRequest{}
	allUsersRes, err := client.GetAllUsers(ctx, allUsersReq)
	if err != nil {
		log.Fatalf("GetAllUsers RPC failed: %v", err)
	}

	var users []dtos.UserDto
	fmt.Println("All Users:")
	for _, u := range allUsersRes.Users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.Id, u.Name, u.Email)
		user := dtos.UserDto{
			ID:    int(u.Id),
			Name:  u.Name,
			Email: u.Email,
		}

		// Append the user to the list
		users = append(users, user)
	}
	return users, err
}

func GetUserByID(id int) (dtos.UserDto, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Example 2: Get a user by ID
	userByIDReq := &pb.GetUserByIdRequest{UserId: int32(id)}
	userByIDRes, err := client.GetUserById(ctx, userByIDReq)
	if err != nil {
		log.Fatalf("GetUserById RPC failed: %v", err)
	}

	fmt.Printf("User by ID: ID: %d, Name: %s, Email: %s\n", userByIDRes.Id, userByIDRes.Name, userByIDRes.Email)

	user := dtos.UserDto{
		ID:    int(userByIDRes.Id),
		Name:  userByIDRes.Name,
		Email: userByIDRes.Email,
	}

	return user, err
}
