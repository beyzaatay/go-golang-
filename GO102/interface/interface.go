package main

import "fmt"

type yuruyen interface {
	Yuru() string
}

type Insan string
type Kedi string

func (i Insan) Yuru() string {
	return "ben bir insanım ve yürüyorum"
}

func (k Kedi) Yuru() string {
	return "ben bir kediyim ve yürüyorum"
}

// interfacedan oluşturulan func
func yurume(y yuruyen) {
	fmt.Println(y.Yuru())
}

func main() {
	var kedi Kedi
	var insan Insan

	yurume(insan)
	yurume(kedi)
}
