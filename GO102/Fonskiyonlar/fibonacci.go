package main

import "fmt"

func fibo(n int) int {
	a := 0
	b := 1

	for i := 0; i <= n; i++ {
		gecici := a
		a = b
		b = gecici + a
	}
	return a
}

func main() {

	for n := 0; n <= 20; n++ {
		sonuc := fibo(n)
		fmt.Printf("%d = %d", n, sonuc)
	}
}
