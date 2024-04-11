package school

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID         string     `json:"id"`
	ZipeCode   string     `json:"zipe_code"`
	City       string     `json:"city"`
	Address    string     `json:"address"`
	Number     int        `json:"number"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at,omitempty"`
}

func NewAddress(
	id *string,
	zipeCode string,
	city string,
	address string,
	number int,
	createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time,
) *Address {
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

	return &Address{
		ID:         newID,
		ZipeCode:   zipeCode,
		City:       city,
		Address:    address,
		Number:     number,
		Created_at: createdAt,
		Updated_at: updatedAt,
		Deleted_at: deletedAt,
	}
}
