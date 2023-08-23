package main

import "fmt"

// bu hatatyı gördüğünüzde kanal kapasitesini arttır

func main() {

	/*
		ch:=make(chan int,2)

		ch<-1

		fmt.Println(<-ch)
	*/

	kanal := make(chan string, 5)

	kanal <- "AUDI"
	kanal <- "FORD"
	kanal <- "KIA"
	kanal <- "LADA"

	fmt.Printf("Kanal Kapasitesi%d\n", cap(kanal))
	fmt.Printf("Gelen Veri Boyutu:%d\n", len(kanal))
	fmt.Printf("Alinan Veri:%s\n", <-kanal)
	fmt.Printf("Yeni Gelen Veri Boyutu:%d\n", len(kanal))

}
