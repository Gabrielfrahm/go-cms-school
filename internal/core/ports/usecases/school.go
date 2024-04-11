package usecases

import "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/school"

type SchoolUseCasePort interface {
	Create(input CreateSchoolInput) (*school.School, error)
}

type AddressInput struct {
	ZipeCode string
	City     string
	Address  string
	Number   int
}

type CreateSchoolInput struct {
	Name       string
	Address    AddressInput
	DirectorID string
}
