package auth

import (
	auth2 "Module/internal/auth"
	"fmt"
	"testing"
)

func TestIntrospectToken(t *testing.T) {
	// This works but make sure the inspectToken has a valid token before testing this so make a post
	// to postman or make it so here gets sends a request to keyclaok?
	// maybe at keycloak to the testing docker

	//auth := auth2.NewAuth()
	//
	//valid, err := auth.IntrospectToken("eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJCR2U4WnlEbGFTWWplakV1UmZxUUJHV3J2REhFdV85cDdGc1p0bFpyRXhvIn0.eyJleHAiOjE3MDA2ODMxMzEsImlhdCI6MTcwMDY4MjgzMSwianRpIjoiN2M2NDdlYzUtMGQ3OC00YzlkLTkyZmYtMjAyNzU1YjVjNGFiIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZWN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6IjNhNDJhMTc1LWFkMTItNGQwOC04ZjFiLTI3NDQ0MTE2ZjM0ZCIsInR5cCI6IkJlYXJlciIsImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiI4ZTdjODRmMS1jYmZlLTRjMDEtYTQyZi04Y2FjZjQ3NTcwOTgiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1lbnQtY2xpZW50Ijp7InJvbGVzIjpbInVwZGF0ZV9zY2hvb2wiLCJnZXRfZXhlcmNpc2VzIiwiZ2V0X2NsYXNzZXMiLCJkZWxldGVfbW9kdWxlIiwiZ2V0X3NjaG9vbHMiLCJkZWxldGVfZXhlcmNpc2UiLCJ1cGRhdGVfZXhlcmNpc2UiLCJnZXRfZXhlcmNpc2UiLCJkZWxldGVfbW9kdWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJkZWxldGVfZXhlcmNpc2VfYWxsIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJjcmVhdGVfbW9kdWxlIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwidXBkYXRlX2V4ZXJjaXNlX2FsbCIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlX2FsbCIsImRlbGV0ZV9zY2hvb2wiLCJ1cGRhdGVfY2xhc3NfYWxsIiwidXBkYXRlX21vZHVsZSIsImdldF9jbGFzcyIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZGVsZXRlX2NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiOGU3Yzg0ZjEtY2JmZS00YzAxLWE0MmYtOGNhY2Y0NzU3MDk4IiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFtaWx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.srO-rpUGifnrGJUpMhLXiYiAVcPjJrmYfY6rl_69ciOLkOQx_M8LWUuAlnYjyky9pbC3erhDVHos-_hsZawKaQPmwxDtRxgXdUFRov5gVhvX-zYpVeqzQ97hu5oE0aUcnB5ISVFicLT2tnkV950uFzJNn7PYFtcs6QmDBe1zYpDrFjGj4ILkPGIRzFBqhd5xfyf0VX6Y642PWNXCNIABUDgLvWKZDpLA2yCyXCqW8IVOZJ7wwzZw9MC1kd4HXxBURrnXvoaoTJ3z5_g4jrMpDrwtzl3UUkyLy9Bs6D6mcz99hci9NgcmQrsFt0Cai6syI4wRLkjLJhASSVCIcVfXXg")
	//
	//fmt.Println("token result:", valid)
	//
	//if err != nil {
	//	t.Errorf("DecodeToken failed: %v", err)
	//	return
	//}
	//
	//fmt.Println("token result:", valid)
}

