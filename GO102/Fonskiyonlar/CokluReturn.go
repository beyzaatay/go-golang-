package main

import "fmt"

func islem(x, y int) (int, int) {
	return x + y, x - y
}

func maxmin(a, b int) (int, int) {
	if a > b {
		return a, b
	} else {
		return b, a
	}
}
func main() {

	topla, cıkar := islem(10, 9)

	fmt.Printf("toplam:%d ,fark:%d \n", topla, cıkar)

	buyuk, kucuk := maxmin(70, 79)

	fmt.Printf("max:%d min:%d", buyuk, kucuk)
}
