# Chapter 3

## Types

Go has several built-in data types.

### Numbers

Go has several different types to represent numbers.
Generally we split numbers into two different kinds:
- integers - numbers without a decimal component (e.g. -3, 0, 2, ...)
- floating-point numbers - numbers with a decimal component (e.g. 1.23, 123.4, 0.000123)

Go has several operators>
| operator | description |
| :---: | :--- |
| + | addition |
| - | substraction |
| * | multiplication |
| / | division |
| % | remainder |

#### Integers

Go's integer types are:
- `uint8` aka `byte` - unsigned 8 bit integer (positive numbers or zero)
- `uint16` - unsigned 16 bit integer
- `uint32` - unsigned 32 bit integer
- `uint64` - unsigned 64 bit integer
- `int8` - signed 8 bit integer (positive, negative numbers, or zero)
- `int16` - signed 16 bit integer
- `int32` aka `rune` - signed 32 bit integer
- `int64` - signed 64 bit integer
- machine dependent - size depends on the architecture type:
  - `uint`
  - `uintptr`
  - `int` - generally recommended

#### Floating Point Numbers

- Floating point numbers are inexact. Occasionally it's not possible to represent a number (e.g. 1.01 - 0.99 = 0.020000000000018).
- Like integers, floating point numbers also have a certain size (32 bit or 64 bit). Using a larger sized floating point number increases it's precision.
- In addition to numbers there are several other numbers which can be represented:
  - "NaN" (not a number) for things like 0/0
  - positive and negative infinity

Go has two floating point types:
- `float32` - single precision
- `float64` - double precision
- `complex64`
- `complex128`

### Strings

A string is a sequence of characters with a definite length used to represent text. Go strings are made up of individual bytes, usually one for each character.

String literals can be created using double quotes (e.g. "Hello World") or back ticks (e.g. `Hello World`). The difference is that double quoted strings cannot contain newlines and they allow special escape sequences. For example `\n` gets replaced with a newline and `\t` gets replaced with a tab character.

There are several common operations on strings:
- to find the length of a string: `len("Hello World")`
- to access an individual character in the string: `"Hello World"[1]`
- to concatenate two strings together: `"Hello " + "World"`

### Booleans

A boolean value is a special 1 bit integer type used to represent true and false (or on and off).

Three logical operators are used with boolean values:

|Operator| Description |
| :---: | :--- |
| `&&` | and |
| `\|\|` | or |
| `!` | not |

| Expression | Value |
| :---: | :--- |
| true && true | true |
| true && false | false |
| false && true | false |
| false && false | false |

| Expression | Value |
| :---: | :--- |
| true \|\| true | true |
| true \|\| false | true |
| false \|\| true | true |
| false \|\| false | false |

| Expression | Value |
| :---: | :--- |
| !true | false |
| !false | true |

## Problems

### 1. How are integers stored on a computer?
Computers use a base-2 binary system - 0 and 1.

### 2. We know that (in base 10) the largest 1 digit number is 9 and the largest 2 digit number is 99. Given that in binarz the largest 2 digit number is 11 (3), the largest 3 digit number is 111 (7) and the largest 4 digit number is 1111 (15). What is the largest 8 digit number?
> Hint: 10^1 - 1 = 9 and 10^2 - 1 = 99

2^8 - 1 = 256 - 1 = 255

### 3. Although overpowered for the task you can use Go as a calculator. Write a program that computes `32132 * 42452` and prints it to the terminal (Use the `*` operator for multiplication)

```go
package main

import "fmt"

func main() {
	fmt.Println("32132 x 42452 =", 32132 * 42452)
}
```

### 4. What is a string? How do you find its length?
A string is a sequence of characters with a definite length used to represent text. A string's length can be found with the `len()` operation.

### 5. What's the value of the expression `(true && false) || (false && true) || !(false && false)`?
(true && false) || (false && true) || !(false && false) = false || false || !(false) = false || true = true
