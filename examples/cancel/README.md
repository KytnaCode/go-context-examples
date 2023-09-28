# go-context-examples - Cancel

In this example, we will create a context that can be canceled using the `context.WithCancel()` function.

## Why use `context.WithCancel()`?

The `context.WithCancel()` function is used when you want to create a context that can be canceled. This is useful when you want to cancel a request or a function call, depending on the context.

## Run the example

There's two ways to run the example. 

### First way, using the `run.sh` script

Go to the root of the repository and run the following command:

```bash
./scripts/run.sh cancel
```

or for running all the examples:

```bash
./scripts/run.sh
```

> Note: If you want to run the example from Windows, use the second way.

### Second way, using `go run`

To run the example, either bash or batch, run the following command:

```bash
go run main.go
```

## General Pseudo Code

```
MAIN
    Create base context
    Create context with cancel
    SomeFunction(Context)
    
    ...
    
    Cancel context
    SomeFunction returns
   
SomeFunction(Context)
    Do something
    Wait for context to be canceled
    Interrupt the function
```

## Example code
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    baseCtx := context.Background()

    ctx, cancel := context.WithCancel(baseCtx) 

    go SomeFunction(ctx)

    time.Sleep(1 * time.Second)

    cancel()
}

func SomeFunction(ctx context.Context) {
    fmt.Println("Doing something")

    select {
    case <-ctx.Done():
        fmt.Println("Context canceled")
    }
}
```
