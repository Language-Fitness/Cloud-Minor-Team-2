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

func TestResolver_UpdateExercise(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise")
	c := helper.CreateClient()

	t.Run("Update exercise", func(t *testing.T) {

		// Call the resolver via the client and modify the context via functional options
		c.MustPost(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInput()),
			helper.AddContext(Token),
		)

		// Perform assertions based on your expected results
		assert.NotEmpty(t, r.UpdateExerciseResponse)
		//fmt.Printf("%+v\n", r.CreateExerciseResponse)
		assert.NotEmpty(t, r.UpdateExerciseResponse.UpdateExercise.ID)
		assert.Equal(t, requests.ClassID, r.UpdateExerciseResponse.UpdateExercise.ClassID)
		assert.Equal(t, requests.ModuleID, r.UpdateExerciseResponse.UpdateExercise.ModuleID)
		assert.Equal(t, requests.Name, r.UpdateExerciseResponse.UpdateExercise.Name)
		assert.Equal(t, requests.Question, r.UpdateExerciseResponse.UpdateExercise.Question)
		assert.Equal(t, requests.Answers[0]["value"], r.UpdateExerciseResponse.UpdateExercise.Answers[0].Value)
		assert.Equal(t, requests.Answers[0]["correct"], r.UpdateExerciseResponse.UpdateExercise.Answers[0].Correct)
		assert.Equal(t, requests.Answers[1]["value"], r.UpdateExerciseResponse.UpdateExercise.Answers[1].Value)
		assert.Equal(t, requests.Answers[1]["correct"], r.UpdateExerciseResponse.UpdateExercise.Answers[1].Correct)
		assert.Equal(t, requests.Difficulty, r.UpdateExerciseResponse.UpdateExercise.Difficulty)
	})
}

func TestResolver_UpdateExercise_InvalidClassID(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise_InvalidClassID")
	c := helper.CreateClient()

	t.Run("Update exercise with invalid ClassID", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidClassId()),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.InvalidClassIDResponseError, errorResponse[0].Message)
	})
}

func TestResolver_UpdateExercise_InvalidModuleID(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise_InvalidModuleID")
	c := helper.CreateClient()

	t.Run("Update exercise with invalid ModuleID", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidModuleId()),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.InvalidModuleIDResponseError, errorResponse[0].Message)
	})
}

func TestResolver_UpdateExercise_InvalidName(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise_InvalidName")
	c := helper.CreateClient()

	t.Run("Update exercise with invalid Name", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidName()),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.InvalidNameResponseError, errorResponse[0].Message)
	})
}

func TestResolver_UpdateExercise_InvalidQuestion(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise_InvalidQuestion")
	c := helper.CreateClient()

	t.Run("Update exercise with invalid Question", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidQuestion()),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.InvalidQuestionResponseError, errorResponse[0].Message)
	})
}

func TestResolver_UpdateExercise_NoCorrectAnswers(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise_NoCorrectAnswers")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid Answers", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInputNoCorrectAnswers()),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.NoCorrectAnswersResponseError, errorResponse[0].Message)
	})
}

func TestResolver_UpdateExercise_NoAnswers(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise_NoAnswers")
	c := helper.CreateClient()

	t.Run("Update exercise with invalid Answers", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInputNoAnswers()),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.NoAnswersResponseError, errorResponse[0].Message)
	})
}

func TestResolver_UpdateExercise_NoIncorrectAnswer(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateExercise_NoIncorrectAnswer")
	c := helper.CreateClient()

	t.Run("Update exercise with invalid Answers", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.UpdateExerciseMutation,
			&r.UpdateExerciseResponse,
			client.Var("exerciseID", "95f964a0-9749-4064-9162-cdd1b7b5d776"),
			client.Var("exerciseInput", requests.GenerateExerciseInputNoIncorrectAnswers()),
			helper.AddContext(Token),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.NoIncorrectAnswersResponseError, errorResponse[0].Message)
	})
}
