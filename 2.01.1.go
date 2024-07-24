package main

import (
	"fmt"
	"strconv"
)

func main() {
	var one string = "104"
	var two int = 35
	oneInt, err := strconv.Atoi(one)
	twoStr := strconv.Itoa(two)
	fmt.Println(oneInt + oneInt)
	fmt.Println(twoStr + twoStr)
	if err != nil {
		fmt.Println("Невозможно сконвертировать строку в число!")
	}
}
