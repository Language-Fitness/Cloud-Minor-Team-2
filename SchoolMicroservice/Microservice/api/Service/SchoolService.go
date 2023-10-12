package Service

import (
	"example/micro/school-microservice/Domain"
)

type ISchoolService interface {
	GetSchool() ([]interface{}, error)
	GetOneSchool(id string) (interface{}, error)
	CreateSchool(obj interface{}) error
	UpdateSchool(id string, obj interface{}) error
	DeleteSchool(id string) error
}

type SchoolService struct {
	repo Domain.ISchoolRepository
}

func NewSchoolService(repo Domain.ISchoolRepository) SchoolService {
	return SchoolService{repo}
}

// GetSchools a list of resource
func (s *SchoolService) GetSchools() ([]Domain.School, error) {
	return s.repo.Get()
}

// GetOneSchool resource based on its ID
func (s *SchoolService) GetOneSchool(id string) (Domain.School, error) {
	return s.repo.GetOne(id)
}

// CreateSchool a new resource
func (s *SchoolService) CreateSchool(obj Domain.School) error {
	return s.repo.Create(obj)
}

// UpdateSchool a resource
func (s *SchoolService) UpdateSchool(id string, obj Domain.School) error {
	return s.repo.Update(id, obj)
}

// DeleteSchool a resource
func (s *SchoolService) DeleteSchool(id string) error {
	return s.repo.Delete(id)
}
