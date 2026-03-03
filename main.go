package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/config"
	"github.com/itpark/market/auth/internal/config/db"
	"github.com/itpark/market/auth/internal/presentation/http/user/router"
	"github.com/itpark/market/auth/internal/telemetry/logging"
	"log"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config load failed: %v", err)
	}

	logging.Init(cfg)

	ctx := context.Background()

	connection := db.InitConnection(ctx, cfg)
	if connection == nil {
		logging.Error(ctx, "Failed to connect to database")
		log.Fatal("DB connection is nil")
	}

	r := gin.Default()
	api := r.Group("/api")
	userRouter := router.NewUserRouter(connection)
	userRouter.RegisterRoutes(api)

	if err := r.Run(":8080"); err != nil {
		logging.Error(ctx, "Failed to start server: %v", err.Error())
		log.Fatalf("Server error: %v", err)
	}

}
