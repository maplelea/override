package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"override/api"
	"override/config"
	"override/internal/repositories"
	"override/middleware"
)

var logger *zap.Logger

func init() {
	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatalf("config initialization failed: %v", err)
	}

	// 初始化日志
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("logger initialization failed: %v", err)
	}
	defer logger.Sync()

	// 初始化数据库
	repositories.InitDB()

	// 初始化Redis
	repositories.InitRedis()

	// 初始化RabbitMQ
	rabbitMQConn, err := repositories.InitRabbitMQ()
	if err != nil {
		log.Fatalf("rabbitmq initialization failed: %v", err)
	}
	defer rabbitMQConn.Close()
}

func main() {
	r := gin.Default()
	r.Use(middleware.Logger(logger)) // 应用自定义的zap日志中间件
	r.Use(middleware.Auth())

	api.SetupRoutes(r)

	port := config.Viper.GetString("server.port")
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
