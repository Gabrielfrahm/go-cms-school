package user

type Permission struct {
	Users    int `json:"users" validate:"required"`
	Classes  int `json:"classes" validate:"required"`
	Profiles int `json:"profiles" validate:"required"`
	Lessons  int `json:"lessons" validate:"required"`
}

type CreateUserRequest struct {
	Name       *string    `json:"name" validate:"omitempty"`
	Email      *string    `json:"email" validate:"omitempty,email"`
	Password   *string    `json:"password" validate:"omitempty,min=6"`
	TypeUser   string     `json:"type_user"  validate:"required,oneof=ADMIN ADMIN_SCHOOL TEACHER STUDENT"`
	ProfileId  string     `json:"profile_id" validate:"required"`
	Permission Permission `json:"permission" validate:"required,dive"`
}

var CreateUserValidationMessages = map[string]string{
	"required": "The field is required",
	"email":    "The field should be email valid",
	"min":      "The field must have at least 6 characters",
}

type ListAllUserRequest struct {
	Page     int     `json:"page" validate:"required"`
	PerPage  int     `json:"perPage" validate:"required"`
	Name     *string `json:"name" validate:"omitempty"`
	Email    *string `json:"email" validate:"omitempty"`
	TypeUser *string `json:"type_user"  validate:"omitempty,oneof=ADMIN ADMIN_SCHOOL TEACHER STUDENT"`
}

var ListAllUserValidationMessages = map[string]string{
	"required": "The field is required",
}
