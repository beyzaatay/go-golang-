package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	k_adi := "beyza"
	k_sifre := "12345"

	giris_hakki := 3

	fmt.Println("***********LOGİN*************")

	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("kullanıcı adını giriniz:")

		scanner.Scan()

		giris_kadi := scanner.Text()

		fmt.Print("kullanıcı şifresini giriniz:")

		scanner.Scan()
		giris_ksifre := scanner.Text()

		if k_adi != giris_kadi && k_sifre != giris_ksifre {
			fmt.Println("kullanıcı adı ve şireniz yanlış")
			giris_hakki--
		} else if k_adi == giris_kadi && k_sifre != giris_ksifre {
			fmt.Println("şifreniz yanlış.")
			giris_hakki--
		} else if k_adi != giris_kadi && k_sifre == giris_ksifre {
			fmt.Println("kullanıcı adınız yanlış")
			giris_hakki--
		} else {
			fmt.Println("hosgeldiniz")
			break
		}

		if giris_hakki == 0 {
			fmt.Println("giriş hakkınız bitmiştir.")
			break
		}

	}
}
