package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"saga/graph/model"
	"saga/internal/auth"
	"saga/internal/repository"
	"saga/internal/validation"
	"saga/proto/pb"
	"time"
)

type ISagaService interface {
	InitSagaSteps(token string, filter *model.SagaFilter) (*model.SuccessMessage, error)
	initializeSagaObject(token string, filter *model.SagaFilter) (*model.SagaObject, error)
	findAllChildren(token string, sagaObject *model.SagaObject) ([]model.SagaObject, error)
	findBottomChildren(sagaObject *model.SagaObject) (*model.SagaObject, error)
	softDeleteItems(items []model.SagaObject) error
	areAllItemsDeleted(items []model.SagaObject) bool
	undeleteItems(items []model.SagaObject) error
}

// SagaService GOLANG STRUCT
// Contains two interfaces for a Validator and a Repo.
type SagaService struct {
	Validator validation.IValidator
	Policy    auth.IPolicy
	Repo      repository.ISagaObjectRepository
}

// NewSagaService GOLANG FACTORY
// Returns a SagaService implementing ISagaService.
func NewSagaService(collection *mongo.Collection) ISagaService {
	return &SagaService{
		Validator: validation.NewValidator(),
		Policy:    auth.NewPolicy(),
		Repo:      repository.NewSagaObjectRepository(collection),
	}
}

func (s SagaService) InitSagaSteps(token string, filter *model.SagaFilter) (*model.SuccessMessage, error) {
	// Step 1: check if saga object exist and if it does then create it
	token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJIaUpNcWZhTGFWQXBiME5JTEpweTlacmdtRzBERElIaWpVZklVWjM2NXJvIn0.eyJleHAiOjE3MDM5NDc5NzMsImlhdCI6MTcwMzk0NzY3MywianRpIjoiNjk1ZWNlNTktZDYwOS00MjExLWFiMjQtNjA2ODMwMTFkMTQ0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjZiMDNiYTVkLTVkMGUtNGRkOC05ZjdmLTkyOGU3NWVhOGVjYSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIzYjYxNGNmNy00NmVjLTQ5NDEtOWU3Zi0wODkzZGRiODA3NmUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZ2V0X2NsYXNzZXNfYWxsIiwidXBkYXRlX3NjaG9vbCIsImZpbHRlcl9zY2hvb2xfbWFkZV9ieSIsImZpbHRlcl9zY2hvb2xfbmFtZSIsImZpbHRlcl9tb2R1bGVfY2F0ZWdvcnkiLCJmaWx0ZXJfY2xhc3NfbWFkZV9ieSIsImZpbHRlcl9tb2R1bGVfc29mdERlbGV0ZSIsImdldF9leGVyY2lzZXMiLCJnZXRfY2xhc3NlcyIsImRlbGV0ZV9tb2R1bGUiLCJkZWxldGVfZXhlcmNpc2UiLCJnZXRfc2Nob29scyIsInVwZGF0ZV9leGVyY2lzZSIsImdldF9leGVyY2lzZSIsImRlbGV0ZV9tb2R1bGVfYWxsIiwiY3JlYXRlX2V4ZXJjaXNlIiwiZ2V0X3NjaG9vbCIsImRlbGV0ZV9leGVyY2lzZV9hbGwiLCJmaWx0ZXJfc2Nob29sX2xvY2F0aW9uIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJmaWx0ZXJfbW9kdWxlX2RpZmZpY3VsdHkiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwidXBkYXRlX2V4ZXJjaXNlX2FsbCIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX3NvZnREZWxldGUiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImdldF9tb2R1bGVzX2FsbCIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJmaWx0ZXJfbW9kdWxlX3NjaG9vbF9pZCIsImZpbHRlcl9tb2R1bGVfbWFkZV9ieSIsImZpbHRlcl9jbGFzc19uYW1lIiwidXBkYXRlX2NsYXNzX2FsbCIsImZpbHRlcl9tb2R1bGVfbmFtZSIsInVwZGF0ZV9tb2R1bGUiLCJnZXRfY2xhc3MiLCJkZWxldGVfc2Nob29sX2FsbCIsImZpbHRlcl9tb2R1bGVfcHJpdmF0ZSIsInVwZGF0ZV9jbGFzcyIsImdldF9zY2hvb2xzX2FsbCIsImZpbHRlcl9jbGFzc19zb2Z0RGVsZXRlIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiM2I2MTRjZjctNDZlYy00OTQxLTllN2YtMDg5M2RkYjgwNzZlIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.ufqx_RD2A5abIdqEvf79lwl9bsQjZnIga54v82OdEKDpH47IF4yfnKIwl5f4sCpVeyaPl_ihPpYLtBwuD2ZLD-O-u6zRnFyVm3sXuAeN2CC3FOEWZtxr0gxECySaW7k3Oj7AWZimn_yxJfxyElRuNhlg4811gFJ1bZgGkl_3vJvg_61FEIBQB74vQA51jx27Y2-kSxdSMxXAkgWVNYjFtjgDyzeGZUHibqw8uLX4NYASprW4lGDVu-A3S_Vj3dJvJJON6Oe_8-IS-LH2Vw6olJNjEonxm9x5HJAWwUcn_Md4ShUB3u-k9jT1MAFkke1p4h5wuRnTa5mY3yjaF8LiRw"

	sagaObject, err := s.initializeSagaObject(token, filter)
	if err != nil {
		return nil, err
	}

	// Step 2: Find all possible children
	children, err := s.findAllChildren(token, sagaObject)
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

	successMessage := &model.SuccessMessage{
		ID:         "1",
		Text:       "Operation successful",
		Status:     model.SagaObjectStatusExist,
		ObjectID:   "123",
		ObjectType: model.SagaObjectTypesModule,
	}

	return successMessage, nil
}

