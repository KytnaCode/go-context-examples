package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	// create a context with cancel
	ctx, cancel := context.WithCancel(ctx)

	// create a channel to print
	printCh := make(chan int)
	go doAnother(ctx, printCh)

	// send 1, 2, 3 to printCh
	for num := 1; num <= 3; num++ {
		time.Sleep(25 * time.Millisecond)
		printCh <- num
	}

	// cancel the context, this will stop the doAnother goroutine
	cancel()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("doSomething: finished")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	// loop until ctx is done
	for {
		select {
		// if ctx is done, return
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Println("doAnother: finished")
			return
		// if printCh has value, print it
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func main() {
	// create an empty context
	ctx := context.Background()

	doSomething(ctx)
}
