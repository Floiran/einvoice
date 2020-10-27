package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
	"testing"

	"github.com/slovak-egov/einvoice/apiserver/db"
)

var flagtests = []struct {
	query  string
	responseLength int
}{
	{"", 1},
	{"?format=d16b", 0},
}
func TestGetInvoices(t *testing.T) {
	// Fill DB
	createTestInvoice(t)
	t.Cleanup(a.manager.Db.ClearData)

	// Run tests
	for _, tt := range flagtests {
		t.Run(tt.query, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/invoices" + tt.query, nil)
			response := executeRequest(req)

			checkResponseCode(t, http.StatusOK, response.Code)

			var parsedResponse []db.Invoice
			json.Unmarshal(response.Body.Bytes(), &parsedResponse)

			if len(parsedResponse) != tt.responseLength {
				t.Errorf("Expected an array of length %d. Got %s", tt.responseLength, response.Body.String())
			}
		})
	}
}

func TestGetInvoice(t *testing.T) {
	t.Cleanup(a.manager.Db.ClearData)
	id := createTestInvoice(t)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/invoices/%d", id), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	var parsedResponse db.Invoice
	json.Unmarshal(response.Body.Bytes(), &parsedResponse)
	if parsedResponse.Id != id {
		t.Errorf("Expected invoice with id %d. Got %d", id, parsedResponse.Id)
	}

	// Try to get nonexistent invoice
	req, _ = http.NewRequest("GET", fmt.Sprintf("/invoices/%d", id + 1), nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestCreateInvoice(t *testing.T) {
	t.Cleanup(a.manager.Db.ClearData)

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)
	formatWriter, _ := multipartWriter.CreateFormField("format")
	formatWriter.Write([]byte(db.UblFormat))

	dataWriter, _ := multipartWriter.CreateFormField("data")
	invoice, _ := ioutil.ReadFile("../../xml/ubl21/example/ubl21_invoice.xml")
	dataWriter.Write(invoice)
	multipartWriter.Close()

	req, _ := http.NewRequest("POST", "/invoices", &requestBody)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var createdResponse db.Invoice
	json.Unmarshal(response.Body.Bytes(), &createdResponse)
	expectedResponse := db.Invoice{
		Id: createdResponse.Id, // No need to assert this param,
		CreatedAt: createdResponse.CreatedAt, // No need to assert this param
		Sender: "Custom Cotter Pins",
		Receiver: "North American Veeblefetzer",
		Price: 100,
		Format: db.UblFormat,
		Attachments: []db.Attachment{},
	}
	if !reflect.DeepEqual(createdResponse, expectedResponse) {
		t.Errorf("Expected created response was %v. Got %v", expectedResponse, createdResponse)
	}

	// Try to get invoice metadata through API
	req, _ = http.NewRequest("GET", fmt.Sprintf("/invoices/%d", createdResponse.Id), nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	var getResponse db.Invoice
	json.Unmarshal(response.Body.Bytes(), &getResponse)
	if !reflect.DeepEqual(createdResponse, getResponse) {
		t.Errorf("Created response was %v. While GET request returned %v", createdResponse, getResponse)
	}

	// Try to get actual invoice through API
	req, _ = http.NewRequest("GET", fmt.Sprintf("/invoices/%d/detail", createdResponse.Id), nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if bytes.Compare(invoice, response.Body.Bytes()) != 0 {
		t.Errorf("Response was incorrect. We expected %s", invoice)
	}
}
