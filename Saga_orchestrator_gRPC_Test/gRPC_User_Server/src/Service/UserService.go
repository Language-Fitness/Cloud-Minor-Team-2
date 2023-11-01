package Service

import "gRPC_User_Server/proto/pb"

func GetUserFromDataSource(userID string) *pb.UserResponse {
	if userID == "1" {
		return &pb.UserResponse{
			UserId:   userID,
			Name:     "John",
			LastName: "Doe",
			Age:      30}
	} else if userID == "2" {
		return &pb.UserResponse{
			UserId:   userID,
			Name:     "Jane",
			LastName: "Doe",
			Age:      25}
	} else {
		return nil
	}
}
