package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var userinput string
	chin := make(chan string)
	chout := make(chan string)
	go shout(chin, chout)
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Please enter something(Q to exit)")
		fmt.Printf("->")
		_, _ = fmt.Scanln(&userinput)
		if strings.ToUpper(userinput) == "Q" {
			break
		}
		chin <- userinput
		fmt.Println(<-chout)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Closing the channels")
	time.Sleep(1 * time.Second)
	close(chin)
	close(chout)
	fmt.Println("Exit")

}

func shout(chin <-chan string, chout chan<- string) {
	for {
		s := <-chin
		chout <- fmt.Sprintf("!!%s!!", strings.ToUpper(s))
	}
}
