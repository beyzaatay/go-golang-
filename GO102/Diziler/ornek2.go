package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	ulkeler := make([]string, 5)

	for i := 0; i <= 6; i++ {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("ülke giriniz:")

		scanner.Scan()
		ulkegiris := scanner.Text()
		ulkeler = append(ulkeler, ulkegiris)
	}

	for index, eleman := range ulkeler {

		if index <= 5 {
			continue
		}
		fmt.Printf("girdiniz %d ulke %s \n", index-5, eleman)
	}
}
