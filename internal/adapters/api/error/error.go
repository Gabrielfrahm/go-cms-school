package error

import (
	"encoding/json"
	"net/http"
)

type ValidationErrorResponse struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

func SendValidationErrorResponse(w http.ResponseWriter, errorCode, errorMessage string, errorDetails map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	errResp := ValidationErrorResponse{
		Code:    errorCode,
		Message: errorMessage,
		Details: errorDetails,
	}
	json.NewEncoder(w).Encode(errResp)
}
