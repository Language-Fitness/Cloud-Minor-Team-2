package service

import (
	"context"
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
	sagaObject, err := s.initializeSagaObject(token, filter)
	if err != nil {
		return nil, err
	}

	err2 := s.findAllChildren(token, sagaObject)
	if err2 != nil {
		return nil, err2
	}

	client, conn4, err := createGRPCClient(sagaObject.ObjectType)
	if err != nil {
		return nil, err
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

	sagaObject, err = s.updateSagaObject(sagaObject, &requestUpdated, sagaObject, err)
	if err != nil {
		return nil, err
	}

	if err := s.softDeleteChildren(token, sagaObject); err != nil {
		// Handle rollback logic or return an error
		return nil, err
	}

	//// Step 6: Loop through items to check if all object_status are Deleted
	if !s.areAllItemsDeleted(token, sagaObject) {
		// Step 7: If not everything is deleted, reloop steps 4 and 5 but undelete every item
		if err := s.undeleteItems(token, sagaObject); err != nil {
			successMessage := &model.SuccessMessage{
				ID:         sagaObject.ID,
				Text:       "Operation not successful",
				Status:     model.SagaObjectStatusExist,
				ObjectID:   sagaObject.ObjectID,
				ObjectType: filter.ObjectType,
			}

			return successMessage, err
		}
	}

	successMessage := &model.SuccessMessage{
		ID:         sagaObject.ID,
		Text:       "Operation successful",
		Status:     model.SagaObjectStatusDeleted,
		ObjectID:   sagaObject.ObjectID,
		ObjectType: filter.ObjectType,
	}

	return successMessage, nil
}

func (s SagaService) initializeSagaObject(token string, filter *model.SagaFilter) (*model.SagaObject, error) {
	client, conn1, err := createGRPCClient(filter.ObjectType)
	if err != nil {
		return nil, err
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

	return object, nil
}

func (s SagaService) findAllChildren(token string, sagaObject *model.SagaObject) error {
	if sagaObject.ObjectType == model.SagaObjectTypesResult {
		return nil
	}

	client, conn2, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	if err != nil {
		return err
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
		return err
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

	objects, err := s.Repo.ListSagaObjects(s.createQuery(sagaObject))
	if err != nil {
		return false
	}

	client, conn3, err := createGRPCClient(getChildTypeByType(sagaObject.ObjectType))
	if err != nil {
		return false
	}

	defer conn3.Close()

	for _, child := range objects {
		request := pb.ObjectRequest{
			BearerToken:  token,
			ObjectId:     child.ObjectID,
			ObjectType:   convertToPBObjectType(child.ObjectType),
			ObjectStatus: pb.SagaObjectStatus_EXIST,
		}

		response, err := client.FindSagaObject(context.Background(), &request)
		if err != nil {
			log.Printf("failed to call FindObject RPC: %v", err)
		}

		if convertToModelObjectStatus(response.ObjectStatus) != child.ObjectStatus {
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
		return err
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
