package main

import "fmt"

func carpma(x, y int) int {

	return x * y
}

func ikiyeKatla(a int) int {

	return a * 2
}

func main() {

	fmt.Println(carpma(10, 5))

	for i := 0; i <= 10; i++ {
		fmt.Println(ikiyeKatla(ikiyeKatla(i)))
	}

}
