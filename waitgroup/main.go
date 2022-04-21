package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3) // Add numbers of goroutine to wait

	start := time.Now()
	go doSomething(wg)
	go doSomething(wg)
	go doSomething(wg)

	wg.Wait() // Tell Main to wait until wg counter is zero

	fmt.Println(time.Since(start))

}

func doSomething(wg *sync.WaitGroup) {
	fmt.Println("Hello, do something...")
	time.Sleep(120 * time.Millisecond)
	wg.Done() // Tell Main that I'm done
}
