# go-context-examples - Causes

In this example, we will learn about the causes of a context's cancellation.

## What is a cause?

A cause is the error that caused the context to be canceled. The cause will be determined by the function that canceled
the context.

## Why is it important to know the cause?

Knowing the cause of a context's cancellation is important because it allows you to know what caused the context to be
canceled. This can be useful when debugging a program, and for logging purposes.

## How to specify the cause of a context's cancellation?

To specify the cause of a context's cancellation, you need to use
the `context.WithCancelCause()`, `context.WithTimeoutCause()` or `context.WithDeadlineCause()` functions. This function
takes a context and an error as arguments and returns a new context with the specified error as the cause of the
context's cancellation.

There's an example of how to use `context.WithCancelCause()`

```go
ctx, cancel := context.WithCancelCause(context.Background())
err := errors.New("some error") // create an error
cancel(err) // specify the cause of the context's cancellation
```

## How to get the cause of a context's cancellation?

To get the cause of a context's cancellation, you need to use the `context.Cause()` function. This function takes a
context as an argument and returns the cause of the context's cancellation.

> Note:
>
> The `context.Cause()` return the cause error passed to cancel function, on the other hand, `ctx.Err()`always
> return `context.Canceled`.

There's an example of how to use `context.Cause()`

```go
// create a context with a cause
// ...

fmt.Println(ctx.Err().Error()) // prints "context.Canceled"
fmt.Println(context.Cause(ctx).Error()) // prints "some error"
```

## Run the example

### With run.sh script (Linux/macOS)

Go to the root of the repository and run the following command to run the example:

```bash
chmod +x ./scripts/run.sh
./scripts/run.sh causes
```

or the following command to run all the examples:

```bash
chmod +x ./scripts/run.sh
./scripts/run.sh
```
> Note: For get help run:
> ```bash
> ./scripts/run.sh -h
> ```

### With `go run` command (Linux/macOS/Windows)

Just run the following command, either bash or batch:

```bash
go run main.go # if you are in the examples/causes directory
# go run examples/causes/main.go # if you are in the root of the repository
```

### With Docker (Linux/macOS/Windows)

#### With run-docker.sh script (Linux/macOS)
Go to the root of the repository and run the following command to build the Docker image:

```bash
chmod +x ./scripts/run-docker.sh
./scripts/run-docker.sh
```

#### With `docker build` and `docker run` commands (Linux/macOS/Windows)

Go to the root of the repository and run the following command to build the Docker image:

```bash
docker build -t go-context-examples .
docker run --rm go-context-examples
```
## License

This project is licensed under the CC0-1.0 License - see the [LICENSE](../../LICENSE) file for details.

## Authors

- [**KytnaCode**](https://github.com/KytnaCode)