package migrations

import (
	"ResultMicroservice/internal/database"
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
	"time"
)

var resultClassID = "962c8541-e4be-4c06-9875-d4184b367dca"
var resultModuleID = "09d6be4b-da77-4be0-9094-445e1a5e639a"

type FakeResult struct {
	ID          string `faker:"uuid_hyphenated"`
	ExerciseID  string
	UserID      string `faker:"uuid_hyphenated"`
	ClassID     string `faker:"uuid_hyphenated"`
	ModuleID    string `faker:"uuid_hyphenated"`
	Input       string `faker:"sentence"`
	Result      bool
	CreatedAt   *string
	UpdatedAt   *string
	SoftDeleted bool
}

func InitResult() {
	randomResults := GenerateRandomResults(10)

	collection, _ := database.GetCollection()
	for _, result := range randomResults {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		_, err := collection.InsertOne(ctx, result)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GenerateRandomResults(n int) []FakeResult {
	var results []FakeResult

	for i := 0; i < n; i++ {
		result := FakeResult{}
		err := faker.FakeData(&result)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		result.ExerciseID = "95f964a0-9749-4064-9162-cdd1b7b5d776"
		result.ClassID = resultClassID
		result.ModuleID = resultModuleID
		result.SoftDeleted = false
		result.CreatedAt = generateRandomDate()
		result.UpdatedAt = generateRandomDate()

		results = append(results, result)
	}

	fakeResultWithIdSet := FakeResult{}
	err := faker.FakeData(&fakeResultWithIdSet)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	fakeResultWithIdSet.ExerciseID = "95f964a0-9749-4064-9162-cdd1b7b5d776"
	fakeResultWithIdSet.ClassID = resultClassID
	fakeResultWithIdSet.ModuleID = resultModuleID
	fakeResultWithIdSet.SoftDeleted = false
	fakeResultWithIdSet.CreatedAt = generateRandomDate()
	fakeResultWithIdSet.UpdatedAt = generateRandomDate()

	results = append(results, fakeResultWithIdSet)

	return results
}

// generateRandomDate generates a random date string
func generateRandomDate() *string {
	rand.Seed(time.Now().UnixNano())
	days := rand.Intn(365) // Random number of days
	randomDate := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
	return &randomDate
}
