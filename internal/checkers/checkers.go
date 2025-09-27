package checkers
type Board struct{
	cells[][]int // Доска, это двумерный массив
}
type Fiegure struct{
	row int		// строка
	col int		// колонна
	player int  // игрок 1 или 2
	isKing bool // Королева или нет(t or f)
}
