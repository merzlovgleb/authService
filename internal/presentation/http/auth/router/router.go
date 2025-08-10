package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/app/utils"
	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/presentation/http/auth/handlers"
)

type AuthRouter struct {
	refreshHandler *handlers.RefreshHandler
}

func NewAuthRouter(connection *db.DbConnection) *AuthRouter {
	return &AuthRouter{
		refreshHandler: handlers.NewRefreshHandler(connection),
	}
}

func (r *AuthRouter) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	auth.Use(utils.JWTMiddleware())

	auth.GET("/info", handlers.UserInfoHandler)
	auth.POST("/refresh", r.refreshHandler.Handle)
}
