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
	"time"

	"github.com/TylerBrock/colorjson"

	"github.com/altafino/mindee-client/models"
)

const (
	apiV2BaseURL    = "https://api-v2.mindee.net"
	enqueueEndpoint = "/v2/inferences/enqueue"
	jobEndpoint     = "/v2/jobs/%s"
	pollInterval    = 1 * time.Second
	initialWait     = 3 * time.Second
	maxPollAttempts = 120 // 2 minutes max
)

// V2Config contains configuration for V2 API
type V2Config struct {
	APIKey  string
	ModelID string
}

// V2JobResponse represents the initial response from enqueue
type V2JobResponse struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	ModelID   string `json:"model_id"`
	CreatedAt string `json:"created_at"`
	ResultURL string `json:"result_url,omitempty"`
}

// V2InferenceResponse represents the final inference result
type V2InferenceResponse struct {
	ID        string                 `json:"id"`
	Status    string                 `json:"status"`
	ModelID   string                 `json:"model_id"`
	Document  map[string]interface{} `json:"document"`
	Result    map[string]interface{} `json:"result,omitempty"`
}

// GetInvoiceDataForFilePathV2 gets invoice data using V2 API with polling
func GetInvoiceDataForFilePathV2(filePath string, config V2Config) (*models.InvoiceData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return getInvoiceDataV2(fileContents, config)
}

// GetInvoiceDataForBase64V2 gets invoice data from base64 using V2 API
func GetInvoiceDataForBase64V2(base64Content string, config V2Config) (*models.InvoiceData, error) {
	fileContents, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	return getInvoiceDataV2(fileContents, config)
}

func getInvoiceDataV2(fileContents []byte, config V2Config) (*models.InvoiceData, error) {
	// Step 1: Enqueue the document
	jobID, err := enqueueDocument(fileContents, config)
	if err != nil {
		return nil, fmt.Errorf("failed to enqueue document: %v", err)
	}

	// Step 2: Wait initial period
	time.Sleep(initialWait)

	// Step 3: Poll for results
	inferenceResult, err := pollForResults(jobID, config.APIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get results: %v", err)
	}

	// Step 4: Convert V2 response to V1 format for backward compatibility
	invoiceData, err := convertV2ToV1Format(inferenceResult)
	if err != nil {
		return nil, fmt.Errorf("failed to convert response: %v", err)
	}

	s, _ := colorjson.Marshal(invoiceData)
	fmt.Println(string(s))

	return invoiceData, nil
}

func enqueueDocument(fileContents []byte, config V2Config) (string, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Add model_id field
	if err := writer.WriteField("model_id", config.ModelID); err != nil {
		return "", fmt.Errorf("failed to write model_id: %v", err)
	}

	// Add file
	part, err := writer.CreateFormFile("file", "document")
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	if _, err := io.Copy(part, bytes.NewReader(fileContents)); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	url := apiV2BaseURL + enqueueEndpoint
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Token "+config.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var jobResp V2JobResponse
	if err := json.NewDecoder(resp.Body).Decode(&jobResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return jobResp.ID, nil
}

func pollForResults(jobID, apiKey string) (*V2InferenceResponse, error) {
	url := fmt.Sprintf(apiV2BaseURL+jobEndpoint, jobID)
	client := &http.Client{}

	for attempt := 0; attempt < maxPollAttempts; attempt++ {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %v", err)
		}

		req.Header.Set("Authorization", "Token "+apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to send request: %v", err)
		}

		// If we get a redirect (302), follow it to get the final result
		if resp.StatusCode == http.StatusFound || resp.StatusCode == http.StatusSeeOther {
			location := resp.Header.Get("Location")
			resp.Body.Close()
			
			if location != "" {
				return getInferenceResult(location, apiKey)
			}
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			time.Sleep(pollInterval)
			continue
		}

		var jobResp V2JobResponse
		if err := json.NewDecoder(resp.Body).Decode(&jobResp); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}
		resp.Body.Close()

		// Check if result_url is available
		if jobResp.ResultURL != "" {
			return getInferenceResult(jobResp.ResultURL, apiKey)
		}

		// Check status
		if jobResp.Status == "failed" {
			return nil, fmt.Errorf("job failed")
		}

		if jobResp.Status == "completed" || jobResp.Status == "processed" {
			// Try to get inference result using job ID
			inferenceURL := fmt.Sprintf("%s/v2/inferences/%s", apiV2BaseURL, jobID)
			return getInferenceResult(inferenceURL, apiKey)
		}

		time.Sleep(pollInterval)
	}

	return nil, fmt.Errorf("polling timeout after %d attempts", maxPollAttempts)
}

func getInferenceResult(url, apiKey string) (*V2InferenceResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Token "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var inferenceResp V2InferenceResponse
	if err := json.NewDecoder(resp.Body).Decode(&inferenceResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &inferenceResp, nil
}

func convertV2ToV1Format(v2Resp *V2InferenceResponse) (*models.InvoiceData, error) {
	// This is a simplified conversion
	// The actual V2 response structure may vary based on the model
	// Users may need to customize this conversion based on their model's output
	
	// For now, return a basic structure indicating V2 is being used
	// In a real implementation, you would map the V2 response fields to V1 fields
	data := &models.InvoiceData{}
	
	// Note: The actual field mapping would depend on the V2 model configuration
	// This is a placeholder that indicates V2 API is working but doesn't do full conversion
	// Users should implement proper field mapping based on their model
	
	return data, nil
}
