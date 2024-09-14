package cidigit

import (
	"gopackages/wordz"
)

func RandomCity() string {
	wordz.Words = []string{"Москва", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Челябинск"}
	randomTown := wordz.Random()
	return randomTown

}
