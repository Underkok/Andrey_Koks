package main

import "fmt"

func main() {
	var a *int
	var b int = 25
	a = &b
	fmt.Println((*a))
	*a = 120
	fmt.Println(b)

}
