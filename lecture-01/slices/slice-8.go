package slices

import "fmt"

// From the first slice, remove all the numbers that are in the second.

func Slice_8(a []int, b []int) []int {
	fmt.Println("From the first slice, remove all the numbers that are in the second.")
	fmt.Println("Slice_8 - Slice a before removing numbers from slice b:", a)
	fmt.Println("Slice_8 - Slice b:", b)

	for _, v := range b {
		for i, vv := range a {
			if v == vv {
				a = append(a[:i], a[i+1:]...)
			}
		}
	}

	fmt.Println("Slice_8 - Slice a after removing numbers from slice b:", a)

	return a
}
