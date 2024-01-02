package migrations

import (
	"Class/internal/database"
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
	"time"
)

type FakeClass struct {
	ID          string `faker:"uuid_hyphenated"`
	ModuleID    string
	Name        string `faker:"word"`
	Description string `faker:"sentence"`
	Difficulty  string `faker:"oneof:A1,A2,B1,B2,C1,C2"`
	MadeBy      string
	CreatedAt   *string
	UpdatedAt   *string
	SoftDeleted bool
}

func Init() {

	randomClasses := GenerateRandomClasses(10)

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

func GenerateRandomClasses(n int) []FakeClass {
	var classes []FakeClass

	for i := 0; i < n; i++ {
		class := FakeClass{}
		err := faker.FakeData(&class)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		class.ModuleID = "09d6be4b-da77-4be0-9094-445e1a5e639a"
		class.MadeBy = "0e520bea-a96b-47cc-96bc-83633e47c58e"
		class.SoftDeleted = false
		class.CreatedAt = generateRandomDate()
		class.UpdatedAt = generateRandomDate()

		classes = append(classes, class)
	}

	fakeClassWithIdSet := FakeClass{}
	err := faker.FakeData(&fakeClassWithIdSet)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	fakeClassWithIdSet.ModuleID = "09d6be4b-da77-4be0-9094-445e1a5e639a"
	fakeClassWithIdSet.MadeBy = "0e520bea-a96b-47cc-96bc-83633e47c58e"
	fakeClassWithIdSet.SoftDeleted = false
	fakeClassWithIdSet.CreatedAt = generateRandomDate()
	fakeClassWithIdSet.UpdatedAt = generateRandomDate()

	classes = append(classes, fakeClassWithIdSet)

	return classes
}

// generateRandomDate generates a random date string
func generateRandomDate() *string {
	rand.Seed(time.Now().UnixNano())
	days := rand.Intn(365) // Random number of days
	randomDate := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
	return &randomDate
}
