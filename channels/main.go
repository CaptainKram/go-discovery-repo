package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " might be down"
		return
	}
	c <- link + " it's up"
}
