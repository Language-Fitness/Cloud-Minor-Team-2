package requests

var (
	Name            = "FakeName"
	Location        = "FakeLocation"
	HasOpenaiAccess = true
	OpenAIKey       = "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"
)

var CreateSchoolMutation = `
    mutation CreateSchool($input: SchoolInput!) {
	  createSchool(input: $input) {
		id
		name
		location
		made_by
		has_openai_access
		openai_key
		join_code
	  }
	}`

var GetSchoolQuery = `
	query GetSchool($getSchoolId: ID!) {
	  getSchool(id: $getSchoolId) {
		id
		name
		location
		made_by
	  }
	}`

var UpdateSchoolMutation = `
	mutation Mutation($input: SchoolInput!, $updateSchoolId: ID!) {
	  updateSchool(input: $input, id: $updateSchoolId) {
		id
		name
		location
		made_by
		openai_key
		has_openai_access
		join_code
	  }
	}`

func GenerateSchoolInput() map[string]interface{} {
	return map[string]interface{}{
		"name":              Name,
		"location":          Location,
		"has_openai_access": HasOpenaiAccess,
		"openai_key":        OpenAIKey,
	}
}

func GenerateSchoolInputNoAccess() map[string]interface{} {
	return map[string]interface{}{
		"name":              Name,
		"location":          Location,
		"has_openai_access": false,
		"openai_key":        OpenAIKey,
	}
}

func GenerateSchoolInputUpdate() map[string]interface{} {
	return map[string]interface{}{
		"name":              "Updated",
		"location":          "Updated",
		"has_openai_access": HasOpenaiAccess,
		"openai_key":        OpenAIKey,
	}
}
