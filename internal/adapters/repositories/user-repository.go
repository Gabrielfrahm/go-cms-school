package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/permission"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/profile"
	entity "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var (
		id                  string
		name                string
		emailsql            string
		password            string
		type_user           string
		profile_id          string
		profile_name        string
		permission_users    int
		permission_classes  int
		permission_profiles int
		permission_lessons  int
		createdAt           sql.NullTime
		updatedAt           sql.NullTime
		deletedAt           sql.NullTime
	)
	err := r.db.QueryRow(
		`SELECT 
			u.id,
			u.name, 
			u.email, 
			u.password, 
			u.type_user,
			u.profile_id ,
			p.name AS profile_name,
			perm.users as profile_users,
			perm.classes as profile_classes,
			perm.profiles as profile_profiles,
			perm.lessons as profile_lessons,
			u.created_at, 
			u.updated_at, 
			u.deleted_at
		FROM 
			users u
		JOIN 
			profiles p ON u.profile_id = p.id
		JOIN 
			user_permissions perm ON u.id = perm.user_id
		WHERE 
    		u.email = $1`,
		email,
	).Scan(&id,
		&name,
		&emailsql,
		&password,
		&type_user,
		&profile_id,
		&profile_name,
		&permission_users,
		&permission_classes,
		&permission_profiles,
		&permission_lessons,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			return nil, errors.New("user not found")
		}
		return nil, err // Outro erro de banco de dados
	}

	permissions := permission.NewPermission(permission_users, permission_classes, permission_profiles, permission_lessons)

	profile := profile.NewProfile(&profile_id, profile_name, *permissions, &createdAt.Time, &updatedAt.Time, &deletedAt.Time)

	return entity.NewUser(&id, &name, &emailsql, &password, type_user, *profile, *permissions, &createdAt.Time, &updatedAt.Time, &deletedAt.Time), nil
}
