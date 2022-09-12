package main

import (
	"fmt"
	"reflect"
)

type Parent interface {
	iamParent() bool
}

type Child struct{}

func (c Child) iamParent() bool {
	return false
}

func checkChild(instance Parent) {
	myType := reflect.TypeOf(instance)

	fmt.Printf("Type: %v\n", myType)

	switch t := instance.(type) {
	default:
		fmt.Println(t)
	}
}

func main() {
	child := Child{}

	checkChild(child)
}
