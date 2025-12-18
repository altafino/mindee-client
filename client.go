package mindee_client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/TylerBrock/colorjson"

	"github.com/altafino/mindee-client/models"
)

const (
	apiURL = "https://api.mindee.net/v2/products/mindee/invoices/v4/predict"
)

func GetInvoiceDataForFilePath(filePath, apiKey string) (*models.InvoiceData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return getInvoiceData(fileContents, apiKey)
}

func GetInvoiceDataForBase64(base64Content, apiKey string) (*models.InvoiceData, error) {
	fileContents, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	return getInvoiceData(fileContents, apiKey)
}

func getInvoiceData(fileContents []byte, apiKey string) (*models.InvoiceData, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", "file")
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %v", err)
	}

	io.Copy(part, bytes.NewReader(fileContents))
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, apiURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Token "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	data := &models.InvoiceData{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	s, _ := colorjson.Marshal(data)
	fmt.Println(string(s))

	return data, nil
}
