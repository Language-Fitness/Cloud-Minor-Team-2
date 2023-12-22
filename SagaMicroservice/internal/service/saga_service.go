package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"saga/graph/model"
	"saga/internal/auth"
	"saga/internal/validation"
)

type ISagaService interface {
	InitSagaSteps(token string, filter *model.SagaFilter) (*model.SuccessMessage, error)
}

// SagaService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type SagaService struct {
	Validator validation.IValidator
	Policy    auth.IPolicy
}

// NewSagaService GOLANG FACTORY
// Returns a SagaService implementing ISagaService.
func NewSagaService(collection *mongo.Collection) ISagaService {
	return &SagaService{
		Validator: validation.NewValidator(),
		Policy:    auth.NewPolicy(),
	}
}

func (s SagaService) InitSagaSteps(token string, filter *model.SagaFilter) (*model.SuccessMessage, error) {

	sagaObject, err := s.initializeSagaObject(token, filter)
	if err != nil {
		return nil, err
	}

	// Step 2: Find all possible children
	children, err := s.findAllChildren(sagaObject)
	if err != nil {
		return nil, err
	}

	// Step 3: Loop through children and find those, if any
	for _, child := range children {
		// Your logic for finding children
		fmt.Println(child)
	}

	// Step 4: Loop through everything starting with the bottom children
	bottomChildren, err := s.findBottomChildren(sagaObject)
	if err != nil {
		return nil, err
	}

	// Step 5: Start soft deleting items and change object_status to Deleted if success
	if err := s.softDeleteItems(bottomChildren); err != nil {
		// Handle rollback logic or return an error
		return nil, err
	}

	// Step 6: Loop through items to check if all object_status are Deleted
	if !s.areAllItemsDeleted(bottomChildren) {
		// Step 7: If not everything is deleted, reloop steps 4 and 5 but undelete every item
		if err := s.undeleteItems(bottomChildren); err != nil {
			// Handle rollback logic or return an error
			return nil, err
		}
	}

	// Step 8: Save the object and return success message
	if err := s.saveSagaObject(sagaObject); err != nil {
		// Handle rollback logic or return an error
		return nil, err
	}

	successMessage := &model.SuccessMessage{
		ID:         "1",
		Text:       "Operation successful",
		Status:     model.SagaObjectStatusExist,
		ObjectID:   "123",
		ObjectType: model.SagaObjectTypesModule,
	}

	return successMessage, nil
}

func (s SagaService) initializeSagaObject(token string, filter *model.SagaFilter) (model.SagaObject, error) {
	//func (s SagaService) initializeSagaObject(token string, filter model.SagaFilter) (model.SagaObject, error) {
	// Step 1 logic here
	// Use s.Validator and s.Policy interfaces as needed
	// Example: return an instance of YourSagaObjectType

	sagaObject := model.SagaObject{
		ID: "1",
	}

	return sagaObject, nil
}

func (s SagaService) findAllChildren(sagaObject model.SagaObject) ([]model.SagaObject, error) {
	// Step 2 logic here
	// Example: return a slice of ChildType
	return nil, nil
}

func (s SagaService) findBottomChildren(sagaObject model.SagaObject) ([]model.SagaObject, error) {
	// Step 4 logic here
	// Example: return a slice of BottomChildType
	return nil, nil
}

func (s SagaService) softDeleteItems(items []model.SagaObject) error {
	// Step 5 logic here
	// Example: soft delete items and update object_status
	return nil
}

func (s SagaService) areAllItemsDeleted(items []model.SagaObject) bool {
	// Step 6 logic here
	// Example: check if all items have object_status set to Deleted
	return true
}

func (s SagaService) undeleteItems(items []model.SagaObject) error {
	// Step 8 logic here
	// Example: undelete items and update object_status
	return nil
}

func (s SagaService) saveSagaObject(sagaObject model.SagaObject) error {
	// Step 8 logic here
	// Example: save saga object to MongoDB
	return nil
}


