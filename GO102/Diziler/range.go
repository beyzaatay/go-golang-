package main

import "fmt"

func main() {

	var a []int = []int{1, 5, 6, 5, 4}

	for index, eleman := range a {
		fmt.Printf("%d: %d \n", index, eleman)
	}

	isim := "beyza atay"

	for i, harf := range isim {
		fmt.Printf("%d =%c \n", i, harf)
	}
}
