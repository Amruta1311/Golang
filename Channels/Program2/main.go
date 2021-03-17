//Status Checker of different websites. Porpose is to verify if the domain is able to respond to traffic

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // created a channel

	for _, link := range links {
		go checkLink(link, c) // This is a parallel approach

		//We only use the go keyword in front of function calls
	}

	// fmt.Println(<-c) //Waits for a value to be sent into the channel. When we get one, log it out immediately
	// for i := 0; i < len(links); i++ {
	// 	//	fmt.Println(<-c) //Until the channel recieves the message, the for loop will be blocked.
	// 	go checkLink(<-c)
	// }

	// for { //Infinite loop
	// 	//	fmt.Println(<-c) //Until the channel recieves the message, the for loop will be blocked.
	// 	go checkLink(<-c, c)
	// }

	for l := range c { //Infinite loop but it waits for the channel to return some value then assign it to l
		//	fmt.Println(<-c) //Until the channel recieves the message, the for loop will be blocked.

		go func(link string) { //for go routine
			time.Sleep(5 * time.Second) //Pauses for 5 secs
			checkLink(link, c)          // This function literal makes sure that the main function does not terminate after 5 seconds
		}(l) //This l is copied to memory for the go routine to work on. Set of parenthesis that actually executes and passes the 'l' as an argument
		//The above syntax takes care that the main function and go routines do access the same memory space and cause trouble
	}
}

//Channels are used to communicate between different running go routines
//Channels make sure that the main routine is aware when each of the child routines have completed their code.
//We create one channel that communicate with all different go routines
//All the essential data that we are trying to communicate via the channel should be necessarily of the same type
//Receiving message from a channel is a blocking operation

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //Takes a link and makes a http request to it . Then checks if the response is coming of the traffic

	if err != nil {
		fmt.Println(link, " might be down!")
		//c <- "Might be down I think!"
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	// c <- "Yep, it is up!"
	c <- link
}
