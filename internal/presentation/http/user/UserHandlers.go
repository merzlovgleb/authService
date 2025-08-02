package user

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	groupRepository "github.com/itpark/market/auth/internal/infrastructure/repository"
	customErrors "github.com/itpark/market/auth/internal/presentation/http/common"
	"github.com/itpark/market/auth/internal/presentation/http/user/dto"
	"github.com/itpark/market/auth/internal/service"
	"net/http"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(connection *db.DbConnection) *UserHandler {
	repository := groupRepository.NewUserRepository(connection)
	userService := service.NewUserService(repository)

	return &UserHandler{
		Service: userService,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var userDto dto.CreateUserDto

	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, customErrors.CreateError("Invalid request body", err))
		return
	}

	if err := h.Service.CreateUser(ctx.Request.Context(), &userDto); err != nil {
		ctx.JSON(http.StatusInternalServerError, customErrors.CreateError("Failed to create user", err))
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	users := h.Service.GetAllUsers(ctx)
	ctx.JSON(http.StatusOK, users)
}
