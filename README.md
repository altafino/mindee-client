# Mindee Client

A simple Golang client for the Mindee API, allowing you to extract invoice data from JPEG, PNG, WEBP, HEIC, TIFF image or PDF files.

## Installation

To use this library in your Go project, run the following command:

```bash
go get github.com/altafino/mindee-client
```

## Usage

Import the `mindee-client` library in your Go code:

```go
import "github.com/altafino/mindee-client"
```

This library provides two methods for getting invoice data from files:

1. `GetInvoiceDataForFilePath(filePath, apiKey string) (*models.InvoiceData, error)` - for getting invoice data from a file path
2. `GetInvoiceDataForBase64(base64Content, apiKey string) (*models.InvoiceData, error)` - for getting invoice data from a base64-encoded file

### Example

```go
package main

import (
    "fmt"
    "os"
    "github.com/altafino/mindee-client"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env file (API_KEY and TEST_FILE_PATH)
    if err := godotenv.Load(); err != nil {
        fmt.Printf("Error loading .env file: %v\n", err)
        return
    }
}
```
### Example Project
I did a small sample app to see usage: https://github.com/altafino/testpdf


## Testing

To run tests for this library, execute the following command in the root directory of the project:

```bash
go test
```

## License

This library is released under the [MIT License](LICENSE).
