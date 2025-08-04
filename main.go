package main

//
//import (
//	"context"
//	"github.com/gin-gonic/gin"
//	"github.com/itpark/market/auth/internal/config"
//	"github.com/itpark/market/auth/internal/config/db"
//	"github.com/itpark/market/auth/internal/presentation/http/user/router"
//	"log"
//)
//
//func main() {
//	// Загрузка конфигурации
//	cfg, err := config.Load()
//	if err != nil {
//		log.Fatalf("Failed to load config: %v", err)
//	}
//
//	// Инициализация подключения к БД
//	ctx := context.Background()
//	connection := db.InitConnection(ctx, cfg)
//	if connection == nil {
//		log.Fatal("Failed to connect to database")
//	}
//
//	// Инициализация Gin
//	r := gin.Default()
//
//	// Создание группы маршрутов
//	api := r.Group("/api")
//
//	// Регистрация роутов пользователей
//	userRouter := router.NewUserRouter(connection)
//	userRouter.RegisterRoutes(api)
//
//	// Запуск сервера
//	port := cfg.Server.HTTPPort
//	if port == "" {
//		port = "8080" // fallback на порт по умолчанию
//	}
//	if err := r.Run(":" + port); err != nil {
//		log.Fatalf("Failed to start server: %v", err)
//	}
//}
