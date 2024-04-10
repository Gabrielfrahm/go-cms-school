package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/permission"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/entities/profile"
	entity "github.com/Gabrielfrahm/go-cms-school/internal/core/entities/user"
	"github.com/Gabrielfrahm/go-cms-school/internal/core/ports/repositories"
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
		profile_type_user   string
		profile_createdAt   sql.NullTime
		profile_updatedAt   sql.NullTime
		profile_deletedAt   sql.NullTime
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
			u.profile_id,
			p.name AS profile_name,
			p.type_user as profile_type_user,
			p.created_at as profile_created_at, 
			p.updated_at as profile_updated_at, 
			p.deleted_at as profile_deleted_at,
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
		&profile_type_user,
		&profile_createdAt,
		&profile_updatedAt,
		&profile_deletedAt,
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
			return nil, errors.New("user not found")
		}
		return nil, err // Outro erro de banco de dados
	}

	permissions := permission.NewPermission(permission_users, permission_classes, permission_profiles, permission_lessons)

	profile := profile.NewProfile(&profile_id, profile_name, profile_type_user, *permissions, &profile_createdAt.Time, &profile_updatedAt.Time, &profile_deletedAt.Time)

	return entity.NewUser(&id, &name, &emailsql, &password, type_user, *profile, *permissions, &createdAt.Time, &updatedAt.Time, &deletedAt.Time), nil
}

func (r *UserRepository) Create(user *entity.User) (*entity.User, error) {

	if user.Email != nil {
		rows, err := r.db.Query("SELECT email FROM users WHERE email = $1", user.Email)
		if err != nil {
			return &entity.User{}, err
		}
		defer rows.Close()

		if rows.Next() {
			return &entity.User{}, errors.New("email already in use")
		}
	}

	tx, err := r.db.Begin()
	if err != nil {
		return &entity.User{}, err
	}
	defer tx.Rollback()
	var userID string
	err = tx.QueryRow("INSERT INTO users (name, email, password, type_user, profile_id, created_at, updated_at, deleted_at) VALUES ($1, $2, $3,$4,$5,$6,$7, $8) RETURNING id", user.Name, user.Email, user.Password, user.Type_user, user.Profile.ID, user.Created_at, user.Updated_at, user.Deleted_at).Scan(&userID)
	if err != nil {
		return &entity.User{}, err
	}

	stmt, err := tx.Prepare("INSERT INTO user_permissions (user_id, users, classes, profiles, lessons) VALUES($1,$2,$3,$4, $5)")
	if err != nil {
		return &entity.User{}, err
	}

	defer stmt.Close()

	stmt.Exec(userID, user.Permissions.Users, user.Permissions.Classes, user.Permissions.Profiles, user.Permissions.Lessons)

	if err := tx.Commit(); err != nil {
		return &entity.User{}, err
	}

	user.ID = userID

	return user, nil
}

