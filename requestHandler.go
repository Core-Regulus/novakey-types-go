package novakeytypes

type ValidationErrorResponse struct {
	Error       bool
	FailedField string
	Value       any
	Tag         string
}

type ErrorResponse struct {
	Error            string `json:"error,omitempty"`
	Code             string `json:"code,omitempty"`
	Status           int    `json:"status,omitempty"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}

