package main

import "fmt"

func main() {

	var isim string
	var yas int

	fmt.Print("adınızı giriniz:")
	fmt.Scan(&isim)
	fmt.Print("yasınızı girin:")
	fmt.Scan(&yas)

	fmt.Printf("merhaba %s yasınız : %d", isim, yas)

	var meyve string
	fmt.Print("meyve girin:")
	fmt.Scanf("%s", &meyve)

	fmt.Println("meyveniz", meyve)

}
