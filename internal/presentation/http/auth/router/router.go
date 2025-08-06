package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/app/utils"
	"github.com/itpark/market/auth/internal/presentation/http/auth/handlers"
)

type AuthRouter struct{}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (r *AuthRouter) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	auth.Use(utils.JWTMiddleware())

	auth.GET("/my", handlers.MyHandler)
}
