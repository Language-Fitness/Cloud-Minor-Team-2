package repository

import (
	"ExerciseMicroservice/graph/model" // Update this with the correct package name
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// IExerciseRepository GOLANG INTERFACE
// Implements CRUD methods for queries and mutations on Exercise.
type IExerciseRepository interface {
	CreateExercise(newExercise *model.Exercise) (*model.Exercise, error)
	UpdateExercise(id string, updatedExercise model.Exercise) (*model.Exercise, error)
	SoftDeleteExerciseByID(id string, existingExercise model.Exercise) error
	HardDeleteExerciseByID(id string) error
	GetExerciseByID(id string) (*model.Exercise, error)
	ListExercises(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.Exercise, error)
}

// ExerciseRepository GOLANG STRUCT
// Contains a model.Exercise list and a mongo.Collection.
type ExerciseRepository struct {
	exercises  []*model.Exercise
	collection *mongo.Collection
}

// NewExerciseRepository GOLANG FACTORY
// Returns an ExerciseRepository implementing IExerciseRepository.
func NewExerciseRepository(collection *mongo.Collection) IExerciseRepository {
	return &ExerciseRepository{
		exercises:  []*model.Exercise{},
		collection: collection,
	}
}

func (r *ExerciseRepository) CreateExercise(newExercise *model.Exercise) (*model.Exercise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newExercise)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": newExercise.ID}
	var fetchedExercise model.Exercise

	err = r.collection.FindOne(ctx, filter).Decode(&fetchedExercise)
	if err != nil {
		return nil, err
	}

	return &fetchedExercise, nil
}

func (r *ExerciseRepository) UpdateExercise(id string, updatedExercise model.Exercise) (*model.Exercise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedExercise}
	var result model.Exercise

	err := r.collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ExerciseRepository) SoftDeleteExerciseByID(id string, existingExercise model.Exercise) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": existingExercise}

	err := r.collection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&existingExercise)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExerciseRepository) HardDeleteExerciseByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}

	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExerciseRepository) GetExerciseByID(id string) (*model.Exercise, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	var result model.Exercise

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ExerciseRepository) ListExercises(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.Exercise, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	var exercises []*model.Exercise

	cursor, err := r.collection.Find(ctx, bsonFilter, paginateOptions)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		cursor.Close(ctx)
	}(cursor, ctx)

	// Decode results
	for cursor.Next(ctx) {
		var exercise model.Exercise
		if err := cursor.Decode(&exercise); err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}

	return exercises, nil
}
