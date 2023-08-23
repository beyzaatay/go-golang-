package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(" toplama icin :1 \n çikarma icin :2 \n carpma icin:3 \nbolme icin:4 \n yapılacak işlemi giriniz:")
	scanner.Scan()
	islem, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("birinci sayiyi girin:")
	scanner.Scan()
	sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("ikinci sayiyi girin:")
	scanner.Scan()
	sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)

	if islem == 1 {
		fmt.Println("toplamı:", sayi1+sayi2)

	} else if islem == 2 && sayi1 > sayi2 {
		fmt.Println("farkı:", sayi1-sayi2)

	} else if islem == 2 && sayi2 > sayi1 {
		fmt.Println("farkı:", sayi2-sayi1)

	} else if islem == 3 {
		fmt.Println("carpimi:", sayi1*sayi2)

	} else if islem == 4 {
		fmt.Println("bolumu:", sayi1/sayi2)

	} else {
		fmt.Println("hatalı giriş yaptınız!!!")
	}

}
