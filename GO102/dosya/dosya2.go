package main

import (
	"log"
	"os"
)

func main() {

	//eskiAd := "veri.txt"
	//yeniAd := "data.txt"

	//os.Rename(eskiAd, yeniAd)

	dosyaKonum := "data.txt"
	yenikonum := "ba/veri.txt"

	err := os.Rename(dosyaKonum, yenikonum)
	if err != nil {
		log.Fatal(err)
	}

}
