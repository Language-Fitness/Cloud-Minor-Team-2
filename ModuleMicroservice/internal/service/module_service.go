package service

import (
	"Module/graph/model"
	"Module/internal/auth"
	"Module/internal/repository"
	"Module/internal/validation"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
	"time"
)

// IModuleService GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Module.
type IModuleService interface {
	CreateModule(token string, newModule model.ModuleInput) (*model.Module, error)
	UpdateModule(token string, id string, updateData model.ModuleInput) (*model.Module, error)
	DeleteModule(token string, id string, filter *model.Filter) error
	GetModuleById(token string, id string) (*model.Module, error)
	ListModules(token string, filter *model.Filter, paginate *model.Paginator) ([]*model.ModuleInfo, error)
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

func (m *ModuleService) CreateModule(token string, newModule model.ModuleInput) (*model.Module, error) {
	sub, err := m.Policy.CreateModule(token)
	if err != nil {
		return nil, err
	}

	m.Validator.Validate(newModule.Name, []string{"IsString", "Length:<25"}, "Name")
	m.Validator.Validate(newModule.Description, []string{"IsString", "Length:<50"}, "Description")
	m.Validator.Validate(newModule.Difficulty, []string{"IsInt"}, "Difficulty")
	m.Validator.Validate(newModule.Category, []string{"IsString"}, "Category")
	m.Validator.Validate(newModule.Private, []string{"IsBoolean"}, "Private")
	if newModule.Private {
		m.Validator.Validate(*newModule.Key, []string{"IsString", "Length:<30"}, "Key")
	}

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

func (m *ModuleService) UpdateModule(token string, id string, updateData model.ModuleInput) (*model.Module, error) {
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

func (m *ModuleService) DeleteModule(token string, id string, filter *model.Filter) error {
	//isAdmin, existingModule, err := m.Policy.DeleteModule(token, id)
	//if err != nil {
	//	return err
	//}

	//if !*existingModule.SoftDeleted {
	//	softDelete := true
	//	existingModule.SoftDeleted = &softDelete
	//
	//	err := m.Repo.SoftDeleteModuleByID(id, *existingModule)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}

	//
	//if isAdmin && filter != nil && !*filter.SoftDelete {
	//	err := m.Repo.HardDeleteModuleByID(id)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}

	return errors.New("module could not be deleted")
}

func (m *ModuleService) GetModuleById(token string, id string) (*model.Module, error) {
	existingModule, err := m.Policy.GetModule(token, id)
	if err != nil {
		return nil, err
	}

	return existingModule, nil
}

func (m *ModuleService) ListModules(token string, filter *model.Filter, paginate *model.Paginator) ([]*model.ModuleInfo, error) {
	//isAdmin, err := m.Policy.ListModules(token)
	//if err != nil {
	//	return nil, err
	//}

	filterNameInput := dereferenceArrayIfNeeded(filter.Name.Input)
	fmt.Println(filterNameInput)
	fmt.Println("$" + string(filter.Name.Type))
	fmt.Println(dereferenceIfNeeded(filter.Difficulty))
	fmt.Println(filter.Difficulty)

	m.Validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")
	m.Validator.Validate(dereferenceArrayIfNeeded(filter.Name.Input), []string{"IsNull", "ArrayType:string"}, "Filter Name input")
	m.Validator.Validate(filter.Private, []string{"IsNull", "IsBoolean"}, "Filter Private")
	m.Validator.Validate(paginate.Amount, []string{"IsInt", "Size:>0", "Size:<101"}, "Paginate Amount")
	m.Validator.Validate(paginate.Step, []string{"IsInt", "Size:>0"}, "Paginate Step")

	validationErrors := m.Validator.GetErrors()
	if len(validationErrors) > 0 {
		errorMessage := "Validation errors: " + strings.Join(validationErrors, ", ")
		m.Validator.ClearErrors()
		return nil, errors.New(errorMessage)
	}

	fmt.Println(dereferenceIfNeeded(filter.SoftDelete))

	bsonFilter := bson.D{}
	if m.Policy.HasPermissions(token, "filter_module_SoftDelete") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: dereferenceIfNeeded(filter.SoftDelete)})
	}

	if m.Policy.HasPermissions(token, "filter_module_Name") == true {
		bsonFilter = addFilter(bsonFilter, "name", string(filter.Name.Type), dereferenceArrayIfNeeded(filter.Name.Input))
		fmt.Println("test")
		fmt.Println(bsonFilter)
	}

	if m.Policy.HasPermissions(token, "filter_module_Difficulty") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "difficulty", Value: dereferenceIfNeeded(filter.Difficulty)})
	}

	if m.Policy.HasPermissions(token, "filter_module_Private") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "private", Value: dereferenceIfNeeded(filter.Private)})
	}

	if m.Policy.HasPermissions(token, "filter_module_Category") == true {
		bsonFilter = append(bsonFilter, bson.E{Key: "category", Value: dereferenceIfNeeded(filter.Category)})
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

func dereferenceArrayIfNeeded(value interface{}) []string {
	var newArray []string

	if myArray, ok := value.([]*string); ok {
		for _, pointer := range myArray {
			if pointer == nil {
				continue
			}

			value := *pointer
			newArray = append(newArray, value)
		}
	}

	return newArray
}

func addFilter(bsonFilter []bson.E, key, op string, values []string) []bson.E {
	switch op {
	case "eq": // works
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$in", values}}})
	case "ne": // works
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$nin", values}}})
	case "starts": // works
		var regexPatterns []string
		for _, prefix := range values {
			regexPatterns = append(regexPatterns, "^"+prefix)
		}
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$regex", strings.Join(regexPatterns, "|")}}})
	case "ends": // works
		var regexPatterns []string
		for _, suffix := range values {
			regexPatterns = append(regexPatterns, suffix+"$")
		}
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$regex", strings.Join(regexPatterns, "|")}}})
	}

	return bsonFilter
}

func dereferenceIfNeeded(value interface{}) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		return reflect.ValueOf(value).Elem().Interface()
	}

	return value
}
