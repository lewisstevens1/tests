package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"-----",
	}

	// Create a channel so the seperate routines can communicate
	c := make(chan string)

	for _, link := range links {
		// Create a routine for each of the links
		go checkLink(link, c)
	}

	for l := range c {
		// Loop through all the channels constantly.
		go func(link string) {
			// Create an anonymous function call which will pass the link into its own seperate memory address so is not overwritten.
			time.Sleep(time.Second * 5)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("[DOWN]", link)
		c <- link
		return
	}

	fmt.Println("[UP]", link)
	c <- link
}
