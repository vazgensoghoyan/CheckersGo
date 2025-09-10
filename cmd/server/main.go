package main

import (
	"checkers/internal/server"
	"checkers/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server.InitHandlers(router)

	logger.Log.Info("Сервер запущен на порту 8080")
	router.Run("localhost:8080")
}
