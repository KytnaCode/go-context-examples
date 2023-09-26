package main

import (
	"context"
	"fmt"
)

func doSomethingElse(ctx context.Context) {
	// the value associated with "myKey" is 'myOtherValue' here
	fmt.Printf("doSomethingElse: myKey's value is: %s\n", ctx.Value("myKey"))
}

func doSomething(ctx context.Context) {
	// print the value associated with "myKey" in the ctx
	fmt.Printf("doSomething: myKey's value is: %s\n", ctx.Value("myKey"))

	// create a new context from the existing one that has a new value associated with "myKey"
	newCtx := context.WithValue(ctx, "myKey", "myOtherValue")
	doSomethingElse(newCtx)

	// the value associated with "myKey" is still 'myValue' here because context.WithValue returns a copy of parent context
	fmt.Printf("doSomething: myKey's value is: %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()

	// context.WithValue returns a copy of parent in which the value associated with key is val.
	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}
