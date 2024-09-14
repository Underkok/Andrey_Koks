package main

import (
	"fmt"

	"github.com/fatih/color"

	"gopackages/cidigit"
	newcolor "gopackages/color"
	. "gopackages/wordz"

	"github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()

	fmt.Println("Hello world")

	color.Red("hello world andrey")

	color.Green("i am not green")

	color.Blue("Ahahaha, yes, i am blue guy")

	fmt.Println(Hello)

	fmt.Println(Random())

	randomDigit := cidigit.RandomDigit()
	fmt.Println("Случайное число:", randomDigit)

	city := cidigit.RandomCity()
	fmt.Println("Случайный город:", city)

	fmt.Println(xstrings.Shuffle(cidigit.RandomCity()))
	fmt.Println(xstrings.Shuffle(cidigit.RandomDigit()))

}
