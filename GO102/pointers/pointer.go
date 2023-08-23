package main

import "fmt"

// pointerlar bellek adresini döndürür.
// pointerlar ile  değişkenin değerini değiştirebiliriz.
func main() {

	x := 10
	p := &x

	fmt.Println(p)

	// * pointerların bellek adresindeki değeri döndürür

	fmt.Println(*p)

	*p = 44

	fmt.Println(x)

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T,%v\n", p, p)

	var z string = "merhaba"

	var pz *string = &z

	fmt.Println(z, pz)
}
