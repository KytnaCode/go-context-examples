package main

import "context"

// an example function that takes a context
func doSomething(ctx context.Context) {
	println("doing something")

	// ... do something with ctx ...
}

func main() {
	// This will work, but it's not the best way to use context.
	doSomething(nil)

	// context.TODO() returns a non-nil, empty Context.
	// Code should use context.TODO when it's unclear which Context to use or it is not yet available
	doSomething(context.TODO())

	// context.Background() returns a non-nil, empty Context.
	// It is never canceled, has no values, and has no deadline.
	// It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests
	doSomething(context.Background())
}
