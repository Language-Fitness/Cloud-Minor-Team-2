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
	MadeBy      string
	Private     bool
	Key         string `faker:"uuid_hyphenated"`
	CreatedAt   *string
	UpdatedAt   *string
	SoftDeleted bool
}

func Init() {
	fmt.Println("test")

	randomModules := GenerateRandomModules(10)

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
		module.MadeBy = "0e520bea-a96b-47cc-96bc-83633e47c58e"
		module.SoftDeleted = false
		module.CreatedAt = generateRandomDate()
		module.UpdatedAt = generateRandomDate()

		modules = append(modules, module)
	}

	fakeModuleWithIdSet := FakeModule{}
	err := faker.FakeData(&fakeModuleWithIdSet)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	fakeModuleWithIdSet.ID = "09d6be4b-da77-4be0-9094-445e1a5e639a"
	fakeModuleWithIdSet.Private = false
	fakeModuleWithIdSet.MadeBy = "0e520bea-a96b-47cc-96bc-83633e47c58e"
	fakeModuleWithIdSet.SoftDeleted = false
	fakeModuleWithIdSet.CreatedAt = generateRandomDate()
	fakeModuleWithIdSet.UpdatedAt = generateRandomDate()

	modules = append(modules, fakeModuleWithIdSet)

	return modules
}

// generateRandomDate generates a random date string
func generateRandomDate() *string {
	rand.Seed(time.Now().UnixNano())
	days := rand.Intn(365) // Random number of days
	randomDate := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
	return &randomDate
}
