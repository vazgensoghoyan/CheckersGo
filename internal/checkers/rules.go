package checkers

import "fmt"

func GetName(row, col int) string {
	letters := "abcdefgh" // для обозначения полей
	if row < 0 || row >= 8 || col < 0 || col >= 8 {
		return "" // Проверка на ввод поля для хода, потом доделать вывод
	}
	letter := string(letters[col]) // мы преобразуем из цифры букву для нумерации коллон
	number := 8 - row              // номер строки, тк первая строка сверху по индексу 0, но в поле для шашек она 8
	return fmt.Sprintf("%s%d", letter, number)
}