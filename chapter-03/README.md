# Chapter 3

## Variables

Variables in Go are created bz first using the `var` keyword, then specifying the variable name (e.g. `x`), the type (e.g. `string`) and finally assigning a value to the variable (e.g. `Hello World`).

### How to Name a Variable

Names must start with a letter and may contain letters, numbers or the `_`(underscore) symbol. The Go compiler doesn't care what you name a variable so the name is meant for the developer's benefit. Pick names which clearly describe the variable's purpose.

According to ["Effective Go: Mixed Caps"](https://go.dev/doc/effective_go#mixed-caps):

```text
Finally, the convention in Go is to use MixedCaps or mixedCaps rather than underscores to write multiword names.
```

### Scope

"Go is lexically scoped using blocks". Basically this means that the variable exists within the nearest curly braces `{}` (a block) including any nested curly braces (blocks), but not outside of them.

### Constants

Go also has support for constants, which are basically vaariables whose values cannot be changed later. They are created in the same way you create variables but instead of using `var` keyword we use the `const` keyword.

```go
package main

import "fmt"

func main() {
    const x string = "Hello World"
    fmt.Println(x)
}
```

Trying to overwite the variable will result in a compile-time error:

```text
./main.go:8:2: cannot assign to x (neither addressable nor a map index expression)
```

Constants are a good way to reuse common values in a program without writing them out each time. For instance, `Pi` in the `math` package is defined as a constant.

### Defining Multiple Variables

Go also has another shorthand when you need to define multiple variables:

```go
var (
    a = 1
    b = 24
    c = 56
)
```

Use the keyword `var` or `const` followed by parentheses with each variable on its own line.

## Problems

### 1. What are two ways to create a new variable?

```go
// first way: var <variableName> <type> = <value>
var x string = "Hello"
// second way: <variableName> := <value>
y := "World"
```

### 2. What is the value of `x` after running: `x := 5; x += 1`?

x = 6

### 3. What is scope and how do you determine the scope of a variable in Go?

"Go is lexically scoped using blocks". Basically this means that the variable exists within the nearest curly braces `{}` (a block) including any nested curly braces (blocks), but not outside of them.

### 4. What is the difference between `var` and `const`?

`var` is used to create variables (values can change), whereas `const` is used to create constants (values cannot change).

### 5. Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (`C = (F - 32) * 5/9`)

```go
package main

import "fmt"

func main() {
    fmt.Print("Enter a degrees in Fahrenheit: ")
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit)

    celsius := (fahrenheit - 32) * 5 / 9
    fmt.Println(celsius)
}
```

### 6. Write another program that converts from feet into meters. (`1 ft = 0.3048m`)

```go
package main

import "fmt"

func main() {
    fmt.Print("Enter length (in ft): ")
    var feet float64
    fmt.Scanf("%f", &feet)

    meters := feet * 0.3048
    fmt.Println(meters)
}
```
