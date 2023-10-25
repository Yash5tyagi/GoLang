package main

import "fmt"

func fibonacci(ch chan int, quit chan bool) {
	x := 0
	y := 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 5
	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch)
		}
		quit <- false
	}(n)
	fibonacci(ch, quit)
}
