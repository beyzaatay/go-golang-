package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Remove("data.txt")

	err := os.Remove("veri.html")
	if err != nil {
		fmt.Println("dosya bulunamadı")
	}

}
