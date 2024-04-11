package school

import (
	"time"

	"github.com/google/uuid"
)

type School struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	AddressID  string     `json:"address_id"`
	DirectorID string     `json:"director_id"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at,omitempty"`
}

func NewSchool(
	id *string,
	name string,
	addressID string,
	directorID *string,
	createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time,
) *School {
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

	if deletedAt == nil {
		now := time.Now()
		deletedAt = &now
	}

	return &School{
		ID:         newID,
		Name:       name,
		AddressID:  addressID,
		DirectorID: *directorID,
		Created_at: createdAt,
		Updated_at: updatedAt,
		Deleted_at: deletedAt,
	}
}
