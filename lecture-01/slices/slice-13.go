package slices

import "fmt"

// Return a copy of the passed slice to the user

func Slice_13(s []int) []int {
	fmt.Println("Return a copy of the passed slice to the user")
	fmt.Println("Slice_13 - Original Slice:", s)

	copyArr := make([]int, len(s))
	copy(copyArr, s)

	fmt.Println("Slice_13 - Copy of the Slice:", copyArr)

	return copyArr
}
