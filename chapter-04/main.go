package main

import (
	"fmt"
	"math"
)

var x string = "global string variable"

func variables() {
	// variable decration & definition
	var a string = "Hello"
	fmt.Println("a =", a)
	// variable declaration
	var b string
	// variable definition
	b = "World"
	fmt.Println("b =", b)
	fmt.Println("a == b ->", a == b)

	var c string
	c = "first"
	fmt.Println(c)
	c += "second"
	fmt.Println(c)

	// shorthand for var d string = "Hello World"
	d := "Hello World"
	fmt.Println(d)
	e := 5
	fmt.Println(e)

	fmt.Println(x)

	pi := math.Pi
	fmt.Println(pi)
}

func doubleNumber() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}

func FahrenheitToCelsius() {
	fmt.Print("Enter a degrees in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println(celsius)
}

func feetToMeters() {
	fmt.Print("Enter length (in ft): ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := feet * 0.3048
	fmt.Println(meters)
}

func main() {
	variables()
	doubleNumber()
	FahrenheitToCelsius()
	feetToMeters()
}
