package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHandlers(router *gin.Engine) {
	router.POST("/games", createGameRequest)
}

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
