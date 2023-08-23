package main

import (
	"fmt"
	"strconv"
)

func main() {

	//TYPE CASTİNG
	var toplam int = 984
	var sayi int = 45

	var sonuc = float32(toplam) / float32(sayi)

	var s1 = int(sonuc)
	fmt.Println(sonuc)
	fmt.Println(s1)

	//TYPE CONVERSİON
	/*
		//String int e donusum
		var str = "1"
		var sayi, _ = strconv.Atoi(str)
		fmt.Println(sayi)

		var sonuc = sayi + 7
		fmt.Println(sonuc)

	*/

	//int stringe donusum

	var sayi = 1
	var str = strconv.Itoa(sayi)
	fmt.Println(str)

	fmt.Println("nasılsın" + str)

}
