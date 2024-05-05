# Chapter 1

## Your First Program

As part of the tradition, the first program will be to print out "Hello World!" in your terminal.

```shell
mkdir chapter-01
cd chapter-01
touch main.go
```

After we have created our first program - "Hello World!", let's compile and run the program:

```shell
go run main.go
```

The `go run` command compiles the program into an executable saved in a temporary directory and then runs it. In case you want to learn more about the comamnd you can simply execute `go help run` in the terminal.

Another useful command is `go doc fmt Println`, which prints the documentation comments associated with `Println` from the `fmt` package:

```text
package fmt // import "fmt"

func Println(a ...any) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
```

## Problems

### 1. What is a whitespace?

Newlines, spaces and tabs are known as whitespace (because you cannot see them).

### 2. What is a comment? What are the two ways of writing a comment?

Comments are ignored by the Go compiler and are there for your own sake.
There are 2 types of comments:

- `//` - single-line commments
- `/* */` - multi-line comments

### 3. Our program began with `package main`. What would the files in the `fmt` package begin with?

`package fmt`

### 4. We used the `Println` function defined in the `fmt` package. If we wanted to use the `Exit` funtion from the `os` package what would we need to do?

`os.Exit()`

### 5. Modify the program we wrote so that instead of printing `Hello World!` it prints `Hello, my name is` followed by your name

```go
package main

import "fmt"

// this is a comment

func main() {
    fmt.Println("Hello, my name is MV.")
}
```
