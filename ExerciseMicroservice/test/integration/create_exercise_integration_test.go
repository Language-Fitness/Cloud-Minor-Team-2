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

var Token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDQ4OTE3NTYsImlhdCI6MTcwNDg5MDg1NiwianRpIjoiMjFiMzE0MmEtZmFjOS00M2RhLTlkMjItMjZkYmNmNjRjYWZhIiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiNmMxY2U0NDgtNjcwZi00N2IyLTgzZjctNGQ3NzFiMDE3NzViIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6ImYxZmQxMGJhLTM2YTEtNGU1Yi05NDFkLTY1YWNiZWE3NTM0ZSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9yZXN1bHRfc29mdERlbGV0ZSIsImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZmlsdGVyX2V4ZXJjaXNlX2RpZmZpY3VsdHkiLCJmaWx0ZXJfc2Nob29sX25hbWUiLCJ1cGRhdGVfcmVzdWx0IiwiZmlsdGVyX2V4ZXJjaXNlX21vZHVsZV9pZCIsImZpbHRlcl9tb2R1bGVfY2F0ZWdvcnkiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX2xvY2F0aW9uIiwiZmlsdGVyX21vZHVsZV9kaWZmaWN1bHR5IiwiZmlsdGVyX3Jlc3VsdF9tb2R1bGVfaWQiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwiZmlsdGVyX3NjaG9vbF9zb2Z0RGVsZXRlIiwiZGVsZXRlX3Jlc3VsdF9hbGwiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJjcmVhdGVfcmVzdWx0IiwiZ2V0X3Jlc3VsdF9hbGwiLCJmaWx0ZXJfbW9kdWxlX21hZGVfYnkiLCJsaXN0X3Jlc3VsdHNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX3F1ZXN0aW9uX3R5cGVfaWQiLCJ1cGRhdGVfY2xhc3NfYWxsIiwiZ2V0X2NsYXNzIiwiZ2V0X3NjaG9vbHNfYWxsIiwiZmlsdGVyX3Jlc3VsdF9leGVyY2lzZV9pZCIsImZpbHRlcl9jbGFzc19zb2Z0RGVsZXRlIiwidXBkYXRlX3Jlc3VsdF9hbGwiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zX2Zyb21fZmlsZSIsImdldF9jbGFzc2VzX2FsbCIsInVwZGF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX21hZGVfYnkiLCJnZXRfZXhlcmNpc2VzX2FsbCIsIm9wZW5haV9nZW5lcmF0ZV9leHBsYW5hdGlvbiIsImZpbHRlcl9jbGFzc19tYWRlX2J5IiwiZmlsdGVyX21vZHVsZV9zb2Z0RGVsZXRlIiwiZ2V0X2V4ZXJjaXNlcyIsImdldF9jbGFzc2VzIiwiZGVsZXRlX21vZHVsZSIsImdldF9zY2hvb2xzIiwiZGVsZXRlX2V4ZXJjaXNlIiwidXBkYXRlX2V4ZXJjaXNlIiwiZ2V0X2V4ZXJjaXNlIiwiZmlsdGVyX3Jlc3VsdF91c2VyX2lkIiwiZmlsdGVyX2V4ZXJjaXNlX25hbWUiLCJmaWx0ZXJfZXhlcmNpc2Vfc29mdERlbGV0ZSIsImRlbGV0ZV9leGVyY2lzZV9hbGwiLCJmaWx0ZXJfcmVzdWx0X2NsYXNzX2lkIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJkZWxldGVfcmVzdWx0IiwiY3JlYXRlX21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJjcmVhdGVfY2xhc3MiLCJjcmVhdGVfc2Nob29sIiwiZ2V0X21vZHVsZXNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX2NsYXNzX2lkIiwibGlzdF9yZXN1bHRzIiwiZmlsdGVyX21vZHVsZV9zY2hvb2xfaWQiLCJmaWx0ZXJfY2xhc3NfbmFtZSIsImdldF9yZXN1bHQiLCJmaWx0ZXJfc2Nob29sX2hhc19vcGVuYWlfYWNjZXNzIiwib3BlbmFpX2dldF9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlIiwiZmlsdGVyX21vZHVsZV9uYW1lIiwiZmlsdGVyX21vZHVsZV9tYWRlX2J5X25hbWUiLCJmaWx0ZXJfZXhlcmNpc2VfbWFkZV9ieSIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZmlsdGVyX21vZHVsZV9wcml2YXRlIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiZjFmZDEwYmEtMzZhMS00ZTViLTk0MWQtNjVhY2JlYTc1MzRlIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.itOZMLAiMpArPe5bte0Aq3wogIEyZPVfpNR1bVte6cWx72nthy-sx4oW05REAWJumfOw3PjiBA66Zm_KN2QfYkcVw-VdtUtNj98pwfrINHgpymRiIkfim2-kAfmw3n8MmjzS2g0s2l0WtlEhIbR2XunJYuZsoYYTaALpmp1vk5Rq-qBNTCRet5CIg9A8ANCRJj3ql6zt4WMiNBHKG4skDA6mftB1NasmT2qA--aRQR_NKToRE7Sbom1Mw1FcZcMJ5HtuQUplsPB0qpFXWw9duUUStzVQVecvmoVT_4QCnu3i5URfIWM3jsm0eZLyFVY71GafK9q07lohi25nkEtSjg"

