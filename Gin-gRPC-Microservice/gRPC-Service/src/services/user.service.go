package services

import (
	"gRPC-Service/proto/pb"
	"gRPC-Service/src/DAL"
)

func GetAllUsers() ([]*pb.User, error) {
	return DAL.GetAllUsers()
}

func GetUserByID(id int) (*pb.User, error) {
	return DAL.GetUserByID(int32(id))
}
