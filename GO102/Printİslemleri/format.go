package main

import "fmt"

func main() {

	var a int = 5
	fmt.Printf("benim yasim: %d", a) //int da %d

	var sayi2 float32 = 3.323435

	fmt.Printf("pı: %.2f", sayi2)

	var dogruMu bool = true
	fmt.Printf("%t", dogruMu)

	var isim string = "beyza"
	var yas int = 20
	fmt.Printf("benim adim %s, %d yasindayim", isim, yas)

	var x int = 54
	var y float32 = 6.387423
	var z bool = true
	var d string = "beyza"

	fmt.Printf("%T %T %T %T", x, y, z, d)

	var mesaj string = "merhaba dünya"
	var yil int = 2023

	var tummesaj = fmt.Sprintf("%d yılına %s", yil, mesaj)
	fmt.Println(tummesaj)
}
