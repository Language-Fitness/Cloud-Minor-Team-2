package service

import (
	"Module/graph/model"
	"Module/internal/auth"
	"Module/internal/helper"
	"Module/internal/repository"
	"Module/internal/validation"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

// IModuleService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Module.
type IModuleService interface {
	CreateModule(token string, newModule model.ModuleInputCreate) (*model.Module, error)
	UpdateModule(token string, id string, updateData model.ModuleInputUpdate) (*model.Module, error)
	DeleteModule(token string, id string, filter *model.ModuleFilter) error
	GetModuleById(token string, id string) (*model.Module, error)
	ListModules(token string, filter *model.ModuleFilter, paginate *model.Paginator) ([]*model.ModuleInfo, error)
}

// ModuleService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type ModuleService struct {
	Validator validation.IValidator
	Repo      repository.IModuleRepository
	Policy    auth.IPolicy
}

// NewModuleService GOLANG FACTORY
// Returns a ModuleService implementing IModuleService.
func NewModuleService(collection *mongo.Collection) IModuleService {
	return &ModuleService{
		Validator: validation.NewValidator(),
		Repo:      repository.NewModuleRepository(collection),
		Policy:    auth.NewPolicy(collection),
	}
}

func (m *ModuleService) CreateModule(token string, newModule model.ModuleInputCreate) (*model.Module, error) {
	sub, err := m.Policy.CreateModule(token)
	if err != nil {
		return nil, err
	}
	m.Validator.Validate(newModule.SchoolID, []string{"IsUUID"}, "Filter School")
	m.Validator.Validate(newModule.Name, []string{"IsString", "Length:<25"}, "Name")
	m.Validator.Validate(newModule.Description, []string{"IsString", "Length:<50"}, "Description")
	m.Validator.Validate(newModule.Difficulty, []string{"IsInt"}, "Difficulty")
	m.Validator.Validate(newModule.Category, []string{"IsString"}, "Category")
	m.Validator.Validate(newModule.Private, []string{"IsBoolean"}, "Private")
	if newModule.Private {
		m.Validator.Validate(*newModule.Key, []string{"IsString", "Length:<30"}, "Key")
	}

	//@TODO check if school exist

	validationErrors := m.Validator.GetErrors()

	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	softDeleted := false

	moduleToInsert := &model.Module{
		ID:          uuid.New().String(),
		SchoolID:    newModule.SchoolID,
		Name:        newModule.Name,
		Description: newModule.Description,
		Difficulty:  newModule.Difficulty,
		Category:    newModule.Category,
		MadeBy:      sub,
		Private:     newModule.Private,
		CreatedAt:   &timestamp,
		SoftDeleted: &softDeleted,
	}

	if newModule.Private {
		moduleToInsert.Key = newModule.Key
	}

	result, err := m.Repo.CreateModule(moduleToInsert)
	if err != nil {
		return nil, err
	}

	m.Validator.ClearErrors()
	return result, nil
}

func (m *ModuleService) UpdateModule(token string, id string, updateData model.ModuleInputUpdate) (*model.Module, error) {
	existingModule, err := m.Policy.UpdateModule(token, id)
	if err != nil {
		return nil, err
	}

	m.Validator.Validate(updateData.Name, []string{"IsString", "Length:<25"}, "Name")
	m.Validator.Validate(updateData.Description, []string{"IsString", "Length:<50"}, "Description")
	m.Validator.Validate(updateData.Difficulty, []string{"IsInt"}, "Difficulty")
	m.Validator.Validate(updateData.Category, []string{"IsString"}, "Category")
	m.Validator.Validate(updateData.Private, []string{"IsBoolean"}, "Private")
	if updateData.Private {
		m.Validator.Validate(*updateData.Key, []string{"IsString", "Length:<30"}, "Key")
	}

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	timestamp := time.Now().String()
	newModule := model.Module{
		ID:          existingModule.ID,
		SchoolID:    existingModule.SchoolID,
		Name:        updateData.Name,
		Description: updateData.Description,
		Difficulty:  updateData.Difficulty,
		Category:    updateData.Category,
		MadeBy:      existingModule.MadeBy,
		Private:     updateData.Private,
		CreatedAt:   existingModule.CreatedAt,
		UpdatedAt:   &timestamp,
		SoftDeleted: existingModule.SoftDeleted,
	}

	if updateData.Private {
		newModule.Key = updateData.Key
	}

	result, err := m.Repo.UpdateModule(id, newModule)
	if err != nil {
		return nil, err
	}

	m.Validator.ClearErrors()
	return result, nil
}

