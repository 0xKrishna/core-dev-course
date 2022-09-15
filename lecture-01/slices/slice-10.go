package slices

import "fmt"

// Same as 9, but shift by user-specified i.

func Slice_10(s []int, i int) []int {
	fmt.Println("Same as 9, but shift by user-specified i")
	fmt.Println("Slice_10 - Slice before shifting:", s)
	fmt.Println("Slice_10 - Shift by:", i)

	if len(s) > 0 {
		newArr := make([]int, len(s))
		for j := 0; j < len(s); j++ {
			newArr[j] = s[(j+i)%len(s)]
		}

		s = newArr
	}

	fmt.Println("Slice_10 - Slice after shifting:", s)

	return s
}
