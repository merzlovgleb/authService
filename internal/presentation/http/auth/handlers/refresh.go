package handlers

import (
	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/infrastructure/repository"
	token "github.com/itpark/market/auth/internal/presentation/http/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/presentation/http/auth/dto"
)

type RefreshHandler struct {
	TokenService *token.JwtTokenService
	UserRepo     *repository.UserRepository
}

func NewRefreshHandler(connection *db.DbConnection) *RefreshHandler {
	userRepository := repository.NewUserRepository(connection)
	tokenService, _ := token.New("secret", 2*time.Minute, 2*time.Minute)
	return &RefreshHandler{TokenService: tokenService,
		UserRepo: userRepository,
	}
}
func (h *RefreshHandler) Handle(c *gin.Context) {
	var req dto.RefreshRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	userID, err := h.TokenService.ParseToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid refresh token"})
		return
	}

	user, err := h.UserRepo.GetUserById(c, userID.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	accessToken, err := h.TokenService.GenerateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create access token"})
		return
	}

	refreshToken, err := h.TokenService.GenerateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create refresh token"})
		return
	}

	c.JSON(http.StatusOK, dto.TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
