package main

import (
	"checkers/internal/server"
	"checkers/pkg/logger"

	_ "checkers/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server.InitHandlers(router)

	server.InitSwagger(router)

	logger.Log.Info("Server is RUNNED on port 8080")
	router.Run(":8080")
}
