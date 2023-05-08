package mindee_client_test

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	mindee_client "github.com/altafino/mindee-client"
)

// How to run the tests:
// 1. Set the apiKey and testFilePath constants below
// 2. Run `go test ./...` from the root of the project

const (
	apiKey       = "7eed5d3f8e8254eef68fe31a8da7d2ac"
	testFilePath = "HORISENAG_Inv_GMS-2200232_31.01.2022_000005668.pdf"
)

func TestGetInvoiceDataForFilePath(t *testing.T) {
	data, err := mindee_client.GetInvoiceDataForFilePath(testFilePath, apiKey)
	if err != nil {
		t.Fatalf("GetInvoiceDataForFilePath failed: %v", err)
	}

	if data == nil {
		t.Fatal("GetInvoiceDataForFilePath returned nil data")
	} else {
		t.Logf("GetInvoiceDataForFilePath returned data: %+v", data)
	}

	// Add more checks to validate the structure of the returned data, e.g.:
	// if data.FieldName == "" {
	//     t.Fatal("GetInvoiceDataForFilePath returned empty FieldName")
	// }
}

func TestGetInvoiceDataForBase64(t *testing.T) {
	file, err := os.Open(testFilePath)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	base64Content := base64.StdEncoding.EncodeToString(fileContents)

	data, err := mindee_client.GetInvoiceDataForBase64(base64Content, apiKey)
	if err != nil {
		t.Fatalf("GetInvoiceDataForBase64 failed: %v", err)
	}

	if data == nil {
		t.Fatal("GetInvoiceDataForBase64 returned nil data")
	} else {
		t.Logf("GetInvoiceDataForBase64 returned data: %+v", data)
	}

	// Add more checks to validate the structure of the returned data, e.g.:
	// if data.FieldName == "" {
	//     t.Fatal("GetInvoiceDataForBase64 returned empty FieldName")
	// }
}
