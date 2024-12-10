package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("GoRoutine 1!")
		waitGroup.Done()
	}()

	go func() {
		escrever("GoRoutine 2!")
		waitGroup.Done()
	}()

	go func() {
		escrever("GoRoutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("GoRoutine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
