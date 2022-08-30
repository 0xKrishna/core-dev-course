package slices

import "fmt"

// Add the number int(N) to the beginning of the slice.

func Slice_3(s []int, N int) []int {
	fmt.Println("Add the number int(N) to the beginning of the slice.")
	fmt.Println("Slice_3 - Slice before adding N:", s)
	fmt.Println("Slice_3 - N:", N)

	if len(s) > 0 {
		s[0] += N
	}

	fmt.Println("Slice_3 - Slice after adding N:", s)

	return s
}
