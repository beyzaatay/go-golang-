package main

import "fmt"

type calisan struct {
	isim      string
	yas       int
	maas      int
	isHaveKid bool
}

type mudur struct {
	calisan
	unvan string
}

func main() {

	a1 := calisan{
		isim:      "beyza",
		yas:       19,
		maas:      79,
		isHaveKid: false,
	}

	fmt.Println(a1)

	y1 := mudur{
		calisan: calisan{
			isim:      "ahmet",
			yas:       29,
			maas:      100,
			isHaveKid: false,
		},
		unvan: "şef",
	}
	fmt.Println(y1)

	kurucu := struct { //anonim struct
		isim    string
		sermaye int
	}{isim: "hasan", sermaye: 68999}

	fmt.Println(kurucu)
}
