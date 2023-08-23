package main

import "fmt"

func isaret(ptr *int) {
	*ptr = 55
}

func onlacarp(sayi *int) int {
	return *sayi * 10
}
func main() {

	x := 93

	isaret(&x)
	fmt.Println(x)

	fmt.Println(onlacarp(&x))

}
