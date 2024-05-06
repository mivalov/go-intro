package main

import "fmt"

func forLoop1() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
}

func forLoop2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func printOddOrEven() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

func problem1() {
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func problem2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func problem3() {
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

func main() {
	forLoop1()
	forLoop2()
	printOddOrEven()
	problem1()
	problem2()
	problem3()
}
