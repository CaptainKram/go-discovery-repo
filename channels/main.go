package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://habitica.com/",
		"https://adventofcode.com",
		"https://github.com/",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// we can add argument link string but there's no need for current version
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
