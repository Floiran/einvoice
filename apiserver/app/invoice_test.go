package app_test

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
	query          string
	responseLength int
}{
	{"", 1},
	{"?format=d16b", 0},
}

func TestGetInvoices(t *testing.T) {
	// Fill DB
	t.Cleanup(cleanDb(t, a))
	createTestInvoice(t)

	// Run tests
	for _, tt := range flagtests {
		t.Run(tt.query, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/invoices"+tt.query, nil)
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
	t.Cleanup(cleanDb(t, a))
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
	req, _ = http.NewRequest("GET", fmt.Sprintf("/invoices/%d", id+1), nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

type testAttachment struct {
	fileName string
	data     string
}

var createInvoiceTest = []struct {
	name        string
	attachments []testAttachment
}{
	{"without attachments", []testAttachment{}},
	{"with attachments", []testAttachment{
		{"at1.pdf", "attachment data1"},
		{"at2.pdf", "attachment data2"},
	}},
}

func TestCreateInvoice(t *testing.T) {
	for _, tt := range createInvoiceTest {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(cleanDb(t, a))

			var requestBody bytes.Buffer
			multipartWriter := multipart.NewWriter(&requestBody)
			formatWriter, _ := multipartWriter.CreateFormField("format")
			formatWriter.Write([]byte(db.UblFormat))

			dataWriter, _ := multipartWriter.CreateFormField("data")
			invoice, _ := ioutil.ReadFile("../../xml/ubl21/example/ubl21_invoice.xml")
			dataWriter.Write(invoice)

			for i, at := range tt.attachments {
				attachmentWriter, _ := multipartWriter.CreateFormFile("attachment"+fmt.Sprint(i), at.fileName)
				attachmentWriter.Write([]byte(at.data))
			}

			multipartWriter.Close()

			req, _ := http.NewRequest("POST", "/invoices", &requestBody)
			req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

			response := executeRequest(req)
			checkResponseCode(t, http.StatusCreated, response.Code)

			var createdResponse db.Invoice
			json.Unmarshal(response.Body.Bytes(), &createdResponse)
			expectedAttachments := []db.Attachment{}
			for i, at := range tt.attachments {
				expectedAttachments = append(expectedAttachments,
					db.Attachment{
						InvoiceId: 0,                                 // Default value, missing in response
						Id:        createdResponse.Attachments[i].Id, // No need to assert this param,
						Name:      at.fileName,
						CreatedAt: createdResponse.Attachments[i].CreatedAt, // No need to assert this param,
					})
			}

			expectedResponse := db.Invoice{
				Id:          createdResponse.Id,        // No need to assert this param,
				CreatedAt:   createdResponse.CreatedAt, // No need to assert this param
				Sender:      "Custom Cotter Pins",
				Receiver:    "North American Veeblefetzer",
				Price:       100,
				Format:      db.UblFormat,
				Attachments: expectedAttachments,
			}
			bytes1, _ := json.Marshal(expectedResponse)
			if !bytes.Equal(bytes1, response.Body.Bytes()) {
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
			if !bytes.Equal(invoice, response.Body.Bytes()) {
				t.Errorf("Response was incorrect. We expected %s", invoice)
			}

			// Try to get attachment through API
			for i, at := range tt.attachments {
				req, _ = http.NewRequest("GET", fmt.Sprintf("/attachments/%d", expectedAttachments[i].Id), nil)
				response = executeRequest(req)

				checkResponseCode(t, http.StatusOK, response.Code)
				if !bytes.Equal([]byte(at.data), response.Body.Bytes()) {
					t.Errorf("Response was incorrect. We expected %s", invoice)
				}
			}
		})
	}
}
