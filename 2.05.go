package main

import "fmt"

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false

}
func getMax(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	max := values[0]
	for _, value := range values[1:] {
		if value > max {
			value = max
		}
	}
	return max
}
func main() {
	a := []string{"kvas", "kefir", "yogurt"}
	x := "kefir"
	if contains(a, x) {
		fmt.Println("Serega est okroshku na", x)
	} else {
		fmt.Println("Serega ne budet s etim est")
	}
	fmt.Println(getMax(1, 5, 10, 19, 24, 28, 18))

}
