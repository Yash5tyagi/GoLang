package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	x := 2
	var mt sync.Mutex

	go func(ch chan int) {
		mt.Lock()
		x++
		fmt.Printf("value sent first:%d\n", x)
		time.Sleep(2 * time.Second)
		ch <- x
		time.Sleep(2 * time.Second)
		fmt.Println("check")
		mt.Unlock()

	}(ch)

	go func(ch chan int) {
		mt.Lock()
		x = x * x
		fmt.Printf("value sent second:%d\n", x)
		ch <- x
		mt.Unlock()

	}(ch)

	fmt.Printf("value received first:%d\n", <-ch)

	fmt.Printf("value received second:%d\n", <-ch)

}
