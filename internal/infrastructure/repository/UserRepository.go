package repository

import (
	"context"
	"database/sql"

	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/domain"
)

type UserRepository struct {
	DbConnection *db.DbConnection
}

func NewUserRepository(dbConnection *db.DbConnection) *UserRepository {
	return &UserRepository{DbConnection: dbConnection}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO dco.users (user_id, email, password, name, surname, role, created_at, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repo.DbConnection.DB.ExecContext(
		ctx,
		query,
		user.ID,
		user.Email,
		[]byte(user.Password),
		user.Name,
		user.Surname,
		user.Role,
		user.CreatedAt,
		user.IsActive,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetAllUsers(ctx context.Context) []domain.User {
	var users []domain.User
	err := repo.DbConnection.DB.SelectContext(ctx, &users, `
		SELECT
			user_id AS id,
			email,
			name,
			surname,
			role,
			created_at,
			is_active
		FROM dco.users`)
	if err != nil {
		return nil
	}
	return users
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT
			user_id AS id,
			email,
			password,
			name,
			surname,
			role,
			created_at,
			is_active
		FROM dco.users
		WHERE email = $1`

	var user domain.User
	err := repo.DbConnection.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.Role,
		&user.CreatedAt,
		&user.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	query := `
		SELECT
			user_id AS id,
			email,
			password,
			name,
			surname,
			role,
			created_at,
			is_active
		FROM dco.users
		WHERE user_id = $1`

	var user domain.User
	err := repo.DbConnection.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.Role,
		&user.CreatedAt,
		&user.IsActive,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) DeleteUserById(ctx context.Context, id string) error {
	query := `DELETE FROM dco.users WHERE user_id = $1`

	result, err := repo.DbConnection.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (repo *UserRepository) UpdateUserById(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		UPDATE dco.users
		SET 
			email = $2,
			password = $3,
			name = $4,
			surname = $5,
			role = $6,
			is_active = $7,
			updated_at = NOW()
		WHERE user_id = $1
		RETURNING
			user_id AS id,
			email,
			password,
			name,
			surname,
			role,
			created_at,
			is_active`

	err := repo.DbConnection.DB.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Email,
		[]byte(user.Password),
		user.Name,
		user.Surname,
		user.Role,
		user.IsActive,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.Role,
		&user.CreatedAt,
		&user.IsActive,
	)
	if err != nil {
		return nil, err // вернёт sql.ErrNoRows, если пользователь не найден
	}

	return user, nil
}
