package apperrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// AppError represents an application error with HTTP status code
type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Err        string `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != "" {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return errors.New(e.Err)
}

// New creates a new AppError
func New(message string, statusCode int) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

// NewWithErr creates a new AppError with an error with parameter
func NewWithErr(message string, statusCode int) func(err error) *AppError {
	return func(err error) *AppError {
		appErr := &AppError{
			Message:    message,
			StatusCode: statusCode,
			Err:        err.Error(),
		}
		//logs the error message
		PrintFormattedError(appErr)
		return appErr
	}
}

// Wrap wraps an existing error with AppError
func Wrap(err error, message string, statusCode int) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err.Error(),
	}
}

// ErrorResponse represents the JSON error response
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains error details for the response
type ErrorDetail struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// WriteError writes an error response to the HTTP response writer
func WriteError(w http.ResponseWriter, err error) {
	var appErr *AppError

	if e, ok := err.(*AppError); ok {
		appErr = e
	} else {
		appErr = Wrap(err, "An unexpected error occurred", http.StatusInternalServerError)
	}

	response := ErrorResponse{
		Error: ErrorDetail{
			Message:    appErr.Message,
			StatusCode: appErr.StatusCode,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func PrintFormattedError(appErr *AppError) {
	fmt.Printf("\nError: %s\n", appErr.Error())
	fmt.Printf("Status Code: %d\n", appErr.StatusCode)
	fmt.Printf("Message: %s\n", appErr.Message)
	fmt.Println("")
}
