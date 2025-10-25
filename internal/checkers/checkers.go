package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Checkers struct {
	board       [][]Figure
	isWhiteTurn bool // Чей ход
}

type Move struct {
	fromRow int
	fromCol int
	toRow   int
	toCol   int
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
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			// Все клетки по умолчанию пустые
			board[row][col] = Figure{isNone: true, isWhite: false, isKing: false}
			
			// Шашки только на черных полях (где row+col нечетное)
			if (row+col)%2 == 1 {
				if row < 3 { // Первые 3 ряда - белые шашки
					board[row][col] = Figure{isNone: false, isWhite: true, isKing: false}
				} else if row > 4 { // Последние 3 ряда - черные шашки
					board[row][col] = Figure{isNone: false, isWhite: false, isKing: false}
				}
				// Ряды 3 и 4 (индексы row=3 и row=4) остаются пустыми
			}
		}
	}

	return &Checkers{board: board, isWhiteTurn: true}
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
				if figure.isKing {
					symbol = " Wk " // Белая дамка
				} else {
					symbol = " W " // Обычная белая
				}
			} else {
				if figure.isKing {
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
	fromFigure := c.board[fromRow][fromCol]
	if fromFigure.isNone {
		return false, "На исходной клетке нет фигуры"
	}

	// Проверка, что ходит правильный игрок
	if fromFigure.isWhite != c.isWhiteTurn {
		if c.isWhiteTurn {
			return false, "Сейчас ход белых"
		}
		return false, "Сейчас ход черных"
	}

	// Проверка, что целевая клетка пуста
	if !c.board[toRow][toCol].isNone {
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
		if !fromFigure.isKing {
			if fromFigure.isWhite && rowDiff < 0 {
				return false, "Белые шашки ходят вниз"
			}
			if !fromFigure.isWhite && rowDiff > 0 {
				return false, "Черные шашки ходят вверх"
			}
		}
		return true, ""
	}

	// Взятие (прыжок через шашку)
	if abs(rowDiff) == 2 {
		middleRow := (fromRow + toRow) / 2
		middleCol := (fromCol + toCol) / 2
		middleFigure := c.board[middleRow][middleCol]

		// Проверка, что есть фигура противника для взятия
		if middleFigure.isNone {
			return false, "Нет фигуры для взятия"
		}
		if middleFigure.isWhite == fromFigure.isWhite {
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
	c.board[toRow][toCol] = c.board[fromRow][fromCol]
	c.board[fromRow][fromCol] = Figure{isNone: true}

	// Если было взятие, убираем съеденную фигуру
	if abs(toRow-fromRow) == 2 {
		middleRow := (fromRow + toRow) / 2
		middleCol := (fromCol + toCol) / 2
		c.board[middleRow][middleCol] = Figure{isNone: true}
	}

	// Превращение в дамку
	if !c.board[toRow][toCol].isKing {
		if c.board[toRow][toCol].isWhite && toRow == 7 {
			c.board[toRow][toCol].isKing = true
			fmt.Println("Белая шашка стала дамкой!")
		} else if !c.board[toRow][toCol].isWhite && toRow == 0 {
			c.board[toRow][toCol].isKing = true
			fmt.Println("Черная шашка стала дамкой!")
		}
	}

	// Переключаем ход
	c.isWhiteTurn = !c.isWhiteTurn
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

		if c.isWhiteTurn {
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
