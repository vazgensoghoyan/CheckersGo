package checkers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Checkers struct {
	Board       [][]Figure
	IsWhiteTurn bool // Чей ход
}

type Move struct {
	fromRow int
	fromCol int
	toRow   int
	toCol   int
}

type Figure struct {
	IsNone  bool
	IsWhite bool
	IsKing  bool
}

func NewCheckers() *Checkers {
	board := make([][]Figure, 8)

	for i := range board {
		board[i] = make([]Figure, 8)
	}

	// Расставляем фигуры
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			// Все клетки по умолчанию пустые
			board[row][col] = Figure{IsNone: true, IsWhite: false, IsKing: false}

			// Шашки только на черных полях (где row+col нечетное)
			if (row+col)%2 == 1 {
				if row < 3 { // Первые 3 ряда - белые шашки
					board[row][col] = Figure{IsNone: false, IsWhite: true, IsKing: false}
				} else if row > 4 { // Последние 3 ряда - черные шашки
					board[row][col] = Figure{IsNone: false, IsWhite: false, IsKing: false}
				}
				// Ряды 3 и 4 (индексы row=3 и row=4) остаются пустыми
			}
		}
	}

	return &Checkers{Board: board, IsWhiteTurn: true}
}

func (checkers *Checkers) PrintBoard() {
	for row := 0; row < 8; row++ {
		fmt.Printf("%d ", 8-row)
		for col := 0; col < 8; col++ {
			figure := checkers.Board[row][col]
			var symbol string
			if figure.IsNone {
				symbol = " . " // Пустая клетка
			} else if figure.IsWhite {
				if figure.IsKing {
					symbol = " Wk " // Белая дамка
				} else {
					symbol = " W " // Обычная белая
				}
			} else {
				if figure.IsKing {
					symbol = " Bk " // Черная дамка
				} else {
					symbol = " B " // Обычная черная
				}
			}
			fmt.Print(symbol) // Печать символа
		}
		fmt.Println() // Печать ряда
	}
	fmt.Println("   a  b  c  d  e  f  g  h") // Вывод букв снизу
}

func (c *Checkers) IsValidMove(move Move) (bool, string) {
	fromRow, fromCol := move.fromRow, move.fromCol
	toRow, toCol := move.toRow, move.toCol

	// Проверка границ
	if fromRow < 0 || fromRow > 7 || fromCol < 0 || fromCol > 7 ||
		toRow < 0 || toRow > 7 || toCol < 0 || toCol > 7 {
		return false, "Координаты вне границ доски"
	}

	// Проверка, что на исходной клетке есть фигура
	fromFigure := c.Board[fromRow][fromCol]
	if fromFigure.IsNone {
		return false, "На исходной клетке нет фигуры"
	}

	// Проверка, что ходит правильный игрок
	if fromFigure.IsWhite != c.IsWhiteTurn {
		if c.IsWhiteTurn {
			return false, "Сейчас ход белых"
		}
		return false, "Сейчас ход черных"
	}

	// Проверка, что целевая клетка пуста
	if !c.Board[toRow][toCol].IsNone {
		return false, "Целевая клетка занята"
	}

	// Проверка, что ход по диагонали
	rowDiff := toRow - fromRow
	colDiff := toCol - fromCol
	if abs(rowDiff) != abs(colDiff) {
		return false, "Шашки ходят только по диагонали"
	}

	// Обычный ход (на одну клетку)
	if abs(rowDiff) == 1 {
		// Обычные шашки ходят только вперед
		if !fromFigure.IsKing {
			if fromFigure.IsWhite && rowDiff < 0 {
				return false, "Белые шашки ходят вниз"
			}
			if !fromFigure.IsWhite && rowDiff > 0 {
				return false, "Черные шашки ходят вверх"
			}
		}
		return true, ""
	}

	// Взятие (прыжок через шашку)
	if abs(rowDiff) == 2 {
		middleRow := (fromRow + toRow) / 2
		middleCol := (fromCol + toCol) / 2
		middleFigure := c.Board[middleRow][middleCol]

		// Проверка, что есть фигура противника для взятия
		if middleFigure.IsNone {
			return false, "Нет фигуры для взятия"
		}
		if middleFigure.IsWhite == fromFigure.IsWhite {
			return false, "Нельзя брать свою фигуру"
		}

		return true, ""
	}

	return false, "Недопустимый ход"
}

func (c *Checkers) MakeMove(move Move) bool {
	valid, msg := c.IsValidMove(move)
	if !valid {
		fmt.Println("Ошибка:", msg)
		return false
	}

	fromRow, fromCol := move.fromRow, move.fromCol
	toRow, toCol := move.toRow, move.toCol

	// Перемещаем фигуру
	c.Board[toRow][toCol] = c.Board[fromRow][fromCol]
	c.Board[fromRow][fromCol] = Figure{IsNone: true}

	// Если было взятие, убираем съеденную фигуру
	if abs(toRow-fromRow) == 2 {
		middleRow := (fromRow + toRow) / 2
		middleCol := (fromCol + toCol) / 2
		c.Board[middleRow][middleCol] = Figure{IsNone: true}
	}

	// Превращение в дамку
	if !c.Board[toRow][toCol].IsKing {
		if c.Board[toRow][toCol].IsWhite && toRow == 7 {
			c.Board[toRow][toCol].IsKing = true
			fmt.Println("Белая шашка стала дамкой!")
		} else if !c.Board[toRow][toCol].IsWhite && toRow == 0 {
			c.Board[toRow][toCol].IsKing = true
			fmt.Println("Черная шашка стала дамкой!")
		}
	}

	// Переключаем ход
	c.IsWhiteTurn = !c.IsWhiteTurn
	return true
}

func ParseMove(input string) (Move, error) {
	parts := strings.Fields(input)
	if len(parts) != 2 {
		return Move{}, fmt.Errorf("формат: a2 b3")
	}

	from, err := parsePosition(parts[0])
	if err != nil {
		return Move{}, err
	}

	to, err := parsePosition(parts[1])
	if err != nil {
		return Move{}, err
	}

	return Move{
		fromRow: from[0],
		fromCol: from[1],
		toRow:   to[0],
		toCol:   to[1],
	}, nil
}

func parsePosition(pos string) ([2]int, error) {
	if len(pos) != 2 {
		return [2]int{}, fmt.Errorf("неверный формат позиции: %s", pos)
	}

	col := int(pos[0] - 'a')
	row := 8 - int(pos[1]-'0')

	if col < 0 || col > 7 || row < 0 || row > 7 {
		return [2]int{}, fmt.Errorf("позиция вне доски: %s", pos)
	}

	return [2]int{row, col}, nil
}

// это будет где то извне происходить, так будет игра начинаться
func StartGame(c *Checkers) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		c.PrintBoard()

		if c.IsWhiteTurn {
			fmt.Print("\nХод белых (например: c3 d4): ")
		} else {
			fmt.Print("\nХод черных (например: c6 d5): ")
		}

		if !scanner.Scan() {
			break
		}

		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if input == "quit" || input == "exit" {
			fmt.Println("Игра завершена!")
			break
		}

		move, err := ParseMove(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		c.MakeMove(move)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
