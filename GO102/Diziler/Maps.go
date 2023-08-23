package main

import (
	"fmt"
)

func main() {

	var ogrenci map[int]string = map[int]string{
		10: "beyza",
		35: "izmir",
		68: "buse",
	}

	fmt.Println(ogrenci)
	fmt.Println(ogrenci[35])

	ogrenci[10] = "atay"
	fmt.Println(ogrenci)

	plaka := make(map[int]string)

	plaka[10] = "balıkesir"
	plaka[54] = "sakarya"

	fmt.Println(plaka)

	unvanlar := make(map[string]string)

	unvanlar["doctor"] = "dr"
	unvanlar["professor"] = "pr"
	unvanlar["docent"] = "doc"

	fmt.Println(unvanlar)

	delete(unvanlar, "doctor")

	fmt.Println(unvanlar)

	for key, value := range unvanlar {
		fmt.Printf("kısaltılması:%v tam hali:%v \n", value, key)
	}

}
