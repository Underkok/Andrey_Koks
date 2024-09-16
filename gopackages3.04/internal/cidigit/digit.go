package cidigit

import (
	"fmt"
	"strconv"

	"gopackages/wordz"
)

func RandomDigit() string {
	index := wordz.Random()
	digit, _ := strconv.ParseInt(index, 10, 64)
	return fmt.Sprintf("%d", digit)
}
