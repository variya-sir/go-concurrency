package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Read ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myCh

		fmt.Printf("ChanelOpen: %v\n", isChanelOpen)
		fmt.Printf("Value from Chanel: %v\n", val)

		//fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
