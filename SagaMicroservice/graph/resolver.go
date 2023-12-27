package graph

import (
	"saga/graph/model"
	"saga/internal/auth"
	"saga/internal/database"
	"saga/internal/service"
	"saga/proto/pb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service     service.ISagaService
	SagaObjects []*model.SagaObject
	Policy      auth.IPolicy
}

func NewResolver(grpcClient pb.GRPCSagaServiceClient) *Resolver {
	collection, _ := database.GetCollection()

	return &Resolver{
		Service:     service.NewSagaService(grpcClient, collection),
		SagaObjects: []*model.SagaObject{},
		Policy:      auth.NewPolicy(),
	}
}
