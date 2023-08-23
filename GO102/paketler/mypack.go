package main

import (
	"GO102/newpack"
	m "GO102/newpack/matematik"
	"fmt"
)

func main() {

	newpack.NewOrder()
	newpack.Staart()

	fmt.Println(m.Carp(4))

}
