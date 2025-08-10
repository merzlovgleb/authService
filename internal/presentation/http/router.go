package http

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	authRouter "github.com/itpark/market/auth/internal/presentation/http/auth/router"
	groupRouter "github.com/itpark/market/auth/internal/presentation/http/group/router"
)

func RegisterRoutes(engine *gin.Engine, db *db.DbConnection) *gin.Engine {
	api := engine.Group("/api/v1")

	api.GET("/health", healthCheck)

	group := groupRouter.NewGroupRouter(db)
	group.RegisterRoutes(api)

	auth := authRouter.NewAuthRouter(db)
	auth.RegisterRoutes(api)

	return engine
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
