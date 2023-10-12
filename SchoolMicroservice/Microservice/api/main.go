package main

import (
	"example/micro/school-microservice/App"
	"example/micro/school-microservice/Domain"
	"example/micro/school-microservice/Service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const webPort = ":8080"

func main() {
	fmt.Println("Starting App")

	var router = mux.NewRouter()

	schoolRepo := Domain.NewSchoolRepository()
	schoolServices := Service.NewSchoolService(schoolRepo)

	var schoolHandlers = App.SchoolHandlers{Service: schoolServices}

	router.HandleFunc("/schools", schoolHandlers.GetAll).
		Methods("GET").
		Name("GetAllSchools")

	router.HandleFunc("/schools/{id}", schoolHandlers.GetOne).
		Methods("GET").
		Name(" GetSchool")

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, router))
}
