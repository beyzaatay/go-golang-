package main

import (
	"fmt"
	"sync"
)

// goroutine i belli noktalarda durdurup çalıştırabilen bir struct yapısı
func main() {

	wg := sync.WaitGroup{}
	wg.Add(1) // 1 tane goroutine i etkileyebilir

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine")
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("ana thread")

}
