package repository

import (
	"context"
	"example/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// IClassRepository GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Class.
type IClassRepository interface {
	CreateClass(newClass *model.Class) (*model.Class, error)
	UpdateClass(id string, updatedClass model.Class) (*model.Class, error)
	DeleteClassByID(id string) error
	GetClassByID(id string) (*model.Class, error)
	ListClasses() ([]*model.ClassInfo, error)
}

// ClassRepository GOLANG STRUCT
// Contains a model.Class list and a mongo.Collection.
type ClassRepository struct {
	Classes    []*model.Class
	collection *mongo.Collection
}

// NewClassRepository GOLANG FACTORY
// Returns a ClassRepository implementing IClassRepository.
func NewClassRepository(collection *mongo.Collection) IClassRepository {
	return &ClassRepository{
		Classes:    []*model.Class{},
		collection: collection,
	}
}

func (r *ClassRepository) CreateClass(newClass *model.Class) (*model.Class, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newClass)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": newClass.ID}
	var fetchedClass model.Class

	err = r.collection.FindOne(ctx, filter).Decode(&fetchedClass)
	if err != nil {
		return nil, err
	}

	return &fetchedClass, nil
}

func (r *ClassRepository) UpdateClass(id string, updatedClass model.Class) (*model.Class, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.GetClassByID(id)
	if err != nil {
		return nil, err // Return the error if it doesn't exist in MongoDB.
	}

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedClass}
	var result model.Class

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

func (r *ClassRepository) DeleteClassByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.GetClassByID(id)
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

func (r *ClassRepository) GetClassByID(id string) (*model.Class, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	var result model.Class

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ClassRepository) ListClasses() ([]*model.ClassInfo, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	var classes []*model.ClassInfo

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var Class model.ClassInfo
		if err := cursor.Decode(&Class); err != nil {
			return nil, err // Return any decoding errors.
		}
		classes = append(classes, &Class)
	}

	return classes, nil
}
