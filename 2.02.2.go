package main

import (
	"fmt"
	"math"
)

func main() {
	var CircumFerence = 35.0
	var CircumRadius = new(float64)
	*CircumRadius = CircumFerence / (2 * math.Pi)
	CircumArea := math.Pi * (*CircumRadius) * (*CircumRadius)
	fmt.Printf("Radius %.2f\n", *CircumRadius)
	fmt.Printf("Area %.2f\n", CircumArea)
}
