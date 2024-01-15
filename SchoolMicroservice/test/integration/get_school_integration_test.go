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

func TestResolver_GetSchoolWithAdminToken(t *testing.T) {
	fmt.Println("\nRunning TestResolver_GetSchoolWithAdminToken")
	c := helper.CreateClient()

	t.Run("Get school with admin token", func(t *testing.T) {
		c.MustPost(
			requests.GetSchoolQuery,
			&r.GetSchoolResponse,
			client.Var("getSchoolId", "09d6be4b-da77-4be0-9094-445e1a5e639a"),
			helper.AddContext(Token),
		)

		assert.NotEmpty(t, r.GetSchoolResponse)
		assert.NotEmpty(t, r.GetSchoolResponse.GetSchool.ID)
		assert.Equal(t, requests.Name, r.GetSchoolResponse.GetSchool.Name)
		assert.Equal(t, requests.Location, r.GetSchoolResponse.GetSchool.Location)
	})
}

//func TestResolver_GetSchoolWithStudentToken(t *testing.T) {
//	fmt.Println("\nRunning TestResolver_GetSchoolWithStudentToken")
//	c := helper.CreateClient()
//
//	t.Run("Get school with student token", func(t *testing.T) {
//		err := c.Post(
//			requests.GetSchoolQuery,
//			&r.GetSchoolResponse,
//			client.Var("getSchoolId", "09d6be4b-da77-4be0-9094-445e1a5e639a"),
//			helper.AddContext(StudentToken),
//		)
//
//		assert.NotNil(t, err)
//
//		var errorResponse []r.ErrorType
//		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
//		assert.Nil(t, err2)
//
//		assert.Equal(t, r.InvalidPermissionResponseError, errorResponse[0].Message)
//	})
//}

func TestResolver_GetSchoolWithNonExistingSchool(t *testing.T) {
	fmt.Println("\nRunning TestResolver_GetSchoolWithNonExistingSchool")
	c := helper.CreateClient()

	t.Run("Get school with admin token", func(t *testing.T) {
		err := c.Post(
			requests.GetSchoolQuery,
			&r.GetSchoolResponse,
			client.Var("getSchoolId", "09d6be4b-da77-4be0-9094-445e1a5e639c"),
			helper.AddContext(Token),
		)

		assert.NotNil(t, err)

		var errorResponse []r.ErrorType
		err2 := json.NewDecoder(strings.NewReader(err.Error())).Decode(&errorResponse)
		assert.Nil(t, err2)

		assert.Equal(t, r.SchoolNotFoundResponseError, errorResponse[0].Message)
	})
}
