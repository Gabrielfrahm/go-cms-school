package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/permission"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/profile"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) FindById(id string) (*profile.Profile, error) {
	var (
		idSql               string
		name                string
		type_user           string
		permission_users    int
		permission_classes  int
		permission_profiles int
		permission_lessons  int
		createdAt           sql.NullTime
		updatedAt           sql.NullTime
		deletedAt           sql.NullTime
	)

	// check if user has token in db.
	err := r.db.QueryRow(
		`SELECT 
			p.id AS profile_is,
			p.name AS profile_name,
			p.type_user as profile_type_user,
			p.created_at as profile_created_at, 
			p.updated_at as profile_updated_at, 
			p.deleted_at as profile_deleted_at,
			perm.users as profile_users,
			perm.classes as profile_classes,
			perm.profiles as profile_profiles,
			perm.lessons as profile_lessons
		FROM 
			profiles p
		JOIN 
			profile_permissions perm ON p.id = perm.profile_id
		WHERE 
    		p.id = $1`,
		id,
	).Scan(&idSql,
		&name,
		&type_user,
		&createdAt,
		&updatedAt,
		&deletedAt,
		&permission_users,
		&permission_classes,
		&permission_profiles,
		&permission_lessons,
	)

	if err != nil {
		return &profile.Profile{}, err
	}
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			return nil, errors.New("profile not found")
		}
		return nil, err // Outro erro de banco de dados
	}

	permissions := permission.NewPermission(permission_users, permission_classes, permission_profiles, permission_lessons)
	profileEntity := profile.NewProfile(&idSql, name, type_user, *permissions, &createdAt.Time, &updatedAt.Time, &deletedAt.Time)

	return profileEntity, nil
}
