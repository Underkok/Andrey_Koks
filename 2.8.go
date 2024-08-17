package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("ini.txt")
	if err != nil {
		fmt.Errorf("Ошибка открытия файлa: %w", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	countStrings := 0

	for scanner.Scan() {
		countStrings++
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("Ошибка сканирования файла: %w", err)
		return
	}

	fmt.Println("Total strings: %d\n", countStrings)

}
