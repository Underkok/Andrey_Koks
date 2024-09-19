package main

import (
	"errors"
	"fmt"
	"myapp/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := internal.NewCustomer("Dmitriy", 24, 10000, 100, true)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	discount, _ := cust.CalcDiscount()
	fmt.Println(discount)

	finalPrice, err := internal.CalcPrice(*cust, 1800)
	if err != nil {
		fmt.Println("Error", err)

	} else {
		fmt.Println("Final Price:", finalPrice)
	}

}
