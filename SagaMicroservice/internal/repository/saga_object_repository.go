package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"saga/graph/model"
	"time"
)

// ISagaObjectRepository GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on SagaObject.
type ISagaObjectRepository interface {
	CreateSagaObject(newSagaObject *model.SagaObject) (*model.SagaObject, error)
	UpdateSagaObject(id string, updatedSagaObject model.SagaObject) (*model.SagaObject, error)
	DeleteSagaObjectByID(id string, existingSagaObject model.SagaObject) error
	GetSagaObjectByIDAndType(id string, sagaType model.SagaObjectTypes) (*model.SagaObject, error)
	ListSagaObjects(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.SagaObject, error)
}

// SagaObjectRepository GOLANG STRUCT
// Contains a model.SagaObject list and a mongo.Collection.
type SagaObjectRepository struct {
	modules    []*model.SagaObject
	collection *mongo.Collection
}

// NewSagaObjectRepository GOLANG FACTORY
// Returns a SagaObjectRepository implementing ISagaObjectRepository.
func NewSagaObjectRepository(collection *mongo.Collection) ISagaObjectRepository {
	return &SagaObjectRepository{
		modules:    []*model.SagaObject{},
		collection: collection,
	}
}

func (r *SagaObjectRepository) CreateSagaObject(newSagaObject *model.SagaObject) (*model.SagaObject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newSagaObject)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": newSagaObject.ID}
	var fetchedSagaObject model.SagaObject

	err = r.collection.FindOne(ctx, filter).Decode(&fetchedSagaObject)
	if err != nil {
		return nil, err
	}

	return &fetchedSagaObject, nil
}

func (r *SagaObjectRepository) UpdateSagaObject(id string, updatedSagaObject model.SagaObject) (*model.SagaObject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedSagaObject}
	var result model.SagaObject

	err := r.collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}

	return &result, nil
}

func (r *SagaObjectRepository) DeleteSagaObjectByID(id string, existingSagaObject model.SagaObject) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": existingSagaObject}
	var result model.SagaObject

	err := r.collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func (r *SagaObjectRepository) GetSagaObjectByIDAndType(id string, sagaType model.SagaObjectTypes) (*model.SagaObject, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id, "sagaType": sagaType}
	var result model.SagaObject

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SagaObjectRepository) ListSagaObjects(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.SagaObject, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	var modules []*model.SagaObject

	fmt.Println("ewa")
	fmt.Println(bsonFilter, paginateOptions)

	cursor, err := r.collection.Find(ctx, bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			// Handle closing error if needed
		}
	}(cursor, ctx)

	// Decode results
	for cursor.Next(ctx) {
		var module model.SagaObject
		if err := cursor.Decode(&module); err != nil {
			return nil, err
		}
		modules = append(modules, &module)
	}

	return modules, nil
}
