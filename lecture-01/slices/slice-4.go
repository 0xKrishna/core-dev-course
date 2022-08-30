package slices

import "fmt"

// Take the last number of slice, return it to the user, and remove this element from slice.

func Slice_4(s []int) ([]int, int) {
	fmt.Println("Take the last number of slice, return it to the user, and remove this element from slice.")
	fmt.Println("Slice_4 - Slice before taking last element:", s)

	if len(s) > 0 {
		last := s[len(s)-1]
		s = s[:len(s)-1]

		fmt.Println("Slice_4 - Slice after taking last element:", s)
		fmt.Println("Slice_4 - Last element:", last)

		return s, last
	}

	return s, 0
}
