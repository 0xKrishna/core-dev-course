package slices

import "fmt"

// Same as 9, but right shift

func Slice_11(s []int) []int {
	fmt.Println("Same as 9, but right shift")
	fmt.Println("Slice_11 - Slice before shifting:", s)

	if len(s) > 0 {
		last := s[len(s)-1]

		for i := len(s) - 1; i > 0; i-- {
			s[i] = s[i-1]
		}

		s[0] = last
	}

	fmt.Println("Slice_11 - Slice after shifting:", s)

	return s
}
