package main

import "fmt"

func main() {

	//break
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}

	//continue

	for a := 1; a <= 5; a++ {
		if a == 3 {
			continue // 3'ü atlar
		}

		fmt.Println(a)
	}

}
