package repositories

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/profile"

type ProfileRepository interface {
	FindById(id string) (*profile.Profile, error)
}
