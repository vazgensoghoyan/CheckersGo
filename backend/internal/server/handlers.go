package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitHandlers(router *gin.Engine) {
	router.POST("/join", joinHandler)
	router.GET("/state", stateHandler)
	router.POST("/move", moveHandler)
	router.POST("/reset", resetHandler)
}

// Присоединиться к игре
func joinHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат"})
		return
	}

	playerID := uuid.New().String()
	color, err := Server.JoinGame(playerID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"player_id": playerID,
		"color":     color,
	})
}

// Получить состояние игры
func stateHandler(c *gin.Context) {
	playerID := c.Query("player_id")
	game, yourTurn, err := Server.GetState(playerID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"board":     game.Board, // метод возвращает [][]Figure
		"isWhiteTurn": game.IsWhiteTurn,
		"yourTurn":  yourTurn,
	})
}

// Сделать ход
func moveHandler(c *gin.Context) {
	var req struct {
		PlayerID string `json:"player_id"`
		From     string `json:"from"`
		To       string `json:"to"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат"})
		return
	}

	success, message := Server.MakeMove(req.PlayerID, req.From, req.To)
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
	})
}

// Сбросить игру
func resetHandler(c *gin.Context) {
	Server.ResetGame()
	c.JSON(http.StatusOK, gin.H{"success": true})
}
