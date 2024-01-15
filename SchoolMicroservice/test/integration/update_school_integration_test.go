package integration

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/stretchr/testify/assert"
	"school/test/integration/helper"
	"school/test/integration/requests"
	r "school/test/integration/responses"
	"strings"
	"testing"
)

func TestResolver_UpdateSchool(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateSchool")
	c := helper.CreateClient()

	t.Run("Update school", func(t *testing.T) {

		c.MustPost(
			requests.UpdateSchoolMutation,
			&r.UpdateSchoolResponse,
			client.Var("updateSchoolId", "09d6be4b-da77-4be0-9094-445e1a5e639a"),
			client.Var("input", requests.GenerateSchoolInputUpdate()),
			helper.AddContext(Token),
		)

		assert.NotEmpty(t, r.UpdateSchoolResponse)
		assert.NotEmpty(t, r.UpdateSchoolResponse.UpdateSchool.ID)
		assert.Equal(t, "Updated", r.UpdateSchoolResponse.UpdateSchool.Name)
		assert.Equal(t, "Updated", r.UpdateSchoolResponse.UpdateSchool.Location)
	})
}

func TestResolver_UpdateSchool_NoAdminToken(t *testing.T) {
	fmt.Println("\nRunning TestResolver_UpdateSchool_NoAdminToken")
	c := helper.CreateClient()

	t.Run("Update school", func(t *testing.T) {

		err := c.Post(
			requests.UpdateSchoolMutation,
			&r.UpdateSchoolResponse,
			client.Var("updateSchoolId", "09d6be4b-da77-4be0-9094-445e1a5e639a"),
			client.Var("input", requests.GenerateSchoolInputUpdate()),
			helper.AddContext(StudentToken),
		)

		assert.NotNil(t, err)

		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.InvalidPermissionResponseError, errorResponse[0].Message)
	})
}
