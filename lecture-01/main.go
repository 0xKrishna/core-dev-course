package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/0xKrishna/core-dev-course/lecture-01/maps"
	"github.com/0xKrishna/core-dev-course/lecture-01/slices"
)

var fixedSlice1 = []int{1, 2, 3, 4, 5}

func main() {
	fmt.Printf("\nCore Dev Course - Lecture 01\n\n")

	slices.Slice_1(randomIntSlice(5), 2)

	outputSeparator()

	slices.Slice_2(randomIntSlice(5), 3)

	outputSeparator()

	slices.Slice_3(randomIntSlice(5), 4)

	outputSeparator()

	slices.Slice_4(randomIntSlice(5))

	outputSeparator()

	slices.Slice_5(randomIntSlice(5))

	outputSeparator()

	slices.Slice_6(randomIntSlice(5), 3)

	outputSeparator()

	slices.Slice_7(fixedSlice1, []int{3, 4, 5, 6, 7})

	outputSeparator()

	slices.Slice_8(fixedSlice1, []int{2, 3, 4})

	outputSeparator()

	slices.Slice_9(randomIntSlice(5))

	outputSeparator()

	slices.Slice_10(randomIntSlice(5), 2)

	outputSeparator()

	slices.Slice_11(randomIntSlice(5))

	outputSeparator()

	slices.Slice_12(randomIntSlice(5), 3)

	outputSeparator()

	slices.Slice_13(randomIntSlice(5))

	outputSeparator()

	slices.Slice_14(randomIntSlice(5))

	outputSeparator()

	slices.Slice_15(randomIntSlice(5), []string{"Krishna", "Sandeep", "Arpit", "Shivam", "Manav"})

	outputSeparator()

	maps.Map_1("Mississippi")

	outputSeparator()

	maps.Map_2(randomIntSlice(10))

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

func randomIntSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = randomNumber()
	}

	return slice
}

func randomNumber() int {
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		log.Fatal("error generating random number:", err)
	}

	return int(n.Int64())
}
