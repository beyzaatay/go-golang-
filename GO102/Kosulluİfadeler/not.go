package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("vize notunu giriniz:")

	scanner.Scan()
	vizenotu, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("final notunu giriniz:")
	scanner.Scan()
	finalnotu, _ := strconv.ParseFloat(scanner.Text(), 64)

	ortalama := (vizenotu * 0.4) + (finalnotu * 0.6)
	fmt.Println("ortalama:", ortalama)

	if ortalama > 85 && ortalama <= 100 {
		println("AA")
	} else if ortalama >= 70 && ortalama < 85 {
		fmt.Println("BB")
	} else if ortalama < 50 && ortalama >= 0 {
		fmt.Println("FF")
	} else {
		fmt.Println("0 ile 100 arasında bir deger girin.")
	}

}
