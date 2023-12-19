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
	DeleteSchool(id string, existingSchool model.School) error
	GetSchoolByID(id string) (*model.School, error)
	ListSchools(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.SchoolInfo, error)
}

// SchoolRepository GOLANG STRUCT
// Contains a model.School list and a mongo.Collection.
type SchoolRepository struct {
	Schools    []*model.School
	collection *mongo.Collection
}

// NewSchoolRepository GOLANG FACTORY
// Returns a SchoolRepository implementing ISchoolRepository.
func NewSchoolRepository(collection *mongo.Collection) ISchoolRepository {
	return &SchoolRepository{
		Schools:    []*model.School{},
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

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedSchool}
	var result model.School

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

func (r *SchoolRepository) DeleteSchool(id string, existingSchool model.School) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10-second timeout
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{"$set": existingSchool}
	var result model.School

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

func (r *SchoolRepository) ListSchools(bsonFilter bson.D, paginateOptions *options.FindOptions) ([]*model.SchoolInfo, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10) // 10-second timeout
	defer cancel()

	var Schools []*model.SchoolInfo

	cursor, err := r.collection.Find(ctx, bsonFilter, paginateOptions)
	if err != nil {
		return nil, err // Return any MongoDB-related errors.
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var School model.SchoolInfo
		if err := cursor.Decode(&School); err != nil {
			return nil, err // Return any decoding errors.
		}
		Schools = append(Schools, &School)
	}

	return Schools, nil
}
