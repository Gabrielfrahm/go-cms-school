package repositories

import (
	"database/sql"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/school"
)

type SchoolRepositoryPort interface {
	Create(tx *sql.Tx, entity *school.School) (*school.School, error)
}
