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
		"http://amazon.ca",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// //inifite loop
	// for {
	// 	go checkLink(<-c,c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //Passing l into the function litearl will make a copy of l in memory so it can be used by the go routine
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
