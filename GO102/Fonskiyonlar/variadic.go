package main

import "fmt"

func merhaba1(hello ...string) {
	fmt.Println(hello)
}

func toplama(param ...int) int {
	sonuc := 0
	for _, sayi := range param {
		sonuc += sayi
	}
	return sonuc

}

func main() {

	merhaba1("ayşe", "fatma", "buse")

	fmt.Println(toplama(5, 9))
}
