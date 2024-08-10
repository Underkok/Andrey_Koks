package main

import (
	"bufio"
	"fmt"
	"os"
)


func main () {
	file,err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла",err)
		return
	} 
	defer file.Close()
	out,err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Ошибка создания файлв",err)
		return
	}
	defer out.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(out, "%d. %s\n", lineNumber, line)
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка сканировки, всё-таки окрошка на кефире не такая вкусная",err)
		return
	}
	fmt.Println("Данные успешно скопированы")
}
