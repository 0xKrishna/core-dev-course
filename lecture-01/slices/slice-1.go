package slices

import "fmt"

// Add given int(N) to each of []int elements.

func Slice_1(s []int, N int) []int {
	fmt.Println("Add given int(N) to each of []int elements.")
	fmt.Println("Slice_1 - Slice before adding N:", s)
	fmt.Println("Slice_1 - N:", N)

	for i := range s {
		s[i] += N
	}

	fmt.Println("Slice_1 - Slice after adding N:", s)

	return s
}
