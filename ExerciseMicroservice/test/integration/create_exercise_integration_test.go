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

var Token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDUzMzQ5NDgsImlhdCI6MTcwNTMzNDA0OCwianRpIjoiNTdiYjQyMDEtNzNmNi00Zjg1LThjYmMtOGRkNjcxMzFhNTEyIiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiNmMxY2U0NDgtNjcwZi00N2IyLTgzZjctNGQ3NzFiMDE3NzViIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6Ijc0YTA3ZjQ2LWQ4MTEtNGI0Yy1hYzc5LTE2MTc1OTJkNjZkMiIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsibWFuYWdlLXVzZXJzIiwidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImZpbHRlcl9yZXN1bHRfc29mdERlbGV0ZSIsImZpbHRlcl9jbGFzc19kaWZmaWN1bHR5IiwiZmlsdGVyX2V4ZXJjaXNlX2RpZmZpY3VsdHkiLCJmaWx0ZXJfc2Nob29sX25hbWUiLCJ1cGRhdGVfcmVzdWx0IiwiZmlsdGVyX2V4ZXJjaXNlX21vZHVsZV9pZCIsImZpbHRlcl9tb2R1bGVfY2F0ZWdvcnkiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX2xvY2F0aW9uIiwiZmlsdGVyX21vZHVsZV9kaWZmaWN1bHR5IiwiZmlsdGVyX3Jlc3VsdF9tb2R1bGVfaWQiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwiZmlsdGVyX3NjaG9vbF9zb2Z0RGVsZXRlIiwiZGVsZXRlX3Jlc3VsdF9hbGwiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImZpbHRlcl9jbGFzc19tb2R1bGVfaWQiLCJjcmVhdGVfcmVzdWx0IiwiZ2V0X3Jlc3VsdF9hbGwiLCJmaWx0ZXJfbW9kdWxlX21hZGVfYnkiLCJsaXN0X3Jlc3VsdHNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX3F1ZXN0aW9uX3R5cGVfaWQiLCJ1cGRhdGVfY2xhc3NfYWxsIiwiZ2V0X2NsYXNzIiwiZ2V0X3NjaG9vbHNfYWxsIiwiZmlsdGVyX3Jlc3VsdF9leGVyY2lzZV9pZCIsImZpbHRlcl9jbGFzc19zb2Z0RGVsZXRlIiwidXBkYXRlX3Jlc3VsdF9hbGwiLCJvcGVuYWlfZ2VuZXJhdGVfcXVlc3Rpb25zX2Zyb21fZmlsZSIsImdldF9jbGFzc2VzX2FsbCIsInVwZGF0ZV9zY2hvb2wiLCJmaWx0ZXJfc2Nob29sX21hZGVfYnkiLCJnZXRfZXhlcmNpc2VzX2FsbCIsIm9wZW5haV9nZW5lcmF0ZV9leHBsYW5hdGlvbiIsImZpbHRlcl9jbGFzc19tYWRlX2J5IiwiZmlsdGVyX21vZHVsZV9zb2Z0RGVsZXRlIiwiZ2V0X2V4ZXJjaXNlcyIsImdldF9jbGFzc2VzIiwiZGVsZXRlX21vZHVsZSIsImdldF9zY2hvb2xzIiwiZGVsZXRlX2V4ZXJjaXNlIiwidXBkYXRlX2V4ZXJjaXNlIiwiZ2V0X2V4ZXJjaXNlIiwiZmlsdGVyX3Jlc3VsdF91c2VyX2lkIiwiZmlsdGVyX2V4ZXJjaXNlX25hbWUiLCJmaWx0ZXJfZXhlcmNpc2Vfc29mdERlbGV0ZSIsImRlbGV0ZV9leGVyY2lzZV9hbGwiLCJmaWx0ZXJfcmVzdWx0X2NsYXNzX2lkIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJkZWxldGVfcmVzdWx0IiwiY3JlYXRlX21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJjcmVhdGVfY2xhc3MiLCJjcmVhdGVfc2Nob29sIiwiZ2V0X21vZHVsZXNfYWxsIiwiZmlsdGVyX2V4ZXJjaXNlX2NsYXNzX2lkIiwibGlzdF9yZXN1bHRzIiwiZmlsdGVyX21vZHVsZV9zY2hvb2xfaWQiLCJmaWx0ZXJfY2xhc3NfbmFtZSIsImdldF9yZXN1bHQiLCJmaWx0ZXJfc2Nob29sX2hhc19vcGVuYWlfYWNjZXNzIiwib3BlbmFpX2dldF9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlIiwiZmlsdGVyX21vZHVsZV9uYW1lIiwiZmlsdGVyX21vZHVsZV9tYWRlX2J5X25hbWUiLCJmaWx0ZXJfZXhlcmNpc2VfbWFkZV9ieSIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZmlsdGVyX21vZHVsZV9wcml2YXRlIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiNzRhMDdmNDYtZDgxMS00YjRjLWFjNzktMTYxNzU5MmQ2NmQyIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.Sj2x-vO9NkMMyf1aYLF6phTD93sm8Z_UslDqjbuVVVIiRSFW4H_IAFtjYx4FRE3HTCA1HA00orh83U90CLNuIS-3w6snGwet1g3gbVtkmFZ5hvJW6Lb1wuEhJumNZbCoXMc2u_XCmk7PlfwpuNkTyytNhW0VprN30JywpK6ofS0sGU07l9n4qIvz9EHgdWbL6wWhWdqWcSWtLwK1FtABGlFt7RP0BC_ypjJ2T5lpYECdsGUWXOuQpq-W4CtBYK9S6mEBPPycXhj9zxRbbsNfOS5VUL-nZVBLCbsPHcf_PWT7vaWuQyVQE7168WiwhYrDIqOeYSjeQaWiRgf7iovRQA"
var StudentToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJuMzNESXZyQUZ0b1JGQ1d2UTMyOF85bXpjeU5JbXptZ1NSNFVKM05rdEdRIn0.eyJleHAiOjE3MDUzMzUwNTMsImlhdCI6MTcwNTMzNDE1MywianRpIjoiNzYxYjhlZDktMzdhYy00YzBhLWI4ZmEtYzFkMjU2NWQ3OTE2IiwiaXNzIjoiaHR0cHM6Ly9leGFtcGxlLWtleWNsb2FrLWJyYW10ZXJsb3V3LWRldi5hcHBzLm9jcDItaW5ob2xsYW5kLmpvcmFuLWJlcmdmZWxkLmNvbS9yZWFsbXMvY2xvdWQtcHJvamVjdCIsImF1ZCI6WyJyZWFsbS1tYW5hZ2VtZW50IiwidXNlci1tYW5hZ2VtZW50LWNsaWVudCIsImFjY291bnQiXSwic3ViIjoiM2U4MDVlOTctZmM3Ni00MzI0LWExOTktNDYzZjYwZTQzYjQ0IiwidHlwIjoiQmVhcmVyIiwiYXpwIjoibG9naW4tY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6IjgzN2IzYzhiLTU2NTMtNDY5Ny1iZTUwLTU3ZDIzOTc0NzdmZiIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsiZGVmYXVsdC1yb2xlcy1jbG91ZC1wcm9qZWN0Iiwib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7InJlYWxtLW1hbmFnZW1lbnQiOnsicm9sZXMiOlsidmlldy11c2VycyIsInF1ZXJ5LWdyb3VwcyIsInF1ZXJ5LXVzZXJzIl19LCJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImRlbGV0ZV9yZXN1bHQiLCJnZXRfbW9kdWxlIiwiZ2V0X21vZHVsZXMiLCJ1cGRhdGVfcmVzdWx0Iiwib3BlbmFpX2dlbmVyYXRlX2V4cGxhbmF0aW9uIiwiY3JlYXRlX3Jlc3VsdCIsImdldF9leGVyY2lzZXMiLCJmaWx0ZXJfZXhlcmNpc2VfY2xhc3NfaWQiLCJnZXRfY2xhc3NlcyIsImxpc3RfcmVzdWx0cyIsImZpbHRlcl9tb2R1bGVfc2Nob29sX2lkIiwiZ2V0X2V4ZXJjaXNlIiwiZmlsdGVyX3Jlc3VsdF91c2VyX2lkIiwiZ2V0X3Jlc3VsdCIsIm9wZW5haV9nZXRfc2Nob29sIiwiZmlsdGVyX3Jlc3VsdF9jbGFzc19pZCIsImZpbHRlcl9tb2R1bGVfbWFkZV9ieV9uYW1lIiwiZ2V0X2NsYXNzIiwiZmlsdGVyX3Jlc3VsdF9leGVyY2lzZV9pZCIsImZpbHRlcl9yZXN1bHRfbW9kdWxlX2lkIl19LCJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiI4MzdiM2M4Yi01NjUzLTQ2OTctYmU1MC01N2QyMzk3NDc3ZmYiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsIm5hbWUiOiJNZXJsaWpuIEJ1c2NoIiwicHJlZmVycmVkX3VzZXJuYW1lIjoibWVybGlqbkBzdHVkZW50LmNvbSIsImdpdmVuX25hbWUiOiJNZXJsaWpuIiwiZmFtaWx5X25hbWUiOiJCdXNjaCIsImVtYWlsIjoibWVybGlqbkBzdHVkZW50LmNvbSJ9.bVwvNayiMx3zboJXo4ZBVm5Ec_3-PpU-t1Ibl7Fq-5NzYFaAL9rY5O4gZvht3Gg_DFLsoQIAAheNtcXANfm8Nqa90yTUE8Ocn6K9hax40z5bOhFkjYakMwo4_PEUQQM0FtsaTXUwDtdIkOBJ78yce0qZaSS1tgfdigmTzLHtuZFp0t55NHYI_YC2JTNZJBnhoaMCPzpY9YGiVHgh_wLA_UyQm2wdbNp3CdPnV_cUNuEC-JFIgK59zxbviNgYZbJxLZDoONORMbc4s-5BxvTqv_4jCLV3LwNlL1yTtwbDwW-p71gHmp-U9thbU-0Zbf2jb6Kj9Jt0YifCzwA449l7xQ"

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
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

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
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

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
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

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
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

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
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

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
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

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
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.NoIncorrectAnswersResponseError, errorResponse[0].Message)
	})
}

func TestResolver_CreateExerciseWithStudentToken_InvalidPermission(t *testing.T) {
	fmt.Println("\nRunning TestResolver_CreateExerciseWithStudentToken_InvalidPermission")
	c := helper.CreateClient()

	t.Run("Create exercise with invalid permission", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		err := c.Post(
			requests.CreateExerciseMutation,
			&r.CreateExerciseResponse,
			client.Var("exerciseInput", requests.GenerateExerciseInput()),
			helper.AddContext(StudentToken),
		)
		assert.NotNil(t, err)

		// In your test, after getting the error
		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)

		assert.Nil(t, err2)

		assert.Equal(t, r.InvalidPermissionResponseError, errorResponse[0].Message)
	})
}
