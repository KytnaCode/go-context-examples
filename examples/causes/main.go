package main

import (
	"context"
	"errors"
	"time"
)

func main() {
	// Create a base context
	baseCtx := context.Background()

	// Create a context with cancel
	ctx, cancel := context.WithCancelCause(baseCtx)
	finished := make(chan struct{})

	// Create a goroutine to wait for the context to be canceled
	go func() {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			// ctx.Err() returns the error that caused the cancellation, or nil if there was no error
			err := ctx.Err()
			// Print the error
			println(err.Error())
			// context.Cause returns the underlying cause of the cancellation, or nil if the Context is not canceled
			cause := context.Cause(ctx)
			// Print the cause
			println(cause.Error())
			finished <- struct{}{}
		}

	}()

	// Cancel the context with an error
	// ctx.Err() always will return context.Canceled error
	// context.Cause(ctx) will return the error passed to cancel
	cancel(errors.New("manually canceled"))
	<-finished
}
