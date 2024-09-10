package models

import (
	"encoding/json"
	"time"
)

type InvoiceData struct {
	APIRequest struct {
		Error      struct{}  `json:"error"`
		Resources  []string  `json:"resources"`
		Status     string    `json:"status"`
		StatusCode int       `json:"status_code"`
		URL        string    `json:"url"`
	} `json:"api_request"`
	Document struct {
		Annotations struct {
			Labels []string `json:"labels"`
		} `json:"annotations"`
		ID        string `json:"id"`
		Inference struct {
			Extras             map[string]interface{} `json:"extras"`
			FinishedAt         time.Time              `json:"finished_at"`
			IsRotationApplied  bool                   `json:"is_rotation_applied"`
			Pages              []Page                 `json:"pages"`
			Prediction         Prediction             `json:"prediction"`
			ProcessingTime     float64                `json:"processing_time"`
			Product            Product                `json:"product"`
			StartedAt          time.Time              `json:"started_at"`
		} `json:"inference"`
		NPages int    `json:"n_pages"`
		Name   string `json:"name"`
		Ocr    struct{} `json:"ocr"`
	} `json:"document"`
}

type Page struct {
	Extras      map[string]interface{} `json:"extras"`
	ID          int                    `json:"id"`
	Orientation struct {
		Value int `json:"value"`
	} `json:"orientation"`
	Prediction PagePrediction `json:"prediction"`
}

type PagePrediction struct {
	BillingAddress               FieldData              `json:"billing_address"`
	CustomerAddress              FieldData              `json:"customer_address"`
	CustomerCompanyRegistrations []CompanyRegistration  `json:"customer_company_registrations"`
	CustomerID                   FieldData              `json:"customer_id"`
	CustomerName                 NameFieldData          `json:"customer_name"`
	Date                         FieldData              `json:"date"`
	DocumentType                 DocumentType           `json:"document_type"`
	DueDate                      FieldData              `json:"due_date"`
	InvoiceNumber                FieldData              `json:"invoice_number"`
	LineItems                    []LineItem             `json:"line_items"`
	Locale                       Locale                 `json:"locale"`
	Orientation                  Orientation            `json:"orientation"`
	ReferenceNumbers             []FieldData            `json:"reference_numbers"`
	ShippingAddress              FieldData              `json:"shipping_address"`
	SupplierAddress              FieldData              `json:"supplier_address"`
	SupplierCompanyRegistrations []CompanyRegistration  `json:"supplier_company_registrations"`
	SupplierEmail                FieldData              `json:"supplier_email"`
	SupplierName                 NameFieldData          `json:"supplier_name"`
	SupplierPaymentDetails       []SupplierPaymentDetail `json:"supplier_payment_details"`
	SupplierPhoneNumber          FieldData              `json:"supplier_phone_number"`
	SupplierWebsite              FieldData              `json:"supplier_website"`
	Taxes                        []Tax                  `json:"taxes"`
	TotalAmount                  AmountFieldData        `json:"total_amount"`
	TotalNet                     AmountFieldData        `json:"total_net"`
	TotalTax                     AmountFieldData        `json:"total_tax"`
}

type Prediction struct {
	BillingAddress               PredictionFieldData              `json:"billing_address"`
	CustomerAddress              PredictionFieldData              `json:"customer_address"`
	CustomerCompanyRegistrations []PredictionCompanyRegistration  `json:"customer_company_registrations"`
	CustomerID                   PredictionFieldData              `json:"customer_id"`
	CustomerName                 PredictionNameFieldData          `json:"customer_name"`
	Date                         PredictionFieldData              `json:"date"`
	DocumentType                 DocumentType                     `json:"document_type"`
	DueDate                      PredictionFieldData              `json:"due_date"`
	InvoiceNumber                PredictionFieldData              `json:"invoice_number"`
	LineItems                    []PredictionLineItem             `json:"line_items"`
	Locale                       PredictionLocale                 `json:"locale"`
	ReferenceNumbers             []PredictionFieldData            `json:"reference_numbers"`
	ShippingAddress              PredictionFieldData              `json:"shipping_address"`
	SupplierAddress              PredictionFieldData              `json:"supplier_address"`
	SupplierCompanyRegistrations []PredictionCompanyRegistration  `json:"supplier_company_registrations"`
	SupplierEmail                PredictionFieldData              `json:"supplier_email"`
	SupplierName                 PredictionNameFieldData          `json:"supplier_name"`
	SupplierPaymentDetails       []PredictionSupplierPaymentDetail `json:"supplier_payment_details"`
	SupplierPhoneNumber          PredictionFieldData              `json:"supplier_phone_number"`
	SupplierWebsite              PredictionFieldData              `json:"supplier_website"`
	Taxes                        []PredictionTax                  `json:"taxes"`
	TotalAmount                  PredictionAmountFieldData        `json:"total_amount"`
	TotalNet                     PredictionAmountFieldData        `json:"total_net"`
	TotalTax                     PredictionAmountFieldData        `json:"total_tax"`
}

