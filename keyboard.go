package keyboard

import (
	"bufio"
	"os"      /* Пакет, который содержит Replacer */
	"strconv" // Содержит ParseFloat - преобразование строки в число и наоборот
	"strings" // Для функции TrimSpace - удаляет все символы-пропуски
)

func GetFloat() (float64, error) {

	reader := bufio.NewReader(os.Stdin)   // Средство чтения для получения текста с клавиатуры
	input, err := reader.ReadString('\n') // Чтение текста, введеного с клавиатуры до Enter
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64) // Преобразование строки в значение float64
	if err != nil {
		return 0, err
	}
	return number, nil
}
