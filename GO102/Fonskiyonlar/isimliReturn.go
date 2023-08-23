package main

import (
	"fmt"
	"math"
)

const pi = 3.14

func hesapla(x, y int) (carp int, bol int) {
	carp = x * y
	bol = x / y
	return
}

func daire(yaricap float64) (cevre float64, alan float64) {
	cevre = 2 * pi * yaricap
	alan = pi * (math.Pow(yaricap, 2))
	return

}

func main() {

	carpim, bolum := hesapla(96, 6)

	fmt.Println("carpımı \n ", carpim)
	fmt.Println("bolum", bolum)

	cevre, alan := daire(8)

	fmt.Println("cevre: \n", cevre)
	fmt.Println("alan:", alan)
}
