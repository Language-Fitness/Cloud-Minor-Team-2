package service

import (
	"Module/graph/model"
	"Module/internal/auth"
	"Module/internal/helper"
	"Module/internal/repository"
	"Module/internal/validation"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"net/http"
	"os"
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

	//subName, err := getUserInfo(token, sub)
	//if err != nil {
	//	return nil, fmt.Errorf("no user information could be extracted")
	//}
	subName := "Merlin"

	m.Validator.Validate(newModule.SchoolID, []string{"IsUUID"}, "Filter School")
	m.Validator.Validate(newModule.Name, []string{"IsString", "Length:<25"}, "Name")
	m.Validator.Validate(newModule.Description, []string{"IsString", "Length:<50"}, "Description")
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
		MadeByName:  subName,
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
		MadeByName:  existingModule.MadeByName,
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

		err := m.Repo.DeleteModuleByID(id, *existingModule)
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

	fmt.Println(filter)
	fmt.Println("init")

	if helper.IsNil(filter.Name) == false {
		m.Validator.Validate(helper.DereferenceArrayIfNeeded(filter.Name.Input), []string{"IsNull", "ArrayType:string"}, "Filter Name input")
	}
	//if helper.IsNil(filter.MadeByName) == false {
	//	m.Validator.Validate(helper.DereferenceArrayIfNeeded(filter.MadeByName.Input), []string{"IsNull", "ArrayType:string"}, "Filter Made by Name input")
	//}

	m.Validator.Validate(filter.SoftDelete, []string{"IsNull", "IsBoolean"}, "Filter softDelete")
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
	fmt.Println("init2")
	fmt.Println(helper.DereferenceIfNeeded(filter.SoftDelete))

	bsonFilter := bson.D{}
	if m.Policy.HasPermissions(token, "filter_module_softDelete") == true && helper.IsNil(filter.SoftDelete) == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: helper.DereferenceIfNeeded(filter.SoftDelete)})
	} else {
		bsonFilter = append(bsonFilter, bson.E{Key: "softdeleted", Value: false})
	}

	if m.Policy.HasPermissions(token, "filter_module_school_id") == true && helper.IsNil(filter.SchoolID) == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "schoolid", Value: helper.DereferenceIfNeeded(filter.SchoolID)})
	}

	fmt.Println(m.Policy.HasPermissions(token, "filter_school_made_by"))
	fmt.Println("filter_module_made_by")
	if m.Policy.HasPermissions(token, "filter_module_made_by") == true && helper.IsNil(filter.MadeBy) == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "madeby", Value: helper.DereferenceIfNeeded(filter.MadeBy)})
	}

	//if m.Policy.HasPermissions(token, "filter_module_made_by_name") == true && helper.IsNil(filter.MadeByName) == false {
	//	bsonFilter = helper.AddFilter(bsonFilter, "madebyname", string(filter.MadeByName.Type), helper.DereferenceArrayIfNeeded(filter.MadeByName.Input))
	//}

	if m.Policy.HasPermissions(token, "filter_module_name") == true && helper.IsNil(filter.Name) == false {
		bsonFilter = helper.AddFilter(bsonFilter, "name", string(filter.Name.Type), helper.DereferenceArrayIfNeeded(filter.Name.Input))
	}

	if m.Policy.HasPermissions(token, "filter_module_difficulty") == true && helper.IsNil(filter.Difficulty) == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "difficulty", Value: helper.DereferenceIfNeeded(filter.Difficulty)})
	}

	if m.Policy.HasPermissions(token, "filter_module_private") == true && helper.IsNil(filter.Private) == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "private", Value: helper.DereferenceIfNeeded(filter.Private)})
	}

	if m.Policy.HasPermissions(token, "filter_module_category") == true && helper.IsNil(filter.Category) == false {
		bsonFilter = append(bsonFilter, bson.E{Key: "category", Value: helper.DereferenceIfNeeded(filter.Category)})
	}

	fmt.Println("init3")
	fmt.Println(bsonFilter)

	paginateOptions := options.Find().
		SetSkip(int64(paginate.Step)).
		SetLimit(int64(paginate.Amount))

	fmt.Println("init4")
	modules, err := m.Repo.ListModules(bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}

	return modules, nil
}

func getUserInfo(token string, userId string) (string, error) {
	baseUrl := os.Getenv("KEYCLOAK_HOST")
	path := "admin/realms/cloud-project/users/"
	fullUrl := baseUrl + path + userId

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	fullName, err := extractNamesFromJson(result)
	if err != nil {
		return "", err
	}

	return fullName, nil
}

func extractNamesFromJson(result map[string]interface{}) (string, error) {
	firstName, ok := result["firstName"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract firstName")
	}

	lastName, ok := result["lastName"].(string)
	if !ok {
		return "", fmt.Errorf("failed to extract lastName")
	}

	fullName := firstName + " " + lastName
	return fullName, nil
}
