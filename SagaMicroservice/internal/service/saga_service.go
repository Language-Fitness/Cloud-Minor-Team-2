package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"saga/graph/model"
	"saga/internal/auth"
	"saga/internal/validation"
	"saga/proto/pb"
	"time"
)

type ISagaService interface {
	InitSagaSteps(token string, filter *model.SagaFilter) (*model.SuccessMessage, error)
	initializeSagaObject(token string, filter *model.SagaFilter) (model.SagaObject, error)
	findAllChildren(sagaObject model.SagaObject) ([]model.SagaObject, error)
	findBottomChildren(sagaObject model.SagaObject) (model.SagaObject, error)
	softDeleteItems(items []model.SagaObject) error
	areAllItemsDeleted(items []model.SagaObject) bool
	undeleteItems(items []model.SagaObject) error
	saveSagaObject(sagaObject model.SagaObject) error
}

// SagaService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type SagaService struct {
	Validator  validation.IValidator
	Policy     auth.IPolicy
	grpcClient pb.GRPCSagaServiceClient
}

// NewSagaService GOLANG FACTORY
// Returns a SagaService implementing ISagaService.
func NewSagaService(grpcClient pb.GRPCSagaServiceClient, collection *mongo.Collection) ISagaService {
	return &SagaService{
		Validator:  validation.NewValidator(),
		Policy:     auth.NewPolicy(),
		grpcClient: grpcClient,
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
	sagaObject, err = s.findBottomChildren(sagaObject)
	if err != nil {
		return nil, err
	}

	//// Step 5: Start soft deleting items and change object_status to Deleted if success
	//if err := s.softDeleteItems(sagaObject); err != nil {
	//	// Handle rollback logic or return an error
	//	return nil, err
	//}
	//
	//// Step 6: Loop through items to check if all object_status are Deleted
	//if !s.areAllItemsDeleted(sagaObject) {
	//	// Step 7: If not everything is deleted, reloop steps 4 and 5 but undelete every item
	//	if err := s.undeleteItems(sagaObject); err != nil {
	//		// Handle rollback logic or return an error
	//		return nil, err
	//	}
	//}

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
		ID:           "1",
		ObjectID:     filter.ObjectID,
		ObjectType:   filter.ObjectType,
		CreatedAt:    time.Now().Format("HH:MM:SS"),
		ObjectStatus: model.SagaObjectStatusExist,
		ActionDoneBy: "1",
	}

	return sagaObject, nil
}

func (s SagaService) findAllChildren(sagaObject model.SagaObject) ([]model.SagaObject, error) {
	// Step 2 logic here
	// Example: return a slice of ChildType

	// Get main object and ID
	// Send GRPC request to search children
	// So if School   	    -> search User
	// So if type User 	    -> search for Module
	// So if type User 	    -> search for Results
	// So if type Module    -> search for Classes
	// So if type Classes   -> search for Exercises
	// So if type Exercises -> search for Results

	// Then push the results in saga object from GRPC request
	// Save SagaObject in DB
	// Return SagaObject

	// Note this function can be reused to find children of the children
	// for example if we first get all classes belonging to module
	// We should be able to recall this function from a for loop to get all the
	// children of the classes with exercises

	return nil, nil
}

func (s SagaService) findBottomChildren(sagaObject model.SagaObject) (model.SagaObject, error) {
	// Step 4 logic here
	// Example: return a slice of BottomChildType
	return sagaObject, nil
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
