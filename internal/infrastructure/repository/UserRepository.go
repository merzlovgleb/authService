package repository

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/domain"
	"github.com/itpark/market/auth/internal/telemetry/logging"
)

type UserRepository struct {
	DbConnection *db.DbConnection
}

func NewUserRepository(dbConnection *db.DbConnection) *UserRepository {
	return &UserRepository{
		DbConnection: dbConnection,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	query := `
		INSERT INTO dco.users (user_id,email,password,name,surname) VALUES ($1,$2,$3,$4,$5)`

	res, err := repo.DbConnection.DB.ExecContext(ctx, query, user.ID, user.Email, []byte(user.Password), user.Name, user.Surname)
	if err != nil {
		logging.Warn(err.Error())
		return uuid.Nil, err
	}
	rowsAffected, _ := res.RowsAffected()
	logging.Debug(fmt.Sprintf("Created %d rows", rowsAffected))
	return user.ID, nil
}

func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := repo.DbConnection.DB.SelectContext(ctx, &users, "SELECT user_id, email, name, surname, role FROM dco.users")
	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}
	return users, nil
}
func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT user_id, email, password, is_active FROM dco.users WHERE email = $1`
	row := repo.DbConnection.DB.QueryRowContext(ctx, query, email)

	var user domain.User

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserById(ctx *gin.Context, id string) (*domain.User, error) {
	query := `SELECT user_id, email, password, is_active FROM dco.users WHERE user_id = $1`
	row := repo.DbConnection.DB.QueryRowContext(ctx, query, id)

	var user domain.User

	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.IsActive); err != nil {
		return nil, err
	}
	return &user, nil
}
