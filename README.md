go-context-examples
===================

This repository contains examples of how to use the [Go Context](https://golang.org/pkg/context/) package.

What is the Go Context package?
-------------------------------
The Go Context package is a package that allows you to pass data and cancelation signals down a call chain. It is a way to pass data and cancelation signals without having to pass them as parameters to every function in the call chain.

What is a Context?
------------------
A context is an object that contains information about a request. It contains data such as the request ID, the user ID, and the user's IP address. It also contains a cancelation signal that can be used to cancel the request.

How to run the examples
-----------------------
To run the examples, you need to have Go installed. You can download Go from [go.dev](https://go.dev/doc/install).
Next, you need to clone this repository and run the examples. You can do this by running the following commands:

```bash
git clone https://github.com/KytnaCode/go-context-examples.git
cd go-context-examples
```
### Run the examples from Linux/macOS
Run the following command to run all the examples:
```bash
chmod +x ./scripts/run.sh
./scripts/run.sh
```

for get help run:
```bash
./scripts/run.sh -h
```

> Note: If you only want to run a specific example, you can run the following command:
>
> for get examples list run:
> ```bash
> ./scripts/run.sh -l
> ```
> 
> then run:
> 
> ```bash
> ./scripts/run.sh <example>
> ```

### Run the examples from Windows
Just run the following command:
```batch
go run .\examples\<example>\main.go
```

### Run the examples in Docker
For this, you need to have Docker installed. You can download Docker from [docker.com](https://www.docker.com/products/docker-desktop).
Next, you need to build the Docker image and run the Docker container. You can do this by running the following commands:

```bash
chmod +x ./scripts/run-docker.sh
./scripts/run-docker.sh
```

or

```batch
docker build -t go-context-examples .
docker run --rm go-context-examples
```
