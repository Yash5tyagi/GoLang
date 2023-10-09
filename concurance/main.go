package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://golang.org",
		"https://facebook.com",
		"https://stackoverflow.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checklink(link, c)

	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}
func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down!")
		c <- "Maybe its down!"
		return
	}
	fmt.Println(link, " is running fine!")
	c <- "its running fine"
}
