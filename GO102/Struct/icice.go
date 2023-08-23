package main

import "fmt"

type adres struct {
	sehir string
	ulke  string
}

type kisi struct {
	isim      string
	yas       int
	kisiadres adres
}

func main() {

	k1 := kisi{
		isim: "beyza",
		yas:  20,
		kisiadres: adres{
			sehir: "istanbul",
			ulke:  "türkiye",
		},
	}

	fmt.Println("isim:", k1.isim)
	fmt.Println("şehir:", k1.kisiadres.sehir)
	fmt.Println("ülke:", k1.kisiadres.ulke)
}
