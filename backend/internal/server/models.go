package server

type joinRequest struct {
	Name string `json:"name"` // Имя игрока
}

type joinResponse struct {
	PlayerID string `json:"player_id"` // Уникальный ID игрока
	Color    string `json:"color"`     // "white" или "black"
}

type errorResponse struct {
	Error string `json:"error"`
}

type figureResponse struct {
	IsNone  bool `json:"is_none"`
	IsWhite bool `json:"is_white"`
	IsKing  bool `json:"is_king"`
}

type stateResponse struct {
	Board       [][]figureResponse `json:"board"`       // Доска 8x8
	IsWhiteTurn bool               `json:"isWhiteTurn"` // Чей ход сейчас
	YourTurn    bool               `json:"yourTurn"`    // Флаг, если ход текущего игрока
}

type moveRequest struct {
	PlayerID string `json:"player_id"`
	From     string `json:"from"` // Пример: "c3"
	To       string `json:"to"`   // Пример: "d4"
}

type moveResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type resetResponse struct {
	Success bool `json:"success"`
}
