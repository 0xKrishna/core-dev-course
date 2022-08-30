package slices

import "fmt"

// Same as 11, but shift by i.

func Slice_12(s []int, i int) []int {
	fmt.Println("Same as 11, but shift by i")
	fmt.Println("Slice_12 - Slice before shifting:", s)
	fmt.Println("Slice_12 - Shift by:", i)

	if len(s) > 0 {
		newArr := make([]int, len(s))
		for j := 0; j < len(s); j++ {
			newArr[j] = s[(j-i+len(s))%len(s)]
		}

		s = newArr
	}

	fmt.Println("Slice_12 - Slice after shifting:", s)

	return s
}
