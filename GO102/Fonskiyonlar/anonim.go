package main

import "fmt"

func merhaba2(sayi int) int {
	return sayi
}
func main() {

	a := func(toplam int) int {
		return toplam * 4
	}(6)

	fmt.Println(a)

	fmt.Println(merhaba2(a))

}