func (s SagaService) initializeSagaObject(token string, filter *model.SagaFilter) (*model.SagaObject, error) {
	client, conn1, err := createGRPCClient(filter.ObjectType)
	if err != nil {
		fmt.Println(err)
	}

	defer conn1.Close()

	request := pb.ObjectRequest{
		BearerToken:  token,
		ObjectId:     filter.ObjectID,
		ObjectType:   convertToPBObjectType(filter.ObjectType),
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	response, err := client.FindSagaObject(context.Background(), &request)
	if err != nil {
		//log.Fatalf("failed to call FindObject RPC: %v", err)
		log.Printf("failed to call FindObject RPC: %v", err)
	}

	fmt.Println("\nResponse 1:", response.ObjectId)
	fmt.Println("\nResponse 2:", response.ObjectStatus)
	fmt.Println("\nResponse 3:", response.ObjectType)

	sagaObject := model.SagaObject{
		ID:           uuid.New().String(),
		ObjectID:     response.ObjectId,
		ObjectType:   convertToModelObjectType(response.ObjectType),
		CreatedAt:    time.Now().Format("HH:MM:SS"),
		ObjectStatus: convertToModelObjectStatus(response.ObjectStatus),
		ActionDoneBy: "1", //@TODO get this from the bearer token sub
	}

	object, err := s.Repo.CreateSagaObject(&sagaObject)
	if err != nil {
		return nil, err
	}

	fmt.Println(object)

	return object, nil
}

func (s SagaService) findAllChildren(token string, sagaObject *model.SagaObject) ([]model.SagaObject, error) {
	fmt.Println("test4")
	//client, conn2, err := createGRPCClient(model.SagaObjectTypesModule)
	fmt.Println(sagaObject.ObjectType)
	fmt.Println(getChildTypeByType(sagaObject.ObjectType))
	client, conn2, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	fmt.Println("test5")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("test")

	defer conn2.Close()

	request := pb.ObjectRequest{
		BearerToken:  token,
		ObjectId:     sagaObject.ObjectID,
		ObjectType:   pb.SagaObjectType_CLASS,
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	fmt.Println("test2")

	fmt.Println("Calling FindObject RPC...")
	response, err := client.FindSagaObjectChildren(context.Background(), &request)
	if err != nil {
		//log.Fatalf("failed to call FindObject RPC: %v", err)
		log.Printf("failed to call FindObject RPC: %v", err)
	}

	fmt.Println("test3")

	fmt.Println("\nResponse 1:", response.Objects[0].ObjectId)
	fmt.Println("\nResponse 2:", response.Objects[0].ObjectStatus)
	fmt.Println("\nResponse 3:", response.Objects[0].ObjectType)

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

func (s SagaService) findBottomChildren(sagaObject *model.SagaObject) (*model.SagaObject, error) {
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

func getHostByType(sagaType model.SagaObjectTypes) string {
	switch sagaType {
	case model.SagaObjectTypesModule:
		return os.Getenv("MODULE_MS_GRPC_HOST")
	case model.SagaObjectTypesClass:
		return os.Getenv("CLASS_MS_GRPC_HOST")
	case model.SagaObjectTypesExercise:
		return os.Getenv("EXERCISE_MS_GRPC_HOST")
	case model.SagaObjectTypesSchool:
		return os.Getenv("SCHOOL_MS_GRPC_HOST")
	case model.SagaObjectTypesResult:
		return os.Getenv("EXERCISE_MS_GRPC_HOST")
	default:
		return os.Getenv("MODULE_MS_GRPC_HOST")
	}
}

func getChildTypeByType(sagaType model.SagaObjectTypes) model.SagaObjectTypes {
	switch sagaType {
	case model.SagaObjectTypesSchool:
		return model.SagaObjectTypesModule
	case model.SagaObjectTypesModule:
		return model.SagaObjectTypesClass
	case model.SagaObjectTypesClass:
		return model.SagaObjectTypesExercise
	case model.SagaObjectTypesExercise:
		return model.SagaObjectTypesResult
	default:
		return model.SagaObjectTypesResult
	}
}

func convertToPBObjectType(objectType model.SagaObjectTypes) pb.SagaObjectType {
	switch objectType {
	case model.SagaObjectTypesSchool:
		return pb.SagaObjectType_SCHOOL
	case model.SagaObjectTypesModule:
		return pb.SagaObjectType_MODULE
	case model.SagaObjectTypesClass:
		return pb.SagaObjectType_CLASS
	case model.SagaObjectTypesExercise:
		return pb.SagaObjectType_EXERCISE
	default:
		return pb.SagaObjectType_RESULT
	}
}

func convertToModelObjectType(objectType pb.SagaObjectType) model.SagaObjectTypes {
	switch objectType {
	case pb.SagaObjectType_SCHOOL:
		return model.SagaObjectTypesSchool
	case pb.SagaObjectType_MODULE:
		return model.SagaObjectTypesModule
	case pb.SagaObjectType_CLASS:
		return model.SagaObjectTypesClass
	case pb.SagaObjectType_EXERCISE:
		return model.SagaObjectTypesExercise
	default:
		return model.SagaObjectTypesResult
	}
}

func convertToModelObjectStatus(objectType pb.SagaObjectStatus) model.SagaObjectStatus {
	switch objectType {
	case pb.SagaObjectStatus_DELETED:
		return model.SagaObjectStatusDeleted
	default:
		return model.SagaObjectStatusExist
	}
}

func createGRPCClient(sagaType model.SagaObjectTypes) (pb.GRPCSagaServiceClient, *grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithReturnConnectionError(), // Add the WithReturnConnectionError option
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	fmt.Println(getHostByType(sagaType))

	conn, err := grpc.DialContext(context.Background(), getHostByType(sagaType), opts...)
	fmt.Println("oops")
	if err != nil {
		fmt.Printf("failed to dial gRPC server: %v\n", err)
		log.Printf("failed to dial gRPC server: %v", err)
		return nil, nil, err
	}

	fmt.Println("Creating gRPC client...")
	// Create a gRPC client using the connection
	client := pb.NewGRPCSagaServiceClient(conn)
	return client, conn, nil
}
