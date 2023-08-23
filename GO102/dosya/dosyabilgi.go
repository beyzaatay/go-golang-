package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	dosyaBilgi, err := os.Stat("yenidosya.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("dosya adı:", dosyaBilgi.Name())
	fmt.Println("dosya izinleri:", dosyaBilgi.Mode())
	fmt.Println("dosya boyutu:", dosyaBilgi.Size())
	fmt.Println("değiştirme tarihi:", dosyaBilgi.ModTime())
	fmt.Println("klasör mü:", dosyaBilgi.IsDir())
	fmt.Println("sistem interface:", dosyaBilgi.Sys())

	if dosyaBilgi.Size() > 90 {
		os.Remove("yenidosya.txt")
	}
}
