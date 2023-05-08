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
    "github.com/altafino/mindee-client"
)

func main() {
    apiKey := "<your_api_key>"
    filePath := "path/to/invoice.pdf"

    // Get invoice data from a file path
    invoiceData, err := mindee_client.GetInvoiceDataForFilePath(filePath, apiKey)
    if err != nil {
	    fmt.Printf("Error getting invoice data: %v\n", err)
	    return
    }

    fmt.Printf("Invoice data: %+v\n", invoiceData)
}
```

## Testing

To run tests for this library, execute the following command in the root directory of the project:

```bash
go test
```

## License

This library is released under the [MIT License](LICENSE).
