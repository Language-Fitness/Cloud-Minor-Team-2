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
	CreateResult(newResult *model.Result) (*model.Result, error)
	UpdateResult(id string, updatedResult model.Result) (*model.Result, error)
	DeleteResultByID(id string) error
	GetResultByID(id string) (*model.Result, error)
	GetResultByExerciseId(id string) (*model.Result, error)
	GetResultByClassId(id string) ([]*model.Result, error)
	GetResultsByUserID(userID string) ([]*model.Result, error)
	DeleteResultByClassID(classID string, userID string) error
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

func (r *ResultRepository) DeleteResultByClassID(classID string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{
		"class_id": classID,
		"user_id":  userID, // Add filter for user_id
	}

	_, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err // Return any MongoDB-related errors.
	}

	return nil
}

func (r *ResultRepository) GetResultByExerciseId(id string) (*model.Result, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"exercise_id": id}
	var result model.Result

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ResultRepository) GetResultByClassId(id string) ([]*model.Result, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"class_id": id}
	var result []*model.Result

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
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

func (r *ResultRepository) GetResultsByUserID(userID string) ([]*model.Result, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"user_id": userID}
	var results []*model.Result

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
