package main

import (
	"fmt"
	"sync"
)

func main() {
	cont := []int{}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		cont = append(cont, 1)
		fmt.Println("one")
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		cont = append(cont, 2)
		fmt.Println("two")
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		cont = append(cont, 3)
		fmt.Println("three")
	}(wg)
	wg.Wait()
	fmt.Println(cont)
}
