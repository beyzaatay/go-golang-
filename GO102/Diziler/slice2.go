package main

import "fmt"

func main() {

	var a []int = []int{11, 33, 44, 55}

	fmt.Println(a)

	b := append(a, 66)
	fmt.Println(b)

	a = append(a, 99)

	fmt.Println(a)

	var sehirler = []string{"istanbul", "sakarya"}
	sehirler = append(sehirler, "izmir")
	sehirler = append(sehirler, "adıyaman")

	fmt.Println(sehirler)

	sayilar := make([]int, 2)
	fmt.Println(sayilar)

	sayilar[0] = 3
	sayilar[1] = 8

	fmt.Println(sayilar)

	sayilar = append(sayilar, 9)
	fmt.Println(sayilar)
}
