package main

import (
	"fmt"

	"gopackages/cidigit"
	newcolor "gopackages/color"
	. "gopackages/wordz"

	"github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()

	fmt.Println("Hello world")

	fmt.Println(Hello)

	fmt.Println(Random())

	randomDigit := cidigit.RandomDigit()
	fmt.Println("Случайное число:", randomDigit)

	city := cidigit.RandomCity()
	fmt.Println("Случайный город:", city)

	fmt.Println(xstrings.Shuffle(cidigit.RandomCity()))
	fmt.Println(xstrings.Shuffle(cidigit.RandomDigit()))

}
