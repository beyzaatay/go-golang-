package main

import "fmt"

func main() {
	// slice lar esnek ve dinamik görünümlüdür.

	dizi := [5]int{1, 3, 5, 7, 9}

	var dilim []int = dizi[0:5]

	fmt.Println(dilim)

	a := [3]string{"beyza", "atay", "istanbul"}
	b := a[1:]

	fmt.Println("slice", b)
	fmt.Println("array", a)
}
