package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.ToLower("BEYZA ATAY")) // tüm harfleri kücüğe cevirir
	fmt.Println(strings.ToUpper("beyza atay"))

	//INDEX

	fmt.Println(strings.Index("turkiye", "u")) //u kacıncı indexte

	// COUNT

	fmt.Println(strings.Count("sakarya", "a")) // a dan kac tane var
	fmt.Println(strings.Count("ankara", ""))

	// Contains

	fmt.Println(strings.Contains("beyza", "a")) // a var mı yok mu

	// compare -> karşılaştırma

	a1 := "hoşgeldin"
	a2 := "hoşgeldin kullanıcı"
	a3 := "hoşgeldin"

	fmt.Println(strings.Compare(a1, a2))
	fmt.Println(strings.Compare(a1, a3))

	// replace ->yer değiştirme

	fmt.Println(strings.Replace("merhaba beyza, merhaba selin", "merhaba", "hoşgeldin", 1))
	fmt.Println(strings.ReplaceAll("merhaba beyza,merhaba selin", "merhaba", "hoşgeldin"))

	//fields --> boşlukları eler

	sehirler := "     istanbul     sakarya      balıkesir  nevşehir"
	fields := strings.Fields(sehirler)
	fmt.Printf("%q\n", fields)

	//split --> ayırmak

	fmt.Println(strings.Split("isim-şehir-hayvan", "-"))
	fmt.Println(strings.Split("22.11.22", "."))

	// split slice

	gunler := "pazartesi,salı,çarşamba"
	ayir := strings.Split(gunler, ",")
	fmt.Println(ayir[1])

	// split after

	fmt.Println(strings.SplitAfter("a.b.c.d", "."))

}
