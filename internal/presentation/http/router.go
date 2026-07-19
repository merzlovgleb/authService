package http

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/presentation/http/group/router"
	userRouter "github.com/itpark/market/auth/internal/presentation/http/user/router"
)

func RegisterRoutes(engine *gin.Engine, db *db.DbConnection) *gin.Engine {
	api := engine.Group("/api/v1")

	{
		api.GET("/health", healthCheck)
	}

	groupRouter := router.NewGroupRouter(db)
	groupRouter.RegisterRoutes(api)

	newUserRouter := userRouter.NewUserRouter(db)
	newUserRouter.RegisterRoutes(api)

	return engine
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
