package group

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	groupRepository "github.com/itpark/market/auth/internal/infrastructure/repository"
	customErrors "github.com/itpark/market/auth/internal/presentation/http/common"
	"github.com/itpark/market/auth/internal/presentation/http/group/dto"
	"github.com/itpark/market/auth/internal/service"
	"net/http"
)

type Handler struct {
	Service *service.GroupService
}

func Init(connection *db.DbConnection) *Handler {
	repository := groupRepository.NewGroupRepository(connection)
	newGroupService := service.NewGroupService(repository)

	return &Handler{
		Service: newGroupService,
	}
}

func (g *Handler) CreateGroup(ctx *gin.Context) {
	var groupDto dto.CreateGroupDto
	if err := ctx.ShouldBindJSON(&groupDto); err != nil {
		ctx.JSON(http.StatusBadRequest, customErrors.CreateError("Invalid Request body", err))
		return
	}

	if err := g.Service.CreateGroup(ctx, groupDto); err != nil {
		ctx.JSON(http.StatusInternalServerError, customErrors.CreateError("Failed to create group", err))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (g *Handler) FindAll(ctx *gin.Context) {
	groups, err := g.Service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, customErrors.CreateError("Failed to fetch groups", err))
		return
	}
	ctx.JSON(http.StatusOK, groups)
}
