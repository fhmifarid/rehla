package apierror

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type FieldViolation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Status     int
	Code       string
	Message    string
	Details    []FieldViolation
	Retryable  bool
	Underlying error
}

type envelope struct {
	Error payload `json:"error"`
}

type payload struct {
	Code      string           `json:"code"`
	Message   string           `json:"message"`
	RequestID string           `json:"request_id,omitempty"`
	Details   []FieldViolation `json:"details,omitempty"`
	Retryable bool             `json:"retryable"`
}

func (e *Error) Error() string {
	if e.Underlying != nil {
		return e.Code + ": " + e.Underlying.Error()
	}
	return e.Code + ": " + e.Message
}

func (e *Error) Unwrap() error {
	return e.Underlying
}

func Write(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		apiErr = &Error{
			Status:     http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    "An unexpected error occurred.",
			Underlying: err,
		}
	}
	if apiErr.Status == 0 {
		apiErr.Status = http.StatusInternalServerError
	}

	requestID := RequestID(r.Context())
	if apiErr.Status >= http.StatusInternalServerError {
		logger.Error("request failed", "request_id", requestID, "error", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.Status)
	if encodeErr := json.NewEncoder(w).Encode(envelope{Error: payload{
		Code:      apiErr.Code,
		Message:   apiErr.Message,
		RequestID: requestID,
		Details:   apiErr.Details,
		Retryable: apiErr.Retryable,
	}}); encodeErr != nil {
		logger.Error("encode error response", "request_id", requestID, "error", encodeErr)
	}
}
