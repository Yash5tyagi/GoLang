package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}
func main() {
	var mutex sync.Mutex
	msg = "hello world"
	wg.Add(2)
	go updateMessage("change 1", &mutex)
	go updateMessage("change 2", &mutex)
	wg.Wait()
	fmt.Println(msg)
}
