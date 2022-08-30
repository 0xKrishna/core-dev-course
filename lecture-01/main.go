package main

import (
	"fmt"

	"github.com/0xKrishna/core-dev-course/lecture-01/maps"
	"github.com/0xKrishna/core-dev-course/lecture-01/slices"
)

var input = []int{1, 2, 3, 4, 5}

func main() {
	fmt.Printf("\nCore Dev Course - Lecture 01\n\n")

	slices.Slice_1(input, 2)

	outputSeparator()

	slices.Slice_2(input, 3)

	outputSeparator()

	slices.Slice_3(input, 4)

	outputSeparator()

	slices.Slice_4(input)

	outputSeparator()

	slices.Slice_5(input)

	outputSeparator()

	slices.Slice_6(input, 3)

	outputSeparator()

	slices.Slice_7(input, []int{3, 4, 5, 6, 7})

	outputSeparator()

	slices.Slice_8(input, []int{2, 3, 4})

	outputSeparator()

	slices.Slice_9(input)

	outputSeparator()

	slices.Slice_10(input, 2)

	outputSeparator()

	slices.Slice_11(input)

	outputSeparator()

	slices.Slice_12(input, 3)

	outputSeparator()

	slices.Slice_13(input)

	outputSeparator()

	slices.Slice_14(input)

	outputSeparator()

	slices.Slice_15(input, []string{"Krishna", "Sandeep", "Arpit", "Shivam", "Manav"})

	outputSeparator()

	maps.Map_1("Mississippi")

	outputSeparator()

	maps.Map_2([]int{434, 23, 232, 121, 234})

	outputSeparator()

	maps.Map_3([]int{434, 23, 232, 121, 234}, []int{232, 334, 342, 121, 212})

	outputSeparator()

	maps.Map_4(10)

	outputSeparator()

	maps.Map_5(map[int]map[int]string{
		3: {
			15: "Krishna",
			3:  "Sandeep",
			11: "Arpit",
		},
		8: {
			54: "Manav",
			98: "Arpit",
			5:  "Shivam",
		},
		2: {
			21: "Krishna",
			24: "Manav",
			15: "Shivam",
		},
	})
}

func outputSeparator() {
	fmt.Printf("\n-----------------------------------------------------\n\n")
}
