//Status Checker of different websites. Porpose is to verify if the domain is able to respond to traffic

package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link) // This is a serial approach
	}
}

func checkLink(link string) {
	_, err := http.Get(link) //Takes a link and makes a http request to it . Then checks if the response is coming of the traffic

	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}

	fmt.Println(link, " is up!")
}
