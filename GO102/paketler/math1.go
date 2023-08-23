package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	pisayisi := math.Pi

	fmt.Println(pisayisi)

	fmt.Println(math.Pow(2, 10)) // üs alma

	fmt.Println(math.Sqrt(64)) // kök alma

	fmt.Println(math.Abs(-89.8)) // mutlak alma

	fmt.Println(math.Sin(45))

	fmt.Println(math.Mod(56, 7))

	fmt.Println(math.Round(34.6)) //tama yuvarlama

	fmt.Println(math.Ceil(67.1)) // tavan değer
	fmt.Println(math.Floor(3.7)) //taban değer

	fmt.Println(math.Max(67, 89))

	fmt.Println(math.IsNaN(89)) //sayı mı?

	rand.Seed(time.Now().UnixNano())

	rastgelesayi := rand.Intn(90) //  90'a kadar rasgele bir sayi verir
	fmt.Println(rastgelesayi)

}
