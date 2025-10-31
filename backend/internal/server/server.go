package server

import (
	"checkers/pkg/checkers"
	"checkers/pkg/logger"
	"errors"
	"sync"
)

type GameServer struct {
	mu      sync.Mutex
	game    *checkers.Checkers
	players map[string]string // playerID -> "white"/"black"
}

var Server = &GameServer{
	game:    checkers.NewCheckers(),
	players: make(map[string]string),
}

// Добавление игрока
func (s *GameServer) JoinGame(playerID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.players) >= 2 {
		return "", errors.New("игра заполнена")
	}

	color := "white"
	if len(s.players) == 1 {
		color = "black"
	}

	s.players[playerID] = color
	logger.Log.Info("Игрок %s присоединился как %s", playerID, color)

	return color, nil
}

// Получение состояния игры
func (s *GameServer) GetState() (*checkers.Checkers, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.game, nil
}

// Сделать ход
func (s *GameServer) MakeMove(playerID string, from, to string) (bool, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	color, ok := s.players[playerID]
	if !ok {
		return false, "Вы не подключены к игре"
	}

	if (color == "white" && !s.game.IsWhiteTurn) || (color == "black" && s.game.IsWhiteTurn) {
		return false, "Сейчас не ваш ход"
	}

	move, err := checkers.ParseMove(from + " " + to)
	if err != nil {
		return false, err.Error()
	}

	success := s.game.MakeMove(move)
	if !success {
		return false, "Некорректный ход"
	}

	return true, "Ход выполнен"
}

// Сброс игры
func (s *GameServer) ResetGame() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.game = checkers.NewCheckers()
	s.players = make(map[string]string)
	logger.Log.Info("Игра сброшена")
}
