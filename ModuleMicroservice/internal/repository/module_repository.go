package repository

import (
	"Module/graph/model"
	"Module/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModuleRepository struct {
	modules []*model.Module
	client  *mongo.Client
}

func NewModuleRepository() *ModuleRepository {
	client, _ := database.GetDBClient()

	return &ModuleRepository{
		modules: []*model.Module{},
		client:  client,
	}
}

func (r *ModuleRepository) CreateModule(newModule *model.Module) error {

	println(newModule.Name)

	r.modules = append(r.modules, newModule)
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
