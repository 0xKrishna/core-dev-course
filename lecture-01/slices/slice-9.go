package slices

import "fmt"

// Shift all slice elements by 1 to the left. Zero index element becomes the last one, the first - zero, the last - penultimate.

func Slice_9(s []int) []int {
	fmt.Println("Shift all slice elements by 1 to the left. Zero index element becomes the last one, the first - zero, the last - penultimate.")
	fmt.Println("Slice_9 - Slice before shifting:", s)

	if len(s) > 0 {
		zero := s[0]

		for i := 0; i < len(s)-1; i++ {
			s[i] = s[i+1]
		}

		s[len(s)-1] = zero
	}

	fmt.Println("Slice_9 - Slice after shifting:", s)

	return s
}
