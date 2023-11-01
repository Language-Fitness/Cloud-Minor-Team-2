package Server

import (
	"context"
	"gRPC_Address_Server/proto/pb"
	"gRPC_Address_Server/src/Service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddressServiceServer struct {
	pb.UnimplementedGRPC_Address_ServerServer
}

func (s *AddressServiceServer) GetUserAddress(ctx context.Context, req *pb.GetUserAddressRequest) (*pb.AddressResponse, error) {
	// Replace with your logic to fetch user address from a data source
	address := Service.GetAddressFromDataSource(req.UserId)

	if address == "" {
		return nil, status.Errorf(codes.NotFound, "Address not found")
	}

	return &pb.AddressResponse{
		UserId:  req.UserId,
		Address: address,
	}, nil
}
