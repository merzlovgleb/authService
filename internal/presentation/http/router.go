package http

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	grouprouter "github.com/itpark/market/auth/internal/presentation/http/group/router"
	userrouter "github.com/itpark/market/auth/internal/presentation/http/user/router"
)

func RegisterRoutes(engine *gin.Engine, db *db.DbConnection) *gin.Engine {
	v1 := engine.Group("/api/v1")
	{
		v1.GET("/health", healthCheck)
	}

	groupRouter := grouprouter.NewGroupRouter(db)
	groupRouter.RegisterRoutes(v1)

	userRouter := userrouter.NewUserRouter(db)
	userRouter.RegisterRoutes(v1)

	return engine
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
