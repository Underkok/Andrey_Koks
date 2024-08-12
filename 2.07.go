package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Ошибка обнаружения файла", err)
		return
	}
	defer file.Close()

	outFile, err := os.Create("data/data_out.txt")
	if err != nil {
		fmt.Println("Ошибка создания файлв", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)

	// обработка паники
	defer func() {
		if panicError := recover(); panicError != nil {
			fmt.Println("Критическая ошибка", panicError)
			fmt.Println("Содержимое файла data/data_out.txt:")
			data, err := os.ReadFile("data/data_out.txt")
			if err != nil {
				fmt.Println("Ошибка чтения файла:", err)
			} else {
				fmt.Println(string(data))
			}
		}
	}()

	row := 1
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) != 3 {
			panic(fmt.Sprintf("parse error: empty field on string %d", len(fields)))
		}

		for symbol, field := range fields {
			if strings.TrimSpace(field) == "" {
				panic(fmt.Sprintf("parse.error: empty field on string %d", symbol+1))
			}
		}

		// Формируем строку для записи
		outLine := fmt.Sprintf("row(%d)\nName: %s\nAddress: %s\nCity: %s\n\n\n",
			row, fields[0], fields[1], fields[2])

		// Записываем строку
		_, err = outFile.WriteString(outLine)
		if err != nil {
			fmt.Println("Ошибка записи в файл", err)
			return
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка сканирования", err)
		return
	}

	fmt.Println(`Данные успешно обработаны и записаны в файл`)

}
