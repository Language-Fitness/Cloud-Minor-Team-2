package graph

import (
	"Module/graph/model"
	"Module/internal/database"
	"Module/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service service.IModuleService
	Modules []*model.Module
}

type AppConfig struct {
	Collection *mongo.Collection
}

func NewResolver() *Resolver {
	collection, _ := database.GetCollection()

	config := &AppConfig{
		Collection: collection,
	}

	return &Resolver{
		Service: service.NewModuleService(config),
		Modules: []*model.Module{},
	}
}
