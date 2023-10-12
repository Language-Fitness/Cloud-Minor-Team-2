package Domain

type ISchoolRepository interface {
	Get() ([]School, error)
	GetOne(id string) (School, error)
	Create(obj School) error
	Update(id string, obj School) error
	Delete(id string) error
}

type SchoolRepository struct {
}

// NewSchoolRepository creates a new instance of MongoRepo
func NewSchoolRepository() *SchoolRepository {
	return &SchoolRepository{}
}

// Get a list of resource
func (r *SchoolRepository) Get() ([]School, error) {
	return nil, nil
}

// GetOne resource based on its ID
func (r *SchoolRepository) GetOne(id string) (School, error) {
	return School{}, nil
}

// Create a new resource
func (r *SchoolRepository) Create(obj School) error {
	return nil
}

// Update a resource
func (r *SchoolRepository) Update(id string, obj School) error {
	return nil
}

// Delete a resource
func (r *SchoolRepository) Delete(id string) error {
	return nil
}
