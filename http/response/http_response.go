package response

import "net/http"

type (
	// ErrorResponse represents the structure of an error response.
	// It includes an HTTP status code and a message describing the error.
	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	// SuccessResponse represents the structure of a successful response.
	// It includes an HTTP status code and the data returned in the response.
	SuccessResponse struct {
		Code int `json:"code"`
		Data any `json:"data"`
	}
)

// StatusOK returns an HTTP 200 OK status and a SuccessResponse containing the provided data.
// Parameters:
//   - data: The data to include in the success response.
//
// Returns:
//   - int: The HTTP status code (200).
//   - SuccessResponse: The success response containing the data.
func StatusOK(data any) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Code: http.StatusOK,
		Data: data,
	}
}

// StatusCreated returns an HTTP 201 Created status and a SuccessResponse containing the provided message.
// Parameters:
//   - message: The message to include in the success response.
//
// Returns:
//   - int: The HTTP status code (201).
//   - SuccessResponse: The success response containing the message.
func StatusCreated(message string) (int, SuccessResponse) {
	return http.StatusCreated, SuccessResponse{
		Code: http.StatusCreated,
		Data: message,
	}
}

// StatusBadRequest returns an HTTP 400 Bad Request status and an ErrorResponse containing the provided message.
// Parameters:
//   - message: The error message to include in the response.
//
// Returns:
//   - int: The HTTP status code (400).
//   - ErrorResponse: The error response containing the message.
func StatusBadRequest(message string) (int, ErrorResponse) {
	return http.StatusBadRequest, ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

// StatusUnauthorized returns an HTTP 401 Unauthorized status and an ErrorResponse containing the provided message.
// Parameters:
//   - message: The error message to include in the response.
//
// Returns:
//   - int: The HTTP status code (401).
//   - ErrorResponse: The error response containing the message.
func StatusUnauthorized(message string) (int, ErrorResponse) {
	return http.StatusUnauthorized, ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

// StatusNotFound returns an HTTP 404 Not Found status and an ErrorResponse containing the provided message.
// Parameters:
//   - message: The error message to include in the response.
//
// Returns:
//   - int: The HTTP status code (404).
//   - ErrorResponse: The error response containing the message.
func StatusNotFound(message string) (int, ErrorResponse) {
	return http.StatusNotFound, ErrorResponse{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// StatusInternalServerError returns an HTTP 500 Internal Server Error status and an ErrorResponse containing the provided message.
// Parameters:
//   - message: The error message to include in the response.
//
// Returns:
//   - int: The HTTP status code (500).
//   - ErrorResponse: The error response containing the message.
func StatusInternalServerError(message string) (int, ErrorResponse) {
	return http.StatusInternalServerError, ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
