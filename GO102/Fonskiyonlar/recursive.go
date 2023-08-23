package main

import (
	"fmt"
)

func faktoriyel(sayi int) int {

	if sayi <= 1 {
		return 1
	}

	return sayi * faktoriyel(sayi-1)
}

func main() {

	giris := 13

	fmt.Printf("%d sayisinin faktöriyeli:%d", giris, faktoriyel(giris))

}
