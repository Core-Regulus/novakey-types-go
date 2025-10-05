package novakeytypes

import "fmt"

type ValidationErrorResponse struct {
	Error       bool
	FailedField string
	Value       any
	Tag         string
}

type ErrorResponse struct {
	Error
}

type Error struct {
	Error            string `json:"error,omitempty"`
	Code             string `json:"code,omitempty"`
	Status           int    `json:"status,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}

func FormatErrorResponse(err Error) string {
  return fmt.Sprintf("status %d, error %s, code %s, description %s", err.Status, err.Error, err.Code, err.ErrorDescription)
}