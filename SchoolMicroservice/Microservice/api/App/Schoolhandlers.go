package App

import (
	"encoding/json"
	"example/micro/school-microservice/Domain"
	"example/micro/school-microservice/Service"
	"net/http"
)

type SchoolHandlers struct {
	Service Service.SchoolService
}

func (s *SchoolHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	_, _ = s.Service.GetSchools()

	school := Domain.School{
		ID:          "aaaa",
		Name:        "Example School",
		Location:    "Somewhere",
		CreatedAt:   "2023-10-12T08:00:00",
		UpdatedAt:   "2023-10-12T08:30:00",
		SoftDeleted: false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(school)
}

func (s *SchoolHandlers) GetOne(w http.ResponseWriter, r *http.Request) {
	_, _ = s.Service.GetOneSchool("uuid")

	school := Domain.School{
		ID:          "aaaa",
		Name:        "Example School",
		Location:    "Somewhere",
		CreatedAt:   "2023-10-12T08:00:00",
		UpdatedAt:   "2023-10-12T08:30:00",
		SoftDeleted: false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(school)
}

func (s *SchoolHandlers) Create(w http.ResponseWriter, r *http.Request) {
	_ = s.Service.CreateSchool(Domain.School{})
}

func (s *SchoolHandlers) Update(w http.ResponseWriter, r *http.Request) {
	_ = s.Service.UpdateSchool("uuid", Domain.School{})
}

func (s *SchoolHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	_ = s.Service.DeleteSchool("uuid")
}
