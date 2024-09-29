package main

import (
	"fmt"
	"slices"
)

func averageTestScoreV1() {
	// array of length 5 of type float64
	var x [5]float64
	// grades
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	// alternative 1-liners
	// var x [5]float64 = [5]float64{98, 93, 77, 82, 83}
	// var x = [5]float64{98, 93, 77, 82, 83}
	// x := [5]float64{98, 93, 77, 82, 83}

	// compute total score
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	// print the average score
	fmt.Println(total / 5)
}

func averageTestScoreV2() {
	// array of length 5 of type float64
	x := [5]float64{98, 93, 77, 82, 83}

	// compute total score
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	// print the average score
	fmt.Println(total / float64(len(x)))
}

func slice_append() {
	slice1 := []int{1, 2, 3}
	slice1 = append(slice1, 4, 5)
	fmt.Println(slice1)
}

func slice_copy() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied_elements := copy(s2, s1)
	fmt.Println(
		"copied", copied_elements,
		"elements from s1", s1,
		"to s2", s2,
	)
}

func task_1() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("The 4th item of arr is", arr[3])
	fmt.Println("The 4th item of s is", s[3])
}

func task_2() {
	s := make([]int, 3, 9)
	// get the length of the slice
	fmt.Println("len(s) =", len(s))
	// get the capacity of the slice
	fmt.Println("cap(s) =", cap(s))
}

func task_3() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("x[2:5] =", x[2:5])
}

func task_4() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var smallest_number int
	for index, value := range x {
		if index == 0 || value < smallest_number {
			smallest_number = value
		}
	}
	fmt.Println("The smallest number in x is", smallest_number)
	fmt.Println("slices.Min(x) =", slices.Min(x))
}

func main() {
	// averageTestScoreV1()
	// averageTestScoreV2()
	// x := make([]float64, 5)
	// fmt.Println(x)

	slice_append()
	// slice_copy()

	task_1()
	task_2()
	task_3()
	task_4()

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
}
