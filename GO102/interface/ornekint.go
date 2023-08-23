package main

import "fmt"

const PI = 3.14

type hesapla interface {
	alan() float64
	cevre() float64
}
type daire struct {
	yaricap float64
}

type dikdortgen struct {
	kisakenar float64
	uzunkenar float64
}

func (d daire) alan() float64 {
	return PI * d.yaricap * d.yaricap
}

func (g dikdortgen) alan() float64 {
	return g.uzunkenar * g.kisakenar
}

func (d daire) cevre() float64 {
	return 2.0 * PI * d.yaricap
}
func (g dikdortgen) cevre() float64 {
	return 2.0 * (g.uzunkenar * g.kisakenar)
}

func hesapYap(h hesapla) {
	fmt.Println(h)
	fmt.Println(h.alan())

	fmt.Printf(" veri tipi:%T\n", h)

	fmt.Println(h.cevre())
}

func main() {
	daire1 := daire{5.0}
	dikdortgen1 := dikdortgen{6.0, 9.0}

	hesapYap(daire1)
	hesapYap(dikdortgen1)

}
