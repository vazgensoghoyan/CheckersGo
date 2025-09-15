package server

import (
	"net/http"

	"checkers/pkg/logger"

	"github.com/gin-gonic/gin"
)

func InitHandlers(router *gin.Engine) {
	logger.Log.Info("POST /games DEFINED")
	router.POST("/games", CreateGame)
}

// TODOOOOO

// CreateGame godoc
// @Summary Создать новую игру
// @Description Стартует новую партию шашек
// @Tags games
// @Accept json
// @Produce json
// @Param request body server.CreateGameRequest true "Игроки"
// @Success 200 {object} server.CreateGameResponce
// @Failure 400 {object} server.CreateGameResponce
// @Router /games [post]
func CreateGame(c *gin.Context) {
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
