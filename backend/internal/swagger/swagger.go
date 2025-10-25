package swagger

import (
	"os"

	_ "checkers/docs"
	"checkers/pkg/logger"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(router *gin.Engine) {
	logger.Log.Info("Swagger TURNED ON")
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)
}

func HasSwagger() bool {
	return os.Getenv("ENABLE_SWAGGER") == "true"
}
