package logout

type LogoutRequest struct {
	Token string `json:"token" validate:"required"`
}

var validationMessages = map[string]string{
	"required": "The field is required",
}
