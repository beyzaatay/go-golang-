package main

import "fmt"

func main() {

	yas := 16

	if yas >= 18 {
		fmt.Println("mekana girebilirsin")

	} else if yas < 18 && yas >= 14 {
		fmt.Println("ailenle mekana girebilirsin")
	} else {
		fmt.Println("mekana giremezsin.")
	}
}
