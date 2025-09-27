package checkers

type Checkers struct {
	board [][]Figure
}

func NewCheckers() *Checkers {
	board := make([][]Figure, 8)

	for i := range board {
		board[i] = make([]Figure, 8)
	}

	// расставляем фигуры

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = Figure{isNone: true, isWhite: false, isKing: false}
		}
	}

	// черные
	// белые

	return &Checkers{board: board}
}

func StartGame(checkers *Checkers) {
	// начинаем обрабатывать действия игроков и менять доску соответствующе
}

type Figure struct {
	isNone  bool
	isWhite bool
	isKing  bool
}

// это будет где то извне происходить, так будет игра начинаться
func main() {
	var game = NewCheckers()

	StartGame(game)
}
