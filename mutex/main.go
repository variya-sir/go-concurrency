package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - LearnCodeonline.in")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Append score 1")
		//mut.Lock()
		score = append(score, 1)
		//mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Append score 2")
		//mut.Lock()
		score = append(score, 2)
		//mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Read score")
		//mut.RLock()
		fmt.Println(score)
		//mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
