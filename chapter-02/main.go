package main

import "fmt"

func main() {
	// numbers
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("32132 x 42452 =", 32132*42452)
	// strings
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	// booleans
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println(
		"(true && false) || (false && true) || !(false && false) =",
		(true && false) || (false && true) || !(false && false),
	)
}
