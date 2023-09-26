# go-context-examples - Empty

In this example, we will create an empty context using the `context.Background()` and `context.TODO()` functions.

## Why not use nil?

The `context.Background()` and `context.TODO()` functions 
return a context that is empty. This means that it does not contain 
any data or cancellation signal. This is useful when you want 
to create a context that does not contain any data or 
cancellation signal.

## `context.TODO()`

`context.TODO()` is used when you want to create a context that does 
not contain any data or cancellation signal. this is useful as a 
placeholder when you want to create a context but you don't know 
what data or cancellation signal to use yet.

## `context.Background()`

`context.Background()` is used when you want an empty context 
either for adding data or

## `context.TODO()` vs `context.Background()`

`context.TODO()` is used as a placeholder or when it's not clear what 
data or cancellation signal to use yet, on the other hand, 
`context.Background()` is used when you want an empty context,
is commonly used for `main` functions and tests.

## Run the example

To run the example, either bash or batch, run the following command:

```bash
go run main.go
```