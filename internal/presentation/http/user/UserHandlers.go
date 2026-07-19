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

func (h *UserHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.Service.GetUserById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, customErrors.CreateError("User not found", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	users, err := h.Service.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, customErrors.CreateError("Failed to fetch users", err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}
