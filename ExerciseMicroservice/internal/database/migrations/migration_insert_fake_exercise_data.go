package migrations

import (
	"ExerciseMicroservice/graph/model"
	"ExerciseMicroservice/internal/database"
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"math/rand"
	"time"
)

var classID = "962c8541-e4be-4c06-9875-d4184b367dca"
var moduleID = "09d6be4b-da77-4be0-9094-445e1a5e639a"

type FakeAnswer struct {
	Value   string `faker:"sentence"`
	Correct bool
}

type FakeExercise struct {
	ID          string `faker:"uuid_hyphenated"`
	ClassID     string `faker:"uuid_hyphenated"`
	ModuleID    string `faker:"uuid_hyphenated"`
	Name        string `faker:"word"`
	Question    string `faker:"sentence"`
	Answers     []*FakeAnswer
	Difficulty  model.LanguageLevel `faker:"oneof:A1,A2,B1,B2,C1,C2"`
	CreatedAt   *string
	UpdatedAt   *string
	SoftDeleted bool
	MadeBy      string `faker:"uuid_hyphenated"`
}

func InitExercise() {
	randomExercises := GenerateRandomExercises(10)

	collection, _ := database.GetCollection()
	for _, exercise := range randomExercises {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		_, err := collection.InsertOne(ctx, exercise)
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GenerateRandomExercises(n int) []FakeExercise {
	var exercises []FakeExercise

	for i := 0; i < n; i++ {
		exercise := FakeExercise{}
		err := faker.FakeData(&exercise)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		exercise.ClassID = classID
		exercise.ModuleID = moduleID
		exercise.SoftDeleted = false
		exercise.CreatedAt = generateRandomDate()
		exercise.UpdatedAt = generateRandomDate()

		exercises = append(exercises, exercise)
	}

	fakeExerciseWithIdSet := FakeExercise{}
	err := faker.FakeData(&fakeExerciseWithIdSet)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	fakeExerciseWithIdSet.ID = "95f964a0-9749-4064-9162-cdd1b7b5d776" // use this id for result
	fakeExerciseWithIdSet.ClassID = classID
	fakeExerciseWithIdSet.ModuleID = moduleID
	fakeExerciseWithIdSet.SoftDeleted = false
	fakeExerciseWithIdSet.CreatedAt = generateRandomDate()
	fakeExerciseWithIdSet.UpdatedAt = generateRandomDate()

	exercises = append(exercises, fakeExerciseWithIdSet)

	return exercises
}

// generateRandomDate generates a random date string
func generateRandomDate() *string {
	rand.Seed(time.Now().UnixNano())
	days := rand.Intn(365) // Random number of days
	randomDate := time.Now().AddDate(0, 0, -days).Format(time.RFC3339)
	return &randomDate
}
