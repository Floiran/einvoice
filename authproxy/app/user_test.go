package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	// Fill DB
	t.Cleanup(cleanData(t))
	user, sessionToken := createTestUser(t)

	var flagtests = []struct {
		name           string
		header         string
		responseStatus int
	}{
		{"unauthorized", "", http.StatusUnauthorized},
		{"authorized", "Bearer "+sessionToken, http.StatusOK},
	}
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/users/me", nil)
			response := executeRequest(req, tt.header)

			checkResponseCode(t, tt.responseStatus, response.Code)
			if tt.responseStatus == http.StatusOK {
				var parsedResponse map[string]string
				json.Unmarshal(response.Body.Bytes(), &parsedResponse)

				expectedResponse := map[string]string{
					"id":                user.Id,
					"name":              user.Name,
					"serviceAccountKey": user.ServiceAccountKey,
					"email":             user.Email,
				}

				if !reflect.DeepEqual(parsedResponse, expectedResponse) {
					t.Errorf("User data should be %v, but is %v", expectedResponse, parsedResponse)
				}
			}
		})
	}
}

func TestPatchUser(t *testing.T) {
	// Fill DB
	t.Cleanup(cleanData(t))
	user, sessionToken := createTestUser(t)

	expectedUserResponse := map[string]string{
		"id":                user.Id,
		"name":              user.Name,
		"serviceAccountKey": user.ServiceAccountKey,
		"email":             user.Email,
	}

	var flagtests = []struct {
		name string
		requestBody map[string]string
	}{
		{"Set email", map[string]string{"email": "a@a.sk"}},
		{"Delete email", map[string]string{"email": ""}},
		{"Set service account key", map[string]string{"serviceAccountKey": "123"}},
		{"Set more props", map[string]string{"email": "b@b.sk", "serviceAccountKey": "1"}},
	}
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Errorf("Request body serialization failed with error %s", err)
			}
			req, _ := http.NewRequest("PATCH", "/users/me", bytes.NewReader(requestBody))
			response := executeRequest(req, "Bearer "+sessionToken)

			checkResponseCode(t, http.StatusOK, response.Code)

			var parsedResponse map[string]string
			json.Unmarshal(response.Body.Bytes(), &parsedResponse)

			for key, value := range tt.requestBody {
				expectedUserResponse[key] = value
			}

			if !reflect.DeepEqual(parsedResponse, expectedUserResponse) {
				t.Errorf("User data should be %v, but is %v", expectedUserResponse, parsedResponse)
			}
		})
	}
}
