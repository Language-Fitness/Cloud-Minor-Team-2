package repository

import (
	"Module/graph/model"
	"Module/internal/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IModuleRepository interface {
	CreateModule(newModule *model.Module) error
	UpdateModule(updatedModule *model.Module) error
	DeleteModuleByID(id string) error
	GetModuleByID(id string) (*model.Module, error)
	ListModules() ([]*model.Module, error)
}

type ModuleRepository struct {
	modules    []*model.Module
	collection *mongo.Collection
}

func NewModuleRepository() *ModuleRepository {
	collection, _ := database.GetCollection()

	return &ModuleRepository{
		modules:    []*model.Module{},
		collection: collection,
	}
}

func (r *ModuleRepository) CreateModule(newModule *model.Module) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newModule)
	if err != nil {
		return err
	}
	return nil
}

func (r *ModuleRepository) UpdateModule(updatedModule *model.Module) error {
	// Check if the module exists in MongoDB.
	_, err := r.GetModuleByID(updatedModule.ID)
	if err != nil {
		return err // Return the error if it doesn't exist in MongoDB.
	}

	// If the module exists in MongoDB, update it in MongoDB.
	filter := bson.M{"id": updatedModule.ID}
	update := bson.M{"$set": updatedModule}
	_, err = r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err // Return any MongoDB-related errors.
	}

	return nil
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
