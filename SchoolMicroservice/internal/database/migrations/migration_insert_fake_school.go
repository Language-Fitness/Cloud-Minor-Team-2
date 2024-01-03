package migrations

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
	"school/internal/database"
	"time"
)

type FakeSchool struct {
	ID              string `faker:"uuid_hyphenated"`
	Name            string `faker:"word"`
	Location        string `faker:"word"`
	MadeBy          string
	HasOpenaiAccess bool
	OpenaiKey       *string
	JoinCode        string
	CreatedAt       *string
	UpdatedAt       *string
	SoftDeleted     bool
}

func Init() {

	randomClasses := GenerateRandomSchool(10)

	collection, _ := database.GetCollection()
	for _, class := range randomClasses {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		_, err := collection.InsertOne(ctx, class)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GenerateRandomSchool(n int) []FakeSchool {
	var schools []FakeSchool

	for i := 0; i < n; i++ {
		school := FakeSchool{}
		err := faker.FakeData(&school)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		key := "ebaaf8d2-4457-4fa7-bfae-b8a39bdd0bae"
		school.MadeBy = "0e520bea-a96b-47cc-96bc-83633e47c58e"
		school.HasOpenaiAccess = true
		school.OpenaiKey = &key
		school.JoinCode = "fa9fa272-b5dc-4153-9a43-69bd09fb57e7"
		school.SoftDeleted = false
		school.CreatedAt = generateRandomDate()
		school.UpdatedAt = generateRandomDate()

		schools = append(schools, school)
	}

	return schools
}

// generateRandomDate generates a random date string
func generateRandomDate() *string {
	rand.Seed(time.Now().UnixNano())
	days := rand.Intn(365) // Random number of days
	randomDate := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
	return &randomDate
}
