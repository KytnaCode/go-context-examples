# go-context-examples - With Value

In this example, we will create a context with a value using the 
`context.WithValue()` function.

## What is a context value?

A context value is a value that is stored in a context.

## Why use a context value?

With a context value, you can pass data down a call chain 
without having to pass it as a parameter to every function 
in the call chain.

## Example

```go
package main

import (
    "context"
    "fmt"
)

func main() {
    ctx := context.WithValue(context.Background(), "key", "value")
    fmt.Println(ctx.Value("key"))
}
```

Output:

```
value
```