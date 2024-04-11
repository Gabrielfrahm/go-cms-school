package repositories

import (
	"database/sql"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/school"
)

type AddressRepoPort interface {
	Create(tx *sql.Tx, entity *school.Address) (*school.Address, error)
}
