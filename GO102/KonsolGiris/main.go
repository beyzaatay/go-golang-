package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*
		tarayici := bufio.NewScanner(os.Stdin)
		fmt.Print("birşeyler yazınız:")

		tarayici.Scan()

		verigirisi := tarayici.Text()

		fmt.Printf("bunu yazdiniz: %s", verigirisi)

	*/

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("hangi yıl dogdunuz:")

	scanner.Scan()
	verigir, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("şu anda %d yaşındasın", 2023-verigir)

}
