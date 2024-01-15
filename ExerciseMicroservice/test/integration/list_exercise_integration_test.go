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
	})
}
