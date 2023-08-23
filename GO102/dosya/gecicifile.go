package main

import (
	"fmt"
	"log"
	"os"
)

//TEMP

func main() {

	geciciKlasor, err := os.MkdirTemp("", "beyza")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("gecici klasör:", geciciKlasor)

	//işlemler

	//defer os.RemoveAll(geciciKlasor)

	geciciDosya, err := os.CreateTemp("", "beyzaatay")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gecici dosya:", geciciDosya.Name())

	defer os.Remove(geciciDosya.Name())

}
