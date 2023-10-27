package repository

import (
	"Module/graph/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// IModuleRepository GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on Module.
type IModuleRepository interface {
	CreateModule(newModule *model.Module) (*model.Module, error)
	UpdateModule(id string, updatedModule model.ModuleInput) (*model.Module, error)
	DeleteModuleByID(id string) error
	GetModuleByID(id string) (*model.Module, error)
	ListModules() ([]*model.Module, error)
}

// ModuleRepository GOLANG STRUCT
// Contains a model.Module list and a mongo.Collection.
type ModuleRepository struct {
	modules    []*model.Module
	collection *mongo.Collection
}

// NewModuleRepository GOLANG FACTORY
// Returns a ModuleRepository implementing IModuleRepository.
func NewModuleRepository(collection *mongo.Collection) IModuleRepository {
	return &ModuleRepository{
		modules:    []*model.Module{},
		collection: collection,
	}
}

func (r *ModuleRepository) CreateModule(newModule *model.Module) (*model.Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newModule)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": newModule.ID}
	var fetchedModule model.Module

	err = r.collection.FindOne(ctx, filter).Decode(&fetchedModule)
	if err != nil {
		return nil, err
	}

	return &fetchedModule, nil
}

func (r *ModuleRepository) UpdateModule(id string, updatedModule model.ModuleInput) (*model.Module, error) {
	// Check if the module exists in MongoDB.
	_, err := r.GetModuleByID(id)
	if err != nil {
		return nil, err // Return the error if it doesn't exist in MongoDB.
	}

	// If the module exists in MongoDB, update it in MongoDB.
	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedModule}
	var result model.Module

	err = r.collection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}

	return &result, nil
}

func (r *ModuleRepository) DeleteModuleByID(id string) error {
	// Check if the module exists in MongoDB.
	_, err := r.GetModuleByID(id)
	if err != nil {
		return err // Return the error if it exists in MongoDB.
	}

	// If the module exists in MongoDB, delete it from MongoDB.
	filter := bson.M{"id": id}
	_, err = r.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err // Return any MongoDB-related errors.
	}

	return nil
}

func (r *ModuleRepository) GetModuleByID(id string) (*model.Module, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()
	// Attempt to fetch the module from MongoDB.
	filter := bson.M{"id": id}
	var result model.Module
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ModuleRepository) ListModules() ([]*model.Module, error) {
	// Define an empty slice to store the modules.
	var modules []*model.Module

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	// Retrieve all modules from MongoDB.
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
		var module model.Module
		if err := cursor.Decode(&module); err != nil {
			return nil, err // Return any decoding errors.
		}
		modules = append(modules, &module)
	}

	return modules, nil
}
