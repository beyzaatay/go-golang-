package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//nil --> null demektir.

	{
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("sayı giriniz:")
		scanner.Scan()

		a, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			// fmt.Println(err.Error())
			log.Fatal("hata", err)
		}

		fmt.Println(a)
	}

	file, er := os.Open("/main.go")
	if er != nil {
		fmt.Println(er)
		//fmt.Println(errors.New("DOSYA YOK!!!!"))
	}

	fmt.Println(file.Name(), "başarılı bir şekilde açıldı.")

}
