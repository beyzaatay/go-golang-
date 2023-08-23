package main

import "fmt"

type Ogrenci struct {
	isim    string
	yas     int
	notlar  []int
	gectiMi bool
}

func (o Ogrenci) AdGoster() string {
	return o.isim

}

func (ogr Ogrenci) BilgiGoster() {
	fmt.Println("öğrenci adı:", ogr.isim)
	fmt.Println("öğrenci yas:", ogr.yas)
	fmt.Println("öğrenci notlar:", ogr.notlar)
	fmt.Println("öğrenci geçti mi:", ogr.gectiMi)
}

func (not Ogrenci) Ortalama() float64 {
	toplam := 0
	for _, i := range not.notlar {
		toplam += i
	}
	return float64(toplam) / float64(len(not.notlar))
}

func (gec Ogrenci) gectimi() bool {
	if gec.Ortalama() <= 50 {
		return false
	} else {
		return true
	}
}

func main() {
	o1 := Ogrenci{isim: "beyza", yas: 20, notlar: []int{90, 40, 89}}

	fmt.Println(o1.AdGoster())

	o1.BilgiGoster()

	fmt.Println(o1.Ortalama())

	o1.gectimi()
}
