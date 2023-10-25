package main

import "fmt"

func main() {
	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Even(ch1)
	go Odd(ch2)

	for num := range ar {
		if num%2 == 0 {
			ch1 <- num
		} else {
			ch2 <- num
		}
	}

}
func Odd(ch chan int) {
	for x := range ch {
		fmt.Printf("Odd:%d\n", x)
	}
}
func Even(ch chan int) {
	for x := range ch {
		fmt.Printf("Even:%d\n", x)
	}
}
