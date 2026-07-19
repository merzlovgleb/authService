package repository

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/domain"
	"github.com/itpark/market/auth/internal/telemetry/logging"
)

type GroupRepository struct {
	DbConnection *db.DbConnection
}

func NewGroupRepository(dbConnection *db.DbConnection) *GroupRepository {
	return &GroupRepository{
		DbConnection: dbConnection,
	}
}

func (repo *GroupRepository) CreateGroup(ctx context.Context, name string) error {
	query := `
		INSERT INTO dco.groups (title) VALUES ($1)
		`

	res, err := repo.DbConnection.DB.ExecContext(ctx, query, name)
	if err != nil {
		logging.Warn(err.Error())
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	logging.Debug(fmt.Sprintf("Created %d rows", rowsAffected))
	return nil
}

func (repo *GroupRepository) GetAll(ctx *gin.Context) ([]domain.Group, error) {
	var groups []domain.Group
	err := repo.DbConnection.DB.SelectContext(ctx, &groups, "SELECT id, title from dco.groups")
	if err != nil {
		logging.Warn(err.Error())
		return nil, err
	}
	return groups, nil
}
