package error

import (
	"net/http"

	"github.com/go-playground/validator"
)

func ValidateRequest(req interface{}, w http.ResponseWriter, validationMessages map[string]string) bool {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			details := make(map[string]string)
			for _, fieldError := range validationErrors {
				if msg, found := validationMessages[fieldError.Tag()]; found {
					details[fieldError.Field()] = msg
				} else {
					details[fieldError.Field()] = fieldError.Tag()
				}
			}
			SendValidationErrorResponse(w, "VALIDATION_ERROR", "Validation error occurred", details)
			return false
		}
		http.Error(w, "Validation error", http.StatusBadRequest)
		return false
	}
	return true
}
