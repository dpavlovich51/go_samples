package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(zigzagConversion("PAHNAPLSIIGYIR", 3))
}

func zigzagConversion(s string, numRows int) string { // Граничные случаи: если строка пустая, длина <= 1 или numRows == 1
	if len(s) <= 1 || numRows <= 1 {
		return s
	}

	// Создаем массив strings.Builder для каждой строки
	rows := make([]strings.Builder, numRows)
	currentRow := 0
	step := 1 // Направление: 1 (вниз), -1 (вверх)

	// Распределяем символы по строкам
	for _, char := range s {
		// Добавляем текущий символ в соответствующую строку
		rows[currentRow].WriteRune(char)

		// Меняем направление на первой и последней строке
		switch currentRow {
		case 0:
			step = 1 // Движемся вниз
		case numRows - 1:
			step = -1 // Движемся вверх
		}
		currentRow += step
	}

	// Объединяем все строки в результат
	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(rows[i].String())
	}

	return result.String()
}
