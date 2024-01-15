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
	GetResultByID(id string) (*model.Result, error)
	ListResults(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.ResultInfo, error)
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

func (r *ResultRepository) ListResults(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.ResultInfo, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	var results []*model.ResultInfo

	cursor, err := r.collection.Find(ctx, bsonFilter, paginateOptions)
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		cursor.Close(ctx)
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var result model.ResultInfo
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}

	return results, nil
}
