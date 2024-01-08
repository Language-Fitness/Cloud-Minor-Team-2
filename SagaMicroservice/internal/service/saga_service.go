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
	"saga/internal/helper"
	"saga/internal/repository"
	"saga/internal/validation"
	"saga/proto/pb"
	"time"
)

type ISagaService interface {
	InitSagaSteps(token string, filter *model.SagaFilter) (*model.SuccessMessage, error)
	initializeSagaObject(token string, filter *model.SagaFilter) (*model.SagaObject, error)
	findAllChildren(token string, sagaObject *model.SagaObject) error
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
	token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI5ck1vbGRRY0pVSFM1alhGNXBLb1M5cVZGMC0yLWZnS0ZreVhRMnZiX0JvIn0.eyJleHAiOjE3MDQ3MjUwODQsImlhdCI6MTcwNDcyNDc4NCwianRpIjoiMWYzMTc2YzQtMWVmYi00MTdjLWE3ZTctNjI0MzMyYjcxNDc1IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjJiOTY2ZTQ2LTg3NDgtNDZhYy1hODNkLTU4MjlmOGI5ZTk0ZiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIwM2Y5YzMxZS0zYTcwLTQ5ZGEtYjVmMS1hMGE5MWVkZGE0NmYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9yZXN1bHRfc29mdERlbGV0ZSIsImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZmlsdGVyX2V4ZXJjaXNlX2RpZmZpY3VsdHkiLCJmaWx0ZXJfc2Nob29sX25hbWUiLCJ1cGRhdGVfcmVzdWx0IiwiZmlsdGVyX2V4ZXJjaXNlX21vZHVsZV9pZCIsImZpbHRlcl9tb2R1bGVfY2F0ZWdvcnkiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX2xvY2F0aW9uIiwiZmlsdGVyX21vZHVsZV9kaWZmaWN1bHR5IiwiZmlsdGVyX3Jlc3VsdF9tb2R1bGVfaWQiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwiZGVsZXRlX3Jlc3VsdF9hbGwiLCJmaWx0ZXJfc2Nob29sX3NvZnREZWxldGUiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9yZXN1bHQiLCJmaWx0ZXJfY2xhc3NfbW9kdWxlX2lkIiwiZ2V0X3Jlc3VsdF9hbGwiLCJmaWx0ZXJfbW9kdWxlX21hZGVfYnkiLCJsaXN0X3Jlc3VsdHNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX3F1ZXN0aW9uX3R5cGVfaWQiLCJ1cGRhdGVfY2xhc3NfYWxsIiwiZ2V0X2NsYXNzIiwiZ2V0X3NjaG9vbHNfYWxsIiwiZmlsdGVyX3Jlc3VsdF9leGVyY2lzZV9pZCIsImZpbHRlcl9jbGFzc19zb2Z0RGVsZXRlIiwidXBkYXRlX3Jlc3VsdF9hbGwiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zX2Zyb21fZmlsZSIsImdldF9jbGFzc2VzX2FsbCIsInVwZGF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX21hZGVfYnkiLCJnZXRfZXhlcmNpc2VzX2FsbCIsImZpbHRlcl9jbGFzc19tYWRlX2J5IiwiZmlsdGVyX21vZHVsZV9zb2Z0RGVsZXRlIiwib3BlbmFpX2dlbmVyYXRlX2V4cGxhbmF0aW9uIiwiZ2V0X2NsYXNzZXMiLCJnZXRfZXhlcmNpc2VzIiwiZGVsZXRlX21vZHVsZSIsImRlbGV0ZV9leGVyY2lzZSIsImdldF9zY2hvb2xzIiwiZ2V0X2V4ZXJjaXNlIiwidXBkYXRlX2V4ZXJjaXNlIiwiZmlsdGVyX3Jlc3VsdF91c2VyX2lkIiwiZmlsdGVyX2V4ZXJjaXNlX25hbWUiLCJmaWx0ZXJfZXhlcmNpc2Vfc29mdERlbGV0ZSIsImRlbGV0ZV9leGVyY2lzZV9hbGwiLCJmaWx0ZXJfcmVzdWx0X2NsYXNzX2lkIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJkZWxldGVfcmVzdWx0IiwiY3JlYXRlX21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJjcmVhdGVfY2xhc3MiLCJjcmVhdGVfc2Nob29sIiwiZ2V0X21vZHVsZXNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX2NsYXNzX2lkIiwiZmlsdGVyX21vZHVsZV9zY2hvb2xfaWQiLCJsaXN0X3Jlc3VsdHMiLCJmaWx0ZXJfY2xhc3NfbmFtZSIsImdldF9yZXN1bHQiLCJvcGVuYWlfZ2V0X3NjaG9vbCIsImZpbHRlcl9tb2R1bGVfbmFtZSIsImZpbHRlcl9tb2R1bGVfbWFkZV9ieV9uYW1lIiwidXBkYXRlX21vZHVsZSIsImZpbHRlcl9leGVyY2lzZV9tYWRlX2J5IiwiZGVsZXRlX3NjaG9vbF9hbGwiLCJ1cGRhdGVfY2xhc3MiLCJmaWx0ZXJfbW9kdWxlX3ByaXZhdGUiLCJkZWxldGVfY2xhc3NfYWxsIl19LCJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiIwM2Y5YzMxZS0zYTcwLTQ5ZGEtYjVmMS1hMGE5MWVkZGE0NmYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5hbWUiOiJjaGFkIGFkbWluIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW5AYWRtaW4uY29tIiwiZ2l2ZW5fbmFtZSI6ImNoYWQiLCJmYW1pbHlfbmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBhZG1pbi5jb20ifQ.a8FEZpvU8E367vxB9z7BtD4fObTef8YNh10KvViuHl6qyp_hbEvGi2BzBl6X438UbMM73JjJ4ngL515CaKIlcVOf39Fkm_RZMsug4l9guACYPQGQaohVCGDeXQVOe0rdaC6O4ir5lPUPycTJOXf6q1JQfpfaVFhDuTAWyTG4fqO1ibuCLEabAtToCo6Y3opfgjSkvQfFniLHpUQQ-yphIS39N5MIZwIYxtIRmi6bQpL4K79ftkjm46mjwzvRs2H_Ri4kL9gaq-gyKbokR7DCOwbxZP0uCjnYS4mZOBSmNtrMSuGH1MBW8afRruCSx8WjBJsVAcmxJNBHPnMt8Ud9ug"

	//@TODO CHECK IF THERE IS ALREADY AN OBJECT IN THE DB WITH THIS TYPE!!! BEFORE WE CONTINUE
	sagaObject, err := s.initializeSagaObject(token, filter)
	if err != nil {
		return nil, err
	}

	// Step 2: Find all possible children
	err2 := s.findAllChildren(token, sagaObject)
	if err2 != nil {
		return nil, err2
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

func (s SagaService) findAllChildren(token string, sagaObject *model.SagaObject) error {
	fmt.Println("find all children:")
	fmt.Println(sagaObject)

	if sagaObject.ObjectType == model.SagaObjectTypesResult {
		return nil
	}

	fmt.Println("get child by type:")
	fmt.Println(getChildTypeByType(sagaObject.ObjectType))
	client, conn2, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	fmt.Println("create client")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("create client after error")

	defer conn2.Close()

	fmt.Println(convertToPBObjectType(sagaObject.ObjectType))
	request := pb.ObjectRequest{
		BearerToken:  token,
		ObjectId:     sagaObject.ObjectID,
		ObjectType:   convertToPBObjectType(sagaObject.ObjectType),
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	fmt.Println("Calling FindObject RPC...")
	response, err := client.FindSagaObjectChildren(context.Background(), &request)
	if err != nil {
		//log.Fatalf("failed to call FindObject RPC: %v", err)
		log.Printf("failed to call FindObject RPC: %v", err)
	}

	fmt.Println("after response")
	if len(response.Objects) != 0 {
		fmt.Println("\nResponse 1:", response.Objects[0].ObjectId)
		fmt.Println("\nResponse 2:", response.Objects[0].ObjectStatus)
		fmt.Println("\nResponse 3:", response.Objects[0].ObjectType)
	}
	for _, child := range response.Objects {
		fmt.Println("childs:")
		fmt.Println(child)

		sagaObject := model.SagaObject{
			ID:           uuid.New().String(),
			ObjectID:     child.ObjectId,
			ObjectType:   convertToModelObjectType(child.ObjectType),
			CreatedAt:    time.Now().Format("HH:MM:SS"),
			ObjectStatus: convertToModelObjectStatus(child.ObjectStatus),
			ParentID:     helper.StringPointer(sagaObject.ObjectID),
			ActionDoneBy: "1", //@TODO get this from the bearer token sub
		}

		object, err := s.Repo.CreateSagaObject(&sagaObject)
		if err != nil {
			return err
		}

		fmt.Println("for loop recursion")
		err2 := s.findAllChildren(token, object)
		fmt.Println(err2)
		if err2 != nil {
			return err2
		}
	}

	return nil
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
		return os.Getenv("RESULT_MS_GRPC_HOST")
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
