package slices

import "fmt"

// Take the i-th number of slice, return it to the user, and delete this element from slice.
// The number i is passed by the user to the function. Try to do it without allocating a new slice.

func Slice_6(s []int, i int) ([]int, int) {
	fmt.Println("Take the i-th number of slice, return it to the user, and delete this element from slice.")
	fmt.Println("Slice_6 - Slice before taking i-th element:", s)
	fmt.Println("Slice_6 - i:", i)

	if len(s) > i {
		number := s[i]
		s = append(s[0:i], s[i+1:]...)

		fmt.Println("Slice_6 - Slice after taking i-th element:", s)
		fmt.Println("Slice_6 - i-th element:", number)

		return s, number
	}

	return s, 0
}
