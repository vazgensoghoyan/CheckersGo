package server

type CreateGameRequest struct {
	Player1 string `json:"player1"`
	Player2 string `json:"player2"`
}

type CreateGameResponce struct {
	ID          string  `json:"id"`
	Status      string  `json:"status"`
	Board       [][]int `json:"board"`
	CurrentTurn string  `json:"current_turn"`
	Winner      *string `json:"winner,omitempty"`
}
