package repository

import (
	"Module/graph/model"
)

type ModuleRepository struct {
	modules []*model.Module
}

func NewModuleRepository() *ModuleRepository {
	return &ModuleRepository{
		modules: []*model.Module{},
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
