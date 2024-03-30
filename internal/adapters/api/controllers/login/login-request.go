package login

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

var validationMessages = map[string]string{
	"required": "The field is required",
	"email":    "The field should be email valid",
	"min":      "The field must have at least 6 characters",
}
