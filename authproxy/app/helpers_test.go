package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/slovak-egov/einvoice/authproxy/db"
)

func executeRequest(req *http.Request, authHeader string) *httptest.ResponseRecorder {
	req.Header.Set("Authorization", authHeader)
	rr := httptest.NewRecorder()
	a.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func createTestUser(t *testing.T) (*db.User, string) {
	t.Helper()

	id := "123"
	user, err := a.manager.CreateUser(id, "Frantisek")
	if err != nil {
		t.Error(err)
	}

	token := a.manager.CreateUserToken(id)

	return user, token
}

func cleanData(t *testing.T) func() {
	return func() {
		if _, err := a.manager.Db.Db.Model(&db.User{}).Where("TRUE").Delete(); err != nil {
			t.Error(err)
		}
		a.manager.Cache.FlushAll()
	}
}
