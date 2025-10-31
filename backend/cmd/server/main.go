package main

import (
	"checkers/internal/server"
	"checkers/internal/swagger"

	"checkers/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	server.InitHandlers(router)

	swagger.Init(router)

	logger.Log.Info("Server is RUNNED on port 8080")
	router.Run(":8080")
}
