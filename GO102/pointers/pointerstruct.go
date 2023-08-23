package main

import "fmt"

type Vatandas struct {
	isim string
	tc   int
}

func (v Vatandas) AyarlaValue() { //methon on value
	v.isim = "yusuf"
	v.tc = 978966
}

func (v *Vatandas) ayarlaPointer() { //method on pointer
	v.isim = "helin"
	v.tc = 55747
}

func main() {

	vat1 := Vatandas{isim: "beyza", tc: 12345}

	var ptr *Vatandas = &vat1

	fmt.Println(ptr)

	fmt.Println(&(ptr.tc))

	ptr.tc = 000

	fmt.Println(vat1)

	fmt.Println()

	ptr.AyarlaValue()
	fmt.Println(vat1)

	ptr.ayarlaPointer()
	fmt.Println(vat1)

}