func (m *ModuleService) DeleteModule(token string, id string, filter *model.ModuleFilter) error {
	existingModule, err := m.Policy.DeleteModule(token, id)
	if err != nil {
		return err
	}

	if !*existingModule.SoftDeleted {
		softDelete := true
		existingModule.SoftDeleted = &softDelete

		err := m.Repo.SoftDeleteModuleByID(id, *existingModule)
		if err != nil {
			return err
		}
		return nil
	}

	if m.Policy.HasPermissions(token, "delete_module_all") && filter != nil && !*filter.SoftDelete {
		err := m.Repo.HardDeleteModuleByID(id)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("module could not be deleted")
}

func (m *ModuleService) GetModuleById(token string, id string) (*model.Module, error) {
	existingModule, err := m.Policy.GetModule(token, id)
	if err != nil {
		return nil, err
	}

	return existingModule, nil
}

func (m *ModuleService) ListModules(token string, filter *model.ModuleFilter, paginate *model.Paginator) ([]*model.ModuleInfo, error) {
	//err := m.Policy.ListModules(token)
	//if err != nil {
	//	return nil, err
	//}

	m.Validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")
	if helper.IsNil(filter.Name) == false {
		m.Validator.Validate(helper.DereferenceArrayIfNeeded(filter.Name.Input), []string{"IsNull", "ArrayType:string"}, "Filter Name input")
	}
	m.Validator.Validate(filter.Private, []string{"IsNull", "IsBoolean"}, "Filter Private")
	m.Validator.Validate(filter.MadeBy, []string{"IsNull", "IsUUID"}, "Filter MadeBy")
	m.Validator.Validate(filter.SchoolID, []string{"IsNull", "IsUUID"}, "Filter School")
	m.Validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	m.Validator.Validate(paginate.Step, []string{"IsInt", "Size:>=0"}, "Paginate Step")

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	fmt.Println(helper.DereferenceIfNeeded(filter.SoftDelete))

	bsonFilter := bson.D{}
	if m.Policy.HasPermissions(token, "filter_module_softDelete") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: helper.DereferenceIfNeeded(filter.SoftDelete)})
	}

	if m.Policy.HasPermissions(token, "filter_module_school_id") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "schoolid", Value: helper.DereferenceIfNeeded(filter.SchoolID)})
	}

	if m.Policy.HasPermissions(token, "filter_module_made_by") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "madeby", Value: helper.DereferenceIfNeeded(filter.MadeBy)})
	}

	if m.Policy.HasPermissions(token, "filter_module_name") == true && helper.IsNil(filter.Name) == false {
		bsonFilter = helper.AddFilter(bsonFilter, "name", string(filter.Name.Type), helper.DereferenceArrayIfNeeded(filter.Name.Input))
	}

	if m.Policy.HasPermissions(token, "filter_module_difficulty") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "difficulty", Value: helper.DereferenceIfNeeded(filter.Difficulty)})
	}

	if m.Policy.HasPermissions(token, "filter_module_private") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "private", Value: helper.DereferenceIfNeeded(filter.Private)})
	}

	if m.Policy.HasPermissions(token, "filter_module_category") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "category", Value: helper.DereferenceIfNeeded(filter.Category)})
	}

	fmt.Println(bsonFilter)

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	modules, err := m.Repo.ListModules(bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}

	return modules, nil
}
