package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go doSomething()
	go doSomething()
	go doSomething()

	time.Sleep(120 * time.Millisecond)
	fmt.Println(time.Since(start))

}

func doSomething() {
	fmt.Println("Hello, do something...")
	time.Sleep(120 * time.Millisecond)
}
