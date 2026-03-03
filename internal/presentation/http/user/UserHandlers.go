package user

import (
	"database/sql"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	repository "github.com/itpark/market/auth/internal/infrastructure/repository"
	customErrors "github.com/itpark/market/auth/internal/presentation/http/common"
	"github.com/itpark/market/auth/internal/presentation/http/user/dto"
	"github.com/itpark/market/auth/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(connection *db.DbConnection) *UserHandler {
	userRepository := repository.NewUserRepository(connection)
	userService := service.NewUserService(userRepository)

	return &UserHandler{Service: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userDto dto.CreateUserDto

	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.CreateError("Invalid request body", err))
		return
	}

	createdUser, err := h.Service.CreateUser(c.Request.Context(), &userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, customErrors.CreateError("Failed to create user", err))
		return
	}

	c.Header("Location", "/api/v1/users/"+createdUser.ID.String())
	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) FindAll(c *gin.Context) {
	users := h.Service.GetAllUsers(c.Request.Context())
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.Service.GetUserById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, customErrors.CreateError("User not found", err))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteUserById(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, customErrors.CreateError("Failed to delete user", err))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) UpdateUserById(c *gin.Context) {
	id := c.Param("id")

	// Валидация формата ID (если используется UUID)
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.CreateError("Invalid user ID format", err))
		return
	}

	var updateDto dto.UpdateUserDto
	if err := c.ShouldBindJSON(&updateDto); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.CreateError("Invalid request body", err))
		return
	}

	updatedUser, err := h.Service.UpdateUserById(c.Request.Context(), id, &updateDto)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, customErrors.CreateError("User not found", err))
			return
		}
		c.JSON(http.StatusInternalServerError, customErrors.CreateError("Failed to update user", err))
		return
	}

	// ✅ Возвращаем обновлённого пользователя
	c.JSON(http.StatusOK, updatedUser)
}
