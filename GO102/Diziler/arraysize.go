package main

import "fmt"

func main() {
	// array statik list dinamik yapılıdır.

	arr := [...]int{4, 5, 7, 4, 3} //kaç eleman gelecek bilmyiorsak.
	fmt.Println(arr)

	arr2 := [4]int{3, 4, 5, 6}

	fmt.Println(len(arr2))
	fmt.Println(cap(arr2)) // dizinin kaç elemeana kadar alacağını gösterir.

	for i := 0; i <= len(arr)-1; i++ {
		fmt.Println(arr2[i])
	}

}
