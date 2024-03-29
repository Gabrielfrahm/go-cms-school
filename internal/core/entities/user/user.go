package user

import (
	"time"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/permission"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/profile"
	"github.com/google/uuid"
)

type User struct {
	ID          string                `json:"id"`
	Name        *string               `json:"name"`
	Email       *string               `json:"email"`
	Password    *string               `json:"password"`
	Type_user   string                `json:"type_user"`
	Profile     profile.Profile       `json:"profile"`
	Permissions permission.Permission `json:"permissions,omitempty"`
	Created_at  *time.Time            `json:"created_at"`
	Updated_at  *time.Time            `json:"updated_at"`
	Deleted_at  *time.Time            `json:"deleted_at,omitempty"`
}

func NewUser(
	id *string,
	name *string,
	email *string,
	password *string,
	typeUser string,
	profile profile.Profile,
	permissions permission.Permission,
	createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time,
) *User {
	var newID string

	if id == nil {
		newID = uuid.New().String()
	} else {
		newID = *id
	}

	if createdAt == nil {
		now := time.Now()
		createdAt = &now
	}

	if updatedAt == nil {
		now := time.Now()
		updatedAt = &now
	}

	return &User{
		ID:          newID,
		Name:        name,
		Email:       email,
		Password:    password,
		Type_user:   typeUser,
		Profile:     profile,
		Permissions: permissions,
		Created_at:  createdAt,
		Updated_at:  updatedAt,
		Deleted_at:  deletedAt,
	}
}

func UpdateUser(
	user *User, name *string,
	email *string,
	password *string,
	profile *profile.Profile,
	permissions *permission.Permission,
) {
	if name != nil {
		user.Name = name
	}

	if email != nil {
		user.Email = email
	}

	if password != nil {
		user.Password = password
	}

	if profile != nil {
		user.Profile = *profile
	}

	if permissions != nil {
		user.Permissions = *permissions
	}

	now := time.Now()
	user.Updated_at = &now
}
