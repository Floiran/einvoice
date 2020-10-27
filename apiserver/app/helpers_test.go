package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/slovak-egov/einvoice/apiserver/db"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func createTestInvoice(t *testing.T) int {
	invoice := &db.Invoice{
		Sender: "sender",
		Receiver: "receiver",
		Format: db.UblFormat,
		Price: 1,
	}

	if err := a.manager.Db.CreateInvoice(invoice); err != nil {
		t.Fatal(err)
	}

	return invoice.Id
}
