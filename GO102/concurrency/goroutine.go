package main

import (
	"fmt"
	"time"
)

// eşzamanlı fonksiyonları çağırmak için kullandığımız bir fonksiyondur.

func yazdir(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
func main() {

	// thread 1 mb goroutine 2kb

	go yazdir("merhaba") // araya sıkışan kücük threadler gibi
	time.Sleep(1 * time.Second)

	yazdir("dünya")

}
