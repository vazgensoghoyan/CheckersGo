package checkers

import "fmt"

type Checkers struct {
	board [][]Figure
}

type Move struct {
	frowRow int
	fromCol int
	toRow int
	toCol int
}

type Figure struct {
	isNone  bool
	isWhite bool
	isKing  bool
}

func NewCheckers() *Checkers {
	board := make([][]Figure, 8)

	for i := range board {
		board[i] = make([]Figure, 8)
	}

	// Расставляем фигуры
	for row := 0; row < 8; row++ { // По строкам
		for col := 0; col < 8; col++ { // По столбцам
			if (row+col)%2 == 1 { // Шашки на черных полях
				if row < 3 { // Белые сверху
					board[row][col] = Figure{isNone: false, isWhite: true, isKing: false}
				} else if row > 4 { // Черные снизу
					board[row][col] = Figure{isNone: false, isWhite: false, isKing: false}
				} else {
					board[row][col] = Figure{isNone: true, isWhite: false, isKing: false} // Белые поля пустые
				}
			}
		}
	}

	return &Checkers{board: board}
}

func (checkers *Checkers) PrintBoard() {
	for row := 0; row < 8; row++ { 
		fmt.Printf("%d ", 8-row)
		for col := 0; col < 8; col++ { 
			figure := checkers.board[row][col]
			var symbol string
			if figure.isNone {
				symbol = " . " // Пустая клетка				
			} else if figure.isWhite {
				if figure.isKing{
					symbol = " Wk " // Белая дамка
				} else {
					symbol = " W " // Обычная белая
				}
			} else {
				if figure.isKing{
					symbol = " Bk " // Черная дамка
				} else {
					symbol = " B " // Обычная черная
				}
			}
			fmt.Print(symbol) // Печать символа
		}
		fmt.Println() // Печать ряда
	}
	fmt.Println("   a  b  c  d  e  f  g  h")  // Вывод букв снизу
}

func (checkers *Checkers) IsValidMove(fromRow int, fromCol int, toRow int, toCol int, ) bool {
	// проверка границ
	// проврека нашего хода
	// проверка хода
	// проверка на взятие фигуры
	return true
}

func (checkers *Checkers) MakeMove(fromRow int, fromCol int, toRow int, toCol int, )  {
	// выполнение хода
	
}

func Parse() {
	// будет обрабатывать ввод в консоль, типо a2 -> в row, col
}

func StartGame(checkers *Checkers) {
	// начинаем обрабатывать действия игроков и менять доску соответствующе
}

// это будет где то извне происходить, так будет игра начинаться
func main() {
	var game = NewCheckers()
	game.PrintBoard()
	StartGame(game)
}
