package slices

import "fmt"

// Swap all even index elements with the nearest odd indices. eg. 0 and 1, 2 and 3, 4 and 5 etc.

func Slice_14(s []int) []int {
	fmt.Println("Swap all even index elements with the nearest odd indices. eg. 0 and 1, 2 and 3, 4 and 5 etc.")
	fmt.Println("Slice_14 - Slice before swaping:", s)

	if len(s) > 0 {
		for i := 0; i < len(s)-1; i += 2 {
			s[i], s[i+1] = s[i+1], s[i]
		}
	}

	fmt.Println("Slice_14 - Slice after swaping:", s)

	return s
}
