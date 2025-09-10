package main

import (
	"checkers/internal/server"
	"checkers/pkg/logger"

	"github.com/gin-gonic/gin"

	docs "checkers/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	server.InitHandlers(router)

	logger.Log.Info("Сервер запущен на порту 8082")
	router.Run("localhost:8082")
}
