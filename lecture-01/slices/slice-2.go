package slices

import "fmt"

// Add the number int(N) to the end of the slice.

func Slice_2(s []int, N int) {
	fmt.Println("Add the number int(N) to the end of the slice.")
	fmt.Println("Slice_2 - Slice before adding N:", s)
	fmt.Println("Slice_2 - N:", N)

	if len(s) > 0 {
		s[len(s)-1] += N
	}

	fmt.Println("Slice_2 - Slice after adding N:", s)
}
