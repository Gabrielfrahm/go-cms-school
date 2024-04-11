package school

import (
	"database/sql"
	"errors"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/school"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/repositories"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/usecases"
)

type SchoolUseCase struct {
	schoolRepo  repositories.SchoolRepositoryPort
	addressRepo repositories.AddressRepoPort
	db          *sql.DB
}

func NewSchoolUseCase(schoolRepo repositories.SchoolRepositoryPort, addressRepo repositories.AddressRepoPort, db *sql.DB) usecases.SchoolUseCasePort {
	return &SchoolUseCase{
		schoolRepo:  schoolRepo,
		addressRepo: addressRepo,
		db:          db,
	}
}

// Create implements usecases.SchoolUseCasePort.
func (s *SchoolUseCase) Create(input usecases.CreateSchoolInput) (*school.School, error) {

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	addressEntity := school.NewAddress(nil, input.Address.ZipeCode, input.Address.City, input.Address.Address, input.Address.Number, nil, nil, nil)

	addressRepo, err := s.addressRepo.Create(tx, addressEntity)
	if err != nil {
		tx.Rollback()
		return &school.School{}, errors.New(err.Error())
	}

	schoolEntity := school.NewSchool(nil, input.Name, addressRepo.ID, &input.DirectorID, nil, nil, nil)

	schoolRepo, err := s.schoolRepo.Create(tx, schoolEntity)
	if err != nil {
		tx.Rollback()
		return &school.School{}, errors.New(err.Error())
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return schoolRepo, nil
}
