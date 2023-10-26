package main

import (
	"fmt"
	"sync"
	"time"
)

func usingMutChan() {
	ordr := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, m *sync.Mutex) {
			defer wg.Done()
			m.Lock()
			fmt.Printf("Go Routine number:%d Order number:%d\n", i, ordr)
			ordr++
			m.Unlock()
		}(i, &mut)

	}
	wg.Wait()
}
func usingChan() {
	ordr := 0
	ch := make(chan int, 10)
	go func(n int, ch chan int) {

		for i := 0; i < n; i++ {
			time.Sleep(2 * time.Second)
			ch <- i
		}
		close(ch)
	}(cap(ch), ch)
	for c := range ch {
		fmt.Printf("Go Routine number:%d Order number:%d\n", c, ordr)
		ordr++
	}
}
func main() {
	usingMutChan()
	fmt.Println()
	usingChan()
}
