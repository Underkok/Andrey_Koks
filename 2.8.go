package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("ini.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var count int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Конец файла")
				break
			} else {
				fmt.Println("Ошибка чтения файла:", err)
				return
			}
		}
		count++
		fmt.Println(line)
	}

	fmt.Printf("Total strings: %d\n", count)
}
