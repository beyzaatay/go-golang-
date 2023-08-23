package main

import "fmt"

func main() {

	/*
		KOŞULSUZ SWİTCH

		sayi := 10

		switch sayi {
		case 5:
			fmt.Println("sayi 5'e eşittir.")
		case 10:
			fmt.Println("sayi 10'a eşittir.")
		default:
			fmt.Println("tanımsız.")
	*/

	switch sayi := 16; {
	case sayi == 5:
		fmt.Println("sayi 5'e eşittir")
		fallthrough
	case sayi <= 10:
		fmt.Println("sayi 10'dan küçüktür.")
	default:
		fmt.Println("tanımsız")
	}

	// KOŞULLU SWİTCH

	switch a := 6; {
	case a <= 10 && a >= 5:
		fmt.Println("sayi 10 ile 5 arasında")
	case a <= 0 && a <= 5:
		fmt.Println("sayi 0 ile 5 arasında")
	default:
		fmt.Println("aralık tanımsız.")
	}

}
