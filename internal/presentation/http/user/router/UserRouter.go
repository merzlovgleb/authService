package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/presentation/http/user"
)

type UserRouter struct {
	DbConnection *db.DbConnection
}

func NewUserRouter(DbConnection *db.DbConnection) *UserRouter {
	return &UserRouter{
		DbConnection: DbConnection,
	}
}

func (userRouter *UserRouter) RegisterRoutes(routerUser *gin.RouterGroup) {
	userGroup := routerUser.Group("users")

	userHandler := user.NewUserHandler(userRouter.DbConnection)

	userGroup.POST("/", userHandler.CreateUser)
	userGroup.DELETE("/", userHandler.DeleteUserById)
	userGroup.GET("/", userHandler.FindAll)
	userGroup.PUT("/", userHandler.UpdateUserById)
	userGroup.GET("/:id", userHandler.GetUserById)
}
