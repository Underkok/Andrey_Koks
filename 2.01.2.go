package main

import (
	"fmt"
	"math"
)

func main() {
	var AmericanVelocity float64 = 130.0
	var EuropeanVelocity float64 = 124.4
	ev := (EuropeanVelocity * 3600) / 1000
	av := (AmericanVelocity * 3600) / 1609
	rounded_ev := math.Round(ev*100) / 100 // * 100
	rounded_av := math.Round(av*100) / 100 // * 100
	fmt.Println(rounded_ev)
	fmt.Println(rounded_av)

}
