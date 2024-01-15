package integration

import (
	"ExerciseMicroservice/test/integration/helper"
	"ExerciseMicroservice/test/integration/requests"
	r "ExerciseMicroservice/test/integration/responses"
	"github.com/99designs/gqlgen/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolver_ListExercise(t *testing.T) {
	c := helper.CreateClient()

	t.Run("List exercise with admin token", func(t *testing.T) {
		// Call the resolver via the client and modify the context via functional options
		c.MustPost(
			requests.ListExerciseQuery,
			&r.ListExerciseResponse,
			client.Var("filter", requests.GenerateExerciseFilterEmpty()),
			client.Var("paginator", requests.GenerateExercisePaginator()),
			helper.AddContext(Token),
		)

		// Perform assertions based on your expected results
		assert.NotEmpty(t, r.ListExerciseResponse)
		assert.NotEmpty(t, r.ListExerciseResponse.ListExercise)
		assert.Equal(t, requests.ClassID, r.ListExerciseResponse.ListExercise[0].ClassID)
		assert.Equal(t, requests.ModuleID, r.ListExerciseResponse.ListExercise[0].ModuleID)
		assert.Equal(t, requests.Name, r.ListExerciseResponse.ListExercise[0].Name)
		assert.Equal(t, requests.Question, r.ListExerciseResponse.ListExercise[0].Question)
		assert.Equal(t, requests.Answers[0]["value"], r.ListExerciseResponse.ListExercise[0].Answers[0].Value)
		assert.Equal(t, requests.Answers[0]["correct"], r.ListExerciseResponse.ListExercise[0].Answers[0].Correct)
		assert.Equal(t, requests.Answers[1]["value"], r.ListExerciseResponse.ListExercise[0].Answers[1].Value)
		assert.Equal(t, requests.Answers[1]["correct"], r.ListExerciseResponse.ListExercise[0].Answers[1].Correct)
		assert.Equal(t, requests.Difficulty, r.ListExerciseResponse.ListExercise[0].Difficulty)
	})
}
