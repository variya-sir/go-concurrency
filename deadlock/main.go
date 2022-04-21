package main

import (
	"fmt"
	"sync"
)

func hello(wgrp *sync.WaitGroup) {
	fmt.Println("Hello")
	wgrp.Done() /////// removing the wgrp.Done will cause a deadlock
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2) ///// 2 as the value will cause a deadlock
	go hello(&wg)
	wg.Wait() ////// blocks execution until the goroutine finishes
	fmt.Println("main function")
}
