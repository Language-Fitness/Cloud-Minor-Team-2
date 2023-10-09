package Base

//type MongoRepo struct {
//	collection  *mongo.Collection
//	constructor func() interface{}
//}

// New creates a new instance of MongoRepo
//func New(coll *mongo.Collection, cons func() interface{}) *MongoRepo {
//	return &MongoRepo{
//		collection:  coll,
//		constructor: cons,
//	}
//}

//// Get a list of resource
//func (r *MongoRepo) Get() ([]interface{}, error) {
//	return nil, nil
//}
//
//// GetOne resource based on its ID
//func (r *MongoRepo) GetOne(id string) (interface{}, error) {
//	return nil, nil
//}
//
//// Create a new resource
//func (r *MongoRepo) Create(obj interface{}) error {
//	return nil
//}
//
//// Update a resource
//func (r *MongoRepo) Update(id string, obj interface{}) error {
//	return nil
//}
//
//// Delete a resource
//func (r *MongoRepo) Delete(id string) error {
//	return nil
//}

type Repository struct {
	constructor func() interface{}
}

// New creates a new instance of MongoRepo
func New(cons func() interface{}) *Repository {
	return &Repository{
		constructor: cons,
	}
}

// Get a list of resource
func (r *Repository) Get() ([]interface{}, error) {
	return nil, nil
}

// GetOne resource based on its ID
func (r *Repository) GetOne(id string) (interface{}, error) {
	return nil, nil
}

// Create a new resource
func (r *Repository) Create(obj interface{}) error {
	return nil
}

// Update a resource
func (r *Repository) Update(id string, obj interface{}) error {
	return nil
}

// Delete a resource
func (r *Repository) Delete(id string) error {
	return nil
}

type Person struct{}
type Enemy struct{}

func _() {
	_ = New(
		func() interface{} {
			return &Person{}
		})

	_ = New(
		func() interface{} {
			return &Enemy{}
		})
}
