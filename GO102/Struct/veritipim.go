package main

import (
	"fmt"
	"time"
)

type ustunf float64

func (y ustunf) Carpİki() ustunf {
	return y * 2
}

type Aylar int

const (
	Ocak Aylar = 1 + iota
	Subat
	Mart
	Nisan
	Mayis
	Haziran
	Temmuz
	Agustos
	Eylul
	Ekim
	Kasim
	Aralik
)

func main() {

	var u1 ustunf = 7.8

	fmt.Printf("%T,%v\n", u1, u1)

	var u2 float64 = 9.4

	fmt.Println(u1 + ustunf(u2))

	y1 := ustunf(4.8)

	fmt.Println(y1.Carpİki())

	_, ay, _ := time.Now().Date()
	fmt.Println(ay)

	fmt.Println(Aylar(ay))

}
