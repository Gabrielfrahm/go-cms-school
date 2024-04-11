package repositories

import (
	"database/sql"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/school"
)

type SchoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{db: db}
}

func (r *SchoolRepository) Create(tx *sql.Tx, entity *school.School) (*school.School, error) {
	rows, err := tx.Query(
		"SELECT * FROM schools WHERE name = $1", entity.Name,
	)
	if err != nil {
		return &school.School{}, err
	}
	defer rows.Close()
}
