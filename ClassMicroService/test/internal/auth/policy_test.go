package auth

import (
	"example/internal/auth"
	"example/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var adminToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJnWWlqam1Zd3Z5a2t3WUNlZUtpVzV3amxVM215dmVoNTRZSHlVZFc5MUFzIn0.eyJleHAiOjE3MDExMDgyMzEsImlhdCI6MTcwMTEwNzkzMSwianRpIjoiNzYxYTQwZjktNTMzMS00Mzc4LWI5OTktZjhjNWM3MGRkYWEzIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjQwN2VjMjNkLWM2ZjQtNDhkYi05YjFlLWZhN2Q3MDBmMjg2NiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiIwMWZhNjNkZi0wNDJmLTRmNTMtYmYzZi03NDNkYjFjMmY0MjYiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbInVwZGF0ZV9zY2hvb2wiLCJnZXRfY2xhc3NlcyIsImdldF9leGVyY2lzZXMiLCJkZWxldGVfbW9kdWxlIiwiZ2V0X3NjaG9vbHMiLCJkZWxldGVfZXhlcmNpc2UiLCJ1cGRhdGVfZXhlcmNpc2UiLCJnZXRfZXhlcmNpc2UiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJkZWxldGVfZXhlcmNpc2VfYWxsIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsInVwZGF0ZV9leGVyY2lzZV9hbGwiLCJnZXRfbW9kdWxlcyIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImRlbGV0ZV9zY2hvb2wiLCJ1cGRhdGVfY2xhc3NfYWxsIiwidXBkYXRlX21vZHVsZSIsImdldF9jbGFzcyIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiMDFmYTYzZGYtMDQyZi00ZjUzLWJmM2YtNzQzZGIxYzJmNDI2IiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.F4WBE9C3Ct17v5broRGPO92YR-lt9CzLprCnrOe4jWIMcMYyjHSBoLC-oQ7GHSoe1MjXe02CWRP98IZqQ5TPhF7nCliYs5qhn2vZRtlLa-QsjrTF2kZ1F_uEdXVekhVSKIRRFwoH8y2KxkaR3SSQ4bXOtJe8UJQs1AvzHPPeVDmRgfQcCZDNwdQTGI9Sb-8-C_dLXmU6W2ORJN1GmKikn9in4IS2kZ6KEiW6qNqOOllNlSQMZdtLQXf8BlymGf6s8z9j1itpg4iVljKeV8X76A8EHy-xQ98ESB188OVOxFHYReT82xOp5pusRjvMf3K71t20jPcOtUj-GiTZpEy9pQ"
var teacherToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJnWWlqam1Zd3Z5a2t3WUNlZUtpVzV3amxVM215dmVoNTRZSHlVZFc5MUFzIn0.eyJleHAiOjE3MDExMDgyNjAsImlhdCI6MTcwMTEwNzk2MCwianRpIjoiYjgzMGNjNGUtNjRhNC00YWYxLWEyYTYtN2JiNDQwMDAwNzYxIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjQwMTkyN2ZkLTEzYWUtNGUwYS1hZWM5LWJlNjRiOWM5Mzg4MyIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiI5N2QxNDAyOS03OGZiLTRlMTQtYWEyYy03MmUyY2MzY2M0YTUiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbInVwZGF0ZV9zY2hvb2wiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwiY3JlYXRlX2NsYXNzIiwiY3JlYXRlX3NjaG9vbCIsImRlbGV0ZV9zY2hvb2wiLCJnZXRfY2xhc3NlcyIsImdldF9leGVyY2lzZXMiLCJkZWxldGVfbW9kdWxlIiwiZ2V0X3NjaG9vbHMiLCJkZWxldGVfZXhlcmNpc2UiLCJ1cGRhdGVfZXhlcmNpc2UiLCJnZXRfZXhlcmNpc2UiLCJjcmVhdGVfZXhlcmNpc2UiLCJnZXRfc2Nob29sIiwidXBkYXRlX21vZHVsZSIsImdldF9jbGFzcyIsImRlbGV0ZV9jbGFzcyIsInVwZGF0ZV9jbGFzcyJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiOTdkMTQwMjktNzhmYi00ZTE0LWFhMmMtNzJlMmNjM2NjNGE1IiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiQnJhbSBUZXJsb3V3IiwicHJlZmVycmVkX3VzZXJuYW1lIjoiYnJhbUB0ZWFjaGVyLmNvbSIsImdpdmVuX25hbWUiOiJCcmFtIiwiZmFtaWx5X25hbWUiOiJUZXJsb3V3IiwiZW1haWwiOiJicmFtQHRlYWNoZXIuY29tIn0.dqj5629HUvdoZD8hsguhOph5wChDHgHIMhcFwP4vf84xBAWi43OJ8cvogJYXiEfW_GgdPQaoVgb1j5Qrxj5YIAvd19oEQ1ZfAXUzS6n4m5Z992eB3IM-1okrYA5pahINXDq6RihFD_e_eVijC4Fq61t-IicqJC5CW5o5XRnLzMgYf6DaQBp6UbzrR3KzX-oAV4zjH8bvs8RiaHNTPLrBJ6iao-7pQebsGAZWi0t1hNqgJ0vPxBTZRAbnolaEDjb7fKZ_vOmSASEbszE_rsWeZhMLFocVeqSO4hfViaIsirT_W2TgFXfeW05utkHvR76S0rL6YcUHeIGZj5F1vUKnZA"
var studentToken = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJnWWlqam1Zd3Z5a2t3WUNlZUtpVzV3amxVM215dmVoNTRZSHlVZFc5MUFzIn0.eyJleHAiOjE3MDExMDgyODgsImlhdCI6MTcwMTEwNzk4OCwianRpIjoiZmEwZTRjOWItMGNkMC00NWRlLWFmMTEtYTEwNDIwMWIxN2UzIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6ImE3YTA5OGU3LThkOTgtNGU2NS1iZGY1LWQxODljNzM2MjE3NSIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiJjNWYzZGIyOC02NTRlLTQ1Y2ItYjBmNy02MmY2ZmZmZjk4ZDIiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbImdldF9leGVyY2lzZSIsImdldF9tb2R1bGUiLCJnZXRfbW9kdWxlcyIsImdldF9zY2hvb2wiLCJnZXRfY2xhc3MiLCJnZXRfY2xhc3NlcyIsImdldF9leGVyY2lzZXMiLCJnZXRfc2Nob29scyJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiYzVmM2RiMjgtNjU0ZS00NWNiLWIwZjctNjJmNmZmZmY5OGQyIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiTWVybGlqbiBCdXNjaCIsInByZWZlcnJlZF91c2VybmFtZSI6Im1lcmxpam5Ac3R1ZGVudC5jb20iLCJnaXZlbl9uYW1lIjoiTWVybGlqbiIsImZhbWlseV9uYW1lIjoiQnVzY2giLCJlbWFpbCI6Im1lcmxpam5Ac3R1ZGVudC5jb20ifQ.V4JbiZCvC-V6myZ8gfQL3d3Po4ZY81qZRkY6rdyypkNZ-jvZrBXMs73BxWuVo-En5xF1izL_lEMyFhYQ7c6kCtRlzyiZgwwIUE2rXe_CM3mKWOmSE5Iba-MReJsz1y7BOKMcU79OCeRjvhH44-6oz2Dy_LlWBIme0wqXWl_ypNfZ18-BNX3UlCN1BK-FNVRfzsoiH7xhYrn6LbWcOsfqmoz3A-lTVqjrTlutG6V58TDuWiCF2trtzgJLLx6IBVo_yCV3kbOTB-cutV4_AKyVp4weQORbMlWaq3hJcAyMG59E6gpWO3EhRBXoo4ubdSDSq_BsbMyjFgM55UUz6RmPQw"

func TestCreateClassWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.CreateClass(adminToken)

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateClassWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.CreateClass(teacherToken)

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateClassWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.CreateClass(studentToken)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid permissions for this action", err.Error())
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUpdateClassWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	_, err := classPolicy.UpdateClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUpdateClassWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.UpdateClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestUpdateClassWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.UpdateClass(studentToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid permissions for this action", err.Error())
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestDeleteClassWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	_, err := classPolicy.DeleteClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestDeleteClassWithTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.DeleteClass(teacherToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestDeleteClassWithTeacherTokenAndWrongUUID(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	sub := "401927fd-13ae-4e0a-aec9-be64b9c93883"
	mockClassWithValidUUID := mocks.MockClass
	mockClassWithValidUUID.MadeBy = sub

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mockClassWithValidUUID, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.DeleteClass(teacherToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestDeleteClassWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	_, err := classPolicy.DeleteClass(studentToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid permissions for this action", err.Error())
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetClassWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	_, err := classPolicy.GetClass(adminToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetClassWithStudentTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	_, err := classPolicy.GetClass(teacherToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetClassWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	mockRepo.
		On(
			"GetClassByID",
			"3a3bd756-6353-4e29-8aba-5b3531bdb9ed").
		Return(&mocks.MockClass, nil)

	_, err := classPolicy.GetClass(studentToken, "3a3bd756-6353-4e29-8aba-5b3531bdb9ed")

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestListClasssWithAdminToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	err := classPolicy.ListClasses(adminToken)

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestListClasssWithStudentTeacherToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	err := classPolicy.ListClasses(teacherToken)

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestListClasssWithStudentToken(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	mockToken := new(mocks.MockToken)

	classPolicy := auth.Policy{
		ClassRepository: mockRepo,
		Token:           mockToken,
	}

	mockToken.
		On("IntrospectToken", mock.AnythingOfType("string")).
		Return(true, nil)

	err := classPolicy.ListClasses(studentToken)

	assert.Nil(t, err)
	mockToken.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
