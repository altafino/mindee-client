# Mindee Invoice API Go Client

A Golang client for the Mindee API, supporting both V1 (legacy) and V2 APIs for extracting invoice data from JPEG, PNG, WEBP, HEIC, TIFF image or PDF files.

## Installation

To use this library in your Go project, run the following command:

```bash
go get github.com/altafino/mindee-client
```

## API Versions

This library supports both Mindee API V1 and V2:

### V1 API (Legacy)
- **Status**: Maintained for backward compatibility
- **Endpoint**: `api.mindee.net/v1/...`
- **Processing**: Synchronous
- **Free tier**: Ends September 15, 2025
- **Use case**: Existing integrations, simple synchronous workflows

### V2 API (Recommended)
- **Status**: Current, actively developed
- **Endpoint**: `api-v2.mindee.net/v2/...`
- **Processing**: Asynchronous with polling
- **Features**: Advanced AI, custom models, RAG support, webhooks
- **Use case**: New integrations, advanced features, custom models

**Note**: V1 and V2 use different API keys and are not interchangeable.

## Usage

Import the `mindee-client` library in your Go code:

```go
import "github.com/altafino/mindee-client"
```

### V1 API Usage (Legacy)

V1 provides simple synchronous methods:

```go
package main

import (
    "fmt"
    "log"
    "os"
    mindee_client "github.com/altafino/mindee-client"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env file (API_KEY and TEST_FILE_PATH)
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    apiKey := os.Getenv("API_KEY")
    filePath := os.Getenv("TEST_FILE_PATH")

    // V1 API - Synchronous
    data, err := mindee_client.GetInvoiceDataForFilePath(filePath, apiKey)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Printf("Invoice data: %+v\n", data)
}
```

**V1 Methods:**
- `GetInvoiceDataForFilePath(filePath, apiKey string) (*models.InvoiceData, error)`
- `GetInvoiceDataForBase64(base64Content, apiKey string) (*models.InvoiceData, error)`

### V2 API Usage (Recommended)

V2 requires a model ID and uses asynchronous processing:

```go
package main

import (
    "fmt"
    "log"
    "os"
    mindee_client "github.com/altafino/mindee-client"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // V2 API requires a config with API key and model ID
    config := mindee_client.V2Config{
        APIKey:  os.Getenv("API_KEY_V2"),
        ModelID: os.Getenv("MODEL_ID"), // Get from Mindee dashboard
    }

    filePath := os.Getenv("TEST_FILE_PATH")

    // V2 API - Asynchronous with polling
    data, err := mindee_client.GetInvoiceDataForFilePathV2(filePath, config)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Printf("Invoice data: %+v\n", data)
}
```

**V2 Methods:**
- `GetInvoiceDataForFilePathV2(filePath string, config V2Config) (*models.InvoiceData, error)`
- `GetInvoiceDataForBase64V2(base64Content string, config V2Config) (*models.InvoiceData, error)`

**Important Note on V2 Response Mapping:**
The V2 API returns data in a structure that depends on your custom model configuration. The current implementation provides the V2 API infrastructure (authentication, enqueuing, polling) but does not automatically map V2 response fields to the V1 `InvoiceData` structure. You may need to:
1. Modify `convertV2ToV1Format()` in `client_v2.go` to map your specific model's fields
2. Or access the raw V2 response data directly by modifying the return type
3. Or create a new response model that matches your V2 model's schema

### Getting Your V2 Model ID

1. Log in to your [Mindee account](https://platform.mindee.com/)
2. Navigate to the Models section
3. Select or create an Invoice model
4. Copy the Model ID from the model details

### Environment Variables

Create a `.env` file with the following variables:

For V1:
```
API_KEY=your_v1_api_key_here
TEST_FILE_PATH=/path/to/your/test/invoice.pdf
```

For V2:
```
API_KEY_V2=your_v2_api_key_here
MODEL_ID=your_model_id_here
TEST_FILE_PATH=/path/to/your/test/invoice.pdf
```

### Example Project
Sample app demonstrating usage: https://github.com/altafino/testpdf

## Migration Guide: V1 to V2

If you're currently using V1 and want to migrate to V2, follow these steps:

### 1. Get V2 API Credentials
- Create a new API key in the [Mindee V2 platform](https://platform.mindee.com/)
- V1 API keys do not work with V2

### 2. Create or Select a Model
- In the Mindee dashboard, create a custom invoice model or use a pre-configured one
- Note your Model ID (required for all V2 API calls)

### 3. Update Your Code

**Before (V1):**
```go
data, err := mindee_client.GetInvoiceDataForFilePath(filePath, apiKey)
```

**After (V2):**
```go
config := mindee_client.V2Config{
    APIKey:  apiKeyV2,
    ModelID: modelID,
}
data, err := mindee_client.GetInvoiceDataForFilePathV2(filePath, config)
```

### 4. Key Differences

| Feature | V1 | V2 |
|---------|----|----|
| **Processing** | Synchronous | Asynchronous (with polling) |
| **API Keys** | V1 keys | V2 keys (separate) |
| **Model** | Pre-defined | Custom or pre-configured |
| **Response Time** | Immediate | 3-5 seconds (with polling) |
| **Customization** | Limited | Fully customizable models |
| **Features** | Basic extraction | RAG, webhooks, advanced AI |

### 5. Testing
Test your integration thoroughly as V2 may have different response structures depending on your model configuration.

## Testing

To run tests for this library, execute the following command in the root directory of the project:

```bash
go test
```

## License

This library is released under the [MIT License](LICENSE).
