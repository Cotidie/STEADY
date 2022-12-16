package main

import (
	"context"
	"errors"
	"log"
)

type myType struct {
	value string
}

func failingUnpacking(ctx context.Context) (*myType, error) {
	rawValue := ctx.Value("NoHello")
	log.Printf("Non existing value: %v\n", rawValue)
	log.Printf("Nil Value: %v\n", nil)
	typeCheck, ok := rawValue.(myType)
	if ok {
		return &typeCheck, nil
	}

	return nil, errors.New("Type Check Failed!")
}

func successUnpacking(ctx context.Context) (*myType, error) {
	rawValue := ctx.Value("Hello")
	typeCheck, ok := rawValue.(myType)

	if ok {
		return &typeCheck, nil
	}

	return nil, errors.New("Type Check Failed!")
}

func main() {
	sample := myType{
		value: "Hello",
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "Hello", sample)

	packFailed, err := failingUnpacking(ctx)
	if err != nil {
		log.Printf("%v\n", err)
	} else {
		log.Printf("Unpacked value: %s\n", packFailed.value)
	}

	packSuccess, err := successUnpacking(ctx)
	if err != nil {
		log.Printf("%v\n", err)
	} else {
		log.Printf("Unpacked value: %s\n", packSuccess.value)
	}
}
