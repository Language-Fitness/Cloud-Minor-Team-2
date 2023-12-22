package migrations

import (
	"Module/internal/database"
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
	"time"
)

type FakeModule struct {
	ID          string `faker:"uuid_hyphenated"`
	SchoolID    string `faker:"uuid_hyphenated"`
	Name        string `faker:"word"`
	Description string `faker:"sentence"`
	Difficulty  string `faker:"oneof:A1,A2,B1,B2,C1,C2"`
	Category    string `faker:"oneof:Grammatica,Spelling,Woordenschat,Uitdrukkingen,Interpunctie,Werkwoordvervoegingen,Fast_Track"`
	MadeBy      string `faker:"uuid_hyphenated"`
	Private     bool
	Key         string `faker:"uuid_hyphenated"`
	CreatedAt   *string
	UpdatedAt   *string
	SoftDeleted bool
}

func Init() {
	fmt.Println("test")

	randomModules := GenerateRandomModules(100000)

	collection, _ := database.GetCollection()
	for _, module := range randomModules {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		_, err := collection.InsertOne(ctx, module)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GenerateRandomModules(n int) []FakeModule {
	var modules []FakeModule

	for i := 0; i < n; i++ {
		module := FakeModule{}
		err := faker.FakeData(&module)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		module.Private = false
		module.SoftDeleted = false
		module.CreatedAt = generateRandomDate()
		module.UpdatedAt = generateRandomDate()

		modules = append(modules, module)
	}

	return modules
}

// generateRandomDate generates a random date string
func generateRandomDate() *string {
	rand.Seed(time.Now().UnixNano())
	days := rand.Intn(365) // Random number of days
	randomDate := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
	return &randomDate
}
