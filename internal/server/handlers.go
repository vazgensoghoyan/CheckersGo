package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	docs "checkers/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitHandlers(router *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/"
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	router.POST("/games", createGameRequest)
}

// CreateGame dogoc
// @Summary Создать новую игру
// @Description Стартует новую партию шашек
// @Tags games
// @Accept json
// @Produce json
// @Param request body CreateGameRequest true "Игроки"
// @Success 200 {object} GameResponse
// @Failure 400 {object} ErrorResponse
// @Router /games [post]
func createGameRequest(c *gin.Context) {
	var req CreateGameRequest

	if err := c.BindJSON(&req); err != nil {
		return
	}

	// do some work
	board := make([][]int, 8)
	for i := range board {
		board[i] = make([]int, 8)
	}

	// end of work

	resp := CreateGameResponce{
		ID:          "142",
		Status:      "waiting_for_move",
		Board:       board,
		CurrentTurn: req.Player1,
		Winner:      nil,
	}

	c.JSON(http.StatusOK, resp)
}
