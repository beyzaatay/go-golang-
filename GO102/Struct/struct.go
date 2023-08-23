package main

import "fmt"

type ogrenci struct {
	isim  string
	soyad string
	yas   int
}

func main() {

	var o1 ogrenci
	o1.soyad = "atay"
	o1.yas = 20
	o1.isim = "beyza"

	fmt.Println(o1)

	o2 := ogrenci{
		isim:  "beyza",
		soyad: "atay",
		yas:   20,
	}

	fmt.Println(o2)

	o3 := ogrenci{"beyza", "atay", 20}
	fmt.Println(o3)
	fmt.Println(o3.soyad)
}
