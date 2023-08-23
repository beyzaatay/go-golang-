package main

import "fmt"

// defer -> en son işlenecek olanı gösterir , geciktirir
func sayim() {
	defer fmt.Println("sayim bitti")
	fmt.Println("sayim basliyor...")
}

func sira() {
	defer fmt.Println("birinci")
	defer fmt.Println("ikinci")
	defer fmt.Println("ücüncü")
}
func main() {
	sayim()

	sira()
}
