package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	softDeleteChildren(token string, object *model.SagaObject) error
	areAllItemsDeleted(token string, sagaObject *model.SagaObject) bool
	undeleteItems(token string, sagaObject *model.SagaObject) error
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

//goland:noinspection SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection
func (s SagaService) InitSagaSteps(token string, filter *model.SagaFilter) (*model.SuccessMessage, error) {
	// Step 1: check if saga object exist and if it does then create it
	token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDUzMjg4MDYsImlhdCI6MTcwNTMyNzkwNiwianRpIjoiODQ1MGE3YTAtYzI0NC00YzJhLThlNDctNzNlYWZjZWY2ZWU1IiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiNmMxY2U0NDgtNjcwZi00N2IyLTgzZjctNGQ3NzFiMDE3NzViIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6IjY2NzAwYmZlLWExZmYtNGE4NC1iMWRmLTkyY2QyMDBlOWQxZCIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsibWFuYWdlLXVzZXJzIiwidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9yZXN1bHRfc29mdERlbGV0ZSIsImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZmlsdGVyX2V4ZXJjaXNlX2RpZmZpY3VsdHkiLCJmaWx0ZXJfc2Nob29sX25hbWUiLCJ1cGRhdGVfcmVzdWx0IiwiZmlsdGVyX2V4ZXJjaXNlX21vZHVsZV9pZCIsImZpbHRlcl9tb2R1bGVfY2F0ZWdvcnkiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX2xvY2F0aW9uIiwiZmlsdGVyX21vZHVsZV9kaWZmaWN1bHR5IiwiZmlsdGVyX3Jlc3VsdF9tb2R1bGVfaWQiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwiZmlsdGVyX3NjaG9vbF9zb2Z0RGVsZXRlIiwiZGVsZXRlX3Jlc3VsdF9hbGwiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJjcmVhdGVfcmVzdWx0IiwiZ2V0X3Jlc3VsdF9hbGwiLCJmaWx0ZXJfbW9kdWxlX21hZGVfYnkiLCJsaXN0X3Jlc3VsdHNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX3F1ZXN0aW9uX3R5cGVfaWQiLCJ1cGRhdGVfY2xhc3NfYWxsIiwiZ2V0X2NsYXNzIiwiZ2V0X3NjaG9vbHNfYWxsIiwiZmlsdGVyX3Jlc3VsdF9leGVyY2lzZV9pZCIsImZpbHRlcl9jbGFzc19zb2Z0RGVsZXRlIiwidXBkYXRlX3Jlc3VsdF9hbGwiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zX2Zyb21fZmlsZSIsImdldF9jbGFzc2VzX2FsbCIsInVwZGF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX21hZGVfYnkiLCJnZXRfZXhlcmNpc2VzX2FsbCIsIm9wZW5haV9nZW5lcmF0ZV9leHBsYW5hdGlvbiIsImZpbHRlcl9jbGFzc19tYWRlX2J5IiwiZmlsdGVyX21vZHVsZV9zb2Z0RGVsZXRlIiwiZ2V0X2V4ZXJjaXNlcyIsImdldF9jbGFzc2VzIiwiZGVsZXRlX21vZHVsZSIsImdldF9zY2hvb2xzIiwiZGVsZXRlX2V4ZXJjaXNlIiwidXBkYXRlX2V4ZXJjaXNlIiwiZ2V0X2V4ZXJjaXNlIiwiZmlsdGVyX3Jlc3VsdF91c2VyX2lkIiwiZmlsdGVyX2V4ZXJjaXNlX25hbWUiLCJmaWx0ZXJfZXhlcmNpc2Vfc29mdERlbGV0ZSIsImRlbGV0ZV9leGVyY2lzZV9hbGwiLCJmaWx0ZXJfcmVzdWx0X2NsYXNzX2lkIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJkZWxldGVfcmVzdWx0IiwiY3JlYXRlX21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJjcmVhdGVfY2xhc3MiLCJjcmVhdGVfc2Nob29sIiwiZ2V0X21vZHVsZXNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX2NsYXNzX2lkIiwibGlzdF9yZXN1bHRzIiwiZmlsdGVyX21vZHVsZV9zY2hvb2xfaWQiLCJmaWx0ZXJfY2xhc3NfbmFtZSIsImdldF9yZXN1bHQiLCJmaWx0ZXJfc2Nob29sX2hhc19vcGVuYWlfYWNjZXNzIiwib3BlbmFpX2dldF9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlIiwiZmlsdGVyX21vZHVsZV9uYW1lIiwiZmlsdGVyX21vZHVsZV9tYWRlX2J5X25hbWUiLCJmaWx0ZXJfZXhlcmNpc2VfbWFkZV9ieSIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZmlsdGVyX21vZHVsZV9wcml2YXRlIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiNjY3MDBiZmUtYTFmZi00YTg0LWIxZGYtOTJjZDIwMGU5ZDFkIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.lH_n5CK7z0XxDGgzD3Td7jOw51kpcTJOBNBKhXNX_wrbHsIJety_55J0EobI27k227agKP9-eDkPt9PFHdh7R_XvEAaDgIxheAmqYpNDMS0rlhiJwxxdXkK9WBR_IPXfWx2td9wMwtBtZxzMwAew5MfUCbtDU4Ull0cz_2xaWMmtSSBHwqzVNuLmAVX1943SY9L0jGm1JYCMi-mPrN-18xnNhrqDLEKBKBfYBq-fke1toE-uuyXNjJDb8Eg-DL2AH3yrD5fL0wwgU46QeZXNpJ6Kh9XyCuC8Kw_j52_VFsFGXxkdCW_WOFoc05UQFThFsb2DHoSUmT96Ttl2OI_niA"
	fmt.Println("step 1")

	//@TODO CHECK IF THERE IS ALREADY AN OBJECT IN THE DB WITH THIS TYPE!!! BEFORE WE CONTINUE
	sagaObject, err := s.initializeSagaObject(token, filter)
	if err != nil {
		return nil, err
	}

	fmt.Println("step 2")
	// Step 2: Find all possible children
	err2 := s.findAllChildren(token, sagaObject)
	if err2 != nil {
		return nil, err2
	}

	fmt.Println("step 3 \n \n \n \n \n \n \n \n ")
	// Step 5: Start soft deleting items and change object_status to Deleted if success
	client, conn4, err := createGRPCClient(sagaObject.ObjectType)
	if err != nil {
		fmt.Println(err)
	}

	defer conn4.Close()

	request := pb.ObjectRequest{
		BearerToken:  token,
		ObjectId:     sagaObject.ObjectID,
		ObjectType:   convertToPBObjectType(sagaObject.ObjectType),
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	response, err := client.DeleteObject(context.Background(), &request)
	if err != nil {
		log.Printf("failed to call FindObject RPC: %v", err)
	}

	requestUpdated := pb.ObjectRequest{
		BearerToken:  token,
		ObjectId:     sagaObject.ObjectID,
		ObjectType:   response.ObjectType,
		ObjectStatus: response.ObjectStatus,
	}

	fmt.Println("step 4")
	sagaObject, err = s.updateSagaObject(sagaObject, &requestUpdated, sagaObject, err)
	if err != nil {
		return nil, err
	}

	fmt.Println("step 5")
	if err := s.softDeleteChildren(token, sagaObject); err != nil {
		// Handle rollback logic or return an error
		return nil, err
	}

	//@TODO WORKS UNTIL STEP 5 FOR SURE AND TESTED

	fmt.Println("step 6")
	//// Step 6: Loop through items to check if all object_status are Deleted
	if !s.areAllItemsDeleted(token, sagaObject) {
		// Step 7: If not everything is deleted, reloop steps 4 and 5 but undelete every item
		if err := s.undeleteItems(token, sagaObject); err != nil {
			// Handle rollback logic or return an error
			return nil, err
		}
	}
	fmt.Println("step 7")
	successMessage := &model.SuccessMessage{
		ID:         "1",
		Text:       "Operation successful",
		Status:     model.SagaObjectStatusExist,
		ObjectID:   "123",
		ObjectType: model.SagaObjectTypesModule,
	}
	//
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
	if sagaObject.ObjectType == model.SagaObjectTypesResult {
		return nil
	}

	client, conn2, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	if err != nil {
		fmt.Println(err)
	}

	defer conn2.Close()

	request := pb.ObjectRequest{
		BearerToken:  token,
		ObjectId:     sagaObject.ObjectID,
		ObjectType:   convertToPBObjectType(sagaObject.ObjectType),
		ObjectStatus: pb.SagaObjectStatus_EXIST,
	}

	response, err := client.FindSagaObjectChildren(context.Background(), &request)
	if err != nil {
		log.Printf("failed to call FindObject RPC: %v", err)
	}

	for _, child := range response.Objects {
		fmt.Println("1234567890")
		fmt.Println(convertToModelObjectStatus(child.ObjectStatus))
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

		err2 := s.findAllChildren(token, object)
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (s SagaService) softDeleteChildren(token string, sagaObject *model.SagaObject) error {
	if sagaObject.ObjectType == model.SagaObjectTypesResult {
		return nil
	}

	objects, err := s.Repo.ListSagaObjects(s.createQuery(sagaObject))
	if err != nil {
		return err
	}

	client, conn3, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	if err != nil {
		fmt.Println(err)
	}

	defer conn3.Close()

	for _, child := range objects {
		request := pb.ObjectRequest{
			BearerToken:  token,
			ObjectId:     child.ObjectID,
			ObjectType:   convertToPBObjectType(child.ObjectType),
			ObjectStatus: pb.SagaObjectStatus_EXIST,
		}

		response, err := client.DeleteObject(context.Background(), &request)
		if err != nil {
			//log.Fatalf("failed to call FindObject RPC: %v", err)
			log.Printf("failed to call FindObject RPC: %v", err)
		}

		requestUpdated := pb.ObjectRequest{
			BearerToken:  token,
			ObjectId:     child.ObjectID,
			ObjectType:   response.ObjectType,
			ObjectStatus: response.ObjectStatus,
		}

		updateSagaObject, err3 := s.updateSagaObject(child, &requestUpdated, sagaObject, err)
		if err3 != nil {
			return err3
		}

		err2 := s.softDeleteChildren(token, updateSagaObject)
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (s SagaService) areAllItemsDeleted(token string, sagaObject *model.SagaObject) bool {
	if sagaObject.ObjectType == model.SagaObjectTypesResult {
		return true
	}

	fmt.Println(getChildTypeByType(sagaObject.ObjectType))

	objects, err := s.Repo.ListSagaObjects(s.createQuery(sagaObject))
	if err != nil {
		return false
	}

	client, conn3, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	if err != nil {
		fmt.Println(err)
	}

	defer conn3.Close()

	for _, child := range objects {
		request := pb.ObjectRequest{
			BearerToken:  token,
			ObjectId:     child.ObjectID,
			ObjectType:   convertToPBObjectType(child.ObjectType),
			ObjectStatus: pb.SagaObjectStatus_EXIST,
		}

		fmt.Println("Calling FindObject RPC...")
		fmt.Println(child)
		response, err := client.FindSagaObject(context.Background(), &request)
		if err != nil {
			log.Printf("failed to call FindObject RPC: %v", err)
		}

		fmt.Println(response.ObjectStatus)
		fmt.Println(child.ObjectStatus)
		fmt.Println(child)
		if convertToModelObjectStatus(response.ObjectStatus) != child.ObjectStatus {
			fmt.Println("failed to convert object status")
			return false
		}

		res := s.areAllItemsDeleted(token, child)
		if !res {
			return res
		}
	}

	return true
}

func (s SagaService) undeleteItems(token string, sagaObject *model.SagaObject) error {
	if sagaObject.ObjectType == model.SagaObjectTypesResult {
		return nil
	}

	objects, err := s.Repo.ListSagaObjects(s.createQuery(sagaObject))
	if err != nil {
		return err
	}

	client, conn3, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	if err != nil {
		fmt.Println(err)
	}

	defer conn3.Close()

	for _, child := range objects {

		request := pb.ObjectRequest{
			BearerToken:  token,
			ObjectId:     child.ObjectID,
			ObjectType:   convertToPBObjectType(child.ObjectType),
			ObjectStatus: pb.SagaObjectStatus_EXIST,
		}

		_, err = client.UnDeleteObject(context.Background(), &request)
		if err != nil {
			log.Printf("failed to call FindObject RPC: %v", err)
			return err
		}

		res := s.undeleteItems(token, child)
		if res != nil {
			log.Printf("failed to call FindObject RPC: %v", err)
			return res
		}
	}

	return nil
}

func (s SagaService) createQuery(sagaObject *model.SagaObject) (bson.D, *options.FindOptions) {
	filterQuery := bson.D{{Key: "parentid", Value: sagaObject.ObjectID}}
	optionsQuery := options.Find().
		SetSkip(int64(0)).
		SetLimit(int64(100))

	return filterQuery, optionsQuery
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

func (s SagaService) updateSagaObject(child *model.SagaObject, response *pb.ObjectRequest, sagaObject *model.SagaObject, err error) (*model.SagaObject, error) {
	fmt.Println("updateSagaObject")
	fmt.Println(helper.StringPointer(sagaObject.ObjectID))
	fmt.Println(sagaObject.ObjectType)
	fmt.Println(child.ObjectType)

	updatedSagaObject := model.SagaObject{
		ID:           child.ID,
		ObjectID:     child.ObjectID,
		ObjectType:   child.ObjectType,
		CreatedAt:    child.CreatedAt,
		UpdatedAt:    helper.StringPointer(time.Now().Format("HH:MM:SS")),
		ObjectStatus: convertToModelObjectStatus(response.ObjectStatus),
		ParentID: func() *string {
			if sagaObject.ObjectType == child.ObjectType {
				return nil
			}
			return helper.StringPointer(sagaObject.ObjectID)
		}(),
		ActionDoneBy: child.ActionDoneBy,
	}

	updateSagaObject, err := s.Repo.UpdateSagaObject(child.ID, updatedSagaObject)
	if err != nil {
		return nil, err
	}
	return updateSagaObject, nil
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

	conn, err := grpc.DialContext(context.Background(), getHostByType(sagaType), opts...)
	if err != nil {
		log.Printf("failed to dial gRPC server: %v", err)
		return nil, nil, err
	}

	client := pb.NewGRPCSagaServiceClient(conn)
	return client, conn, nil
}
