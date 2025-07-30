package repository

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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

func (repo *UserRepository) CreateUser(ctx context.Context, name string, email string, password string) {
	query := `
		INSERT INTO dco.users (title) VALUES ($1)
		`

	res, err := repo.DbConnection.DB.ExecContext(ctx, query, name, email, password)
	if err != nil {
		logging.Error(err.Error())
	}
	rowsAffected, _ := res.RowsAffected()
	logging.Debug(fmt.Sprintf("Created %d rows", rowsAffected))
}

func (repo *UserRepository) GetAllUsers(ctx *gin.Context) []domain.User {
	var users []domain.User
	err := repo.DbConnection.DB.SelectContext(ctx, &users, "SELECT id, title from dco.users")
	if err != nil {
		logging.Error(err.Error())
	}
	return users
}
