package repository

import (
	"Module/graph/model"
	"Module/internal/database"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ModuleRepository struct {
	modules    []*model.Module
	collection *mongo.Collection
}

func NewModuleRepository() *ModuleRepository {
	collection, _ := database.GetCollection()

	return &ModuleRepository{
		modules:    []*model.Module{},
		collection: collection,
	}
}

func (r *ModuleRepository) CreateModule(newModule *model.Module) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	println("test")

	_, err := r.collection.InsertOne(ctx, newModule)
	if err != nil {
		return err
	}
	return nil
}

func (r *ModuleRepository) UpdateModule(updatedModule *model.Module) error {
	for i, module := range r.modules {
		if module.ID == updatedModule.ID {
			r.modules[i] = updatedModule
			return nil
		}
	}
	return nil
}

func (r *ModuleRepository) DeleteModuleByID(id string) error {
	for i, module := range r.modules {
		if module.ID == id {
			r.modules = append(r.modules[:i], r.modules[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *ModuleRepository) GetModuleByID(id string) (*model.Module, error) {
	for _, module := range r.modules {
		if module.ID == id {
			return module, nil
		}
	}
	return nil, nil
}

func (r *ModuleRepository) ListModules() ([]*model.Module, error) {
	return r.modules, nil
}
