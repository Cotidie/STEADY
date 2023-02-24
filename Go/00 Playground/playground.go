package main

import (
	"fmt"
	"log"
	"net/url"
)

// Defer keyword works in function scope.
func deferTiming() {
	if true {
		defer fmt.Println("Defer executed!")
	}
	fmt.Println("if statement ended")
}

// Conclusion: it needs double slash(//) to parse Host. Otherwise Host will be empty
// Example: localhost/some/path => Host: "", Path: "localhost/some/path"
func urlParsing() {
	endpoint := "localhost/some/path"
	url, err := url.Parse(endpoint)
	if err != nil {
		log.Panicf("Filed to parse URL: %v\n", err)
	}

	fmt.Printf("Host: %s\n Path: %s\n", url.Host, url.Path)
}

func main() {
	deferTiming()
}
