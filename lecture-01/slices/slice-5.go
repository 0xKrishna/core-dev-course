package slices

import "fmt"

// Take the first slice number, return it to the user, and delete this element from slice.

func Slice_5(s []int) int {
	fmt.Println("Take the first slice number, return it to the user, and delete this element from slice.")
	fmt.Println("Slice_5 - Slice before taking first element:", s)

	if len(s) > 0 {
		first := s[0]
		s = s[1:]

		fmt.Println("Slice_5 - Slice after taking first element:", s)
		fmt.Println("Slice_5 - First element:", first)

		return first
	}

	return 0
}
