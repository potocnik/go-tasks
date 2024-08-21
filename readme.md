# Go Task List

## Requirements

```
go get -u github.com/jstemmer/go-junit-report
```

Append to `~/.zshrc` file:
```
# GO
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOROOT/bin
```

## Test

```
cd tests
go test ./...
```

Alternatively, `cd` into any directory under `./tests/`, such as `./tests/unit/` or `./tests/integration/` before executing test command to run granular tests.

Generate test report

```
go test -v 2>&1 ./... | $GOBIN/go-junit-report > test-report.xml
```

## Command line

```
cd cmd/console/
```

Run the program with the following options

1. Print tasks:
```go run .```
2. Push:
```go run . push "Task contents"```
3. Pop:
```go run . pop```

## HTTP Server

```
cd cmd/http/
go run ./...
```

## Requirements

### Part 1

10. Create a program using a variadic function to print a list of 10 things To Do. [Variadic Functions][Structures]
11. Create a program to output a list of 10 things To Do in JSON format. [Variadic Functions][Structures][JSON]
12. Create a program using a variadic function to output a list of 10 things To Do to a JSON format file. [Variadic Functions][Structures][JSON]
13. Create a console program to read a list of 10 things To Do from a JSON format file and display. [Variadic Functions][Structures][JSON]
14. Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers, while another updates the same data variable with even numbers. After each update , attempt to display the data contained in the data variable to screen. [Goroutines][Concurrency][Race Conditions]
15. Refactor the program created in exercise 14 to use channels, mutexes to synchronise all actions. [Concurrency][Waitgroups][Workerpools][Mutexes]
16. Create a program that prints a list of things To Do and the current status of the To Do item using two goroutines which alternate between To Do Items and To Do statuses [Concurrency][Waitgroups][Workerpools][Mutexes]

## References

Project structure: [Go Repo Structure](https://gist.github.com/ayoubzulfiqar/9f1a34049332711fddd4d4b2bfd46096#file-folder_structure-md)