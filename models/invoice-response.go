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
			// ... (other fields)
			Prediction struct {
				// ... (prediction fields)
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
