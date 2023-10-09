package main

import "fmt"

type bot interface {
	getgreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getgreeting() string {
	return "Hi There!"
}
func (spanishBot) getgreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getgreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
