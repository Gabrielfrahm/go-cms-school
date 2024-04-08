package profile

import (
	"time"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/permission"
	"github.com/google/uuid"
)

type Profile struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	TypeUser    string                `json:"type_user"`
	Permissions permission.Permission `json:"permissions,omitempty"`
	Created_at  *time.Time            `json:"created_at"`
	Updated_at  *time.Time            `json:"updated_at"`
	Deleted_at  *time.Time            `json:"deleted_at,omitempty"`
}

func NewProfile(id *string, name string, typeUser string,
	permissions permission.Permission, createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time) *Profile {
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

	return &Profile{
		ID:          newID,
		Name:        name,
		TypeUser:    typeUser,
		Permissions: permissions,
		Created_at:  createdAt,
		Updated_at:  updatedAt,
		Deleted_at:  deletedAt,
	}
}

func UpdateProfile(profile *Profile, name *string, permissions *permission.Permission) {
	if name != nil {
		profile.Name = *name
	}

	if permissions != nil {
		profile.Permissions = *permissions
	}

	now := time.Now()
	profile.Updated_at = &now
}

func DeleteProfile(profile *Profile) {
	deletedAt := time.Now()
	profile.Deleted_at = &deletedAt
}