func TestDecodeToken(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJCR2U4WnlEbGFTWW" +
		"plakV1UmZxUUJHV3J2REhFdV85cDdGc1p0bFpyRXhvIn0.eyJleHAiOjE3MDA2ODA0NDQsIml" +
		"hdCI6MTcwMDY4MDE0NCwianRpIjoiYmI0OTE1YzItNTk4MC00MzViLTg3ODEtZWUxNTdiZG" +
		"U1Y2Q0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4ODg4L3JlYWxtcy9jbG91ZC1wcm9qZ" +
		"WN0IiwiYXVkIjpbInVzZXItbWFuYWdlbWVudC1jbGllbnQiLCJhY2NvdW50Il0sInN1YiI6" +
		"IjNhNDJhMTc1LWFkMTItNGQwOC04ZjFiLTI3NDQ0MTE2ZjM0ZCIsInR5cCI6IkJlYXJlciI" +
		"sImF6cCI6ImxvZ2luLWNsaWVudCIsInNlc3Npb25fc3RhdGUiOiJhMDJkZTMzMy03YTI2LT" +
		"Q0NzktOWMwZi04YTdjNTBkMmY0NjMiLCJhY3IiOiIxIiwicmVhbG1fYWNjZXNzIjp7InJvb" +
		"GVzIjpbImRlZmF1bHQtcm9sZXMtY2xvdWQtcHJvamVjdCIsIm9mZmxpbmVfYWNjZXNzIiwi" +
		"dW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ1c2VyLW1hbmFnZW1" +
		"lbnQtY2xpZW50Ijp7InJvbGVzIjpbInVwZGF0ZV9zY2hvb2wiLCJnZXRfZXhlcmNpc2VzIi" +
		"wiZ2V0X2NsYXNzZXMiLCJkZWxldGVfbW9kdWxlIiwiZ2V0X3NjaG9vbHMiLCJkZWxldGVfZ" +
		"XhlcmNpc2UiLCJ1cGRhdGVfZXhlcmNpc2UiLCJnZXRfZXhlcmNpc2UiLCJkZWxldGVfbW9k" +
		"dWxlX2FsbCIsImNyZWF0ZV9leGVyY2lzZSIsImdldF9zY2hvb2wiLCJkZWxldGVfZXhlcmN" +
		"pc2VfYWxsIiwidXBkYXRlX3NjaG9vbF9hbGwiLCJkZWxldGVfY2xhc3MiLCJjcmVhdGVfbW" +
		"9kdWxlIiwiZ2V0X21vZHVsZSIsImdldF9tb2R1bGVzIiwidXBkYXRlX2V4ZXJjaXNlX2Fsb" +
		"CIsImNyZWF0ZV9jbGFzcyIsImNyZWF0ZV9zY2hvb2wiLCJ1cGRhdGVfbW9kdWxlX2FsbCIs" +
		"ImRlbGV0ZV9zY2hvb2wiLCJ1cGRhdGVfY2xhc3NfYWxsIiwidXBkYXRlX21vZHVsZSIsImd" +
		"ldF9jbGFzcyIsImRlbGV0ZV9zY2hvb2xfYWxsIiwidXBkYXRlX2NsYXNzIiwiZGVsZXRlX2" +
		"NsYXNzX2FsbCJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hb" +
		"mFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBw" +
		"cm9maWxlIiwic2lkIjoiYTAyZGUzMzMtN2EyNi00NDc5LTljMGYtOGE3YzUwZDJmNDYzIiw" +
		"iZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiY2hhZCBhZG1pbiIsInByZWZlcnJlZF" +
		"91c2VybmFtZSI6ImFkbWluQGFkbWluLmNvbSIsImdpdmVuX25hbWUiOiJjaGFkIiwiZmFta" +
		"Wx5X25hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIn0.waJwhv0b8GUI" +
		"zOWhkvCQHxVUwH-yOIkWt860xCde412AocH-ux6Nou9V-vUPe_RLjSPVh3EGO4BKtquhlNC" +
		"ndwVT6ZEJ5whg8H6EdD5hYSGYBkWMF6rFTsdZRm08_F2wPuv5h6a0uy1n78mW-3h8Sph7Qn" +
		"oc_Xzw9Qzx9rvxzcRycX3zQELy7zi77PW64IbE-CsdIVWyny4cWQBhffvEUVuuwOCYdqtJv" +
		"cu2gjOfnbUA0x1q_a40c43y3JSB2_JHIlvH8KtlAzrTz9pH8ctz_AYYRJhA3PifhSuL6tJN" +
		"WM1s96g5N0c1zKw0icIPUTfLAksFuRfIHpmvxTUGT1YRVg"

	auth := auth2.NewAuth()

	claims, err := auth.DecodeToken(token)
	if err != nil {
		t.Errorf("DecodeToken failed: %v", err)
	}

	resourceAccess, ok := claims["resource_access"].(map[string]interface{})
	if !ok {
		// Handle the case where 'resource_access' is not a map
		t.Errorf("Error: 'resource_access' is not a map")
		return
	}

	// Access the 'user-management-client' map
	userManagementClient, ok := resourceAccess["user-management-client"].(map[string]interface{})
	if !ok {
		// Handle the case where 'user-management-client' is not a map
		t.Errorf("Error: 'user-management-client' is not a map")
		return
	}

	// Access the 'roles' array within 'user-management-client'
	roles, ok := userManagementClient["roles"].([]interface{})
	if !ok {
		// Handle the case where 'roles' is not an array
		t.Errorf("Error: 'roles' is not an array")
		return
	}

	test := hasRole(roles, "update_school")

	if test == false {
		t.Errorf("Error: role doesnt exist in here")
	}

	fmt.Println("Roles:", test)
	fmt.Println("Roles:", roles)
}

func hasRole(roles []interface{}, targetRole string) bool {
	for _, role := range roles {
		if role == targetRole {
			return true
		}
	}
	return false
}
