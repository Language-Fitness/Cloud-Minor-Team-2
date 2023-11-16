package repository

import (
	"context"
	"example/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// ISchoolRepository GOLANG INTERFACE
// Implements five CRUD methods for query's and mutations on School.
type ISchoolRepository interface {
	CreateSchool(newSchool *model.School) (*model.School, error)
	UpdateSchool(id string, updatedSchool model.School) (*model.School, error)
	DeleteSchoolByID(id string) error
	GetSchoolByID(id string) (*model.School, error)
	ListSchools() ([]*model.School, error)
}

// SchoolRepository GOLANG STRUCT
// Contains a model.School list and a mongo.Collection.
type SchoolRepository struct {
	Schooles   []*model.School
	collection *mongo.Collection
}

// NewSchoolRepository GOLANG FACTORY
// Returns a SchoolRepository implementing ISchoolRepository.
func NewSchoolRepository(collection *mongo.Collection) ISchoolRepository {
	return &SchoolRepository{
		Schooles:   []*model.School{},
		collection: collection,
	}
}

func (r *SchoolRepository) CreateSchool(newSchool *model.School) (*model.School, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.collection.InsertOne(ctx, newSchool)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"id": newSchool.ID}
	var fetchedSchool model.School

	err = r.collection.FindOne(ctx, filter).Decode(&fetchedSchool)
	if err != nil {
		return nil, err
	}

	return &fetchedSchool, nil
}

func (r *SchoolRepository) UpdateSchool(id string, updatedSchool model.School) (*model.School, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.GetSchoolByID(id)
	if err != nil {
		return nil, err // Return the error if it doesn't exist in MongoDB.
	}

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedSchool}
	var result model.School

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

func (r *SchoolRepository) DeleteSchoolByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	_, err := r.GetSchoolByID(id)
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

func (r *SchoolRepository) GetSchoolByID(id string) (*model.School, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	var result model.School

	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SchoolRepository) ListSchools() ([]*model.School, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	var Schooles []*model.School

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
		var School model.School
		if err := cursor.Decode(&School); err != nil {
			return nil, err // Return any decoding errors.
		}
		Schooles = append(Schooles, &School)
	}

	return Schooles, nil
}
