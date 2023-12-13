package repository

import (
	"ResultMicroservice/graph/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// IResultRepository GOLANG INTERFACE
// Implements five CRUD methods for queries and mutations on Result.
type IResultRepository interface {
	//Todo only return items that are not soft deleted
	CreateResult(newResult *model.Result) (*model.Result, error)
	UpdateResult(id string, updatedResult model.Result) (*model.Result, error)
	DeleteResultByID(id string) error
	GetResultByID(id string) (*model.Result, error)
	ListResults() ([]*model.Result, error) //TODO: implement
	//Saga GRPC
	//Todo has to be soft deleted, before it can be hard deleted
	SoftDeleteByUser(userID string) error
	SoftDeleteByClass(classID string) error
	SoftDeleteByModule(moduleID string) error
	DeleteByUser(userID string) error
	DeleteByClass(classID string) error
	DeleteByModule(moduleID string) error
}

// ResultRepository GOLANG STRUCT
// Contains a model.Result list and a mongo.Collection.
type ResultRepository struct {
	results    []*model.Result
	collection *mongo.Collection
}

// NewResultRepository GOLANG FACTORY
// Returns a ResultRepository implementing IResultRepository.
func NewResultRepository(collection *mongo.Collection) IResultRepository {
	return &ResultRepository{
		results:    []*model.Result{},
		collection: collection,
	}
}

func (r *ResultRepository) ListResults() ([]*model.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ResultRepository) CreateResult(newResult *model.Result) (*model.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newResult)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": newResult.ID}
	var fetchedResult model.Result

	err = r.collection.FindOne(ctx, filter).Decode(&fetchedResult)
	if err != nil {
		return nil, err
	}

	return &fetchedResult, nil
}

func (r *ResultRepository) UpdateResult(id string, updatedResult model.Result) (*model.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.GetResultByID(id)
	if err != nil {
		return nil, err // Return the error if it doesn't exist in MongoDB.
	}

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedResult}
	var result model.Result

	err = r.collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}

	return &result, nil
}

func (r *ResultRepository) DeleteResultByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.GetResultByID(id)
	if err != nil {
		return err // Return the error if it exists in MongoDB.
	}

	filter := bson.M{"id": id}

	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err // Return any MongoDB-related errors.
	}

	return nil
}

func (r *ResultRepository) GetResultByID(id string) (*model.Result, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	var result model.Result

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SoftDeleteByClass GOLANG METHOD
// Soft deletes all results associated with a given class ID.
func (r *ResultRepository) SoftDeleteByClass(classID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"class_id": classID}
	update := bson.M{"$set": bson.M{"soft_deleted": true, "updated_at": time.Now().Format(time.RFC3339)}}

	_, err := r.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteByUser GOLANG METHOD
// Soft deletes all results associated with a given user ID.
func (r *ResultRepository) SoftDeleteByUser(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"user_id": userID}
	update := bson.M{"$set": bson.M{"soft_deleted": true, "updated_at": time.Now().Format(time.RFC3339)}}

	_, err := r.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// SoftDeleteByModule GOLANG METHOD
// Soft deletes all results associated with a given module ID.
func (r *ResultRepository) SoftDeleteByModule(moduleID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"module_id": moduleID}
	update := bson.M{"$set": bson.M{"soft_deleted": true, "updated_at": time.Now().Format(time.RFC3339)}}

	_, err := r.collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByClass GOLANG METHOD
// Deletes all results associated with a given class ID.
func (r *ResultRepository) DeleteByClass(classID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"class_id": classID}

	_, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByUser GOLANG METHOD
// Deletes all results associated with a given user ID.
func (r *ResultRepository) DeleteByUser(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"user_id": userID}

	_, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *ResultRepository) DeleteByModule(moduleID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"module_id": moduleID}

	_, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
