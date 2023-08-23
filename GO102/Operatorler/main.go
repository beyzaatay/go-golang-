package main

import "fmt"

func main() {

	var x = 28
	var y = 4

	fmt.Println("X+Y", x+y)

	x++

	fmt.Println(x) //bir arttırılmıs hali

	var sayi1 = 6
	var sayi2 = 18

	var toplam = sayi2 + sayi1
	fmt.Println(toplam)

	toplam += sayi1 // toplam=toplam+sayi1
	fmt.Println(toplam)

	var a, b = 3, 4

	fmt.Println(a == b)
	fmt.Println(a != b)

	var s, d, z, t = 1, 1, 6, 8
	fmt.Println(s == d && z != t) //ve operatörü
}
