package server

import (
	"net/http"

	"checkers/pkg/checkers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @title Checkers API
// @version 1.0
// @description API для игры в шашки онлайн.
// @BasePath /

// InitHandlers инициализирует все маршруты
func InitHandlers(router *gin.Engine) {
	router.POST("/join", joinHandler)
	router.GET("/state", stateHandler)
	router.POST("/move", moveHandler)
	router.POST("/reset", resetHandler)
}

// @Summary Присоединиться к игре
// @Description Игрок присоединяется к текущей партии и получает ID и цвет фигур.
// @Tags Game
// @Accept json
// @Produce json
// @Param request body joinRequest true "Имя игрока"
// @Success 200 {object} joinResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /join [post]
func joinHandler(c *gin.Context) {
	var req = joinRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат"})
		return
	}

	playerID := uuid.New().String()
	color, err := Server.JoinGame(playerID)
	if err != nil {
		c.JSON(http.StatusForbidden, errorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, joinResponse{
		PlayerID: playerID,
		Color:    color,
	})
}

// @Summary Получить состояние игры
// @Description Возвращает текущее состояние доски, чей сейчас ход и информацию о игроке.
// @Tags Game
// @Accept json
// @Produce json
// @Success 200 {object} stateResponse
// @Failure 403 {object} errorResponse
// @Router /state [get]
func stateHandler(c *gin.Context) {
	game, err := Server.GetState()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stateResponse{
		Board:       convertBoard(game.Board),
		IsWhiteTurn: game.IsWhiteTurn,
	})
}

func convertBoard(board [][]checkers.Figure) [][]figureResponse {
	result := make([][]figureResponse, len(board))
	for i := range board {
		result[i] = make([]figureResponse, len(board[i]))
		for j := range board[i] {
			fig := board[i][j]
			result[i][j] = figureResponse{
				IsNone:  fig.IsNone,
				IsWhite: fig.IsWhite,
				IsKing:  fig.IsKing,
			}
		}
	}
	return result
}

// @Summary Сделать ход
// @Description Игрок делает ход с клетки "from" на клетку "to".
// @Tags Game
// @Accept json
// @Produce json
// @Param request body moveRequest true "Данные хода"
// @Success 200 {object} moveResponse
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Router /move [post]
func moveHandler(c *gin.Context) {
	var req = moveRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат"})
		return
	}

	success, message := Server.MakeMove(req.PlayerID, req.From, req.To)
	c.JSON(http.StatusOK, moveResponse{
		Success: success,
		Message: message,
	})
}

// @Summary Сбросить игру
// @Description Полностью сбрасывает текущую партию, удаляя игроков и заново создавая доску.
// @Tags Game
// @Accept json
// @Produce json
// @Success 200 {object} resetResponse
// @Router /reset [post]
func resetHandler(c *gin.Context) {
	Server.ResetGame()
	c.JSON(http.StatusOK, resetResponse{Success: true})
}