type FieldData struct {
	Confidence float64   `json:"confidence"`
	Polygon    [][]float64 `json:"polygon"`
	Value      string    `json:"value"`
}

type PredictionFieldData struct {
	Confidence float64   `json:"confidence"`
	PageID     int       `json:"page_id"`
	Polygon    [][]float64 `json:"polygon"`
	Value      string    `json:"value"`
}

type NameFieldData struct {
	Confidence float64   `json:"confidence"`
	Polygon    [][]float64 `json:"polygon"`
	RawValue   string    `json:"raw_value"`
	Value      string    `json:"value"`
}

type PredictionNameFieldData struct {
	Confidence float64   `json:"confidence"`
	PageID     int       `json:"page_id"`
	Polygon    [][]float64 `json:"polygon"`
	RawValue   string    `json:"raw_value"`
	Value      string    `json:"value"`
}

type AmountFieldData struct {
	Confidence float64   `json:"confidence"`
	Polygon    [][]float64 `json:"polygon"`
	Value      float64   `json:"value"`
}

type PredictionAmountFieldData struct {
	Confidence float64   `json:"confidence"`
	PageID     int       `json:"page_id"`
	Polygon    [][]float64 `json:"polygon"`
	Value      float64   `json:"value"`
}

type CompanyRegistration struct {
	Confidence float64   `json:"confidence"`
	Polygon    [][]float64 `json:"polygon"`
	Type       string    `json:"type"`
	Value      string    `json:"value"`
}

type PredictionCompanyRegistration struct {
	Confidence float64   `json:"confidence"`
	PageID     int       `json:"page_id"`
	Polygon    [][]float64 `json:"polygon"`
	Type       string    `json:"type"`
	Value      string    `json:"value"`
}

type LineItem struct {
	Confidence   float64   `json:"confidence"`
	Description  string    `json:"description"`
	Polygon      [][]float64 `json:"polygon"`
	ProductCode  string    `json:"product_code"`
	Quantity     *float64  `json:"quantity"`
	TaxAmount    *float64  `json:"tax_amount"`
	TaxRate      *float64  `json:"tax_rate"`
	TotalAmount  *float64  `json:"total_amount"`
	UnitMeasure  *string   `json:"unit_measure"`
	UnitPrice    *float64  `json:"unit_price"`
}

type PredictionLineItem struct {
	Confidence   float64   `json:"confidence"`
	Description  string    `json:"description"`
	PageID       int       `json:"page_id"`
	Polygon      [][]float64 `json:"polygon"`
	ProductCode  string    `json:"product_code"`
	Quantity     *float64  `json:"quantity"`
	TaxAmount    *float64  `json:"tax_amount"`
	TaxRate      *float64  `json:"tax_rate"`
	TotalAmount  *float64  `json:"total_amount"`
	UnitMeasure  *string   `json:"unit_measure"`
	UnitPrice    *float64  `json:"unit_price"`
}

type Locale struct {
	Confidence float64 `json:"confidence"`
	Currency   string  `json:"currency"`
	Language   string  `json:"language"`
}

type PredictionLocale struct {
	Confidence float64 `json:"confidence"`
	Currency   string  `json:"currency"`
	Language   string  `json:"language"`
}

type Orientation struct {
	Confidence float64 `json:"confidence"`
	Degrees    int     `json:"degrees"`
}

type DocumentType struct {
	Value string `json:"value"`
}

type SupplierPaymentDetail struct {
	AccountNumber  string    `json:"account_number"`
	Confidence     float64   `json:"confidence"`
	IBAN           string    `json:"iban"`
	Polygon        [][]float64 `json:"polygon"`
	RoutingNumber  string    `json:"routing_number"`
	SWIFT          string    `json:"swift"`
}

type PredictionSupplierPaymentDetail struct {
	AccountNumber  string    `json:"account_number"`
	Confidence     float64   `json:"confidence"`
	IBAN           string    `json:"iban"`
	PageID         int       `json:"page_id"`
	Polygon        [][]float64 `json:"polygon"`
	RoutingNumber  string    `json:"routing_number"`
	SWIFT          string    `json:"swift"`
}

type Tax struct {
	Base       float64   `json:"base"`
	Confidence float64   `json:"confidence"`
	Polygon    [][]float64 `json:"polygon"`
	Rate       *float64  `json:"rate"`
	Value      float64   `json:"value"`
}

type PredictionTax struct {
	Base       float64   `json:"base"`
	Confidence float64   `json:"confidence"`
	PageID     int       `json:"page_id"`
	Polygon    [][]float64 `json:"polygon"`
	Rate       *float64  `json:"rate"`
	Value      float64   `json:"value"`
}

type Product struct {
	Features []string `json:"features"`
	Name     string   `json:"name"`
	Version  string   `json:"version"`
}

type CustomTime struct {
	Time time.Time
}

const customTimeFormat = "2006-01-02T15:04:05.999999"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var timeString string
	if err := json.Unmarshal(b, &timeString); err != nil {
		return err
	}

	t, err := time.Parse(customTimeFormat, timeString)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}
