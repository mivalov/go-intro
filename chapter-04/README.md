# Chapter 4

## Control Structures

### For

The `for` statement allows us to repeat a list of statements (a block of code) multiple times.

```go
func main() {
    i := 1  // store number to print
    for i <= 10 {  // conditional expression
        fmt.Println(i)
        i++  // variable incrementation (same as i += 1)
    }
}
```

Other programming languages have a lot of different types of loops (while, do, until, foreach, etc.) but Go only has one that can be used in a variety of different ways.

Alternative approach to the `for` loop is the following:

```go
func main() {
    // variable initialization; condition; variable incrementation
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

### If

The `if` statement provides a way of choosing to do different things (execute code blocks) based on a condition. It also has an optional `else` part. If the condition evaluates to `true` then the block after the condition is run, otherwise either the block is skipped or if the `else` block is present that block is run. Furthermore, `if` statements can also have one or more `else if` parts as further conditions.

```go
func main() {
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
    }
}
```

### Switch

A switch statement starts with the keyword `switch` followed by an expression and then a series of `cases`. The value of the expression is compared to the expression following each `case` keyword. If they are equivalent then the statement(s) following the `:` is executed.

```go
switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    case 6: fmt.Println("Six")
    case 7: fmt.Println("Seven")
    case 8: fmt.Println("Eight")
    case 9: fmt.Println("Nine")
}
```

Just like an `if` statement, each case is checked from top to bottom and the first one to succeed is chosen. A switch also supports a default case which will happen if none of the cases matches the value (again similar to the `else` part in an `if` statement).

## Problems

### 1. What does the following program print

```go
i := 10
if i > 10 {
    fmt.Println("Big")
} else {
    fmt.Println("Small")
}
```

Output: Small

### 2. Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
}
```

### 3. Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz"

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
```
