package main

import "fmt"

func main() {

	//2 BOYUTLU DİZİ
	var dizi = [2][2]int{{0, 1}, {2, 4}}
	fmt.Println(dizi)

	fmt.Println(dizi[0][1])

	var dizi2D = [5][2]int{{0, 5}, {8, 9}, {0, 4}, {6, 7}, {0, 0}}
	fmt.Println(dizi2D[4][1])

	for i := 0; i <= i; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("dizi2d[%d][%d]=%d\n", i, j, dizi2D[i][j])

		}

	}

}
