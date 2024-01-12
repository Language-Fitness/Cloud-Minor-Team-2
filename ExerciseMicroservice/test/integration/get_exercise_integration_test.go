package integration

import (
	"ExerciseMicroservice/test/integration/helper"
	"ExerciseMicroservice/test/integration/requests"
	r "ExerciseMicroservice/test/integration/responses"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestResolver_GetExerciseWithAdminToken(t *testing.T) {
	fmt.Println("\nRunning TestResolver_GetExerciseWithAdminToken")
	c := helper.CreateClient()

	t.Run("Get exercise with admin token", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		c.MustPost(
			requests.GetExerciseQuery,
			&r.GetExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			helper.AddContext(Token),
		)

		// Perform assertions based on your expected results
		assert.NotEmpty(t, r.GetExerciseResponse)
		assert.NotEmpty(t, r.GetExerciseResponse.GetExercise.ID)
		assert.Equal(t, requests.ClassID, r.GetExerciseResponse.GetExercise.ClassID)
		assert.Equal(t, requests.ModuleID, r.GetExerciseResponse.GetExercise.ModuleID)
		assert.Equal(t, requests.Name, r.GetExerciseResponse.GetExercise.Name)
		assert.Equal(t, requests.Question, r.GetExerciseResponse.GetExercise.Question)
		assert.Equal(t, requests.Answers[0]["value"], r.GetExerciseResponse.GetExercise.Answers[0].Value)
		assert.Equal(t, requests.Answers[0]["correct"], r.GetExerciseResponse.GetExercise.Answers[0].Correct)
		assert.Equal(t, requests.Answers[1]["value"], r.GetExerciseResponse.GetExercise.Answers[1].Value)
		assert.Equal(t, requests.Answers[1]["correct"], r.GetExerciseResponse.GetExercise.Answers[1].Correct)
		assert.Equal(t, requests.Difficulty, r.GetExerciseResponse.GetExercise.Difficulty)
	})
}

func TestResolver_GetExerciseWithStudentToken(t *testing.T) {
	fmt.Println("\nRunning TestResolver_GetExerciseWithStudentToken")
	c := helper.CreateClient()

	t.Run("Get exercise with student token", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.GetExerciseQuery,
			&r.GetExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			helper.AddContext(StudentToken),
		)
		assert.Nil(t, err)
		assert.NotEmpty(t, r.GetExerciseResponse)
		assert.NotEmpty(t, r.GetExerciseResponse.GetExercise.ID)
		assert.Equal(t, requests.ClassID, r.GetExerciseResponse.GetExercise.ClassID)
		assert.Equal(t, requests.ModuleID, r.GetExerciseResponse.GetExercise.ModuleID)
		assert.Equal(t, requests.Name, r.GetExerciseResponse.GetExercise.Name)
		assert.Equal(t, requests.Question, r.GetExerciseResponse.GetExercise.Question)
		assert.Equal(t, requests.Answers[0]["value"], r.GetExerciseResponse.GetExercise.Answers[0].Value)
		assert.Equal(t, requests.Answers[0]["correct"], r.GetExerciseResponse.GetExercise.Answers[0].Correct)
		assert.Equal(t, requests.Answers[1]["value"], r.GetExerciseResponse.GetExercise.Answers[1].Value)
		assert.Equal(t, requests.Answers[1]["correct"], r.GetExerciseResponse.GetExercise.Answers[1].Correct)
		assert.Equal(t, requests.Difficulty, r.GetExerciseResponse.GetExercise.Difficulty)
	})
}

func TestResolver_GetExerciseInvalidUuid(t *testing.T) {
	fmt.Println("\nRunning TestResolver_GetExerciseInvalidUuid")
	c := helper.CreateClient()

	t.Run("Get exercise with invalid uuid", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.GetExerciseQuery,
			&r.GetExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d76"),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.GetExerciseInvalidUUIDResponseError, errorResponse[0].Message)
	})
}

func TestResolver_GetExerciseUuidDoesNotExist(t *testing.T) {
	fmt.Println("\nRunning TestResolver_GetExerciseUuidDoesNotExist")
	c := helper.CreateClient()

	t.Run("Get exercise with uuid that does not exist", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.GetExerciseQuery,
			&r.GetExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd7b7b7d777"),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.ExerciseNotFoundResponseError, errorResponse[0].Message)
	})
}
