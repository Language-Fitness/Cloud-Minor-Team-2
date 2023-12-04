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
	UpdateModule(id string, updatedModule model.Module) (*model.Module, error)
	SoftDeleteModuleByID(id string, existingModule model.Module) error
	HardDeleteModuleByID(id string) error
	GetModuleByID(id string) (*model.Module, error)
	ListModules() ([]*model.ModuleInfo, error)
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

func (r *ModuleRepository) UpdateModule(id string, updatedModule model.Module) (*model.Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedModule}
	var result model.Module

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

func (r *ModuleRepository) SoftDeleteModuleByID(id string, existingModule model.Module) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": existingModule}
	var result model.Module

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

func (r *ModuleRepository) HardDeleteModuleByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err // Return any MongoDB-related errors.
	}

	return nil
}

func (r *ModuleRepository) GetModuleByID(id string) (*model.Module, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	var result model.Module

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ModuleRepository) ListModules() ([]*model.ModuleInfo, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	var modules []*model.ModuleInfo

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
		var module model.ModuleInfo
		if err := cursor.Decode(&module); err != nil {
			return nil, err // Return any decoding errors.
		}
		modules = append(modules, &module)
	}

	return modules, nil
}
