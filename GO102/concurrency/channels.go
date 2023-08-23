package main

import (
	"fmt"
	"time"
)

//KANALLAR

func main() {

	/*
			k := make(chan string)

			go func() {
				k <- "kanaldan gelen veri"
			}()

			mesaj := <-k

			fmt.Println(mesaj)



			a := make(chan int)


		go func() {
			a <- 5

			time.Sleep(1 * time.Second)

			a <- 9

			time.Sleep(1 * time.Second)

			a <- 18

		}()

		fmt.Println(<-a, <-a, <-a)
	*/

	isimler := []string{"beyza", "yusuf", "ibrahim", "helin"}

	isimKanal := make(chan string)

	// gönderen yapı
	go func(dizi []string) {
		for _, isim := range dizi {
			isimKanal <- isim
		}
	}(isimler)

	go func() {
		for i := 0; i < len(isimler); i++ {
			alinan := <-isimKanal
			fmt.Println("kanaldan gelenler", alinan)

		}
	}()

	<-time.After(time.Second * 3) // kanaldan veriler geldikten 3 sn sonra başlat,

}
