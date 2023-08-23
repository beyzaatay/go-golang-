package main

import "fmt"

func main() {

	fmt.Println("merhaba\n\n\nDünya")
	fmt.Println("merhaba\t\t\tDünya")

	fmt.Println("beyza atay \"yazilimci\"")

	fmt.Println("beyza" + "atay")
	fmt.Println("beyza", "atay") //bosluk birakir

	fmt.Println("go" +
		"programlama" +
		"dilidir.")

	isim := "beyza"
	soyisim := "atay"

	fmt.Println(isim, soyisim)
	fmt.Println(len(isim))

}
