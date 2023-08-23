package main

import "fmt"

func main() {

	sozluk := make(map[string]string)

	var eng string
	var tr string

	for i := 1; i <= 5; i++ {
		fmt.Print("ingilizce kelime giriniz:")
		fmt.Scan(&eng)

		fmt.Print("türkçe karşılığını giriniz:")
		fmt.Scan(&tr)

		sozluk[eng] = tr

	}
	fmt.Println(sozluk)
	fmt.Print("silmek istediğinizi ingilizce kelimeyi giriniz")
	fmt.Scan(&eng)

	delete(sozluk, eng)

	for k, v := range sozluk {
		fmt.Printf("eng:%s \t TR:%s \n", k, v)
	}

}
