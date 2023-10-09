package Base

//type Repository struct {
//	collection  *mongo.Collection
//	constructor func() interface{}
//}
//
//New creates a new instance of MongoRepo
//func New(coll *mongo.Collection, cons func() interface{}) *Repository {
//	return &Repository{
//		collection:  coll,
//		constructor: cons,
//	}
//}
//
//// Get a list of resource
//func (r *Repository) Get() ([]interface{}, error) {
//	return nil, nil
//}
//
//// GetOne resource based on its ID
//func (r *Repository) GetOne(id string) (interface{}, error) {
//	return nil, nil
//}
//
//// Create a new resource
//func (r *Repository) Create(obj interface{}) error {
//	return nil
//}
//
//// Update a resource
//func (r *Repository) Update(id string, obj interface{}) error {
//	return nil
//}
//
//// Delete a resource
//func (r *Repository) Delete(id string) error {
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
	personRepo := New(
		func() interface{} {
			return &Person{}
		})

	_, _ = personRepo.Get()
	_, _ = personRepo.GetOne("uuid")
	_ = personRepo.Create(Person{})
	_ = personRepo.Update("uuid", Person{})
	_ = personRepo.Delete("uuid")

	enemyRepo := New(
		func() interface{} {
			return &Enemy{}
		})

	_, _ = enemyRepo.Get()
	_, _ = enemyRepo.GetOne("uuid")
	_ = enemyRepo.Create(Enemy{})
	_ = enemyRepo.Update("uuid", Enemy{})
	_ = enemyRepo.Delete("uuid")
}
