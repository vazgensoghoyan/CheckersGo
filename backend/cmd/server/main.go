package main

import (
	"checkers/internal/server"
	"checkers/internal/swagger"

	"checkers/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server.InitHandlers(router)

	swagger.Init(router)

	logger.Log.Info("Server is RUNNED on port 8080")
	router.Run(":8080")
}
