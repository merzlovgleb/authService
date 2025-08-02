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
		INSERT INTO dco.users (id,email,password,name,surname) VALUES ($1,$2,$3,$4)`

	res, err := repo.DbConnection.DB.ExecContext(ctx, query, user.ID, user.Email, user.Password, user.Name, user.Surname)
	if err != nil {
		logging.Error(err.Error())
	}
	rowsAffected, _ := res.RowsAffected()
	logging.Debug(fmt.Sprintf("Created %d rows", rowsAffected))
	return user.ID, nil
}

func (repo *UserRepository) GetAllUsers(ctx context.Context) []domain.User {
	var users []domain.User
	err := repo.DbConnection.DB.SelectContext(ctx, &users, "SELECT id, title from dco.users")
	if err != nil {
		logging.Error(nil, err.Error())
		//panic(err)
		return nil
	}
	return users
}
func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password, is_active FROM users WHERE email = $1`
	row := repo.DbConnection.DB.QueryRowContext(ctx, query, email)

	var user domain.User

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserById(ctx *gin.Context, id string) (*domain.User, error) {
	query := `SELECT id, email, password, is_active FROM users WHERE id = $1`
	row := repo.DbConnection.DB.QueryRowContext(ctx, query, id)

	var user domain.User

	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.IsActive); err != nil {
		return nil, err
	}
	return &user, nil
}
