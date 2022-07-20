package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // = contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func printStruct() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "daily142857@gmail.com",
			zipCode: 94061,
		},
	}

	fmt.Printf("%+v", alex)
}

func printMap() {
	colors := make(map[string]int)
	colors["red"] = 1
	colors["green"] = 2
	colors["white"] = 3

	for color, value := range colors {
		fmt.Println(color, ":", value)
	}
}

func main() {
	getGoogleHTMLandLog()
}