func (r *UserRepository) ListAllUser(input repositories.ListAllUserInput) ([]entity.User, int, error) {
	// todo tem que retornar as permissoes do perfil.
	var users []entity.User
	var total int
	query := `SELECT 
	u.id,
	u.name, 
	u.email, 
	u.password, 
	u.type_user,
	u.profile_id,
	p.name AS profile_name,
	p.type_user as profile_type_user,
	p.created_at as profile_created_at, 
	p.updated_at as profile_updated_at, 
	p.deleted_at as profile_deleted_at,
	perm.users as user_users,
	perm.classes as user_classes,
	perm.profiles as user_profiles,
	perm.lessons as user_lessons,
	pp.users as permission_users,
	pp.classes as permission_classes,
	pp.profiles as permission_profiles,
	pp.lessons as permission_lessons, 
	u.created_at, 
	u.updated_at, 
	u.deleted_at
	FROM 
		users u
	JOIN 
		profiles p ON u.profile_id = p.id
	JOIN 
		user_permissions perm ON u.id = perm.user_id
	JOIN
		profile_permissions pp ON p.id = pp.profile_id 
	WHERE 
	1 = 1`

	// parametros do filtros
	if input.Name != nil {
		query += " AND u.name ILIKE '%" + *input.Name + "%'"
	}
	if input.Email != nil {
		query += " AND u.email ILIKE '%" + *input.Email + "%'"
	}
	if input.TypeUser != nil {
		query += " AND u.type_user ILIKE '%" + *input.TypeUser + "%'"
	}

	// Adicione paginação a consulta
	offset := (*input.Page - 1) * *input.PerPage
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", input.PerPage, offset)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		var profile profile.Profile
		var permission permission.Permission
		var createdAt, updatedAt, deletedAt sql.NullTime
		var profileCreatedAt, profileUpdatedAt, profileDeletedAt sql.NullTime
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Type_user,
			&profile.ID,
			&profile.Name,
			&profile.TypeUser,
			&profileCreatedAt,
			&profileUpdatedAt,
			&profileDeletedAt,
			&permission.Users,
			&permission.Classes,
			&permission.Profiles,
			&permission.Lessons,
			&profile.Permissions.Users,
			&profile.Permissions.Classes,
			&profile.Permissions.Profiles,
			&profile.Permissions.Lessons,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			return nil, 0, err
		}

		if createdAt.Valid {
			user.Created_at = &createdAt.Time
		}
		if updatedAt.Valid {
			user.Updated_at = &updatedAt.Time
		}
		if deletedAt.Valid {
			user.Deleted_at = &deletedAt.Time
		}
		if profileCreatedAt.Valid {
			profile.Created_at = &profileCreatedAt.Time
		}
		if profileUpdatedAt.Valid {
			profile.Updated_at = &profileUpdatedAt.Time
		}
		if profileDeletedAt.Valid {
			profile.Deleted_at = &profileDeletedAt.Time
		}

		user.Profile = profile
		user.Permissions = permission
		users = append(users, user)
	}

	// count dos totais de items.
	countQuery := `SELECT COUNT(*) FROM users WHERE 1=1`
	if input.Name != nil {
		countQuery += " AND name ILIKE '%" + *input.Name + "%'"
	}
	if input.Email != nil {
		countQuery += " AND email ILIKE '%" + *input.Email + "%'"
	}

	err = r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepository) FindById(input string) (*entity.User, error) {
	var user entity.User
	var profile profile.Profile
	var permission permission.Permission
	var createdAt, updatedAt, deletedAt sql.NullTime
	var profileCreatedAt, profileUpdatedAt, profileDeletedAt sql.NullTime

	err := r.db.QueryRow(
		`SELECT 
		u.id,
		u.name, 
		u.email, 
		u.password, 
		u.type_user,
		u.profile_id,
		p.name AS profile_name,
		p.type_user as profile_type_user,
		p.created_at as profile_created_at, 
		p.updated_at as profile_updated_at, 
		p.deleted_at as profile_deleted_at,
		perm.users as user_users,
		perm.classes as user_classes,
		perm.profiles as user_profiles,
		perm.lessons as user_lessons,
		pp.users as permission_users,
		pp.classes as permission_classes,
		pp.profiles as permission_profiles,
		pp.lessons as permission_lessons, 
		u.created_at, 
		u.updated_at, 
		u.deleted_at
		FROM 
			users u
		JOIN 
			profiles p ON u.profile_id = p.id
		JOIN 
			user_permissions perm ON u.id = perm.user_id
		JOIN
			profile_permissions pp ON p.id = pp.profile_id 
		WHERE 
		u.id = $1`,
		input,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Type_user,
		&profile.ID,
		&profile.Name,
		&profile.TypeUser,
		&profileCreatedAt,
		&profileUpdatedAt,
		&profileDeletedAt,
		&permission.Users,
		&permission.Classes,
		&permission.Profiles,
		&permission.Lessons,
		&profile.Permissions.Users,
		&profile.Permissions.Classes,
		&profile.Permissions.Profiles,
		&profile.Permissions.Lessons,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err // Outro erro de banco de dados
	}
	user.Profile = profile
	user.Permissions = permission
	return &user, nil
}
