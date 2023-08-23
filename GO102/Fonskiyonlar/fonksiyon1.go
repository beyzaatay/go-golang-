package main

import "fmt"

func merhaba() {
	fmt.Println("merhaba")

}

func merhabaPar(mesaj string) {
	fmt.Println("merhaba", mesaj)
}
func main() {

	merhaba()
	merhaba()

	merhabaPar("beyza")

}
