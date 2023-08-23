package main

import (
	"log"
	"os"
)

func main() {
	// err := os.Mkdir("a", os.ModePerm)
	//err := os.Mkdir("a/b", os.ModePerm)
	err := os.Mkdir("../xml", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}
