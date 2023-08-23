package main

import "fmt"

func fiboz(a int) int {

	if a <= 1 {
		return 1
	}
	return fiboz(a-1) + fiboz(a-2)
}

func main() {

	for i := 0; i <= 20; i++ {
		sonuc := fiboz(i)
		fmt.Printf("%d : %d", i, sonuc)
	}
}
