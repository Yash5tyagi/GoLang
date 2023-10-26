package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go func(ch chan int) {
		for i := 0; i < 50; i++ {
			ch <- i
			fmt.Printf("value sent:%d\n", i)

		}
		close(ch)
	}(ch)
	for i := range ch {
		fmt.Printf("value received:%d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Done!!")
}
