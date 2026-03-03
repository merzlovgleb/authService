package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/infrastructure/repository"
	"net/http"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: repo}
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.UserRepo.GetUserById(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
