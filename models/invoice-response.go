package models

import (
	"encoding/json"
	"time"
)

type InvoiceData struct {
	APIRequest struct {
		Error struct {
		} `json:"error"`
		Resources  []string `json:"resources"`
		Status     string   `json:"status"`
		StatusCode int      `json:"status_code"`
		URL        string   `json:"url"`
	} `json:"api_request"`
	Document struct {
		Annotations struct {
			Labels []string `json:"labels"`
		} `json:"annotations"`
		ID        string `json:"id"`
		Inference struct {
			Extras struct {
			} `json:"extras"`
			FinishedAt CustomTime `json:"finished_at"`
			Pages      []struct {
				Extras struct {
				} `json:"extras"`
				ID         int `json:"id"`
				Prediction struct {
					CustomerAddress struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"customer_address"`
					CustomerCompanyRegistrations []struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Type       string      `json:"type"`
						Value      string      `json:"value"`
					} `json:"customer_company_registrations"`
					CustomerName struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"customer_name"`
					Date struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"date"`
					DocumentType struct {
						Value string `json:"value"`
					} `json:"document_type"`
					DueDate struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"due_date"`
					InvoiceNumber struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"invoice_number"`
					LineItems []struct {
						Confidence  float64     `json:"confidence"`
						Description string      `json:"description"`
						Polygon     [][]float64 `json:"polygon"`
						ProductCode string      `json:"product_code"`
						Quantity    float64     `json:"quantity"`
						TaxAmount   float64     `json:"tax_amount"`
						TaxRate     float64     `json:"tax_rate"`
						TotalAmount float64     `json:"total_amount"`
						UnitPrice   float64     `json:"unit_price"`
					} `json:"line_items"`
					Locale struct {
						Confidence float64 `json:"confidence"`
						Currency   string  `json:"currency"`
						Language   string  `json:"language"`
					} `json:"locale"`
					Orientation struct {
						Confidence float64 `json:"confidence"`
						Degrees    int     `json:"degrees"`
					} `json:"orientation"`
					ReferenceNumbers []struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"reference_numbers"`
					SupplierAddress struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"supplier_address"`
					SupplierCompanyRegistrations []struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Type       string      `json:"type"`
						Value      string      `json:"value"`
					} `json:"supplier_company_registrations"`
					SupplierName struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      string      `json:"value"`
					} `json:"supplier_name"`
					SupplierPaymentDetails []struct {
						AccountNumber string      `json:"account_number"`
						Confidence    float64     `json:"confidence"`
						Iban          string      `json:"iban"`
						Polygon       [][]float64 `json:"polygon"`
						RoutingNumber string      `json:"routing_number"`
						Swift         string      `json:"swift"`
					} `json:"supplier_payment_details"`
					Taxes []struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Rate       float64     `json:"rate"`
						Value      float64     `json:"value"`
					} `json:"taxes"`
					TotalAmount struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      float64     `json:"value"`
					} `json:"total_amount"`
					TotalNet struct {
						Confidence float64     `json:"confidence"`
						Polygon    [][]float64 `json:"polygon"`
						Value      float64     `json:"value"`
					} `json:"total_net"`
				} `json:"prediction"`
			} `json:"pages"`
			Prediction struct {
				CustomerAddress struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"customer_address"`
				CustomerCompanyRegistrations []struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Type       interface{} `json:"type"`
					Value      interface{} `json:"value"`
				} `json:"customer_company_registrations"`
				CustomerName struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"customer_name"`
				Date struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"date"`
				DocumentType struct {
					Value string `json:"value"`
				} `json:"document_type"`
				DueDate struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"due_date"`
				InvoiceNumber struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"invoice_number"`
				LineItems []struct {
					Confidence  float64     `json:"confidence"`
					Description string      `json:"description"`
					PageID      int         `json:"page_id"`
					Polygon     [][]float64 `json:"polygon"`
					ProductCode string      `json:"product_code"`
					Quantity    float64     `json:"quantity"`
					TaxAmount   float64     `json:"tax_amount"`
					TaxRate     float64     `json:"tax_rate"`
					TotalAmount float64     `json:"total_amount"`
					UnitPrice   float64     `json:"unit_price"`
				} `json:"line_items"`
				Locale struct {
					Confidence float64 `json:"confidence"`
					Currency   string  `json:"currency"`
					Language   string  `json:"language"`
				} `json:"locale"`
				ReferenceNumbers []struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"reference_numbers"`
				SupplierAddress struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"supplier_address"`
				SupplierCompanyRegistrations []struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Type       string      `json:"type"`
					Value      string      `json:"value"`
				} `json:"supplier_company_registrations"`
				SupplierName struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      string      `json:"value"`
				} `json:"supplier_name"`
				SupplierPaymentDetails []struct {
					AccountNumber string      `json:"account_number"`
					Confidence    float64     `json:"confidence"`
					Iban          string      `json:"iban"`
					PageID        int         `json:"page_id"`
					Polygon       [][]float64 `json:"polygon"`
					RoutingNumber string      `json:"routing_number"`
					Swift         string      `json:"swift"`
				} `json:"supplier_payment_details"`
				Taxes []struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Rate       float64     `json:"rate"`
					Value      float64     `json:"value"`
				} `json:"taxes"`
				TotalAmount struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      float64     `json:"value"`
				} `json:"total_amount"`
				TotalNet struct {
					Confidence float64     `json:"confidence"`
					PageID     int         `json:"page_id"`
					Polygon    [][]float64 `json:"polygon"`
					Value      float64     `json:"value"`
				} `json:"total_net"`
			} `json:"prediction"`
			ProcessingTime float64 `json:"processing_time"`
			Product        struct {
				Features []string `json:"features"`
				Name     string   `json:"name"`
				Version  string   `json:"version"`
			} `json:"product"`
			StartedAt CustomTime `json:"started_at"`
		} `json:"inference"`
		NPages int    `json:"n_pages"`
		Name   string `json:"name"`
		Ocr    struct {
		} `json:"ocr"`
	} `json:"document"`
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
