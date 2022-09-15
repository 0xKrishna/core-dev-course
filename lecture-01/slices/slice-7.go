package slices

import "fmt"

// Merge two slices and return a new one with all elements of the first and second without duplicates.

func Slice_7(a []int, b []int) []int {
	fmt.Println("Merge two slices and return a new one with all elements of the first and second without duplicates.")
	fmt.Println("Slice_7 - Slice a:", a)
	fmt.Println("Slice_7 - Slice b:", b)

	result := make([]int, len(a)+len(b))
	i := 0

	for _, e := range a {
		if !contains(result, e) {
			result[i] = e
			i++
		}
	}

	for _, e := range b {
		if !contains(result, e) {
			result[i] = e
			i++
		}
	}

	fmt.Println("Slice_7 - Result:", result[:i])

	return result[:i]
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
