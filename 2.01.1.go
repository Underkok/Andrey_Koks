package main

import (
	"fmt"
	"strconv"
)

func main() {
	var one string = "104"
	var two int = 35
	oneInt, _ := strconv.Atoi(one)
	twoStr := strconv.Itoa(two)
	fmt.Println(oneInt + oneInt)
	fmt.Println(twoStr + twoStr)
}