func TestResolver_CreateExercise(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise")
	c := helper.CreateClient()

	t.Run("Create exercise", func(t *testing.T) {

		// Call the resolver via the client and modify the context via functional options
		c.MustPost(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInput()),
			helper.AddContext(Token),
		)

		// Perform assertions based on your expected results
		assert.NotEmpty(t, r.CreateExerciseResponse)
		//fmt.Printf("%+v\n", r.CreateExerciseResponse)
		assert.NotEmpty(t, r.CreateExerciseResponse.CreateExercise.ID)
		assert.Equal(t, requests.ClassID, r.CreateExerciseResponse.CreateExercise.ClassID)
		assert.Equal(t, requests.ModuleID, r.CreateExerciseResponse.CreateExercise.ModuleID)
		assert.Equal(t, requests.Name, r.CreateExerciseResponse.CreateExercise.Name)
		assert.Equal(t, requests.Question, r.CreateExerciseResponse.CreateExercise.Question)
		assert.Equal(t, requests.Answers[0]["value"], r.CreateExerciseResponse.CreateExercise.Answers[0].Value)
		assert.Equal(t, requests.Answers[0]["correct"], r.CreateExerciseResponse.CreateExercise.Answers[0].Correct)
		assert.Equal(t, requests.Answers[1]["value"], r.CreateExerciseResponse.CreateExercise.Answers[1].Value)
		assert.Equal(t, requests.Answers[1]["correct"], r.CreateExerciseResponse.CreateExercise.Answers[1].Correct)
		assert.Equal(t, requests.Difficulty, r.CreateExerciseResponse.CreateExercise.Difficulty)
	})
}

func TestResolver_CreateExercise_InvalidClassID(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise_InvalidClassID")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid ClassID", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidClassId()),
			helper.AddContext(Token),
		)
		if err != nil {
			assert.NotNil(t, err)
		}

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.InvalidClassIDResponseError, errorResponse[0].Message)
	})
}

func TestResolver_CreateExercise_InvalidModuleID(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise_InvalidModuleID")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid ModuleID", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidModuleId()),
			helper.AddContext(Token),
		)
		if err != nil {
			assert.NotNil(t, err)
		}

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.InvalidModuleIDResponseError, errorResponse[0].Message)
	})
}

func TestResolver_CreateExercise_InvalidName(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise_InvalidName")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid Name", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidName()),
			helper.AddContext(Token),
		)
		if err != nil {
			assert.NotNil(t, err)
		}

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.InvalidNameResponseError, errorResponse[0].Message)
	})
}

func TestResolver_CreateExercise_InvalidQuestion(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise_InvalidQuestion")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid Question", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInputInvalidQuestion()),
			helper.AddContext(Token),
		)
		if err != nil {
			assert.NotNil(t, err)
		}

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.InvalidQuestionResponseError, errorResponse[0].Message)
	})
}

func TestResolver_CreateExercise_NoCorrectAnswers(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise_NoCorrectAnswers")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid Answers", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInputNoCorrectAnswers()),
			helper.AddContext(Token),
		)
		if err != nil {
			assert.NotNil(t, err)
		}

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.NoCorrectAnswersResponseError, errorResponse[0].Message)
	})
}

func TestResolver_CreateExercise_NoAnswers(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise_NoAnswers")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid Answers", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInputNoAnswers()),
			helper.AddContext(Token),
		)
		if err != nil {
			assert.NotNil(t, err)
		}

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.NoAnswersResponseError, errorResponse[0].Message)
	})
}

func TestResolver_CreateExercise_NoIncorrectAnswer(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExercise_NoIncorrectAnswer")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid Answers", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInputNoIncorrectAnswers()),
			helper.AddContext(Token),
		)
		if err != nil {
			assert.NotNil(t, err)
		}

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		if err2 != nil {
			assert.Nil(t, err2)
		}

		assert.Equal(t, r.NoIncorrectAnswersResponseError, errorResponse[0].Message)
	})
}